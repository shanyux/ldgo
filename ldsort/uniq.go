/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"fmt"
	"reflect"
)

func uniq(s Interface) int {
	l := s.Len()
	if l <= 0 {
		return 0
	}

	i := 0
	for j := 1; j < l; j++ {
		if s.Compare(i, j) == 0 {
			continue
		}
		i++
		s.Swap(i, j)
	}
	return i + 1
}

// Uniq uniqs slice with compareFunc
// slice type must be slice
// compareFunc type must be:
//
//	func compare(a, b TypeOfSliceElement) int
func Uniq(slice interface{}, compareFunc interface{}) int {
	if slice == nil {
		return 0
	}
	i := getInterface(slice, compareFunc)
	return uniq(i)
}

type uniqStruct struct {
	size    int
	array   reflect.Value
	cmpFunc reflect.Value
}

func (s *uniqStruct) Len() int { return s.size }

func (s *uniqStruct) Swap(i, j int) {
	if i == j {
		return
	}
	a := s.array.Index(i)
	b := s.array.Index(j)
	t := a.Interface()
	a.Set(b)
	b.Set(reflect.ValueOf(t))
}

func (s *uniqStruct) Compare(i, j int) int {
	args := [2]reflect.Value{s.array.Index(i), s.array.Index(j)}
	res := s.cmpFunc.Call(args[:])
	switch r := res[0].Int(); {
	case r == 0:
		return 0
	case r < 0:
		return -1
	default:
		return 1
	}
}

func getInterface(slice interface{}, compareFunc interface{}) Interface {
	aVal := reflect.ValueOf(slice)
	aType := aVal.Type()
	switch aType.Kind() {
	case reflect.Slice:
		break
	default:
		panic(fmt.Sprintf("parameter array must be slice, type: %s", aType))
	}
	eType := aType.Elem()

	fVal := reflect.ValueOf(compareFunc)
	fType := fVal.Type()
	if fType.Kind() != reflect.Func {
		panic(fmt.Sprintf("compare func must be func, type: %s", fType))
	}

	if n := fType.NumIn(); n != 2 {
		panic(fmt.Sprintf("compare func must have 2 input parameters, type: %s", fType))
	}

	type0 := fType.In(0)
	type1 := fType.In(1)
	if type0 != type1 || !(type0 == eType || (type0.Kind()) == reflect.Interface && eType.Implements(type0)) {
		panic(fmt.Sprintf("second parameters of compare func must be or implements %s", eType))
	}

	if n := fType.NumOut(); n != 1 {
		panic(fmt.Sprintf("compare func must have 1 output parameter, type: %s", fType))
	}

	switch typ := fType.Out(0); typ.Kind() {
	default:
		panic(fmt.Sprintf("parameter of compare func must be bool/int, type: %s", typ))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &uniqStruct{
			size:    aVal.Len(),
			array:   aVal,
			cmpFunc: fVal,
		}
	}
}
