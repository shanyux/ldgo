/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func RandInt() int     { return ldcore.RandInt() }
func RandInt31() int32 { return ldcore.RandInt31() }
func RandInt63() int64 { return ldcore.RandInt63() }

func RandIntn(n int) int       { return ldcore.RandIntn(n) }
func RandInt31n(n int32) int32 { return ldcore.RandInt31n(n) }
func RandInt63n(n int64) int64 { return ldcore.RandInt63n(n) }

func RandUint() uint     { return ldcore.RandUint() }
func RandUint32() uint32 { return ldcore.RandUint32() }
func RandUint64() uint64 { return ldcore.RandUint64() }

// RandFloat64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func RandFloat64() float64 { return ldcore.RandFloat64() }

// RandFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func RandFloat32() float32 { return ldcore.RandFloat32() }

func RandString(n int) string {
	return ldcore.RandString(n)
}
