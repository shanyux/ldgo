/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"unsafe"

	"github.com/distroy/ldgo/lderr"
)

type any = interface{}

func Merge(target, source any) lderr.Error {
	c := &mergeContext{
		IsDeep: false,
	}
	return mergeWithContext(c, target, source)
}

func DeepMerge(target, source any) lderr.Error {
	c := &mergeContext{
		IsDeep: true,
	}
	return mergeWithContext(c, target, source)
}

type mergeContext struct {
	IsDeep           bool
	IsMergeSliceElme bool
}

func mergeWithContext(c *mergeContext, target, source any) lderr.Error {
	tVal := valueOf(target)
	sVal := valueOf(source)

	tTyp := tVal.Type()
	sTyp := sVal.Type()

	if tTyp.Kind() != reflect.Ptr {
		return lderr.ErrReflectTargetNotPtr
	}

	if tVal.IsNil() {
		return lderr.ErrReflectTargetNilPtr
	}

	tElemType := tTyp.Elem()
	switch {
	default:
		return lderr.ErrReflectTypeNotEqual

	case tTyp == sTyp ||
		(tElemType.Kind() == reflect.Interface && sTyp.Kind() == reflect.Ptr && sTyp.Elem().Implements(tElemType)):

		if sVal.IsNil() {
			// do not need to merge
			return nil
		}

		tVal = tVal.Elem()
		sVal = sVal.Elem()

	case tElemType == sTyp ||
		(tElemType.Kind() == reflect.Interface && sTyp.Implements(tElemType)):
		tVal = tVal.Elem()
	}

	mergeReflect(c, tVal, sVal)
	return nil
}

func cloneForMerge(c *mergeContext, x reflect.Value) reflect.Value {
	v := x
	if c.IsDeep {
		v = deepClone(v)
	}
	return v
}

func mergeReflect(c *mergeContext, target, source reflect.Value) {
	switch target.Kind() {
	default:
		mergeReflectNormal(c, target, source)

	case reflect.Invalid:
		break

	case reflect.Ptr, reflect.Func, reflect.Chan, reflect.Interface:
		mergeReflectPtr(c, target, source)

	case reflect.Map:
		mergeReflectMap(c, target, source)

	case reflect.Slice:
		mergeReflectSlice(c, target, source)

	case reflect.Array:
		mergeReflectArray(c, target, source)

	case reflect.Struct:
		mergeReflectStruct(c, target, source)
	}
}

func mergeReflectPtr(c *mergeContext, target, source reflect.Value) {
	if target.IsNil() {
		source = cloneForMerge(c, source)
		target.Set(source)
	}
}

func mergeReflectMap(c *mergeContext, target, source reflect.Value) {
	if target.IsNil() {
		source = cloneForMerge(c, source)
		target.Set(source)
		return
	}

	for it := source.MapRange(); it.Next(); {
		key := it.Key()
		sVal := it.Value()
		if !sVal.IsValid() {
			continue
		}

		tVal := target.MapIndex(key)
		if !tVal.IsValid() {
			sVal = cloneForMerge(c, sVal)
			target.SetMapIndex(key, sVal)
			continue
		}

		mergeReflect(c, tVal, sVal)
	}
}

func mergeReflectSlice(c *mergeContext, target, source reflect.Value) {
	if target.IsNil() {
		source = cloneForMerge(c, source)
		target.Set(source)
		return
	}

	if !c.IsMergeSliceElme {
		return
	}

	tLen := target.Len()
	sLen := source.Len()

	if tLen == 0 {
		source = cloneForMerge(c, source)
		target.Set(source)
		return
	}

	resizeSliceReflect(target, sLen)

	for i := 0; i < sLen; i++ {
		tVal := target.Index(i)
		sVal := source.Index(i)

		if i < tLen {
			mergeReflect(c, tVal, sVal)
			continue
		}

		sVal = cloneForMerge(c, sVal)
		tVal.Set(sVal)
	}
}

func mergeReflectArray(c *mergeContext, target, source reflect.Value) {
	// if target.IsNil() {
	// 	source = cloneForMerge(c, source)
	// 	target.Set(source)
	// 	return
	// }
	//
	// if !c.IsMergeSliceElme {
	// 	return
	// }

	len := source.Len()

	for i := 0; i < len; i++ {
		tVal := target.Index(i)
		sVal := source.Index(i)

		mergeReflect(c, tVal, sVal)
	}
}

func mergeReflectStruct(c *mergeContext, target, source reflect.Value) {
	n := target.NumField()
	for i := 0; i < n; i++ {

		tField := target.Field(i)
		sField := source.Field(i)

		tFieldAddr := unsafe.Pointer(tField.UnsafeAddr())
		tField = reflect.NewAt(tField.Type(), tFieldAddr).Elem()

		mergeReflect(c, tField, sField)
	}
}

func mergeReflectNormal(c *mergeContext, target, source reflect.Value) {
	if target.IsZero() {
		target.Set(source)
	}
}
