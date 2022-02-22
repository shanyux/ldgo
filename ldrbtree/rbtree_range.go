/*
 * Copyright (C) distroy
 */

package ldrbtree

type RBTreeRange struct {
	begin RBTreeIterator
	end   RBTreeIterator
}

func (p *RBTreeRange) IsValid() bool {
	return p.begin != p.end
}

func (p *RBTreeRange) Next() {
	p.begin = p.begin.Next()
}

func (p *RBTreeRange) Data() interface{} {
	return p.begin.Data()
}

func (p *RBTreeRange) Iterator() RBTreeIterator {
	return p.begin
}

type RBTreeReverseRange struct {
	begin RBTreeReverseIterator
	end   RBTreeReverseIterator
}

func (p *RBTreeReverseRange) IsValid() bool {
	return p.begin != p.end
}

func (p *RBTreeReverseRange) Next() {
	p.begin = p.begin.Next()
}

func (p *RBTreeReverseRange) Data() interface{} {
	return p.begin.Data()
}

func (p *RBTreeReverseRange) Iterator() RBTreeReverseIterator {
	return p.begin
}
