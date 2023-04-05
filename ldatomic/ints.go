/*
 * Copyright (C) distroy
 */

package ldatomic

type Ints struct {
	d Interface
}

func NewInts(d []int) *Ints {
	p := &Ints{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Ints) get() *Interface    { return &p.d }
func (p *Ints) pack(d []int) []int { return d }
func (p *Ints) unpack(i interface{}) []int {
	d, _ := i.([]int)
	return d
}

func (p *Ints) Store(d []int)              { p.get().Store(p.pack(d)) }
func (p *Ints) Load() []int                { return p.unpack(p.get().Load()) }
func (p *Ints) Swap(new []int) (old []int) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Ints) CompareAndSwap(old, new []int) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Int8s struct {
	d Interface
}

func NewInt8s(d []int8) *Int8s {
	p := &Int8s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Int8s) get() *Interface      { return &p.d }
func (p *Int8s) pack(d []int8) []int8 { return d }
func (p *Int8s) unpack(i interface{}) []int8 {
	d, _ := i.([]int8)
	return d
}

func (p *Int8s) Store(d []int8)               { p.get().Store(p.pack(d)) }
func (p *Int8s) Load() []int8                 { return p.unpack(p.get().Load()) }
func (p *Int8s) Swap(new []int8) (old []int8) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Int8s) CompareAndSwap(old, new []int8) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Int16s struct {
	d Interface
}

func NewInt16s(d []int16) *Int16s {
	p := &Int16s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Int16s) get() *Interface        { return &p.d }
func (p *Int16s) pack(d []int16) []int16 { return d }
func (p *Int16s) unpack(i interface{}) []int16 {
	d, _ := i.([]int16)
	return d
}

func (p *Int16s) Store(d []int16)                { p.get().Store(p.pack(d)) }
func (p *Int16s) Load() []int16                  { return p.unpack(p.get().Load()) }
func (p *Int16s) Swap(new []int16) (old []int16) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Int16s) CompareAndSwap(old, new []int16) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Int32s struct {
	d Interface
}

func NewInt32s(d []int32) *Int32s {
	p := &Int32s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Int32s) get() *Interface        { return &p.d }
func (p *Int32s) pack(d []int32) []int32 { return d }
func (p *Int32s) unpack(i interface{}) []int32 {
	d, _ := i.([]int32)
	return d
}

func (p *Int32s) Store(d []int32)                { p.get().Store(p.pack(d)) }
func (p *Int32s) Load() []int32                  { return p.unpack(p.get().Load()) }
func (p *Int32s) Swap(new []int32) (old []int32) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Int32s) CompareAndSwap(old, new []int32) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Int64s struct {
	d Interface
}

func NewInt64s(d []int64) *Int64s {
	p := &Int64s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Int64s) get() *Interface        { return &p.d }
func (p *Int64s) pack(d []int64) []int64 { return d }
func (p *Int64s) unpack(i interface{}) []int64 {
	d, _ := i.([]int64)
	return d
}

func (p *Int64s) Store(d []int64)                { p.get().Store(p.pack(d)) }
func (p *Int64s) Load() []int64                  { return p.unpack(p.get().Load()) }
func (p *Int64s) Swap(new []int64) (old []int64) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Int64s) CompareAndSwap(old, new []int64) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}
