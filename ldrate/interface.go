/*
 * Copyright (C) distroy
 */

package ldrate

import (
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
	Wait(ctx ldctx.Context) lderr.Error
	WaitN(ctx ldctx.Context, n int) lderr.Error

	Reserve(ctx ldctx.Context) (*Reservation, lderr.Error)
	ReserveN(ctx ldctx.Context, n int) (*Reservation, lderr.Error)
}

type reserver interface {
	ReserveN(ctx ldctx.Context, n int) (*Reservation, lderr.Error)
}

func wait(ctx ldctx.Context, l reserver, n int) lderr.Error {
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
		r.Cancel()
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
		r.Cancel()
		return ldctx.GetError(ctx)
	}

	// We can proceed.
	return nil
}
