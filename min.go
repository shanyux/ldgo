/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func MinInt(n int, args ...int) int         { return ldcore.MinInt(n, args...) }
func MinInt8(n int8, args ...int8) int8     { return ldcore.MinInt8(n, args...) }
func MinInt16(n int16, args ...int16) int16 { return ldcore.MinInt16(n, args...) }
func MinInt32(n int32, args ...int32) int32 { return ldcore.MinInt32(n, args...) }
func MinInt64(n int64, args ...int64) int64 { return ldcore.MinInt64(n, args...) }

func MinUint(n uint, args ...uint) uint         { return ldcore.MinUint(n, args...) }
func MinUint8(n uint8, args ...uint8) uint8     { return ldcore.MinUint8(n, args...) }
func MinUint16(n uint16, args ...uint16) uint16 { return ldcore.MinUint16(n, args...) }
func MinUint32(n uint32, args ...uint32) uint32 { return ldcore.MinUint32(n, args...) }
func MinUint64(n uint64, args ...uint64) uint64 { return ldcore.MinUint64(n, args...) }
