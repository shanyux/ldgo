/*
 * Copyright (C) distroy
 */

package ldatomic

type Error Interface

func NewError(d error) *Error {
	p := &Error{}
	if d != nil {
		p.Store(d)
	}
	return p
}

func (p *Error) get() *Interface { return (*Interface)(p) }

func (p *Error) Store(d error)              { p.get().Store(d) }
func (p *Error) Load() error                { return p.toErr(p.get().Load()) }
func (p *Error) Swap(new error) (old error) { return p.toErr(p.get().Swap(new)) }
func (p *Error) CompareAndSwap(old, new error) (swapped bool) {
	return p.get().CompareAndSwap(p.toErr(old), p.toErr(new))
}

func (p *Error) toErr(i interface{}) error {
	if i == nil {
		return nil
	}
	e, _ := i.(error)
	return e
}
