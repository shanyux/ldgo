/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"github.com/distroy/ldgo/v2/ldcmp"
)

func DefaultCompare[T any](a, b T) int {
	return ldcmp.CompareInterface(a, b)
}
