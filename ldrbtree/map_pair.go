/*
 * Copyright (C) distroy
 */

package ldrbtree

type Pair struct {
	Key   interface{}
	Value interface{}
}

func wrapPairCompare(m *Map) func(a, b interface{}) int {
	return func(a, b interface{}) int {
		aa, bb := a.(Pair), b.(Pair)
		return m.KeyCompare(aa.Key, bb.Key)
	}
}
