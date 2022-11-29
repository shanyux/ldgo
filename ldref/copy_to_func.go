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
		{To: reflect.Func, From: reflect.Invalid}: copyReflectToFuncFromInvalid,
		{To: reflect.Func, From: reflect.Func}:    copyReflectToFuncFromFunc,
		// {To: reflect.Func, From: reflect.UnsafePointer}: copyReflectToFuncFromUnsafePointer,
	})
}

func copyReflectToFuncFromInvalid(c *context, target, source reflect.Value) bool {
	target.Set(reflect.Zero(target.Type()))
	return true
}

func copyReflectToFuncFromFunc(c *context, target, source reflect.Value) bool {
	if target.Type() != source.Type() {
		return false
	}

	target.Set(source)
	return true
}

func copyReflectToFuncFromUnsafePointer(c *context, target, source reflect.Value) bool {
	tAddr := unsafe.Pointer(target.UnsafeAddr())
	tTemp := reflect.NewAt(source.Type(), tAddr).Elem()
	tTemp.Set(source)
	return true
}
