/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"time"

	"golang.org/x/time/rate"
)

type Reservation struct {
	reservations []*rate.Reservation
}

func (r *Reservation) CancelAt(t time.Time) {
	for _, v := range r.reservations {
		v.CancelAt(t)
	}
}

func (r *Reservation) DelayFrom(t time.Time) time.Duration {
	delay := time.Duration(0)
	for _, v := range r.reservations {
		d := v.DelayFrom(t)
		if delay < d {
			delay = d
		}
	}
	return delay
}
