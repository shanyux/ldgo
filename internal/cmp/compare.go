/*
 * Copyright (C) distroy
 */

package cmp

import (
	"strings"
	"time"
)

func CompareInterface(a, b interface{}) int {
	aa := reflectValueOf(a)
	bb := reflectValueOf(b)
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
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareInt8(a, b int8) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareInt16(a, b int16) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareInt32(a, b int32) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareInt64(a, b int64) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareUint(a, b uint) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareUint8(a, b uint8) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareUint16(a, b uint16) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareUint32(a, b uint32) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}
func CompareUint64(a, b uint64) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareUintptr(a, b uintptr) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return 1
	default:
		return -1
	}
}

func CompareFloat32(a, b float32) int { return CompareOrderable(a, b) }
func CompareFloat64(a, b float64) int { return CompareOrderable(a, b) }

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

func CompareString(a, b string) int {
	return strings.Compare(a, b)
	// return bytes.Compare(ldconv.StrToBytesUnsafe(a), ldconv.StrToBytesUnsafe(b))
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
