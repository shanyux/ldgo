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

func (l *Limiters) Wait(c context.Context) error {
	return l.WaitN(c, 1)
}

func (l *Limiters) WaitN(c context.Context, n int) error {
	return wait(c, l, n)
}

func (l *Limiters) Allow(c context.Context) bool         { return l.AllowN(c, 1) }
func (l *Limiters) AllowN(c context.Context, n int) bool { return allow(c, l, n) }

func (l *Limiters) Reserve(c context.Context) (*Reservation, error) {
	return l.ReserveN(c, 1)
}

func (l *Limiters) ReserveN(c context.Context, n int) (*Reservation, error) {
	limiters := l.limiters

	now := time.Now()
	r := &Reservation{
		reservations: make([]*rate.Reservation, 0, len(limiters)),
	}

	for _, l := range limiters {
		v, err := l.ReserveN(c, n)
		if err != nil {
			r.CancelAt(now)
			return nil, err
		}

		r.reservations = append(r.reservations, v.reservations...)
	}

	return r, nil
}
