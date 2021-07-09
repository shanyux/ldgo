/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

func ToByte(v interface{}) (byte, error) { return ldcore.ToByte(v) }

func ToBool(v interface{}) (bool, error) { return ldcore.ToBool(v) }

func ToInt(v interface{}) (int, error)     { return ldcore.ToInt(v) }
func ToInt8(v interface{}) (int8, error)   { return ldcore.ToInt8(v) }
func ToInt16(v interface{}) (int16, error) { return ldcore.ToInt16(v) }
func ToInt32(v interface{}) (int32, error) { return ldcore.ToInt32(v) }
func ToInt64(v interface{}) (int64, error) { return ldcore.ToInt64(v) }

func ToUint(v interface{}) (uint, error)     { return ldcore.ToUint(v) }
func ToUint8(v interface{}) (uint8, error)   { return ldcore.ToUint8(v) }
func ToUint16(v interface{}) (uint16, error) { return ldcore.ToUint16(v) }
func ToUint32(v interface{}) (uint32, error) { return ldcore.ToUint32(v) }
func ToUint64(v interface{}) (uint64, error) { return ldcore.ToUint64(v) }

func ToFloat32(v interface{}) (float32, error) { return ldcore.ToFloat32(v) }
func ToFloat64(v interface{}) (float64, error) { return ldcore.ToFloat64(v) }

func ToString(v interface{}) (string, error) { return ldcore.ToString(v) }
