/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"testing"
)

func testSearchInt64(t testing.TB, l []int64, x int64) {
	n := SearchInt64(l, x)
	t.Logf("%5d pos: %d", x, n)
}

func TestSortInt64(t *testing.T) {
	l := []int64{223, 562, 424, 642, 223, 123, 496, 623, 845, 375}
	if IsSortedInt64(l) {
		t.Fatal("is sorted: ", l)
	}

	SortInt64(l)
	if !IsSortedInt64(l) {
		t.Fatal("is not sorted: ", l)
	}
	t.Log("size: ", len(l))
	t.Log("sorted: ", l)

	testSearchInt64(t, l, 123)
	testSearchInt64(t, l, 223)
	testSearchInt64(t, l, 300)
	testSearchInt64(t, l, 1)
	testSearchInt64(t, l, 10000)
}