/*
 * Copyright (C) distroy
 */

package ldref

import "reflect"

func init() {
	registerCopyFunc(map[copyPair]copyFuncType{
		{To: reflect.Interface, From: reflect.Invalid}:   copyReflectToIfaceFromInvalid,
		{To: reflect.Interface, From: reflect.Interface}: copyReflectToIfaceFromStruct,
		{To: reflect.Interface, From: reflect.Ptr}:       copyReflectToIfaceFromPtr,

		{To: reflect.Interface, From: reflect.Struct}: copyReflectToIfaceFromStruct,
		{To: reflect.Interface, From: reflect.Map}:    copyReflectToIfaceFromStruct,
		{To: reflect.Interface, From: reflect.Slice}:  copyReflectToIfaceFromStruct,
		{To: reflect.Interface, From: reflect.Array}:  copyReflectToIfaceFromStruct,

		{To: reflect.Interface, From: reflect.Bool}:       copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.String}:     copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Float32}:    copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Float64}:    copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Complex64}:  copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Complex128}: copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Int}:        copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Int8}:       copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Int16}:      copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Int32}:      copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Int64}:      copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uint}:       copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uint8}:      copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uint16}:     copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uint32}:     copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uint64}:     copyReflectToIfaceFromOthers,
		{To: reflect.Interface, From: reflect.Uintptr}:    copyReflectToIfaceFromOthers,
	})
}

func copyReflectToIfaceFromInvalid(c *context, target, source reflect.Value) bool {
	target.Set(reflect.Zero(target.Type()))
	return true
}

func copyReflectToIfaceFromPtr(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()

	if !sTyp.Implements(tTyp) {
		return false
	}

	sVal := source
	if c.IsDeep && !sVal.IsNil() {
		sVal = deepClone(sVal)
	}

	target.Set(sVal)
	return true
}

func copyReflectToIfaceFromStruct(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()

	if !sTyp.Implements(tTyp) {
		return false
	}

	sVal := source
	if c.IsDeep {
		sVal = deepClone(sVal)
	}

	target.Set(sVal)
	return true
}

func copyReflectToIfaceFromOthers(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()

	if !sTyp.Implements(tTyp) {
		return false
	}

	target.Set(source)
	return true
}
