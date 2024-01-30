/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"time"

	"github.com/distroy/ldgo/v2/internal/cmp"
)

func CompareInterface(a, b interface{}) int {
	return cmp.CompareInterface(a, b)
}

func CompareBool(a, b bool) int {
	return cmp.CompareBool(a, b)
}

func CompareByte(a, b byte) int {
	return cmp.CompareByte(a, b)
}

func CompareRune(a, b rune) int {
	return cmp.CompareRune(a, b)
}

func CompareInt(a, b int) int {
	return cmp.CompareInt(a, b)
}
func CompareInt8(a, b int8) int {
	return cmp.CompareInt8(a, b)

}
func CompareInt16(a, b int16) int {
	return cmp.CompareInt16(a, b)
}
func CompareInt32(a, b int32) int {
	return cmp.CompareInt32(a, b)
}
func CompareInt64(a, b int64) int {
	return cmp.CompareInt64(a, b)
}

func CompareUint(a, b uint) int {
	return cmp.CompareUint(a, b)
}
func CompareUint8(a, b uint8) int {
	return cmp.CompareUint8(a, b)
}
func CompareUint16(a, b uint16) int {
	return cmp.CompareUint16(a, b)
}
func CompareUint32(a, b uint32) int {
	return cmp.CompareUint32(a, b)
}
func CompareUint64(a, b uint64) int {
	return cmp.CompareUint64(a, b)
}

func CompareUintptr(a, b uintptr) int {
	return cmp.CompareUintptr(a, b)
}

func CompareFloat32(a, b float32) int {
	return cmp.CompareFloat32(a, b)
}

func CompareFloat64(a, b float64) int {
	return cmp.CompareFloat64(a, b)
}

func CompareComplex64(a, b complex64) int {
	return cmp.CompareComplex64(a, b)
}

func CompareComplex128(a, b complex128) int {
	return cmp.CompareComplex128(a, b)
}

func CompareString(a, b string) int {
	return cmp.CompareString(a, b)
}

func CompareDuration(a, b time.Duration) int {
	return cmp.CompareDuration(a, b)
}

func CompareTime(a, b time.Time) int {
	return cmp.CompareTime(a, b)
}
