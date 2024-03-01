/*
 * Copyright (C) distroy
 */

package ldrbtree

import "github.com/distroy/ldgo/v2/lditer"

var (
	_ lditer.Range[int] = (*RBTreeRange[int])(nil)
	_ lditer.Range[int] = (*RBTreeReverseRange[int])(nil)
)

type RBTreeRange[T any] struct {
	Begin RBTreeIterator[T] // [begin, end)
	End   RBTreeIterator[T] // [begin, end)
}

func (p *RBTreeRange[T]) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *RBTreeRange[T]) Next() {
	p.Begin = p.Begin.Next()
}

func (p *RBTreeRange[T]) Data() T {
	return p.Begin.Data()
}

type RBTreeReverseRange[T any] struct {
	Begin RBTreeReverseIterator[T]
	End   RBTreeReverseIterator[T]
}

func (p *RBTreeReverseRange[T]) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *RBTreeReverseRange[T]) Next() {
	p.Begin = p.Begin.Next()
}

func (p *RBTreeReverseRange[T]) Data() T {
	return p.Begin.Data()
}
