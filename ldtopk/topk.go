/*
 * Copyright (C) distroy
 */

package ldtopk

import "github.com/distroy/ldgo/v2/ldcmp"

type LessFunc = func(a, b interface{}) bool

func DefaultLess(a, b interface{}) bool {
	return ldcmp.CompareInterface(a, b) < 0
}

// TopK will keep at most size elements for which less returns true
type TopK struct {
	Size int           // K
	Less LessFunc      // less func
	data []interface{} //
}

func (p *TopK) Data() []interface{} { return p.data }

func (p *TopK) Add(d interface{}) bool {
	k := p.Size
	if k <= 0 {
		return false
	}

	p.init()
	if pos := len(p.data); pos < k {
		headAppendTail(&p.data, p.Less, d)
		return true
	}

	if !p.Less(d, p.data[0]) {
		return false
	}
	p.data[0] = d

	heapFixupHead(&p.data, p.Less)
	return true
}

func (p *TopK) init() {
	if p.Less == nil {
		p.Less = DefaultLess
	}
	if p.data == nil && p.Size > 0 {
		p.data = make([]interface{}, 0, p.Size)
	}
}

func headAppendTail(heap *[]interface{}, less LessFunc, d interface{}) {
	pos := len(*heap)

	*heap = append(*heap, d)
	for parent := 0; pos > 0; pos = parent {
		parent = (pos - 1) / 2
		if !less((*heap)[parent], (*heap)[pos]) {
			break
		}
		(*heap)[parent], (*heap)[pos] = (*heap)[pos], (*heap)[parent]
	}
}

func heapFixupHead(heap *[]interface{}, less LessFunc) {
	size := len(*heap)
	for pos := 0; ; {
		lChild := (pos * 2) + 1
		rChild := (pos * 2) + 2
		if lChild >= size {
			break
		}

		child := lChild
		if rChild < size && less((*heap)[lChild], (*heap)[rChild]) {
			child = rChild
		}

		if !less((*heap)[pos], (*heap)[child]) {
			break
		}

		(*heap)[pos], (*heap)[child] = (*heap)[child], (*heap)[pos]
		pos = child
	}
}
