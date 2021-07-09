/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

type SortSliceString = ldcore.SortSliceString

func SortStrings(a []string)          { ldcore.SortStrings(a) }
func IsSortedStrings(a []string) bool { return ldcore.IsSortedStrings(a) }
func SearchStrings(a []string, x string) int {
	return ldcore.SearchStrings(a, x)
}
