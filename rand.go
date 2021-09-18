/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldrand"

func RandInt() int     { return ldrand.Int() }
func RandInt31() int32 { return ldrand.Int31() }
func RandInt63() int64 { return ldrand.Int63() }

func RandIntn(n int) int       { return ldrand.Intn(n) }
func RandInt31n(n int32) int32 { return ldrand.Int31n(n) }
func RandInt63n(n int64) int64 { return ldrand.Int63n(n) }

func RandUint() uint     { return ldrand.Uint() }
func RandUint32() uint32 { return ldrand.Uint32() }
func RandUint64() uint64 { return ldrand.Uint64() }

// RandFloat64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func RandFloat64() float64 { return ldrand.Float64() }

// RandFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func RandFloat32() float32 { return ldrand.Float32() }

func RandString(n int) string {
	return ldrand.String(n)
}
