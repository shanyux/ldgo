/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"golang.org/x/time/rate"
)

var (
	_ LimiterIface = (*Limiter)(nil)
	_ LimiterIface = (*Limiters)(nil)
)

type LimiterIface interface {
	Wait(ctx ldctx.Context) lderr.Error
	WaitN(ctx ldctx.Context, n int) lderr.Error

	Reserve(ctx ldctx.Context) (*Reservation, lderr.Error)
	ReserveN(ctx ldctx.Context, n int) (*Reservation, lderr.Error)
}

type Limiters struct {
	limiters []LimiterIface
}

func NewLimiters(limiters ...LimiterIface) *Limiters {
	l := &Limiters{}
	l.AddLimiter(limiters...)
	return l
}

func (l *Limiters) AddLimiter(limiters ...LimiterIface) *Limiters {
	// l.limiters = append(l.limiters, limiters...)

	count := 0
	for _, limiter := range limiters {
		n := 1
		if v, ok := limiter.(*Limiters); ok {
			n = len(v.limiters)
		}
		count += n
	}

	buf := make([]LimiterIface, 0, len(l.limiters)+count)
	buf = append(buf, l.limiters...)

	for _, limiter := range limiters {
		if v, ok := limiter.(*Limiters); ok {
			buf = append(buf, v.limiters...)
		} else {
			buf = append(buf, limiter)
		}
	}

	l.limiters = buf
	return l
}

func (l *Limiters) Wait(ctx ldctx.Context) lderr.Error {
	return l.WaitN(ctx, 1)
}

func (l *Limiters) WaitN(ctx ldctx.Context, n int) lderr.Error {
	select {
	case <-ctx.Done():
		return lderr.ErrCtxCanceled
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
		return lderr.ErrCtxDeadlineExceeded
	}

	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		// We can proceed.
		return nil

	case <-ctx.Done():
		break
	}

	// Context was canceled before we could proceed.  Cancel the
	// reservation, which may permit other events to proceed sooner.
	r.Cancel()
	return lderr.ErrCtxCanceled
}

func (l *Limiters) Reserve(ctx ldctx.Context) (*Reservation, lderr.Error) {
	return l.ReserveN(ctx, 1)
}

func (l *Limiters) ReserveN(ctx ldctx.Context, n int) (*Reservation, lderr.Error) {
	limiters := l.limiters
	r := &Reservation{
		reservations: make([]*rate.Reservation, 0, len(limiters)),
	}

	for _, l := range limiters {
		v, err := l.ReserveN(ctx, n)
		if err != nil {
			r.Cancel()
			return nil, err
		}

		r.reservations = append(r.reservations, v.reservations...)
	}

	return r, nil
}
