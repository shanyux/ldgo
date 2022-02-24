/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"testing"

	"github.com/distroy/ldgo/ldrand"
	"github.com/distroy/ldgo/ldsort"
)

const (
	_count = 100
)

var (
	_numsUnordered []int
	_nums          []int
)

func TestMain(m *testing.M) {
	_numsUnordered = ldrand.Perm(_count)
	_nums = make([]int, 0, len(_numsUnordered))
	_nums = append(_nums, _numsUnordered...)
	ldsort.SortInts(_nums)

	// log.Printf("unordered: %v", _numsUnordered)
	// log.Printf("nums: %v", _nums)

	m.Run()
}

func testNewRBTree() *RBTree {
	rbtree := &RBTree{}
	for _, n := range _numsUnordered {
		rbtree.Insert(n)
	}
	return rbtree
}

func testRBTreeDeleteAll(rbtree *RBTree, d interface{}) {
	for it := rbtree.Search(d); it != rbtree.End(); it = rbtree.Search(d) {
		rbtree.Delete(it)
	}
}

func testRBTreeRDeleteAll(rbtree *RBTree, d interface{}) {
	for it := rbtree.RSearch(d); it != rbtree.REnd(); it = rbtree.RSearch(d) {
		rbtree.RDelete(it)
	}
}
