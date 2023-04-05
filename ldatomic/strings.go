/*
 * Copyright (C) distroy
 */

package ldatomic

type Strings struct {
	d Interface
}

func NewStrings(d []string) *Strings {
	p := &Strings{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Strings) get() *Interface          { return &p.d }
func (p *Strings) pack(d []string) []string { return d }
func (p *Strings) unpack(i interface{}) []string {
	d, _ := i.([]string)
	return d
}

func (p *Strings) Store(d []string)                 { p.get().Store(p.pack(d)) }
func (p *Strings) Load() []string                   { return p.unpack(p.get().Load()) }
func (p *Strings) Swap(new []string) (old []string) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Strings) CompareAndSwap(old, new []string) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}

type Bytes struct {
	d Interface
}

func NewBytes(d []byte) *Bytes {
	p := &Bytes{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Bytes) get() *Interface      { return &p.d }
func (p *Bytes) pack(d []byte) []byte { return d }
func (p *Bytes) unpack(i interface{}) []byte {
	d, _ := i.([]byte)
	return d
}

func (p *Bytes) Store(d []byte)               { p.get().Store(p.pack(d)) }
func (p *Bytes) Load() []byte                 { return p.unpack(p.get().Load()) }
func (p *Bytes) Swap(new []byte) (old []byte) { return p.unpack(p.get().Swap(p.pack(new))) }
func (p *Bytes) CompareAndSwap(old, new []byte) (swapped bool) {
	return p.get().CompareAndSwap(p.pack(old), p.pack(new))
}
