/*
 * Copyright (C) distroy
 */

package ldrbtree

type MapRange struct {
	Begin MapIterator // [begin, end)
	End   MapIterator // [begin, end)
}

func (p *MapRange) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *MapRange) Next() {
	p.Begin = p.Begin.Next()
}

func (p *MapRange) Key() interface{} {
	return MapIterator(p.Begin).Key()
}

func (p *MapRange) Value(new ...interface{}) (old interface{}) {
	return MapIterator(p.Begin).Value(new...)
}

type MapReverseRange struct {
	Begin MapReverseIterator
	End   MapReverseIterator
}

func (p *MapReverseRange) HasNext() bool {
	return p.Begin.tree != nil && p.Begin != p.End
}

func (p *MapReverseRange) Next() {
	p.Begin = p.Begin.Next()
}

func (p *MapReverseRange) Key() interface{} {
	return MapReverseIterator(p.Begin).Key()
}

func (p *MapReverseRange) Value(new ...interface{}) (old interface{}) {
	return MapReverseIterator(p.Begin).Value(new...)
}
