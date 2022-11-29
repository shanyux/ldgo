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
		{To: reflect.UnsafePointer, From: reflect.Invalid}:       copyReflectToUnsafePointerFromInvalid,
		{To: reflect.UnsafePointer, From: reflect.UnsafePointer}: copyReflectToUnsafePointerFromUnsafePointer,
		{To: reflect.UnsafePointer, From: reflect.Ptr}:           copyReflectToUnsafePointerFromPtr,
		{To: reflect.UnsafePointer, From: reflect.Func}:          copyReflectToUnsafePointerFromPtr,
	})
}

func copyReflectToUnsafePointerFromInvalid(c *context, target, source reflect.Value) bool {
	target.Set(reflect.Zero(target.Type()))
	return true
}

func copyReflectToUnsafePointerFromUnsafePointer(c *context, target, source reflect.Value) bool {
	target.Set(source)
	return true
}

func copyReflectToUnsafePointerFromPtr(c *context, target, source reflect.Value) bool {
	target.SetPointer(unsafe.Pointer(source.Pointer()))
	return true
}
