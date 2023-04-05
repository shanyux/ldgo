/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"time"
)

var _zeroTime time.Time

type Time struct {
	d Interface
}

func NewTime(d time.Time) *Time {
	p := &Time{}
	p.Store(d)
	return p
}

func (p *Time) Store(d time.Time)                  { p.d.Store(p.pack(d)) }
func (p *Time) Load() time.Time                    { return p.unpack(p.d.Load()) }
func (p *Time) Swap(new time.Time) (old time.Time) { return p.unpack(p.d.Swap(p.pack(new))) }
func (p *Time) Add(d time.Duration) (new time.Time) {
	return p.MustChange(func(old time.Time) (new time.Time) {
		return old.Add(d)
	})
}

func (p *Time) AddDate(years int, months int, days int) (new time.Time) {
	return p.MustChange(func(old time.Time) (new time.Time) {
		return old.AddDate(years, months, days)
	})
}

func (p *Time) MustChange(change func(old time.Time) (new time.Time)) (new time.Time) {
	for {
		new, swapped := p.Change(change)
		if swapped {
			return new
		}
	}
}

func (p *Time) Change(change func(old time.Time) (new time.Time)) (new time.Time, swapped bool) {
	oldAddr := p.d.Load()
	old := p.unpack(oldAddr)
	new = change(old)
	newAddr := p.pack(new)
	return new, p.d.CompareAndSwap(oldAddr, newAddr)
}

func (p *Time) pack(d time.Time) time.Time { return d }

func (p *Time) unpack(addr interface{}) time.Time {
	x, ok := addr.(time.Time)
	if !ok {
		return _zeroTime
	}
	return x
}
