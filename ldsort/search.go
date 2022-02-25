/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"sort"
)

func internalSearch(n int, f func(i int) bool) int { return sort.Search(n, f) }

// Search uses binary search to find and return the smallest index in [0, n)
// at which f(i) == true, or return n if there is no f(i) == true
func Search(n int, f func(i int) bool) int {
	return internalSearch(n, f)
}

// Index uses binary search to find and return the smallest index in [0, n)
// at which compare(i) == 0, or return -1 if there is no compare(i) == 0
func Index(n int, compare func(i int) int) int {
	idx := internalSearch(n, func(i int) bool {
		return compare(i) >= 0
	})
	if idx < n && compare(idx) == 0 {
		return idx
	}
	return -1
}
