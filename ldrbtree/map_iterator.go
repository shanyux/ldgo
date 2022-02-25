/*
 * Copyright (C) distroy
 */

package ldrbtree

type mapIterator = rbtreeIterator

type MapIterator mapIterator

func (i MapIterator) base() mapIterator {
	return mapIterator(i)
}

func (i MapIterator) Key() interface{} {
	return i.node.Data.(Pair).Key
}

func (i MapIterator) Value(new ...interface{}) (old interface{}) {
	oldData := i.node.Data.(Pair)
	if len(new) > 0 {
		i.node.Data = Pair{Key: oldData.Key, Value: new[0]}
	}
	return oldData.Value
}

func (i MapIterator) Next() MapIterator {
	return MapIterator(i.base().next("map iterator", forward(i.tree)))
}

func (i MapIterator) Prev() MapIterator {
	return MapIterator(i.base().prev("map iterator", forward(i.tree)))
}

type MapReverseIterator mapIterator

func (i MapReverseIterator) base() mapIterator {
	return mapIterator(i)
}

func (i MapReverseIterator) Key() interface{} {
	return i.node.Data.(Pair).Key
}

func (i MapReverseIterator) Value(new ...interface{}) (old interface{}) {
	oldData := i.node.Data.(Pair)
	if len(new) > 0 {
		i.node.Data = Pair{Key: oldData.Key, Value: new[0]}
	}
	return oldData.Value
}

func (i MapReverseIterator) Next() MapReverseIterator {
	return MapReverseIterator(i.base().next("map reverse iterator", reverse(i.tree)))
}

func (i MapReverseIterator) Prev() MapReverseIterator {
	return MapReverseIterator(i.base().prev("map reverse iterator", reverse(i.tree)))
}
