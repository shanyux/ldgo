/*
 * Copyright (C) distroy
 */

package ldmath

func Sum[T Number](args ...T) T {
	var sum T = 0
	for _, v := range args {
		sum += v
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumInt(args ...int) int64 {
	sum := int64(0)
	for _, v := range args {
		sum += int64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumInt8(args ...int8) int64 {
	sum := int64(0)
	for _, v := range args {
		sum += int64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumInt16(args ...int16) int64 {
	sum := int64(0)
	for _, v := range args {
		sum += int64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumInt32(args ...int32) int64 {
	sum := int64(0)
	for _, v := range args {
		sum += int64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumInt64(args ...int64) int64 {
	sum := int64(0)
	for _, v := range args {
		sum += int64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumUint(args ...uint) uint64 {
	sum := uint64(0)
	for _, v := range args {
		sum += uint64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumUint8(args ...uint8) uint64 {
	sum := uint64(0)
	for _, v := range args {
		sum += uint64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumUint16(args ...uint16) uint64 {
	sum := uint64(0)
	for _, v := range args {
		sum += uint64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumUint32(args ...uint32) uint64 {
	sum := uint64(0)
	for _, v := range args {
		sum += uint64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumUint64(args ...uint64) uint64 {
	sum := uint64(0)
	for _, v := range args {
		sum += uint64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumFloat32(args ...float32) float64 {
	sum := float64(0)
	for _, v := range args {
		sum += float64(v)
	}
	return sum
}

// Deprecated: use `Sum[Type]` instead.
func SumFloat64(args ...float64) float64 {
	sum := float64(0)
	for _, v := range args {
		sum += float64(v)
	}
	return sum
}
