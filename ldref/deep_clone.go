/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"unsafe"
)

func DeepClone[T any](d T) T {
	var i interface{} = d
	if x, ok := i.(reflect.Value); ok {
		var r interface{} = deepClone(x)
		return r.(T)
	}

	x := deepClone(reflect.ValueOf(i))
	if x.Kind() == reflect.Invalid {
		var zero T
		return zero
	}

	return x.Interface().(T)
}

func deepClone(x0 reflect.Value) reflect.Value {
	if x0.Kind() == reflect.Interface {
		if x0.IsNil() {
			return x0
		}

		x0 = reflect.ValueOf(x0.Interface())
	}

	switch x0.Kind() {
	case reflect.Invalid:
		return x0

	case reflect.Struct:
		return deepCloneStruct(x0)

	case reflect.Ptr:
		return deepClonePtr(x0)

	case reflect.Array:
		return deepCloneArray(x0)

	case reflect.Slice:
		return deepCloneSlice(x0)

	case reflect.Map:
		return deepCloneMap(x0)
	}

	return x0
}

func deepClonePtr(x0 reflect.Value) reflect.Value {
	if x0.IsNil() {
		return x0
	}

	x1 := reflect.New(x0.Type().Elem())
	x1.Elem().Set(deepClone(x0.Elem()))
	return x1
}

func deepCloneStruct(x0 reflect.Value) reflect.Value {
	x1 := reflect.New(x0.Type()).Elem()
	if !x0.CanAddr() {
		x1.Set(x0)
		x0 = x1
	}

	for i, n := 0, x0.NumField(); i < n; i++ {
		f0 := x0.Field(i)
		f1 := x1.Field(i)
		// f1.Set(deepClone(f0))

		a0 := unsafe.Pointer(f0.UnsafeAddr())
		o0 := reflect.NewAt(f0.Type(), a0).Elem()

		a1 := unsafe.Pointer(f1.UnsafeAddr())
		o1 := reflect.NewAt(f1.Type(), a1).Elem()
		o1.Set(deepClone(o0))
	}
	return x1
}

func deepCloneArray(x0 reflect.Value) reflect.Value {
	l0 := x0.Len()
	x1 := reflect.New(reflect.ArrayOf(l0, x0.Type().Elem())).Elem()
	for i := 0; i < l0; i++ {
		v0 := x0.Index(i)
		v0 = deepClone(v0)
		x1.Index(i).Set(v0)
	}
	return x1
}

func deepCloneSlice(x0 reflect.Value) reflect.Value {
	if x0.IsNil() {
		return x0
	}
	l0 := x0.Len()
	x1 := reflect.MakeSlice(reflect.SliceOf(x0.Type().Elem()), l0, l0)
	for i := 0; i < l0; i++ {
		v0 := x0.Index(i)
		v0 = deepClone(v0)
		x1.Index(i).Set(v0)
	}
	return x1
}

func deepCloneMap(x0 reflect.Value) reflect.Value {
	x1 := reflect.MakeMap(x0.Type())
	for it := x0.MapRange(); it.Next(); {
		key := it.Key()
		val := it.Value()

		key = deepClone(key)
		val = deepClone(val)
		x1.SetMapIndex(key, val)
	}
	return x1
}
