/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"reflect"
	"unsafe"

	_ "github.com/agiledragon/gomonkey"
)

//go:linkname replace github.com/agiledragon/gomonkey.replace
func replace(target, double uintptr) []byte

//go:linkname modifyBinary github.com/agiledragon/gomonkey.modifyBinary
func modifyBinary(target uintptr, bytes []byte)

//go:linkname getPointer github.com/agiledragon/gomonkey.getPointer
func getPointer(v reflect.Value) unsafe.Pointer

type patchFunc struct {
	target reflect.Value
	double reflect.Value
	origin []byte
}

type patchVariable struct {
	target reflect.Value
	origin reflect.Value
}

func applyFunc(target, double reflect.Value) patchFunc {
	origin := replace(*(*uintptr)(getPointer(target)), uintptr(getPointer(double)))
	return patchFunc{
		target: target,
		double: double,
		origin: origin,
	}
}

func resetFunc(v patchFunc) {
	modifyBinary(*(*uintptr)(getPointer(v.target)), v.origin)
}

func applyVariable(target, double reflect.Value) patchVariable {
	origin := reflect.ValueOf(target.Elem().Interface())
	target.Elem().Set(double)
	return patchVariable{
		target: target,
		origin: origin,
	}
}

func resetVariable(v patchVariable) {
	v.target.Elem().Set(v.origin)
}
