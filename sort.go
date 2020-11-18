/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"fmt"
	"reflect"
	"sort"
)

type sortReflect struct {
	Size     int
	Array    reflect.Value
	LessFunc reflect.Value
}

func (s sortReflect) Len() int { return s.Size }

func (s sortReflect) Swap(i, j int) {
	if i == j {
		return
	}
	a := s.Array.Index(i)
	b := s.Array.Index(j)
	t := a.Interface()
	a.Set(b)
	b.Set(reflect.ValueOf(t))
}

func (s sortReflect) Less(i, j int) bool {
	args := [2]reflect.Value{s.Array.Index(i), s.Array.Index(j)}
	res := s.LessFunc.Call(args[:])
	return res[0].Bool()
}

// Sort sorts slice with lessFunc
// slice type must be slice
// lessFunc type must be func (a, b TypeOfSliceElement) bool
func Sort(slice interface{}, lessFunc interface{}) {
	aVal := reflect.ValueOf(slice)
	aType := aVal.Type()
	switch aType.Kind() {
	case reflect.Slice:
		break
	default:
		panic(fmt.Sprintf("parameter array must be slice, type: %s", aType))
	}
	eType := aType.Elem()

	fVal := reflect.ValueOf(lessFunc)
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
		panic(fmt.Sprintf("2 parameters of less func must be %s or implements %s", eType, eType))
	}

	if n := fType.NumOut(); n != 1 {
		panic(fmt.Sprintf("less func must have 1 output parameter, type: %s", fType))
	}
	if typ := fType.Out(0); typ.Kind() != reflect.Bool {
		panic(fmt.Sprintf("parameter of less func must be bool, type: %s", typ))
	}

	sort.Sort(sortReflect{
		Size:     aVal.Len(),
		Array:    aVal,
		LessFunc: fVal,
	})
}
