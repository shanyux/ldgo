/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func AsByte(val interface{}, def ...byte) byte { return ldcore.AsByte(val, def...) }

func AsBool(val interface{}, def ...bool) bool {
	return ldcore.AsBool(val, def...)
}

func AsInt(val interface{}, def ...int) int       { return ldcore.AsInt(val, def...) }
func AsInt8(val interface{}, def ...int8) int8    { return ldcore.AsInt8(val, def...) }
func AsInt16(val interface{}, def ...int16) int16 { return ldcore.AsInt16(val, def...) }
func AsInt32(val interface{}, def ...int32) int32 { return ldcore.AsInt32(val, def...) }
func AsInt64(val interface{}, def ...int64) int64 { return ldcore.AsInt64(val, def...) }

func AsUint(val interface{}, def ...uint) uint       { return ldcore.AsUint(val, def...) }
func AsUint8(val interface{}, def ...uint8) uint8    { return ldcore.AsUint8(val, def...) }
func AsUint16(val interface{}, def ...uint16) uint16 { return ldcore.AsUint16(val, def...) }
func AsUint32(val interface{}, def ...uint32) uint32 { return ldcore.AsUint32(val, def...) }
func AsUint64(val interface{}, def ...uint64) uint64 { return ldcore.AsUint64(val, def...) }

func AsFloat32(val interface{}, def ...float32) float32 { return ldcore.AsFloat32(val, def...) }
func AsFloat64(val interface{}, def ...float64) float64 { return ldcore.AsFloat64(val, def...) }

func AsString(val interface{}, def ...string) string { return ldcore.AsString(val, def...) }
