/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"fmt"
	"reflect"
	"sort"
)

type sortStructForCompare struct {
	Size    int
	Array   reflect.Value
	CmpFunc reflect.Value
}

func (s *sortStructForCompare) Len() int { return s.Size }

func (s *sortStructForCompare) Swap(i, j int) {
	if i == j {
		return
	}
	a := s.Array.Index(i)
	b := s.Array.Index(j)
	t := a.Interface()
	a.Set(b)
	b.Set(reflect.ValueOf(t))
}

func (s *sortStructForCompare) Compare(i, j int) int {
	args := [2]reflect.Value{s.Array.Index(i), s.Array.Index(j)}
	res := s.CmpFunc.Call(args[:])
	return int(res[0].Int())
}

type sortStructForLess struct {
	Size     int
	Array    reflect.Value
	LessFunc reflect.Value
}

func (s *sortStructForLess) Len() int { return s.Size }

func (s *sortStructForLess) Swap(i, j int) {
	if i == j {
		return
	}
	a := s.Array.Index(i)
	b := s.Array.Index(j)
	t := a.Interface()
	a.Set(b)
	b.Set(reflect.ValueOf(t))
}

func (s *sortStructForLess) Less(i, j int) bool {
	args := [2]reflect.Value{s.Array.Index(i), s.Array.Index(j)}
	res := s.LessFunc.Call(args[:])
	return res[0].Bool()
}

// Sort sorts slice with lessFunc
// slice type must be slice
// lessOrCompareFunc type must be:
//		func less(a, b TypeOfSliceElement) bool
//		func compare(a, b TypeOfSliceElement) int
func Sort(slice interface{}, lessOrCompareFunc interface{}) {
	aVal := reflect.ValueOf(slice)
	aType := aVal.Type()
	switch aType.Kind() {
	case reflect.Slice:
		break
	default:
		panic(fmt.Sprintf("parameter array must be slice, type: %s", aType))
	}
	eType := aType.Elem()

	fVal := reflect.ValueOf(lessOrCompareFunc)
	fType := fVal.Type()
	if fType.Kind() != reflect.Func {
		panic(fmt.Sprintf("less func must be func, type: %s", fType))
	}

	if n := fType.NumIn(); n != 2 {
		panic(fmt.Sprintf("less func must have 2 input parameters, type: %s", fType))
	}

	type0 := fType.In(0)
	type1 := fType.In(1)
	if type0 != type1 || !(type0 == eType || (type0.Kind()) == reflect.Interface && eType.Implements(type0)) {
		panic(fmt.Sprintf("second parameters of less func must be or implements %s", eType))
	}

	if n := fType.NumOut(); n != 1 {
		panic(fmt.Sprintf("less func must have 1 output parameter, type: %s", fType))
	}

	switch typ := fType.Out(0); typ.Kind() {
	default:
		panic(fmt.Sprintf("parameter of less/compare func must be bool/int, type: %s", typ))

	case reflect.Bool:
		sort.Sort(&sortStructForLess{
			Size:     aVal.Len(),
			Array:    aVal,
			LessFunc: fVal,
		})

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		internalSort(&sortStructForCompare{
			Size:    aVal.Len(),
			Array:   aVal,
			CmpFunc: fVal,
		})
	}
}

type sortIface struct {
	Interface
}

func (s sortIface) Less(i, j int) bool {
	return s.Compare(i, j) < 0
}

func internalSort(slice Interface)          { sort.Sort(sortIface{Interface: slice}) }
func internalIsSorted(slice Interface) bool { return sort.IsSorted(sortIface{Interface: slice}) }

// Search uses binary search to find and return the smallest index in [0, n) at which f(i) is true
func Search(n int, f func(i int) bool) int { return sort.Search(n, f) }
