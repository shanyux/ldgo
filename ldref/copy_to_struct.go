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
		// {To: reflect.Struct, From: reflect.Invalid}: copyReflectToStructFromInvalid,
		{To: reflect.Struct, From: reflect.Struct}: copyReflectToStructFromStruct,
		{To: reflect.Struct, From: reflect.Map}:    copyReflectToStructFromMap,
	})
}

func clearCopyStructIgnoreField(c *copyContext, v reflect.Value, info *copyStructInfo) {
	for _, f := range info.Ignores {
		field := v.Field(f.Index)
		field.Set(reflect.Zero(f.Type))
	}
}

func copyReflectToStructFromStruct(c *copyContext, target, source reflect.Value) bool {
	tTyp := target.Type()
	tInfo := getCopyTypeInfo(tTyp)

	sTyp := source.Type()
	sInfo := getCopyTypeInfo(sTyp)
	if !c.Clone && tTyp == sTyp {
		target.Set(source)
		clearCopyStructIgnoreField(c, target, tInfo)
		return true
	}

	for _, sFieldInfo := range sInfo.Fields {
		tFieldInfo := tInfo.Fields[sFieldInfo.Name]
		if tFieldInfo == nil {
			continue
		}

		tFieldAddr := unsafe.Pointer(target.Field(tFieldInfo.Index).UnsafeAddr())
		tField := reflect.NewAt(tFieldInfo.Type, tFieldAddr).Elem()
		sField := source.Field(sFieldInfo.Index)

		c.PushField(tFieldInfo.Name)
		copyReflect(c, tField, sField)
		c.PopField()
	}

	return true
}

func copyReflectToStructFromMap(c *copyContext, target, source reflect.Value) bool {
	tTyp := target.Type()
	tInfo := getCopyTypeInfo(tTyp)

	sTyp := source.Type()
	if sTyp.Key().Kind() != reflect.String {
		return false
	}

	it := source.MapRange()
	for it.Next() {
		key := it.Key().String()
		tFieldInfo := tInfo.Fields[key]
		if tFieldInfo == nil {
			continue
		}

		tField := target.Field(tFieldInfo.Index)
		value := it.Value()

		c.PushField(tFieldInfo.Name)
		copyReflect(c, tField, value)
		c.PopField()
	}
	return true
}
