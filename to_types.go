/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldconv"
)

func ToByte(v interface{}) (byte, error) { return ldconv.ToByte(v) }

func ToBool(v interface{}) (bool, error) { return ldconv.ToBool(v) }

func ToInt(v interface{}) (int, error)     { return ldconv.ToInt(v) }
func ToInt8(v interface{}) (int8, error)   { return ldconv.ToInt8(v) }
func ToInt16(v interface{}) (int16, error) { return ldconv.ToInt16(v) }
func ToInt32(v interface{}) (int32, error) { return ldconv.ToInt32(v) }
func ToInt64(v interface{}) (int64, error) { return ldconv.ToInt64(v) }

func ToUint(v interface{}) (uint, error)     { return ldconv.ToUint(v) }
func ToUint8(v interface{}) (uint8, error)   { return ldconv.ToUint8(v) }
func ToUint16(v interface{}) (uint16, error) { return ldconv.ToUint16(v) }
func ToUint32(v interface{}) (uint32, error) { return ldconv.ToUint32(v) }
func ToUint64(v interface{}) (uint64, error) { return ldconv.ToUint64(v) }

func ToFloat32(v interface{}) (float32, error) { return ldconv.ToFloat32(v) }
func ToFloat64(v interface{}) (float64, error) { return ldconv.ToFloat64(v) }

func ToString(v interface{}) (string, error) { return ldconv.ToString(v) }
