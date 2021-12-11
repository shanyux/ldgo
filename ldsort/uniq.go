/*
 * Copyright (C) distroy
 */

package ldsort

func uniq(a Interface) int {
	if a.Len() <= 0 {
		return a.Len()
	}

	i := 0
	for j := 1; j < a.Len(); j++ {
		if a.Compare(i, j) == 0 {
			continue
		}
		i++
		a.Swap(i, j)
	}
	return i + 1
}
