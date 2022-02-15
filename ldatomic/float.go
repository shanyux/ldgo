/*
 * Copyright (C) distroy
 */

package ldatomic

import "math"

type Float32 struct {
	v Uint32
}

func (p *Float32) Store(v float32) { p.v.Store(math.Float32bits(v)) }
func (p *Float32) Load() float32   { return math.Float32frombits(p.v.Load()) }
func (p *Float32) Swap(new float32) (old float32) {
	return math.Float32frombits(p.v.Swap(math.Float32bits(new)))
}
func (p *Float32) CompareAndSwap(old, new float32) bool {
	return p.v.CompareAndSwap(math.Float32bits(old), math.Float32bits(new))
}
func (p *Float32) Add(delta float32) (new float32) {
	for {
		old := p.Load()
		new := old + delta
		if p.CompareAndSwap(old, new) {
			return new
		}
	}
}
func (p *Float32) Sub(delta float32) (new float32) {
	return p.Add(-delta)
}

type Float64 struct {
	v Uint64
}

func (p *Float64) Store(v float64) { p.v.Store(math.Float64bits(v)) }
func (p *Float64) Load() float64   { return math.Float64frombits(p.v.Load()) }
func (p *Float64) Swap(new float64) (old float64) {
	return math.Float64frombits(p.v.Swap(math.Float64bits(new)))
}
func (p *Float64) CompareAndSwap(old, new float64) bool {
	return p.v.CompareAndSwap(math.Float64bits(old), math.Float64bits(new))
}
func (p *Float64) Add(delta float64) (new float64) {
	for {
		old := p.Load()
		new := old + delta
		if p.CompareAndSwap(old, new) {
			return new
		}
	}
}
func (p *Float64) Sub(delta float64) (new float64) {
	return p.Add(-delta)
}
