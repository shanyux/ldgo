/*
 * Copyright (C) distroy
 */

package ldatomic

var _zeroString string

type String struct {
	v Interface
}

func (p *String) Store(x string)               { p.v.Store(x) }
func (p *String) Load() string                 { return p.toStr(p.v.Load()) }
func (p *String) Swap(new string) (old string) { return p.toStr(p.v.Swap(new)) }
func (p *String) String() string               { return p.String() }

func (p *String) toStr(i interface{}) string {
	if i == nil {
		return _zeroString
	}
	x, _ := i.(string)
	return x
}
