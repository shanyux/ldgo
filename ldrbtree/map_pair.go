/*
 * Copyright (C) distroy
 */

package ldrbtree

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func wrapPairCompare[K any, V any](m *Map[K, V]) func(a, b Pair[K, V]) int {
	compare := m.KeyCompare
	return func(a, b Pair[K, V]) int {
		return compare(a.Key, b.Key)
	}
}
