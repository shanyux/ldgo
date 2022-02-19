/*
 * Copyright (C) distroy
 */

package ldatomic

type Bool uint32

func (p *Bool) get() *Uint32 { return (*Uint32)(p) }

func (p *Bool) Store(val bool)           { p.get().Store(p.toInt(val)) }
func (p *Bool) Load() bool               { return p.toBool(p.get().Load()) }
func (p *Bool) Swap(new bool) (old bool) { return p.toBool(p.get().Swap(p.toInt(new))) }
func (p *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	return p.get().CompareAndSwap(p.toInt(old), p.toInt(new))
}

func (p *Bool) Toggle() (old bool) {
	for {
		old := p.Load()
		if p.CompareAndSwap(old, !old) {
			return old
		}
	}
}

func (p *Bool) toInt(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
func (p *Bool) toBool(v uint32) bool {
	return v != 0
}
