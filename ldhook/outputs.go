/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"fmt"
	"reflect"

	"github.com/agiledragon/gomonkey"
)

type Values = []interface{}
type OutputCell = gomonkey.OutputCell

type wrapOutputs struct {
	FuncType reflect.Type
	Outputs  []OutputCell
	Length   int
	Position int
	Times    int
}

func newWrap(funcType reflect.Type, outputs []OutputCell) *wrapOutputs {
	return &wrapOutputs{
		FuncType: funcType,
		Outputs:  outputs,
		Length:   len(outputs),
		Position: 0,
		Times:    0,
	}
}

func (w *wrapOutputs) getValues() Values {
	if w.Position >= w.Length {
		return w.Outputs[w.Length-1].Values
	}

	cell := w.Outputs[w.Position]
	w.Times++
	if w.Times >= cell.Times {
		w.Position++
		w.Times = 0
	}
	return cell.Values
}

func (w *wrapOutputs) GetOutputs() []reflect.Value {
	funcType := w.FuncType
	results := w.getValues()

	var resultValues []reflect.Value
	for i, r := range results {
		var resultValue reflect.Value
		if r == nil {
			resultValue = reflect.Zero(funcType.Out(i))

		} else {
			vf := reflect.New(funcType.Out(i)).Elem()
			vr := reflect.ValueOf(r)
			if vf.Kind() != vr.Kind() && vr.Type().ConvertibleTo(vf.Type()) {
				vr = vr.Convert(vf.Type())
			}
			vf.Set(vr)
			resultValue = vf
		}
		resultValues = append(resultValues, resultValue)
	}
	return resultValues
}

func (w *wrapOutputs) MakeFunc() reflect.Value {
	funcType := w.FuncType
	outputs := w.Outputs
	if funcType.NumOut() != len(outputs[0].Values) {
		panic(fmt.Sprintf("func type has %v return values, but only %v values provided as double, func type: %s",
			funcType.NumOut(), len(outputs[0].Values), funcType))
	}

	return reflect.MakeFunc(funcType, func(_ []reflect.Value) []reflect.Value {
		return w.GetOutputs()
	})
}
