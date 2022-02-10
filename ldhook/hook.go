/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"reflect"

	"github.com/agiledragon/gomonkey"
)

type Hook interface {
	hook(patches *gomonkey.Patches)
}

type FuncHook struct {
	Target interface{}
	Double interface{} // must be func/[]OutputCell/OutputCell/[]Values/Values
}

func (h FuncHook) hook(patches *gomonkey.Patches) {
	funcType := reflect.TypeOf(h.Target)
	patches.ApplyFunc(h.Target, getDoubleInterface(funcType, h.Double))
}

type MethodHook struct {
	Target interface{}
	Method string
	Double interface{} // must be func/[]OutputCell/OutputCell/[]Params/Params
}

func (h MethodHook) hook(patches *gomonkey.Patches) {
	tType := reflect.TypeOf(h.Target)
	method, ok := tType.MethodByName(h.Method)
	if !ok {
		panic("retrieve method by name fail")
	}
	mType := method.Type
	patches.ApplyMethod(tType, h.Method, getDoubleInterface(mType, h.Double))
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
