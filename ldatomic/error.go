/*
 * Copyright (C) distroy
 */

package ldatomic

type Error Interface

func (p *Error) get() *Interface { return (*Interface)(p) }

func (p *Error) Store(v error)              { p.get().Store(v) }
func (p *Error) Load() error                { return p.toErr(p.get().Load()) }
func (p *Error) Swap(new error) (old error) { return p.toErr(p.get().Swap(new)) }

func (p *Error) toErr(i interface{}) error {
	if i == nil {
		return nil
	}
	e, _ := i.(error)
	return e
}
