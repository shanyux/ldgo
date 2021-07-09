/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

// Sort sorts slice with lessFunc
// slice type must be slice
// lessFunc type must be func (a, b TypeOfSliceElement) bool
func Sort(slice interface{}, lessFunc interface{}) {
	ldcore.Sort(slice, lessFunc)
}
