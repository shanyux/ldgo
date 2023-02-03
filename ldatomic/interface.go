/*
 * Copyright (C) distroy
 */

package ldatomic

// Interface provides an atomic load and store of an any typed value.
// The zero value for a Value returns nil from Load.
type Interface Value

func NewInterface(d interface{}) *Interface {
	p := &Interface{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Interface) get() *Value { return (*Value)(p) }

func (p *Interface) Store(d interface{}) { p.get().Store(p.pack(d)) }
func (p *Interface) Load() interface{}   { return p.unpack(p.get().Load()) }
func (p *Interface) Swap(new interface{}) (old interface{}) {
	return p.unpack(p.get().Swap(p.pack(new)))
}

func (p *Interface) CompareAndSwap(old, new interface{}) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

func (p *Interface) pack(d interface{}) ifaceData {
	return ifaceData{Data: d}
}
func (p *Interface) unpack(addr interface{}) interface{} {
	if addr == nil {
		return nil
	}
	x := addr.(ifaceData)
	return x.Data
}

type ifaceData struct {
	Data interface{}
}
