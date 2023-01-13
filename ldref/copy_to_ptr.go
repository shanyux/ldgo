/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"unsafe"
)

func init() {
	registerCopyFunc(map[copyPair]copyFuncType{
		{To: reflect.Ptr, From: reflect.Invalid}:       copyReflectToPtrFromInvalid,
		{To: reflect.Ptr, From: reflect.Ptr}:           copyReflectToPtrFromPtr,
		{To: reflect.Ptr, From: reflect.UnsafePointer}: copyReflectToPtrFromUnsafePointer,
		{To: reflect.Ptr, From: reflect.Interface}:     copyReflectToPtrFromIface,

		{To: reflect.Ptr, From: reflect.Func}:       copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Map}:        copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Slice}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Array}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Struct}:     copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Bool}:       copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.String}:     copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Float32}:    copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Float64}:    copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Complex64}:  copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Complex128}: copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Int}:        copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Int8}:       copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Int16}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Int32}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Int64}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uint}:       copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uint8}:      copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uint16}:     copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uint32}:     copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uint64}:     copyReflectToPtrFromOthers,
		{To: reflect.Ptr, From: reflect.Uintptr}:    copyReflectToPtrFromOthers,
	})
}

func copyReflectToPtrFromInvalid(c *copyContext, target, source reflect.Value) bool {
	target.Set(reflect.Zero(target.Type()))
	return true
}

func copyReflectToPtrFromPtr(c *copyContext, target, source reflect.Value) bool {
	sVal, sPtrLvl := indirectCopySource(source)
	if sVal.Kind() == reflect.Ptr {
		target.Set(reflect.Zero(target.Type()))
		return true
	}

	tTyp, tPtrLvl := indirectType(target.Type())

	if !c.Clone && tTyp == sVal.Type() {
		tVal := target
		sVal := source
		for i := 0; i+sPtrLvl < tPtrLvl; i++ {
			if tVal.IsNil() {
				tVal.Set(reflect.New(tVal.Type().Elem()))
			}
			tVal = tVal.Elem()
		}
		for i := 0; i+tPtrLvl < sPtrLvl; i++ {
			sVal = sVal.Elem()
		}

		tVal.Set(sVal)
		return true
	}

	tVal, _ := indirectCopyTarget(target)
	return copyReflect(c, tVal, sVal)
}

func copyReflectToPtrFromUnsafePointer(c *copyContext, target, source reflect.Value) bool {
	tAddr := unsafe.Pointer(target.UnsafeAddr())
	tTemp := reflect.NewAt(source.Type(), tAddr).Elem()
	tTemp.Set(source)
	return true
}

func copyReflectToPtrFromIface(c *copyContext, target, source reflect.Value) bool {
	sVal := reflect.ValueOf(source.Interface())
	tVal := target
	return copyReflect(c, tVal, sVal)
}

func copyReflectToPtrFromOthers(c *copyContext, target, source reflect.Value) bool {
	sVal := source
	tVal, _ := indirectCopyTarget(target)
	return copyReflect(c, tVal, sVal)
}
