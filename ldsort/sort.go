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

func (s *sortStructForCompare) Less(i, j int) bool {
	args := [2]reflect.Value{s.Array.Index(i), s.Array.Index(j)}
	res := s.CmpFunc.Call(args[:])
	return res[0].Int() < 0
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

func getSortInterface(slice interface{}, lessOrCompareFunc interface{}) sort.Interface {
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
		panic(fmt.Sprintf("less/compare func must be func, type: %s", fType))
	}

	if n := fType.NumIn(); n != 2 {
		panic(fmt.Sprintf("less/compare func must have 2 input parameters, type: %s", fType))
	}

	type0 := fType.In(0)
	type1 := fType.In(1)
	if type0 != type1 || !(type0 == eType || (type0.Kind()) == reflect.Interface && eType.Implements(type0)) {
		panic(fmt.Sprintf("second parameters of less/compare func must be or implements %s", eType))
	}

	if n := fType.NumOut(); n != 1 {
		panic(fmt.Sprintf("less/compare func must have 1 output parameter, type: %s", fType))
	}

	switch typ := fType.Out(0); typ.Kind() {
	default:
		panic(fmt.Sprintf("parameter of less/compare func must be bool/int, type: %s", typ))

	case reflect.Bool:
		return &sortStructForLess{
			Size:     aVal.Len(),
			Array:    aVal,
			LessFunc: fVal,
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &sortStructForCompare{
			Size:    aVal.Len(),
			Array:   aVal,
			CmpFunc: fVal,
		}
	}
}

// Sort sorts slice with lessOrCompareFunc
// slice type must be slice
// lessOrCompareFunc type must be:
//
//	func compare(a, b TypeOfSliceElement) int
//	func less(a, b TypeOfSliceElement) bool
func Sort(slice interface{}, lessOrCompareFunc interface{}) {
	if slice == nil {
		return
	}
	i := getSortInterface(slice, lessOrCompareFunc)
	sort.Sort(i)
}

// IsSorted reports whether data is sorted.
// slice type must be slice
// lessOrCompareFunc type must be:
//
//	func compare(a, b TypeOfSliceElement) int
//	func less(a, b TypeOfSliceElement) bool
func IsSorted(slice interface{}, lessOrCompareFunc interface{}) bool {
	if slice == nil {
		return true
	}
	i := getSortInterface(slice, lessOrCompareFunc)
	return sort.IsSorted(i)
}

type sortIface struct {
	Interface
}

func (s sortIface) Less(i, j int) bool {
	return s.Compare(i, j) < 0
}

func internalSort(s Interface)          { sort.Sort(sortIface{Interface: s}) }
func internalIsSorted(s Interface) bool { return sort.IsSorted(sortIface{Interface: s}) }
