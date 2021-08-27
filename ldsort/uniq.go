/*
 * Copyright (C) distroy
 */

package ldsort

import "sort"

func compare(a sort.Interface, i, j int) int {
	iLess := a.Less(i, j)
	jLess := a.Less(j, i)
	if iLess == jLess {
		return 0
	}
	if iLess {
		return -1
	}
	return 1
}

func uniq(a sort.Interface) int {
	if a.Len() <= 0 {
		return a.Len()
	}

	i := 0
	for j := 1; j < a.Len(); j++ {
		if compare(a, i, j) == 0 {
			continue
		}
		i++
		a.Swap(i, j)
	}
	return i + 1
}
