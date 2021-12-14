/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"strings"
)

type Strings []string

func (s Strings) Len() int      { return len(s) }
func (s Strings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Strings) Compare(i, j int) int {
	return strings.Compare(s[i], s[j])
}

func SortStrings(a []string)          { internalSort(Strings(a)) }
func UniqStrings(a []string) []string { return a[:uniq(Strings(a))] }
func IsSortedStrings(a []string) bool { return internalIsSorted(Strings(a)) }
func SearchStrings(a []string, x string) int {
	return internalSearch(len(a), func(i int) bool { return a[i] >= x })
}
func IndexStrings(a []string, x string) int {
	if idx := SearchStrings(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}
