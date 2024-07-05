/*
 * Copyright (C) distroy
 */

package ldmath

import "math"

type Absable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Abs[T Absable](n T) T {
	if n < 0 {
		n = -n
	}
	return n
}

// Deprecated: use `Abs[Type]` instead.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Deprecated: use `Abs[Type]` instead.
func AbsInt8(n int8) int8 {
	if n < 0 {
		return -n
	}
	return n
}

// Deprecated: use `Abs[Type]` instead.
func AbsInt16(n int16) int16 {
	if n < 0 {
		return -n
	}
	return n
}

// Deprecated: use `Abs[Type]` instead.
func AbsInt32(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

// Deprecated: use `Abs[Type]` instead.
func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// AbsFloat64 returns the absolute value of x.
// Special cases are:
//
//   - Abs(±Inf) = +Inf
//   - Abs(NaN) = NaN
//
// Deprecated: use `Abs[Type]` instead.
func AbsFloat64(n float64) float64 { return math.Abs(n) }

// Deprecated: use `Abs[Type]` instead.
func AbsFloat32(n float32) float32 { return float32(AbsFloat64(float64(n))) }
