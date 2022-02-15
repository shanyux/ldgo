/*
 * Copyright (C) distroy
 */

package ldatomic

import "sync/atomic"

type Uint32 uint32

func (p *Uint32) get() *uint32 { return (*uint32)(p) }

func (p *Uint32) Add(delta uint32) (new uint32) { return atomic.AddUint32(p.get(), delta) }
func (p *Uint32) Sub(delta uint32) (new uint32) { return atomic.AddUint32(p.get(), ^(delta - 1)) }
func (p *Uint32) Store(val uint32)              { atomic.StoreUint32(p.get(), val) }
func (p *Uint32) Load() uint32                  { return atomic.LoadUint32(p.get()) }
func (p *Uint32) Swap(new uint32) (old uint32)  { return atomic.SwapUint32(p.get(), new) }
func (p *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(p.get(), old, new)
}

type Uint64 uint64

func (p *Uint64) get() *uint64 { return (*uint64)(p) }

func (p *Uint64) Add(delta uint64) (new uint64) { return atomic.AddUint64(p.get(), delta) }
func (p *Uint64) Sub(delta uint64) (new uint64) { return atomic.AddUint64(p.get(), ^(delta - 1)) }
func (p *Uint64) Store(val uint64)              { atomic.StoreUint64(p.get(), val) }
func (p *Uint64) Load() uint64                  { return atomic.LoadUint64(p.get()) }
func (p *Uint64) Swap(new uint64) (old uint64)  { return atomic.SwapUint64(p.get(), new) }
func (p *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(p.get(), old, new)
}

type Uintptr uintptr

func (p *Uintptr) get() *uintptr { return (*uintptr)(p) }

func (p *Uintptr) Add(delta uintptr) (new uintptr) { return atomic.AddUintptr(p.get(), delta) }
func (p *Uintptr) Sub(delta uintptr) (new uintptr) { return atomic.AddUintptr(p.get(), ^(delta - 1)) }
func (p *Uintptr) Store(val uintptr)               { atomic.StoreUintptr(p.get(), val) }
func (p *Uintptr) Load() uintptr                   { return atomic.LoadUintptr(p.get()) }
func (p *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(p.get(), old, new)
}
