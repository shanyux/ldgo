/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldmath"

func MinInt(n int, args ...int) int         { return ldmath.MinInt(n, args...) }
func MinInt8(n int8, args ...int8) int8     { return ldmath.MinInt8(n, args...) }
func MinInt16(n int16, args ...int16) int16 { return ldmath.MinInt16(n, args...) }
func MinInt32(n int32, args ...int32) int32 { return ldmath.MinInt32(n, args...) }
func MinInt64(n int64, args ...int64) int64 { return ldmath.MinInt64(n, args...) }

func MinUint(n uint, args ...uint) uint         { return ldmath.MinUint(n, args...) }
func MinUint8(n uint8, args ...uint8) uint8     { return ldmath.MinUint8(n, args...) }
func MinUint16(n uint16, args ...uint16) uint16 { return ldmath.MinUint16(n, args...) }
func MinUint32(n uint32, args ...uint32) uint32 { return ldmath.MinUint32(n, args...) }
func MinUint64(n uint64, args ...uint64) uint64 { return ldmath.MinUint64(n, args...) }
