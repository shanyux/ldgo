/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"sync"
)

func Clone[T any](d T) T {
	var i interface{} = d
	if x, ok := i.(reflect.Value); ok {
		var r interface{} = clone(x)
		return r.(T)
	}

	x := clone(reflect.ValueOf(i))
	if x.Kind() == reflect.Invalid {
		var zero T
		return zero
	}

	return x.Interface().(T)
}

func clone(x0 reflect.Value) reflect.Value {
	if x0.Kind() == reflect.Interface {
		if x0.IsNil() {
			return x0
		}

		x0 = reflect.ValueOf(x0.Interface())
	}

	switch x0.Kind() {
	case reflect.Invalid:
		return reflect.Value{}

	case reflect.Struct:
		return x0

	case reflect.Ptr:
		return clonePtr(x0)

	case reflect.Array:
		return cloneArray(x0)

	case reflect.Slice:
		return cloneSlice(x0)

	case reflect.Map:
		return cloneMap(x0)
	}

	return x0
}

func clonePtr(x0 reflect.Value) reflect.Value {
	if x0.IsNil() {
		return x0
	}

	x1 := reflect.New(x0.Type().Elem())
	if isSyncType(x0.Interface()) {
		return x1
	}

	x1.Elem().Set(x0.Elem())
	return x1
}

func isSyncType(i interface{}) bool {
	switch i.(type) {
	case *sync.Mutex, *sync.RWMutex, *sync.Cond:
		return true
	}
	return false
}

func cloneArray(x0 reflect.Value) reflect.Value {
	l0 := x0.Len()
	x1 := reflect.New(reflect.ArrayOf(l0, x0.Type().Elem())).Elem()
	for i := 0; i < l0; i++ {
		v0 := x0.Index(i)
		x1.Index(i).Set(v0)
	}
	return x1
}

func cloneSlice(x0 reflect.Value) reflect.Value {
	l0 := x0.Len()
	x1 := reflect.MakeSlice(reflect.SliceOf(x0.Type().Elem()), 0, l0)
	for i := 0; i < l0; i++ {
		v0 := x0.Index(i)
		x1 = reflect.Append(x1, v0)
		// x1.Index(i).Set(v0)
	}
	return x1
}

func cloneMap(x0 reflect.Value) reflect.Value {
	x1 := reflect.MakeMap(x0.Type())
	for it := x0.MapRange(); it.Next(); {
		key := it.Key()
		val := it.Value()
		x1.SetMapIndex(key, val)
	}
	return x1
}
