/*
 * Copyright (C) distroy
 */

package ldmath

import "math"

func NaN32() float32 { return float32(NaN64()) }
func NaN64() float64 { return math.NaN() }

func IsNaN32(n float32) bool { return n != n }
func IsNaN64(n float64) bool { return n != n }

func Inf32(sign int) float32 { return float32(Inf64(sign)) }
func Inf64(sign int) float64 { return math.Inf(sign) }

func IsInf32(f float32, sign int) bool { return IsInf64(float64(f), sign) }
func IsInf64(f float64, sign int) bool { return math.IsInf(f, sign) }
