/*
 * Copyright (C) distroy
 */

package ldcore

type TopkInterface interface {
	Len() int
	Less(i, j interface{}) bool
	Swap(i, j int)

	Push(x interface{})
	Get(i int) interface{}
	Set(i int, x interface{})
}

func TopkAdd(b TopkInterface, k int, x interface{}) bool {
	if pos := b.Len(); pos < k {
		b.Push(x)

		for parent := 0; pos > 0; pos = parent {
			parent = (pos - 1) / 2
			if !b.Less(b.Get(parent), b.Get(pos)) {
				break
			}
			b.Swap(parent, pos)
		}

		return true
	}

	if !b.Less(x, b.Get(0)) {
		return false
	}

	b.Set(0, x)

	for pos := 0; ; {
		lChild := (pos * 2) + 1
		rChild := (pos * 2) + 2
		if lChild >= k {
			break
		}

		child := lChild
		if rChild < k && b.Less(b.Get(lChild), b.Get(rChild)) {
			child = rChild
		}

		if !b.Less(b.Get(pos), b.Get(child)) {
			break
		}

		b.Swap(pos, child)
		pos = child
	}

	return true
}
