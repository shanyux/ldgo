/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"strconv"
)

func init() {
	registerCopyFunc(map[copyPair]copyFuncType{
		{To: reflect.Array, From: reflect.Invalid}: copyReflectToArrayFromInvalid,
		{To: reflect.Array, From: reflect.String}:  copyReflectToArrayFromString,
		{To: reflect.Array, From: reflect.Slice}:   copyReflectToArrayFromSlice,
		{To: reflect.Array, From: reflect.Array}:   copyReflectToArrayFromArray,
		{To: reflect.Array, From: reflect.Map}:     copyReflectToArrayFromMapWithEmptyStructValue,
	})
}

func copyReflectToArrayFromInvalid(c *context, target, source reflect.Value) bool {
	target.Set(reflect.Zero(target.Type()))
	return true
}

func copyReflectToArrayFromString(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()

	sVal := source
	switch target.Type() {
	default:
		return false

	case typeOfByteSlice:
		sVal = reflect.ValueOf([]byte(sVal.String()))

	case typeOfRuneSlice:
		sVal = reflect.ValueOf([]rune(sVal.String()))
	}

	l := sVal.Len()
	if l > target.Len() {
		c.AddErrorf("%s has %d elements(%s), can not convert to %s",
			sTyp.String(), l, sTyp.Elem().String(), tTyp.String())
		l = target.Len()
	}
	for i := 0; i < l; i++ {
		tItem := target.Index(i)
		sItem := sVal.Index(i)
		tItem.Set(sItem)
	}

	return true
}

func copyReflectToArrayFromArray(c *context, target, source reflect.Value) bool {
	sVal := source.Slice(0, source.Len())
	return copyReflectToArrayFromSlice(c, target, sVal)
}

func copyReflectToArrayFromSlice(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()
	if !c.IsDeep && tTyp == sTyp {
		target.Set(source)
		return true
	}

	if source.Kind() != reflect.Slice || !isCopyTypeConvertible(tTyp.Elem(), sTyp.Elem()) {
		return false
	}

	l := source.Len()
	if l > target.Len() {
		c.AddErrorf("%s has %d elements, can not convert to %s", sTyp.String(), l, tTyp.String())
		l = target.Len()
	}

	for i := 0; i < l; i++ {
		tItem := target.Index(i)
		sItem := source.Index(i)

		c.PushField(strconv.Itoa(i))
		copyReflect(c, tItem, sItem)
		c.PopField()
	}

	return true
}

func copyReflectToArrayFromMapWithEmptyStructValue(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	sTyp := source.Type()

	if !isEmptyStruct(sTyp.Elem()) {
		return false
	}

	if isCopyTypeConvertible(sTyp.Elem(), sTyp.Key()) {
		return false
	}

	l := source.Len()
	if l > target.Len() {
		if target.Kind() == reflect.Array {
			c.AddErrorf("%s has %d elements, can not convert to %s", sTyp.String(), l, tTyp.String())
			l = target.Len()

		} else {
			target.Set(reflect.MakeSlice(tTyp, l, l))
		}
	}

	// sTyp.Comparable()
	i := 0
	for it := source.MapRange(); i < l && it.Next(); i++ {
		tItem := target.Index(i)
		sItem := it.Key()

		c.PushField(strconv.Itoa(i))
		copyReflect(c, tItem, sItem)
		c.PopField()
	}

	return true
}
