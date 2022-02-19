/*
 * Copyright (C) distroy
 */

package ldatomic

var _zeroString string

type String struct {
	d Interface
}

func NewString(d string) *String {
	p := &String{}
	p.Store(d)
	return p
}

func (p *String) Store(d string)               { p.d.Store(d) }
func (p *String) Load() string                 { return p.toStr(p.d.Load()) }
func (p *String) Swap(new string) (old string) { return p.toStr(p.d.Swap(new)) }
func (p *String) String() string               { return p.Load() }

func (p *String) toStr(i interface{}) string {
	if i == nil {
		return _zeroString
	}
	x, _ := i.(string)
	return x
}
