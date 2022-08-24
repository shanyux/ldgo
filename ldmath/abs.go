/*
 * Copyright (C) distroy
 */

package ldmath

import "math"

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt8(n int8) int8 {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt16(n int16) int16 {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt32(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}
func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

// AbsFloat64 returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func AbsFloat64(n float64) float64 { return math.Abs(n) }
func AbsFloat32(n float32) float32 { return float32(AbsFloat64(float64(n))) }
