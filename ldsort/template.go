/*
 * Copyright (C) distroy
 */

package ldsort

type sortable interface {
	string |
		float32 | float64 |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

func compare[T sortable](a, b T) int {
	switch {
	case a == b:
		return 0
	case a > b:
		return +1
	default:
		return -1
	}
}

type Slice[T sortable] []T

func (s Slice[T]) Len() int      { return len(s) }
func (s Slice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Slice[T]) Compare(i, j int) int {
	return compare[T](s[i], s[j])
}

func templateSearch[T sortable](a []T, x T) int {
	return internalSearch(len(a), func(i int) bool { return compare[T](a[i], x) >= 0 })
}

func templateIndex[T sortable](a []T, x T) int {
	return internalIndex(len(a), func(i int) int {
		return compare[T](a[i], x)
	})
}
