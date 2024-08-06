/*
 * Copyright (C) distroy
 */

package ldmath

import "math"

func NaN32() float32 { return float32(NaN64()) }
func NaN64() float64 { return math.NaN() }

// Deprecated: use `IsNaN[Type]` instead.
func IsNaN32(n float32) bool { return n != n }

// Deprecated: use `IsNaN[Type]` instead.
func IsNaN64(n float64) bool { return n != n }

func Inf32(sign int) float32 { return float32(Inf64(sign)) }
func Inf64(sign int) float64 { return math.Inf(sign) }

// Deprecated: use `IsInf[Type]` instead.
func IsInf32(f float32, sign int) bool { return IsInf64(float64(f), sign) }

// Deprecated: use `IsInf[Type]` instead.
func IsInf64(f float64, sign int) bool { return math.IsInf(f, sign) }

func IsNaN[T ~float32 | ~float64](n T) bool { return n != n }

func IsInf[T ~float32 | ~float64](f T, sign int) bool { return math.IsInf(float64(f), sign) }
