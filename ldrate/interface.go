/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldctx"
	"github.com/distroy/ldgo/v2/lderr"
	"golang.org/x/time/rate"
)

var (
	_ ILimiter = (*Limiter)(nil)
	_ ILimiter = (*Limiters)(nil)
)

type ILimiter interface {
	Wait(ctx context.Context) error
	WaitN(ctx context.Context, n int) error

	Reserve(ctx context.Context) (*Reservation, error)
	ReserveN(ctx context.Context, n int) (*Reservation, error)
}

type reserver interface {
	ReserveN(ctx context.Context, n int) (*Reservation, error)
}

func wait(ctx context.Context, l reserver, n int) error {
	select {
	case <-ctx.Done():
		// return lderr.ErrCtxCanceled
		return ldctx.GetError(ctx)
	default:
	}

	now := time.Now()

	// Determine wait limit
	waitLimit := rate.InfDuration
	if deadline, ok := ctx.Deadline(); ok {
		waitLimit = deadline.Sub(now)
	}

	r, err := l.ReserveN(ctx, n)
	if err != nil {
		return err
	}

	// Wait if necessary
	delay := r.DelayFrom(now)
	if delay <= 0 {
		return nil
	}

	if delay >= waitLimit {
		r.CancelAt(now)
		return lderr.ErrCtxDeadlineNotEnough
	}

	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		break

	case <-ctx.Done():
		// Context was canceled before we could proceed.  Cancel the
		// reservation, which may permit other events to proceed sooner.
		r.CancelAt(now)
		return ldctx.GetError(ctx)
	}

	// We can proceed.
	return nil
}

func allow(ctx context.Context, l reserver, n int) bool {
	r, err := l.ReserveN(ctx, n)
	if err != nil {
		return false
	}

	// Wait if necessary
	now := time.Now()
	delay := r.DelayFrom(now)
	// ldctx.LogI(ctx, " ** allow delay", zap.Duration("delay", delay))
	if delay <= 0 {
		return true
	}

	r.CancelAt(now)
	return false
}
