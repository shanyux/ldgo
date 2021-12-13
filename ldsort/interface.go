/*
 * Copyright (C) distroy
 */

package ldsort

type Interface interface {
	Len() int
	Swap(i, j int)
	Compare(i, j int) int
}

type reverse struct {
	Interface
}

func (r reverse) Compare(i, j int) int { return -r.Interface.Compare(j, i) }

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	return &reverse{data}
}
