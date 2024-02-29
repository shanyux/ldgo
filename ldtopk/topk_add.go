/*
 * Copyright (C) distroy
 */

package ldtopk

type sortable interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func less[T sortable](a, b T) bool {
	return a < b
}

func TopkAdd[T sortable](b []T, k int, x T) ([]T, bool) {
	if k <= 0 {
		return b, false
	}

	if pos := len(b); pos < k {
		topkAppendTail(&b, less[T], x)
		return b, true
	}

	if !less[T](x, b[0]) {
		return b, false
	}

	b[0] = x
	topkFixupHead[T](&b, less[T])
	return b, true
}
