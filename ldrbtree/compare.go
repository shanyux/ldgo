/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"github.com/distroy/ldgo/ldcmp"
)

func DefaultCompare(a, b interface{}) int {
	return ldcmp.CompareInterface(a, b)
}
