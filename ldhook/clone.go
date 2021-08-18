/*
 * Copyright (C) distroy
 */

package ldhook

import "reflect"

func Clone(i interface{}) interface{} {
	x := cloneReflectValue(reflect.ValueOf(i))
	if isValueNil(x) {
		return nil
	}
	return x.Interface()
}

func DeepClone(i interface{}) interface{} {
	x := deepCloneValue(reflect.ValueOf(i))
	if isValueNil(x) {
		return nil
	}
	return x.Interface()
}

func isValueNil(x reflect.Value) bool {
	switch x.Kind() {
	case reflect.Invalid:
		return true

	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		fallthrough
	case reflect.Slice, reflect.Interface:
		return x.IsNil()
	}
	return false
}

func cloneReflectValue(x0 reflect.Value) reflect.Value {
	if isValueNil(x0) {
		return reflect.Value{}
	}
	switch x0.Kind() {
	case reflect.Struct:
		return x0

	case reflect.Ptr:
		x1 := reflect.New(x0.Type().Elem())
		x1.Elem().Set(x0.Elem())
		return x1

	case reflect.Array:
		l0 := x0.Len()
		x1 := reflect.New(reflect.ArrayOf(l0, x0.Type().Elem())).Elem()
		for i := 0; i < l0; i++ {
			v0 := x0.Index(i)
			if isValueNil(v0) {
				continue
			}
			x1.Index(i).Set(v0)
		}
		return x1

	case reflect.Slice:
		l0 := x0.Len()
		x1 := reflect.MakeSlice(reflect.SliceOf(x0.Type().Elem()), l0, l0)
		for i := 0; i < l0; i++ {
			v0 := x0.Index(i)
			if isValueNil(v0) {
				continue
			}
			x1.Index(i).Set(v0)
		}
		return x1

	case reflect.Map:
		x1 := reflect.MakeMap(x0.Type())
		for it := x0.MapRange(); it.Next(); {
			key := it.Key()
			val := it.Value()
			if isValueNil(key) || isValueNil(val) {
				continue
			}
			x1.SetMapIndex(key, val)
		}
		return x1
	}

	return x0
}

func deepCloneValue(x0 reflect.Value) reflect.Value {
	if isValueNil(x0) {
		return reflect.Value{}
	}
	switch x0.Kind() {
	case reflect.Struct:
		x1 := reflect.New(x0.Type()).Elem()
		for i, n := 0, x0.NumField(); i < n; i++ {
			f0 := x0.Field(i)
			if isValueNil(f0) {
				continue
			}
			f1 := x1.Field(i)
			f1.Set(deepCloneValue(f0))
		}
		return x1

	case reflect.Ptr:
		x1 := reflect.New(x0.Type().Elem())
		x1.Elem().Set(deepCloneValue(x0.Elem()))
		return x1

	case reflect.Array:
		l0 := x0.Len()
		x1 := reflect.New(reflect.ArrayOf(l0, x0.Type().Elem())).Elem()
		for i := 0; i < l0; i++ {
			v0 := x0.Index(i)
			if isValueNil(v0) {
				continue
			}
			v0 = deepCloneValue(v0)
			x1.Index(i).Set(v0)
		}
		return x1

	case reflect.Slice:
		l0 := x0.Len()
		x1 := reflect.MakeSlice(reflect.SliceOf(x0.Type().Elem()), l0, l0)
		for i := 0; i < l0; i++ {
			v0 := x0.Index(i)
			if isValueNil(v0) {
				continue
			}
			v0 = deepCloneValue(v0)
			x1.Index(i).Set(v0)
		}
		return x1

	case reflect.Map:
		x1 := reflect.MakeMap(x0.Type())
		for it := x0.MapRange(); it.Next(); {
			key := it.Key()
			val := it.Value()
			if isValueNil(key) || isValueNil(val) {
				continue
			}
			key = deepCloneValue(key)
			val = deepCloneValue(val)
			x1.SetMapIndex(key, val)
		}
		return x1
	}

	return x0
}
