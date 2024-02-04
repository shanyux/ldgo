/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"time"
)

var _zeroTime time.Time

type Time struct {
	d Any[time.Time]
}

func NewTime(d time.Time) *Time {
	p := &Time{}
	p.Store(d)
	return p
}

func (p *Time) Store(d time.Time)                  { p.d.Store(d) }
func (p *Time) Load() time.Time                    { return p.d.Load() }
func (p *Time) Swap(new time.Time) (old time.Time) { return p.d.Swap(new) }
func (p *Time) CompareAndSwap(old, new time.Time) (swapped bool) {
	return p.d.CompareAndSwap(old, new)
}

func (p *Time) MustChange(change func(old time.Time) (new time.Time)) (new time.Time) {
	for {
		new, swapped := p.Change(change)
		if swapped {
			return new
		}
	}
}

func (p *Time) Change(change func(old time.Time) (new time.Time)) (new time.Time, changed bool) {
	old := p.Load()
	new = change(old)
	return new, p.d.CompareAndSwap(old, new)
}
