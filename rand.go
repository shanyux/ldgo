/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldrand"

func RandInt() int     { return ldrand.RandInt() }
func RandInt31() int32 { return ldrand.RandInt31() }
func RandInt63() int64 { return ldrand.RandInt63() }

func RandIntn(n int) int       { return ldrand.RandIntn(n) }
func RandInt31n(n int32) int32 { return ldrand.RandInt31n(n) }
func RandInt63n(n int64) int64 { return ldrand.RandInt63n(n) }

func RandUint() uint     { return ldrand.RandUint() }
func RandUint32() uint32 { return ldrand.RandUint32() }
func RandUint64() uint64 { return ldrand.RandUint64() }

// RandFloat64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func RandFloat64() float64 { return ldrand.RandFloat64() }

// RandFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func RandFloat32() float32 { return ldrand.RandFloat32() }

func RandString(n int) string {
	return ldrand.RandString(n)
}
