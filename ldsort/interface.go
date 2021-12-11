/*
 * Copyright (C) distroy
 */

package ldsort

type Interface interface {
	Len() int
	Swap(i, j int)
	Compare(i, j int) int
}
