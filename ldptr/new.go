/*
 * Copyright (C) distroy
 */

package ldptr

import "time"

func New[T any](d T) *T { return &d }

// Deprecated: use `New[Type]` instead.
func NewBool(d bool) *bool { return &d }

// Deprecated: use `New[Type]` instead.
func NewByte(d byte) *byte { return &d }

// Deprecated: use `New[Type]` instead.
func NewRune(d rune) *rune { return &d }

// Deprecated: use `New[Type]` instead.
func NewInt(d int) *int { return &d }

// Deprecated: use `New[Type]` instead.
func NewInt8(d int8) *int8 { return &d }

// Deprecated: use `New[Type]` instead.
func NewInt16(d int16) *int16 { return &d }

// Deprecated: use `New[Type]` instead.
func NewInt32(d int32) *int32 { return &d }

// Deprecated: use `New[Type]` instead.
func NewInt64(d int64) *int64 { return &d }

// Deprecated: use `New[Type]` instead.
func NewUint(d uint) *uint { return &d }

// Deprecated: use `New[Type]` instead.
func NewUint8(d uint8) *uint8 { return &d }

// Deprecated: use `New[Type]` instead.
func NewUint16(d uint16) *uint16 { return &d }

// Deprecated: use `New[Type]` instead.
func NewUint32(d uint32) *uint32 { return &d }

// Deprecated: use `New[Type]` instead.
func NewUint64(d uint64) *uint64 { return &d }

// Deprecated: use `New[Type]` instead.
func NewUintptr(d uintptr) *uintptr { return &d }

// Deprecated: use `New[Type]` instead.
func NewFloat32(d float32) *float32 { return &d }

// Deprecated: use `New[Type]` instead.
func NewFloat64(d float64) *float64 { return &d }

// Deprecated: use `New[Type]` instead.
func NewString(d string) *string { return &d }

// Deprecated: use `New[Type]` instead.
func NewComplex64(d complex64) *complex64 { return &d }

// Deprecated: use `New[Type]` instead.
func NewComplex128(d complex128) *complex128 { return &d }

// Deprecated: use `New[Type]` instead.
func NewTime(d time.Time) *time.Time { return &d }

// Deprecated: use `New[Type]` instead.
func NewDuration(d time.Duration) *time.Duration { return &d }
