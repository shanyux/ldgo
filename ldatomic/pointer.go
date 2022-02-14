/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"sync/atomic"
	"unsafe"
)

type Pointer struct {
	v unsafe.Pointer
}

func (p *Pointer) get() *unsafe.Pointer { return &p.v }

func (p *Pointer) Store(val unsafe.Pointer) { atomic.StorePointer(p.get(), val) }
func (p *Pointer) Load() unsafe.Pointer     { return atomic.LoadPointer(p.get()) }
func (p *Pointer) Swap(new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(p.get(), new)
}
func (p *Pointer) CompareAndSwap(old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(p.get(), old, new)
}
