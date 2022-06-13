/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"math"
	"reflect"
	"sort"

	"github.com/distroy/ldgo/ldmath"
)

type kind int

const (
	kindNil kind = iota
	kindBool
	kindNumber
	kindComplex
	kindString
	kindInvalid
)

func reflectValueOf(v interface{}) reflect.Value {
	if vv, ok := v.(reflect.Value); ok {
		return vv
	}

	vv := reflect.ValueOf(v)
	return vv
}

func CompareReflect(a, b reflect.Value) int {
	if r := compareReflectType(a, b); r != 0 {
		return r
	}

	switch a.Kind() {
	default:
		// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
		panic("bad type in compare: " + a.Type().String())

	case reflect.Invalid:
		return 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		aa := a.Int()
		return compareReflectNumberLeftInt(aa, b)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		aa := a.Uint()
		return compareReflectNumberLeftUint(aa, b)

	case reflect.Float32, reflect.Float64:
		aa := a.Float()
		return compareReflectNumberLeftFloat(aa, b)

	case reflect.String:
		aa := a.String()
		bb := b.String()
		return CompareString(aa, bb)

	case reflect.Complex64, reflect.Complex128:
		aa := a.Complex()
		bb := b.Complex()
		return CompareComplex128(aa, bb)

	case reflect.Bool:
		aa := a.Bool()
		bb := b.Bool()
		return CompareBool(aa, bb)

	case reflect.Ptr, reflect.UnsafePointer:
		aa := a.Pointer()
		bb := b.Pointer()
		return CompareUintptr(aa, bb)

	case reflect.Chan, reflect.Func:
		aa := a.Pointer()
		bb := b.Pointer()
		return CompareUintptr(aa, bb)

	case reflect.Map:
		return compareReflectMap(a, b)

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
	case reflect.Invalid:
		return kindNil

	case reflect.Bool:
		return kindBool

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return kindNumber

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return kindNumber

	case reflect.Float32, reflect.Float64:
		return kindNumber

	case reflect.Complex64, reflect.Complex128:
		return kindComplex

	case reflect.String:
		return kindString

	default:
		return kindInvalid
	}
}

func compareReflectType(a, b reflect.Value) int {
	aKind := convertKind(a.Kind())
	bKind := convertKind(b.Kind())
	if aKind != kindInvalid && aKind == bKind {
		return 0
	}

	if aKind != kindInvalid || bKind != kindInvalid {
		return CompareInt(int(aKind), int(bKind))
	}

	if a.Type() == b.Type() {
		return 0
	}

	aName := a.String()
	bName := b.String()
	if r := CompareInt(len(aName), len(bName)); r != 0 {
		return r
	}
	if aName < bName {
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
		aa := a.Field(i)
		bb := b.Field(i)
		// aa, bb = reflect.ValueOf(aa.Interface()), reflect.ValueOf(bb.Interface())
		if r := CompareReflect(aa, bb); r != 0 {
			return r
		}
	}
	return 0
}

func compareReflectArray(a, b reflect.Value) int {
	for i := 0; i < a.Len(); i++ {
		aa := a.Index(i)
		bb := b.Index(i)
		if r := CompareReflect(aa, bb); r != 0 {
			return r
		}
	}
	return 0
}

func compareReflectMap(a, b reflect.Value) int {
	if r := CompareInt(a.Len(), b.Len()); r != 0 {
		return r
	}

	aKeys := a.MapKeys()
	bKeys := b.MapKeys()
	sort.Sort(sortedReflects(aKeys))
	sort.Sort(sortedReflects(bKeys))

	for i := range aKeys {
		aKey, bKey := aKeys[i], bKeys[i]
		if r := CompareReflect(aKey, bKey); r != 0 {
			return r
		}

		aVal, bVal := a.MapIndex(aKey), b.MapIndex(bKey)
		if r := CompareReflect(aVal, bVal); r != 0 {
			return r
		}
	}

	return 0
}

func compareReflectSlice(a, b reflect.Value) int {
	al := a.Len()
	bl := b.Len()
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
	return CompareReflect(a.Elem(), b.Elem())
}

func compareReflectNumberLeftInt(aa int64, b reflect.Value) int {
	switch b.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		bb := b.Uint()
		if bb > math.MaxInt64 {
			return -1
		}
		return CompareInt64(aa, int64(bb))

	case reflect.Float32, reflect.Float64:
		bb := b.Float()
		return CompareFloat64(float64(aa), bb)
	}

	bb := b.Int()
	return CompareInt64(aa, bb)
}

func compareReflectNumberLeftUint(aa uint64, b reflect.Value) int {
	switch b.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bb := b.Int()
		if aa > math.MaxInt64 {
			return 1
		}
		return CompareInt64(int64(aa), bb)

	case reflect.Float32, reflect.Float64:
		bb := b.Float()
		return CompareFloat64(float64(aa), bb)
	}

	bb := b.Uint()
	return CompareUint64(aa, bb)
}

func compareReflectNumberLeftFloat(aa float64, b reflect.Value) int {
	switch b.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bb := b.Int()
		if aa > math.MaxInt64 {
			return 1
		}
		return CompareFloat64(aa, float64(bb))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		bb := b.Uint()
		return CompareFloat64(aa, float64(bb))
	}

	bb := b.Float()
	return CompareFloat64(aa, bb)
}
