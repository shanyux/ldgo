/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldsort"
)

// Sort sorts slice with lessOrCompareFunc
// slice type must be slice
// lessOrCompareFunc type must be:
//
//	func less(a, b TypeOfSliceElement) bool
//	func compare(a, b TypeOfSliceElement) int
func Sort(slice interface{}, lessOrCompareFunc interface{}) {
	ldsort.Sort(slice, lessOrCompareFunc)
}
