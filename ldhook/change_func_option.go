/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"reflect"
)

type ChangeFuncOption func(p *funcChanger)

func SwapInput(a, b int) ChangeFuncOption {
	return func(p *funcChanger) {
		p.SwapInput(a, b)
	}
}

func SwapOutput(a, b int) ChangeFuncOption {
	return func(p *funcChanger) {
		p.SwapOutput(a, b)
	}
}

func AppendInput(val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		p.AppendInput(val, typ...)
	}
}

func AppendOutput(val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		p.AppendOutput(val, typ...)
	}
}

func InsertInput(pos int, val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		p.InsertInput(pos, val, typ...)
	}
}

func InsertOutput(pos int, val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		p.InsertOutput(pos, val, typ...)
	}
}
