/*
 * Copyright (C) distroy
 */

package ldatomic

type Float32s Interface

func NewFloat32s(d []float32) *Float32s {
	p := &Float32s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Float32s) get() *Interface            { return (*Interface)(p) }
func (p *Float32s) pack(d []float32) []float32 { return d }
func (p *Float32s) unpack(i interface{}) []float32 {
	d, _ := i.([]float32)
	return d
}

func (p *Float32s) Store(d []float32)                  { p.get().Store(p.pack(d)) }
func (p *Float32s) Load() []float32                    { return p.unpack(p.get().Load()) }
func (p *Float32s) Swap(new []float32) (old []float32) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Float32s) CompareAndSwap(old, new []float32) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Float64s Interface

func NewFloat64s(d []float64) *Float64s {
	p := &Float64s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Float64s) get() *Interface            { return (*Interface)(p) }
func (p *Float64s) pack(d []float64) []float64 { return d }
func (p *Float64s) unpack(i interface{}) []float64 {
	d, _ := i.([]float64)
	return d
}

func (p *Float64s) Store(d []float64)                  { p.get().Store(p.pack(d)) }
func (p *Float64s) Load() []float64                    { return p.unpack(p.get().Load()) }
func (p *Float64s) Swap(new []float64) (old []float64) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Float64s) CompareAndSwap(old, new []float64) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}
