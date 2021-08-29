/*
 * Copyright (C) distroy
 */

package ldnew

func Byte(d byte) *byte { return &d }
func Rune(d rune) *rune { return &d }

func Int(d int) *int       { return &d }
func Int8(d int8) *int8    { return &d }
func Int16(d int16) *int16 { return &d }
func Int32(d int32) *int32 { return &d }
func Int64(d int64) *int64 { return &d }

func Uint(d uint) *uint       { return &d }
func Uint8(d uint8) *uint8    { return &d }
func Uint16(d uint16) *uint16 { return &d }
func Uint32(d uint32) *uint32 { return &d }
func Uint64(d uint64) *uint64 { return &d }

func Float32(d float32) *float32 { return &d }
func Float64(d float64) *float64 { return &d }

func String(d string) *string { return &d }
