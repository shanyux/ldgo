/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"fmt"
	"reflect"
)

func ChangeFunc(fn interface{}, options ...ChangeFuncOption) interface{} {
	val := reflect.ValueOf(fn)
	typ := val.Type()
	if val.Kind() != reflect.Func {
		panic(fmt.Sprintf("should be func, but be `%s`", typ.String()))
	}

	p := &funcChanger{}
	p.Init(val, typ)

	for _, opt := range options {
		opt(p)
	}

	return p.MakeFunc().Interface()
}

type funcChangerParameter struct {
	typ   reflect.Type
	value reflect.Value
	pos   int
}

type funcChanger struct {
	value    reflect.Value
	typ      reflect.Type
	variadic bool
	inputCnt int
	inputs   []*funcChangerParameter
	outputs  []*funcChangerParameter
}

func (p *funcChanger) Init(val reflect.Value, typ reflect.Type) {
	inputs := make([]*funcChangerParameter, 0, typ.NumIn())
	for i, n := 0, typ.NumIn(); i < n; i++ {
		inT := typ.In(i)

		inputs = append(inputs, &funcChangerParameter{
			typ: inT,
			pos: i,
		})
	}

	outputs := make([]*funcChangerParameter, 0, typ.NumOut())
	for i, n := 0, typ.NumOut(); i < n; i++ {
		inT := typ.Out(i)

		outputs = append(outputs, &funcChangerParameter{
			typ: inT,
			pos: i,
		})
	}

	*p = funcChanger{
		value:    val,
		typ:      typ,
		variadic: typ.IsVariadic(),
		inputCnt: len(inputs),
		inputs:   inputs,
		outputs:  outputs,
	}
}

func (p *funcChanger) getParamterTypes(l []*funcChangerParameter) []reflect.Type {
	res := make([]reflect.Type, 0, len(l))
	for _, v := range l {
		res = append(res, v.typ)
	}

	return res
}

func (p *funcChanger) MakeFunc() reflect.Value {
	inTypes := p.getParamterTypes(p.inputs)
	outTypes := p.getParamterTypes(p.outputs)
	funcType := reflect.FuncOf(inTypes, outTypes, p.variadic)

	target := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		return p.Do(args)
	})

	return target
}

func (p *funcChanger) Do(ins1 []reflect.Value) []reflect.Value {
	ins0 := make([]reflect.Value, p.inputCnt)
	for i, v := range p.inputs {
		if v.pos < 0 {
			continue
		}

		ins0[v.pos] = ins1[i]
	}

	if p.variadic {
		// log.Println()
		// for i, v := range ins1 {
		// 	log.Printf(" *** %d: %v \n", i, v.Interface())
		// }

		// ins0 = append(ins0, ins1[p.inputCnt+1:]...)

		ins0 = ins0[:p.inputCnt-1]
		last := ins1[len(ins1)-1]
		for i, n := 0, last.Len(); i < n; i++ {
			ins0 = append(ins0, last.Index(i))
		}
	}

	outs0 := p.value.Call(ins0)
	outs1 := make([]reflect.Value, len(p.outputs))
	for i, v := range p.outputs {
		if v.pos < 0 {
			outs1[i] = v.value
			continue
		}

		outs1[i] = outs0[v.pos]
	}

	return outs1
}

func (p *funcChanger) reflectOf(val interface{}, typ ...reflect.Type) (reflect.Value, reflect.Type) {
	v := reflect.ValueOf(val)

	if len(typ) == 0 {
		t := v.Type()
		return v, t
	}

	t := typ[0]
	if !v.IsValid() {
		v = reflect.Zero(t)

	} else if !v.Type().Implements(t) {
		panic(fmt.Sprintf("`%s` should implement `%s`", v.Type().String(), t.String()))
	}

	return v, t
}
