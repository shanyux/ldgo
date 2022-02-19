/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"time"
)

type Duration int64

func (p *Duration) get() *Int64 { return (*Int64)(p) }

func (p *Duration) Store(v time.Duration) { p.get().Store(int64(v)) }
func (p *Duration) Load() time.Duration   { return time.Duration(p.get().Load()) }
func (p *Duration) Swap(old time.Duration) (new time.Duration) {
	return time.Duration(p.get().Swap(int64(old)))
}
func (p *Duration) CompareAndSwap(old, new time.Duration) (swapped bool) {
	return p.get().CompareAndSwap(int64(old), int64(new))
}
func (p *Duration) Add(delta time.Duration) (new time.Duration) {
	return time.Duration(p.get().Add(int64(delta)))
}
func (p *Duration) Sub(delta time.Duration) (new time.Duration) {
	return time.Duration(p.get().Sub(int64(delta)))
}
