/*
 * Copyright (C) distroy
 */

package ldsort

import "sort"

type Strings []string

func (s Strings) Len() int           { return len(s) }
func (s Strings) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Strings) Less(i, j int) bool { return s[i] < s[j] }

func SortStrings(a []string)          { sort.Sort(Strings(a)) }
func IsSortedStrings(a []string) bool { return sort.IsSorted(Strings(a)) }
func SearchStrings(a []string, x string) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}
func IndexStrings(a []string, x string) int {
	if idx := SearchStrings(a, x); idx < len(a) && a[idx] == x {
		return idx
	}
	return -1
}
