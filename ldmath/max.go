/*
 * Copyright (C) distroy
 */

package ldmath

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Max[T Number](n T, args ...T) T {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxInt(n int, args ...int) int {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxInt8(n int8, args ...int8) int8 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxInt16(n int16, args ...int16) int16 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxInt32(n int32, args ...int32) int32 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxInt64(n int64, args ...int64) int64 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxUint(n uint, args ...uint) uint {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxUint8(n uint8, args ...uint8) uint8 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxUint16(n uint16, args ...uint16) uint16 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxUint32(n uint32, args ...uint32) uint32 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxUint64(n uint64, args ...uint64) uint64 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxFloat32(n float32, args ...float32) float32 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}

// Deprecated: use `Max[Type]` instead.
func MaxFloat64(n float64, args ...float64) float64 {
	for _, v := range args {
		if n < v {
			n = v
		}
	}
	return n
}
