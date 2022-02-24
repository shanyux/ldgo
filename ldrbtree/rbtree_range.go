/*
 * Copyright (C) distroy
 */

package ldrbtree

type RBTreeRange struct {
	Begin RBTreeIterator // [begin, end)
	End   RBTreeIterator // [begin, end)
}

func (p *RBTreeRange) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *RBTreeRange) Next() {
	p.Begin = p.Begin.Next()
}

func (p *RBTreeRange) Data() interface{} {
	return p.Begin.Data()
}

type RBTreeReverseRange struct {
	Begin RBTreeReverseIterator
	End   RBTreeReverseIterator
}

func (p *RBTreeReverseRange) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *RBTreeReverseRange) Next() {
	p.Begin = p.Begin.Next()
}

func (p *RBTreeReverseRange) Data() interface{} {
	return p.Begin.Data()
}
