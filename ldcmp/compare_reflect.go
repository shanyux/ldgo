/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"reflect"

	"github.com/distroy/ldgo/ldmath"
)

type kind int

const (
	kindNil kind = iota
	kindInt
	kindUint
	kindFloat
	kindComplex
	kindPtr
)

func CompareReflect(a, b reflect.Value) int {
	aType, bType := a.Type(), b.Type()

	if r := compareReflectType(aType, bType); r != 0 {
		return 4
	}

	switch a.Kind() {
	default:
		// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
		panic("bad type in compare: " + aType.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a, b := a.Int(), b.Int()
		return CompareInt64(a, b)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		a, b := a.Uint(), b.Uint()
		return CompareUint64(a, b)

	case reflect.String:
		a, b := a.String(), b.String()
		return CompareString(a, b)

	case reflect.Float32, reflect.Float64:
		return CompareFloat64(a.Float(), b.Float())

	case reflect.Complex64, reflect.Complex128:
		a, b := a.Complex(), b.Complex()
		return CompareComplex128(a, b)

	case reflect.Bool:
		a, b := a.Bool(), b.Bool()
		return CompareBool(a, b)

	case reflect.Ptr, reflect.UnsafePointer:
		a, b := a.Pointer(), b.Pointer()
		return CompareUintptr(a, b)

	case reflect.Chan, reflect.Func:
		// if r, ok := compareNilReflect(a, b); ok {
		// 	return r
		// }
		ap, bp := a.Pointer(), b.Pointer()
		return CompareUintptr(ap, bp)

	case reflect.Struct:
		return compareReflectStruct(a, b)

	case reflect.Slice:
		return compareReflectSlice(a, b)

	case reflect.Array:
		return compareReflectArray(a, b)

	case reflect.Interface:
		return compareReflectIface(a, b)
	}
}

func convertKind(k reflect.Kind) kind {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kindInt

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return kindUint

	case reflect.Float32, reflect.Float64:
		return kindFloat

	case reflect.Complex64, reflect.Complex128:
		return kindComplex

	case reflect.Ptr, reflect.UnsafePointer:
		return kindPtr

	default:
		return kindNil
	}
}

func compareReflectType(a, b reflect.Type) int {
	if a == b {
		return 0
	}
	if aa, bb := convertKind(a.Kind()), convertKind(b.Kind()); aa != kindNil && aa == bb {
		return 0
	}
	if a.String() < b.String() {
		return -1
	}
	return 1
}

func compareNilReflect(a, b reflect.Value) (int, bool) {
	if a.IsNil() {
		if b.IsNil() {
			return 0, true
		}
		return -1, true
	}
	if b.IsNil() {
		return 1, true
	}
	return 0, false
}

func compareReflectStruct(a, b reflect.Value) int {
	for i := 0; i < a.NumField(); i++ {
		aa, bb := a.Field(i), b.Field(i)
		// aa, bb = reflect.ValueOf(aa.Interface()), reflect.ValueOf(bb.Interface())
		if r := CompareReflect(aa, bb); r != 0 {
			return r
		}
	}
	return 0
}

func compareReflectArray(a, b reflect.Value) int {
	for i := 0; i < a.Len(); i++ {
		aa, bb := a.Index(i), b.Index(i)
		if r := CompareReflect(aa, bb); r != 0 {
			return r
		}
	}
	return 0
}

func compareReflectSlice(a, b reflect.Value) int {
	al, bl := a.Len(), b.Len()
	l := ldmath.MinInt(al, bl)
	for i := 0; i < l; i++ {
		aa, bb := a.Index(i), b.Index(i)
		if r := CompareReflect(aa, bb); r != 0 {
			return r
		}
	}
	return CompareInt(al, bl)
}

func compareReflectIface(a, b reflect.Value) int {
	if r, ok := compareNilReflect(a, b); ok {
		return r
	}
	if r := compareReflectType(a.Elem().Type(), b.Elem().Type()); r != 0 {
		return r
	}
	return CompareReflect(a.Elem(), b.Elem())
}
