/*
 * Copyright (C) distroy
 */

package ldref

import (
	"math"
	"reflect"
)

func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0

	case reflect.Array:
		return isArrayZero(v)

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()

	case reflect.String:
		return v.Len() == 0

	case reflect.Struct:
		return isStructZero(v)

	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{Method: "ldref.IsZero", Kind: v.Kind()})
	}
}

func isArrayZero(v reflect.Value) bool {
	for i := 0; i < v.Len(); i++ {
		if !IsZero(v.Index(i)) {
			return false
		}
	}
	return true
}

func isStructZero(v reflect.Value) bool {
	for i := 0; i < v.NumField(); i++ {
		if !IsZero(v.Field(i)) {
			return false
		}
	}
	return true
}
