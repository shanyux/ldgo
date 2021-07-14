/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldsort"
)

type SortSliceString = ldsort.SortSliceString

func SortStrings(a []string)          { ldsort.SortStrings(a) }
func IsSortedStrings(a []string) bool { return ldsort.IsSortedStrings(a) }
func SearchStrings(a []string, x string) int {
	return ldsort.SearchStrings(a, x)
}
