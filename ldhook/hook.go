/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"reflect"
)

type Hook interface {
	hook(patches *patches)
}

type FuncHook struct {
	Target interface{}
	Double interface{} // must be func/[]OutputCell/OutputCell/[]Values/Values
}

func (h FuncHook) hook(patches *patches) {
	target := h.Target
	funcType := reflect.TypeOf(h.Target)
	double := getDoubleInterface(funcType, h.Double)

	patches.coreApplyFunc(reflect.ValueOf(target), reflect.ValueOf(double))
}

type MethodHook struct {
	Target interface{}
	Method string
	Double interface{} // must be func/[]OutputCell/OutputCell/[]Params/Params
}

func (h MethodHook) hook(patches *patches) {
	tType, ok := h.Target.(reflect.Type)
	if !ok || tType == nil {
		tType = reflect.TypeOf(h.Target)
	}
	method, ok := tType.MethodByName(h.Method)
	if !ok {
		panic("retrieve method by name fail")
	}

	mType := method.Type
	double := getDoubleInterface(mType, h.Double)

	patches.coreApplyFunc(method.Func, reflect.ValueOf(double))
}

func getDoubleInterface(funcType reflect.Type, double interface{}) interface{} {
	switch v := double.(type) {
	default:
		return double

	case []OutputCell:
		outs := make([]ResultCell, 0, len(v))
		for _, out := range v {
			outs = append(outs, ResultCell{
				Outputs: out.Values,
				Times:   out.Times,
			})
		}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case []*OutputCell:
		outs := make([]ResultCell, 0, len(v))
		for _, out := range v {
			outs = append(outs, ResultCell{
				Outputs: out.Values,
				Times:   out.Times,
			})
		}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case OutputCell:
		outs := []ResultCell{{
			Outputs: v.Values,
			Times:   v.Times,
		}}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case *OutputCell:
		outs := []ResultCell{{
			Outputs: v.Values,
			Times:   v.Times,
		}}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case []ResultCell:
		return newWrap(funcType, v).MakeFunc().Interface()

	case []*ResultCell:
		outs := make([]ResultCell, 0, len(v))
		for _, out := range v {
			outs = append(outs, *out)
		}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case ResultCell:
		outs := []ResultCell{v}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case *ResultCell:
		outs := []ResultCell{*v}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case []Values:
		outs := make([]ResultCell, 0, len(v))
		for _, out := range v {
			outs = append(outs, ResultCell{Outputs: out})
		}
		return newWrap(funcType, outs).MakeFunc().Interface()

	case Values:
		outs := []ResultCell{
			{Outputs: v},
		}
		return newWrap(funcType, outs).MakeFunc().Interface()
	}
}
