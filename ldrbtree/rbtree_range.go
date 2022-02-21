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
