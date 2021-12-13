/*
 * Copyright (C) distroy
 */

package ldsort

func uniq(a Interface) int {
	l := a.Len()
	if l <= 0 {
		return 0
	}

	i := 0
	for j := 1; j < l; j++ {
		if a.Compare(i, j) == 0 {
			continue
		}
		i++
		a.Swap(i, j)
	}
	return i + 1
}
