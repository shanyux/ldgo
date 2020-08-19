/*
 * Copyright (C) distroy
 */

package ldgo

type topkSliceString []string

func (s topkSliceString) Len() int                   { return len(s) }
func (s topkSliceString) Swap(i, j int)              { s[i], s[j] = s[j], s[i] }
func (s topkSliceString) Less(i, j interface{}) bool { return i.(string) < j.(string) }
func (s topkSliceString) Get(i int) interface{}      { return s[i] }
func (s topkSliceString) Set(i int, x interface{})   { s[i] = x.(string) }
func (s *topkSliceString) Push(x interface{})        { *s = append(*s, x.(string)) }

func TopkStringsAdd(b []string, k int, x string) ([]string, bool) {
	ok := TopkAdd((*topkSliceString)(&b), k, x)
	return b, ok
}
