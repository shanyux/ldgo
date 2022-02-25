/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"reflect"
	"strings"
	"time"
)

func CompareInterface(a, b interface{}) int {
	aa, bb := reflect.ValueOf(a), reflect.ValueOf(b)
	return CompareReflect(aa, bb)
}

func CompareBool(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return 1
	default:
		return -1
	}
}

func CompareByte(a, b byte) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareRune(a, b rune) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareInt(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareInt8(a, b int8) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareInt16(a, b int16) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareInt32(a, b int32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareInt64(a, b int64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func CompareUint(a, b uint) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareUint8(a, b uint8) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareUint16(a, b uint16) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareUint32(a, b uint32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func CompareUint64(a, b uint64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func CompareUintptr(a, b uintptr) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func CompareFloat32(a, b float32) int {
	switch {
	case isNaNFloat32(a):
		return -1 // No good answer if b is a NaN so don't bother checking.
	case isNaNFloat32(b):
		return 1
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

func CompareFloat64(a, b float64) int {
	switch {
	case isNaNFloat64(a):
		return -1 // No good answer if b is a NaN so don't bother checking.
	case isNaNFloat64(b):
		return 1
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

func CompareComplex64(a, b complex64) int {
	if r := CompareFloat32(real(a), real(b)); r != 0 {
		return r
	}
	return CompareFloat32(imag(a), imag(b))
}

func CompareComplex128(a, b complex128) int {
	if r := CompareFloat64(real(a), real(b)); r != 0 {
		return r
	}
	return CompareFloat64(imag(a), imag(b))
}

func isNaNFloat64(a float64) bool {
	return a != a
}

func isNaNFloat32(a float32) bool {
	return a != a
}

func CompareString(a, b string) int {
	return strings.Compare(a, b)
}

func CompareDuration(a, b time.Duration) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareTime(a, b time.Time) int {
	aa, bb := a.UnixNano(), b.UnixNano()
	return CompareInt64(aa, bb)
}
