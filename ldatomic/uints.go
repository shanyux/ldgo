/*
 * Copyright (C) distroy
 */

package ldatomic

type Uints Interface

func NewUints(d []uint) *Uints {
	p := &Uints{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Uints) get() *Interface      { return (*Interface)(p) }
func (p *Uints) pack(d []uint) []uint { return d }
func (p *Uints) unpack(i interface{}) []uint {
	d, _ := i.([]uint)
	return d
}

func (p *Uints) Store(d []uint)               { p.get().Store(p.pack(d)) }
func (p *Uints) Load() []uint                 { return p.unpack(p.get().Load()) }
func (p *Uints) Swap(new []uint) (old []uint) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Uints) CompareAndSwap(old, new []uint) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Uint8s Interface

func NewUint8s(d []uint8) *Uint8s {
	p := &Uint8s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Uint8s) get() *Interface        { return (*Interface)(p) }
func (p *Uint8s) pack(d []uint8) []uint8 { return d }
func (p *Uint8s) unpack(i interface{}) []uint8 {
	d, _ := i.([]uint8)
	return d
}

func (p *Uint8s) Store(d []uint8)                { p.get().Store(p.pack(d)) }
func (p *Uint8s) Load() []uint8                  { return p.unpack(p.get().Load()) }
func (p *Uint8s) Swap(new []uint8) (old []uint8) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Uint8s) CompareAndSwap(old, new []uint8) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Uint16s Interface

func NewUint16s(d []uint16) *Uint16s {
	p := &Uint16s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Uint16s) get() *Interface          { return (*Interface)(p) }
func (p *Uint16s) pack(d []uint16) []uint16 { return d }
func (p *Uint16s) unpack(i interface{}) []uint16 {
	d, _ := i.([]uint16)
	return d
}

func (p *Uint16s) Store(d []uint16)                 { p.get().Store(p.pack(d)) }
func (p *Uint16s) Load() []uint16                   { return p.unpack(p.get().Load()) }
func (p *Uint16s) Swap(new []uint16) (old []uint16) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Uint16s) CompareAndSwap(old, new []uint16) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Uint32s Interface

func NewUint32s(d []uint32) *Uint32s {
	p := &Uint32s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Uint32s) get() *Interface          { return (*Interface)(p) }
func (p *Uint32s) pack(d []uint32) []uint32 { return d }
func (p *Uint32s) unpack(i interface{}) []uint32 {
	d, _ := i.([]uint32)
	return d
}

func (p *Uint32s) Store(d []uint32)                 { p.get().Store(p.pack(d)) }
func (p *Uint32s) Load() []uint32                   { return p.unpack(p.get().Load()) }
func (p *Uint32s) Swap(new []uint32) (old []uint32) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Uint32s) CompareAndSwap(old, new []uint32) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Uint64s Interface

func NewUint64s(d []uint64) *Uint64s {
	p := &Uint64s{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Uint64s) get() *Interface          { return (*Interface)(p) }
func (p *Uint64s) pack(d []uint64) []uint64 { return d }
func (p *Uint64s) unpack(i interface{}) []uint64 {
	d, _ := i.([]uint64)
	return d
}

func (p *Uint64s) Store(d []uint64)                 { p.get().Store(p.pack(d)) }
func (p *Uint64s) Load() []uint64                   { return p.unpack(p.get().Load()) }
func (p *Uint64s) Swap(new []uint64) (old []uint64) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Uint64s) CompareAndSwap(old, new []uint64) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}
