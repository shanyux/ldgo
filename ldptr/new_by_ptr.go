/*
 * Copyright (C) distroy
 */

package ldptr

import "time"

func NewByPtr[T any](d *T, def ...T) *T {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewBoolByPtr(d *bool, def ...bool) *bool {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewByteByPtr(d *byte, def ...byte) *byte {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewRuneByPtr(d *rune, def ...rune) *rune {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewIntByPtr(d *int, def ...int) *int {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewInt8ByPtr(d *int8, def ...int8) *int8 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewInt16ByPtr(d *int16, def ...int16) *int16 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewInt32ByPtr(d *int32, def ...int32) *int32 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewInt64ByPtr(d *int64, def ...int64) *int64 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUintByPtr(d *uint, def ...uint) *uint {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUint8ByPtr(d *uint8, def ...uint8) *uint8 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUint16ByPtr(d *uint16, def ...uint16) *uint16 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUint32ByPtr(d *uint32, def ...uint32) *uint32 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUint64ByPtr(d *uint64, def ...uint64) *uint64 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewUintptrByPtr(d *uintptr, def ...uintptr) *uintptr {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewFloat32ByPtr(d *float32, def ...float32) *float32 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewFloat64ByPtr(d *float64, def ...float64) *float64 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewStringByPtr(d *string, def ...string) *string {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewComplex64ByPtr(d *complex64, def ...complex64) *complex64 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewComplex128ByPtr(d *complex128, def ...complex128) *complex128 {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewTimeByPtr(d *time.Time, def ...time.Time) *time.Time {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}

// Deprecated: use `NewByPtr[Type]` instead.
func NewDurationByPtr(d *time.Duration, def ...time.Duration) *time.Duration {
	if d == nil {
		if len(def) > 0 {
			return &def[0]
		}
		return nil
	}
	cp := *d
	return &cp
}
