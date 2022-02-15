/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"time"
	"unsafe"
)

var _zeroTime time.Time

type Time struct {
	v Pointer
}

func (p *Time) Store(v time.Time)                  { p.v.Store(p.toAddr(v)) }
func (p *Time) Load() time.Time                    { return p.toTime(p.v.Load()) }
func (p *Time) Swap(new time.Time) (old time.Time) { return p.toTime(p.v.Swap(p.toAddr(new))) }
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
	oldAddr := p.v.Load()
	old := p.toTime(oldAddr)
	new = change(old)
	newAddr := p.toAddr(new)
	return new, p.v.CompareAndSwap(oldAddr, newAddr)
}

func (p *Time) toAddr(v time.Time) unsafe.Pointer { return unsafe.Pointer(&v) }

func (p *Time) toTime(addr unsafe.Pointer) time.Time {
	if addr == nil {
		return _zeroTime
	}
	x := (*time.Time)(addr)
	return *x
}
