/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"sync/atomic"
)

type Int32 int32

func (p *Int32) get() *int32 { return (*int32)(p) }

func (p *Int32) Add(delta int32) (new int32) { return atomic.AddInt32(p.get(), delta) }
func (p *Int32) Sub(delta int32) (new int32) { return atomic.AddInt32(p.get(), -delta) }
func (p *Int32) Store(val int32)             { atomic.StoreInt32(p.get(), val) }
func (p *Int32) Load() int32                 { return atomic.LoadInt32(p.get()) }
func (p *Int32) Swap(new int32) (old int32)  { return atomic.SwapInt32(p.get(), new) }
func (p *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(p.get(), old, new)
}

type Int64 int64

func (p *Int64) get() *int64 { return (*int64)(p) }

func (p *Int64) Add(delta int64) (new int64) { return atomic.AddInt64(p.get(), delta) }
func (p *Int64) Sub(delta int64) (new int64) { return atomic.AddInt64(p.get(), -delta) }
func (p *Int64) Store(val int64)             { atomic.StoreInt64(p.get(), val) }
func (p *Int64) Load() int64                 { return atomic.LoadInt64(p.get()) }
func (p *Int64) Swap(new int64) (old int64)  { return atomic.SwapInt64(p.get(), new) }
func (p *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(p.get(), old, new)
}

type Int int64

func (p *Int) get() *Int64 { return (*Int64)(p) }

func (p *Int) Add(delta int) (new int) { return int(p.get().Add(int64(delta))) }
func (p *Int) Sub(delta int) (new int) { return int(p.get().Sub(int64(delta))) }
func (p *Int) Store(val int)           { p.get().Store(int64(val)) }
func (p *Int) Load() int               { return int(p.get().Load()) }
func (p *Int) Swap(new int) (old int)  { return int(p.get().Swap(int64(new))) }
func (p *Int) CompareAndSwap(old, new int) (swapped bool) {
	return p.get().CompareAndSwap(int64(old), int64(new))
}
