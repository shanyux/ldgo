/*
 * Copyright (C) distroy
 */

package cmp

func isNaN[T comparable](a T) bool {
	return a != a
}

type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func CompareOrderable[T Orderable](a, b T) int {
	switch {
	case isNaN(a):
		if isNaN(b) {
			return 0
		}
		return -1 // No good answer if b is a NaN so don't bother checking.
	case isNaN(b):
		return 1
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}
