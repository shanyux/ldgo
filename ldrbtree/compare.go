/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"fmt"
	"reflect"
	"strings"
)

func reverseCompare(compare CompareFunc) CompareFunc {
	return func(a, b interface{}) int {
		return -compare(a, b)
	}
}

func DefaultCompare(a, b interface{}) int {
	aRef, bRef := reflect.ValueOf(a), reflect.ValueOf(b)
	return compareReflect(aRef, bRef)
}

func reflectKind(k reflect.Kind) reflect.Kind {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.Int64

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflect.Uint64

	case reflect.Float32, reflect.Float64:
		return reflect.Float64

	case reflect.Complex64, reflect.Complex128:
		return reflect.Complex128

	case reflect.Ptr, reflect.UnsafePointer:
		return reflect.Ptr

	default:
		return k
	}
}

func compareReflect(aVal, bVal reflect.Value) int {
	aType, bType := aVal.Type(), bVal.Type()
	aKind, bKind := aType.Kind(), aType.Kind()

	if reflectKind(aKind) != reflectKind(bKind) || (aKind == reflect.Ptr && aType != bType) {
		// No good answer possible, but don't return 0: they're not equal.
		panic(fmt.Sprintf("the data type is not equal. left:%s, right:%s", aType.String(), bType.String()))
	}

	switch aVal.Kind() {
	default:
		// Certain types cannot appear as keys (maps, funcs, slices), but be explicit.
		panic("bad type in compare: " + aType.String())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		a, b := aVal.Int(), bVal.Int()
		return compareInt(a, b)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		a, b := aVal.Uint(), bVal.Uint()
		return compareUint(a, b)

	case reflect.String:
		a, b := aVal.String(), bVal.String()
		return strings.Compare(a, b)

	case reflect.Float32, reflect.Float64:
		return compareFloat(aVal.Float(), bVal.Float())

	case reflect.Complex64, reflect.Complex128:
		a, b := aVal.Complex(), bVal.Complex()
		return compareComplex(a, b)

	case reflect.Bool:
		a, b := aVal.Bool(), bVal.Bool()
		return compareBool(a, b)

	case reflect.Ptr, reflect.UnsafePointer:
		a, b := aVal.Pointer(), bVal.Pointer()
		return comparePointer(a, b)

	case reflect.Chan:
		if c, ok := compareNil(aVal, bVal); ok {
			return c
		}
		ap, bp := aVal.Pointer(), bVal.Pointer()
		return comparePointer(ap, bp)

	case reflect.Struct:
		return compareStruct(aVal, bVal)

	case reflect.Array:
		return compareArray(aVal, bVal)

	case reflect.Interface:
		return compareIface(aVal, bVal)
	}
}

func compareInt(a, b int64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func compareUint(a, b uint64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func compareBool(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return 1
	default:
		return -1
	}
}

func comparePointer(a, b uintptr) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func compareNil(aVal, bVal reflect.Value) (int, bool) {
	if aVal.IsNil() {
		if bVal.IsNil() {
			return 0, true
		}
		return -1, true
	}
	if bVal.IsNil() {
		return 1, true
	}
	return 0, false
}

func compareFloat(a, b float64) int {
	switch {
	case isNaN(a):
		return -1 // No good answer if b is a NaN so don't bother checking.
	case isNaN(b):
		return 1
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

func compareStruct(aVal, bVal reflect.Value) int {
	for i := 0; i < aVal.NumField(); i++ {
		if c := compareReflect(aVal.Field(i), bVal.Field(i)); c != 0 {
			return c
		}
	}
	return 0
}

func compareArray(aVal, bVal reflect.Value) int {
	for i := 0; i < aVal.Len(); i++ {
		if c := compareReflect(aVal.Index(i), bVal.Index(i)); c != 0 {
			return c
		}
	}
	return 0
}

func compareIface(aVal, bVal reflect.Value) int {
	if c, ok := compareNil(aVal, bVal); ok {
		return c
	}
	c := compareReflect(reflect.ValueOf(aVal.Elem().Type()), reflect.ValueOf(bVal.Elem().Type()))
	if c != 0 {
		return c
	}
	return compareReflect(aVal.Elem(), bVal.Elem())
}

func compareComplex(a, b complex128) int {
	if c := compareFloat(real(a), real(b)); c != 0 {
		return c
	}
	return compareFloat(imag(a), imag(b))
}

func isNaN(a float64) bool {
	return a != a
}
