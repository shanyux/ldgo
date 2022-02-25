/*
 * Copyright (C) distroy
 */

package ldptr

import "time"

func NewBoolByPtr(d *bool) *bool {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewByteByPtr(d *byte) *byte {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewRuneByPtr(d *rune) *rune {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewIntByPtr(d *int) *int {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewInt8ByPtr(d *int8) *int8 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewInt16ByPtr(d *int16) *int16 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewInt32ByPtr(d *int32) *int32 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewInt64ByPtr(d *int64) *int64 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewUintByPtr(d *uint) *uint {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewUint8ByPtr(d *uint8) *uint8 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewUint16ByPtr(d *uint16) *uint16 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewUint32ByPtr(d *uint32) *uint32 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewUint64ByPtr(d *uint64) *uint64 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewUintptrByPtr(d *uintptr) *uintptr {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewFloat32ByPtr(d *float32) *float32 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewFloat64ByPtr(d *float64) *float64 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewStringByPtr(d *string) *string {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewComplex64ByPtr(d *complex64) *complex64 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewComplex128ByPtr(d *complex128) *complex128 {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}

func NewTimeByPtr(d *time.Time) *time.Time {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
func NewDurationByPtr(d *time.Duration) *time.Duration {
	if d == nil {
		return nil
	}
	cp := *d
	return &cp
}
