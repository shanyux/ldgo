/*
 * Copyright (C) distroy
 */

package ldrbtree

type Map struct {
	tree       RBTree
	KeyCompare CompareFunc
	// ValueCompare CompareFunc
}

func (m *Map) init() {
	if m.KeyCompare == nil {
		m.KeyCompare = DefaultCompare
	}
	// if m.ValueCompare == nil {
	// 	m.ValueCompare = DefaultCompare
	// }
	if m.tree.Compare == nil {
		m.tree.Compare = wrapPairCompare(m)
	}

	m.tree.init()
}

func (m *Map) Clear() {
	m.init()
	m.tree.Clear()
}

func (m *Map) Insert(key, value interface{}) MapIterator {
	m.init()
	return MapIterator(m.tree.Insert(Pair{Key: key, Value: value}))
}

func (m *Map) InsertOrSearch(key, value interface{}) MapIterator {
	m.init()
	return MapIterator(m.tree.InsertOrSearch(Pair{Key: key, Value: value}))
}

func (m *Map) InsertOrAssign(key, value interface{}) MapIterator {
	m.init()
	return MapIterator(m.tree.InsertOrAssign(Pair{Key: key, Value: value}))
}

func (m *Map) Delete(it MapIterator) MapIterator {
	m.init()
	return MapIterator(m.tree.Delete(RBTreeIterator(it)))
}

func (m *Map) Count(key interface{}) int {
	m.init()
	return m.tree.Count(Pair{Key: key})
}

// Search returns the first pair.key == key
func (m *Map) Search(key interface{}) MapIterator {
	m.init()
	return MapIterator(m.tree.Search(Pair{Key: key}))
}

// SearchRange returns the node range [first pair.key == key, first pair.key > key)
func (m *Map) SearchRange(key interface{}) *MapRange {
	m.init()
	it := m.tree.SearchRange(Pair{Key: key})
	return &MapRange{
		Begin: MapIterator(it.Begin),
		End:   MapIterator(it.End),
	}
}

// LowerBound returns the first pair.key >= key
func (m *Map) LowerBound(key interface{}) MapIterator {
	m.init()
	return MapIterator{
		tree: &m.tree,
		node: rbtreeLowerBound(Pair{Key: key}, forward(&m.tree)),
	}
}

// UpperBound returns the first pair.key > key
func (m *Map) UpperBound(key interface{}) MapIterator {
	m.init()
	return MapIterator{
		tree: &m.tree,
		node: rbtreeUpperBound(Pair{Key: key}, forward(&m.tree)),
	}
}

func (m *Map) Range() *MapRange {
	m.init()
	return &MapRange{
		Begin: MapIterator(rbtreeBeginIterator(forward(&m.tree))),
		End:   MapIterator(rbtreeEndIterator(forward(&m.tree))),
	}
}

func (m *Map) Begin() MapIterator {
	m.init()
	return MapIterator(rbtreeBeginIterator(forward(&m.tree)))
}

func (m *Map) End() MapIterator {
	m.init()
	return MapIterator(rbtreeEndIterator(forward(&m.tree)))
}

// RDelete is reverse delete
func (m *Map) RDelete(it MapReverseIterator) MapReverseIterator {
	m.init()
	return MapReverseIterator(m.tree.RDelete(RBTreeReverseIterator(it)))
}

// RSearch is reverse search
// RSearch returns the last pair.key == key
func (m *Map) RSearch(key interface{}) MapReverseIterator {
	m.init()
	return MapReverseIterator(m.RSearch(Pair{Key: key}))
}

// RSearchRange is reverse search range
// RSearchRange returns the node range [last pair.key == key, last pair.key < key)
func (m *Map) RSearchRange(key interface{}) *MapReverseRange {
	m.init()
	it := m.tree.RSearchRange(Pair{Key: key})
	return &MapReverseRange{
		Begin: MapReverseIterator(it.Begin),
		End:   MapReverseIterator(it.End),
	}
}

// RLowerBound is reverse lower bound
// RLowerBound returns the last pair.key <= key
func (m *Map) RLowerBound(key interface{}) MapReverseIterator {
	m.init()
	return MapReverseIterator{
		tree: &m.tree,
		node: rbtreeLowerBound(Pair{Key: key}, reverse(&m.tree)),
	}
}

// RUpperBound is reverse upper bound
// RUpperBound returns the last pair.key < key
func (m *Map) RUpperBound(key interface{}) MapReverseIterator {
	m.init()
	return MapReverseIterator{
		tree: &m.tree,
		node: rbtreeUpperBound(Pair{Key: key}, reverse(&m.tree)),
	}
}

// RRange is reverse range
func (m *Map) RRange() *MapReverseRange {
	m.init()
	return &MapReverseRange{
		Begin: MapReverseIterator(rbtreeBeginIterator(reverse(&m.tree))),
		End:   MapReverseIterator(rbtreeEndIterator(reverse(&m.tree))),
	}
}

// RBegin is reverse begin
func (m *Map) RBegin() MapReverseIterator {
	m.init()
	return MapReverseIterator(rbtreeBeginIterator(reverse(&m.tree)))
}

// REnd is reverse end
func (m *Map) REnd() MapReverseIterator {
	m.init()
	return MapReverseIterator(rbtreeEndIterator(reverse(&m.tree)))
}
