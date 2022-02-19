/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"unsafe"
)

// Interface provides an atomic load and store of an any typed value.
// The zero value for a Value returns nil from Load.
type Interface Pointer

func NewInterface(d interface{}) *Interface {
	p := &Interface{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Interface) get() *Pointer { return (*Pointer)(p) }

func (p *Interface) Store(d interface{}) { p.get().Store(p.toPointer(d)) }
func (p *Interface) Load() interface{}   { return p.toIface(p.get().Load()) }
func (p *Interface) Swap(new interface{}) (old interface{}) {
	return p.toIface(p.get().Swap(p.toPointer(new)))
}

func (p *Interface) toPointer(d interface{}) unsafe.Pointer {
	if d == nil {
		return nil
	}
	return unsafe.Pointer(&d)
}
func (p *Interface) toIface(addr unsafe.Pointer) interface{} {
	if addr == nil {
		return nil
	}
	x := (*interface{})(addr)
	return *x
}
