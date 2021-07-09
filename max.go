/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func MaxInt(n int, args ...int) int         { return ldcore.MaxInt(n, args...) }
func MaxInt8(n int8, args ...int8) int8     { return ldcore.MaxInt8(n, args...) }
func MaxInt16(n int16, args ...int16) int16 { return ldcore.MaxInt16(n, args...) }
func MaxInt32(n int32, args ...int32) int32 { return ldcore.MaxInt32(n, args...) }
func MaxInt64(n int64, args ...int64) int64 { return ldcore.MaxInt64(n, args...) }

func MaxUint(n uint, args ...uint) uint         { return ldcore.MaxUint(n, args...) }
func MaxUint8(n uint8, args ...uint8) uint8     { return ldcore.MaxUint8(n, args...) }
func MaxUint16(n uint16, args ...uint16) uint16 { return ldcore.MaxUint16(n, args...) }
func MaxUint32(n uint32, args ...uint32) uint32 { return ldcore.MaxUint32(n, args...) }
func MaxUint64(n uint64, args ...uint64) uint64 { return ldcore.MaxUint64(n, args...) }
