/*
 * Copyright (C) distroy
 */

package ldgo

import "testing"

func TestSortString(t *testing.T) {
	l := []string{"223", "562", "424", "642", "223", "abc", "aab", "22", "cbd", "abc"}
	if IsSortedString(l) {
		t.Fatal("is sorted: ", l)
	}

	SortString(l)
	if !IsSortedString(l) {
		t.Fatal("is not sorted: ", l)
	}
	t.Log("size: ", len(l))
	t.Log("sorted: ", l)
}
