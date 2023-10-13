/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"fmt"
	"reflect"
)

type ChangeFuncOption func(p *funcChanger)

func SwapInput(a, b int) ChangeFuncOption {
	return func(p *funcChanger) {
		n := len(p.inputs)
		if p.variadic && (a >= n-1 || b >= n-1) {
			panic(fmt.Sprintf("variadic func must not swap the last input parameter"))
		}

		p.inputs[a], p.inputs[b] = p.inputs[b], p.inputs[a]
	}
}

func SwapOutput(a, b int) ChangeFuncOption {
	return func(p *funcChanger) {
		p.outputs[a], p.outputs[b] = p.outputs[b], p.outputs[a]
	}
}

func AppendInput(val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		if p.variadic {
			panic(fmt.Sprintf("variadic func must not append input parameter"))
		}

		v, t := p.reflectOf(val, typ...)

		p.inputs = append(p.inputs, &funcChangerParameter{
			value: v,
			typ:   t,
			pos:   -1,
		})
	}
}

func AppendOutput(val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		v, t := p.reflectOf(val, typ...)

		p.outputs = append(p.outputs, &funcChangerParameter{
			value: v,
			typ:   t,
			pos:   -1,
		})
	}
}

func AddInput(pos int, val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		n := len(p.inputs)
		if pos > n {
			panic(fmt.Sprintf("the input parameter position for add is out of range. pos:%d, len:%d", pos, n))
		}

		if pos == n && p.variadic {
			panic(fmt.Sprintf("variadic func must not add input parameter after last"))
		}

		v, t := p.reflectOf(val, typ...)

		p.inputs = append(p.inputs[:pos+1], p.inputs[pos:]...)
		p.inputs[pos] = &funcChangerParameter{
			value: v,
			typ:   t,
			pos:   -1,
		}
	}
}

func AddOutput(pos int, val interface{}, typ ...reflect.Type) ChangeFuncOption {
	return func(p *funcChanger) {
		n := len(p.outputs)
		if pos > n {
			panic(fmt.Sprintf("the output parameter position for add is out of range. pos:%d, len:%d", pos, n))
		}

		v, t := p.reflectOf(val, typ...)

		p.outputs = append(p.outputs[:pos+1], p.outputs[pos:]...)
		p.outputs[pos] = &funcChangerParameter{
			value: v,
			typ:   t,
			pos:   -1,
		}
	}
}
