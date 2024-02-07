/*
 * Copyright (C) distroy
 */

package ldatomic

func NewAny[T any](v T) *Any[T] {
	p := &Any[T]{}
	p.Store(v)
	return p
}

type Any[T any] struct {
	d     Value
	cased Bool
}

func (p *Any[T]) Store(v T)          { p.d.Store(p.pack(v)) }
func (p *Any[T]) Load() T            { return p.unpack(p.d.Load()) }
func (p *Any[T]) Swap(new T) (old T) { return p.unpack(p.d.Swap(p.pack(new))) }
func (p *Any[T]) CompareAndSwap(old, new T) (swapped bool) {
	// during the first compare and swap, the value may be nil
	if !p.cased.Load() {
		z := zeroOfType[anyData[T]]()
		p.d.CompareAndSwap(nil, z)
		p.cased.Store(true)
	}
	return p.d.CompareAndSwap(p.pack(old), p.pack(new))
}

func (p *Any[T]) pack(d T) anyData[T] { return anyData[T]{Data: d} }
func (p *Any[T]) unpack(i interface{}) T {
	if i == nil {
		return zeroOfType[T]()
	}
	x := i.(anyData[T])
	return x.Data
}

func zeroOfType[T any]() T {
	var x T
	return x
}

type anyData[T any] struct {
	Data T
}
