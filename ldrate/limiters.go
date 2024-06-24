/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type Limiters struct {
	limiters []ILimiter
}

func NewLimiters(limiters ...ILimiter) *Limiters {
	l := &Limiters{}
	l.AddLimiter(limiters...)
	return l
}

func (l *Limiters) AddLimiter(limiters ...ILimiter) *Limiters {
	// l.limiters = append(l.limiters, limiters...)

	count := 0
	for _, limiter := range limiters {
		n := 1
		if v, ok := limiter.(*Limiters); ok {
			n = len(v.limiters)
		}
		count += n
	}

	buf := make([]ILimiter, 0, len(l.limiters)+count)
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

func (l *Limiters) Wait(ctx context.Context) error {
	return l.WaitN(ctx, 1)
}

func (l *Limiters) WaitN(ctx context.Context, n int) error {
	return wait(ctx, l, n)
}

func (l *Limiters) Allow(ctx context.Context) bool         { return l.AllowN(ctx, 1) }
func (l *Limiters) AllowN(ctx context.Context, n int) bool { return allow(ctx, l, n) }

func (l *Limiters) Reserve(ctx context.Context) (*Reservation, error) {
	return l.ReserveN(ctx, 1)
}

func (l *Limiters) ReserveN(ctx context.Context, n int) (*Reservation, error) {
	limiters := l.limiters

	now := time.Now()
	r := &Reservation{
		reservations: make([]*rate.Reservation, 0, len(limiters)),
	}

	for _, l := range limiters {
		v, err := l.ReserveN(ctx, n)
		if err != nil {
			r.CancelAt(now)
			return nil, err
		}

		r.reservations = append(r.reservations, v.reservations...)
	}

	return r, nil
}
