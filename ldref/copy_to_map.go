/*
 * Copyright (C) distroy
 */

package ldref

import (
	"fmt"
	"reflect"
)

func copyReflectToMap(c *context, target, source reflect.Value) bool {
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		return false

	case reflect.Invalid:
		target.Set(reflect.Zero(target.Type()))

	case reflect.Struct:
		return copyReflectToMapFromStruct(c, target, source)

	case reflect.Map:
		return copyReflectToMapFromMap(c, target, source)

	case reflect.Array:
		source = source.Slice(0, source.Len())
		fallthrough

	case reflect.Slice:
		return copyReflectToMapFromSlice(c, target, source)
	}

	return true
}

func isStructFieldNilPtr(v reflect.Value) bool {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}

		v = v.Elem()
	}

	return false
}

func copyReflectToMapFromStruct(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	keyTyp := tTyp.Key()
	valTyp := tTyp.Elem()

	if keyTyp.Kind() != reflect.String && keyTyp.Kind() != reflect.Interface {
		return false
	}

	// if target.IsNil() {
	target.Set(reflect.MakeMap(target.Type()))
	// }

	sInfo := getCopyTypeInfo(source.Type())
	for _, sFieldInfo := range sInfo.Fields {
		sField := source.Field(sFieldInfo.Index)
		if isStructFieldNilPtr(sField) {
			continue
		}

		key := reflect.ValueOf(sFieldInfo.Name)
		value := reflect.New(valTyp).Elem()

		c.PushField(sFieldInfo.Name)
		copyReflect(c, value, sField)
		c.PopField()

		target.SetMapIndex(key, value)
	}

	return true
}

func copyReflectToMapFromMap(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	tKeyTyp := tTyp.Key()
	tValTyp := tTyp.Elem()

	sTyp := source.Type()
	sKeyTyp := sTyp.Key()
	sValTyp := sTyp.Elem()

	if !c.IsDeep && tTyp == sTyp {
		target.Set(source)
		return true
	}

	if !isCopyTypeConvertible(tKeyTyp, sKeyTyp) {
		return false
	}

	if !isCopyTypeConvertible(tValTyp, sValTyp) {
		return false
	}

	// if target.IsNil() {
	target.Set(reflect.MakeMap(target.Type()))
	// }

	for sIt := source.MapRange(); sIt.Next(); {
		sKey := sIt.Key()
		sVal := sIt.Value()

		tKey := reflect.New(tKeyTyp).Elem()
		tVal := reflect.New(tValTyp).Elem()

		c.PushField(fmt.Sprintf("%v(key)", sKey.Interface()))
		keyEnd := copyReflect(c, tKey, sKey)
		c.PopField()
		if !keyEnd {
			continue
		}

		c.PushField(fmt.Sprintf("%v(val)", sKey.Interface()))
		valEnd := copyReflect(c, tVal, sVal)
		c.PopField()
		if !valEnd {
			continue
		}

		target.SetMapIndex(tKey, tVal)
	}

	return true
}

func copyReflectToMapFromSlice(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	tKeyTyp := tTyp.Key()
	tValTyp := tTyp.Elem()

	if tValTyp == typeOfEmptyStruct {
		return copyReflectToMapFromSliceWithEmptyStructValue(c, target, source)
	}

	sTyp := source.Type()
	sKeyTyp := reflect.TypeOf(int(0))
	sValTyp := sTyp.Elem()

	if !isCopyTypeConvertible(tKeyTyp, sKeyTyp) {
		return false
	}

	if !isCopyTypeConvertible(tValTyp, sValTyp) {
		return false
	}

	// if target.IsNil() {
	target.Set(reflect.MakeMap(target.Type()))
	// }

	for i, l := 0, source.Len(); i < l; i++ {
		sKey := reflect.ValueOf(i)
		sVal := source.Index(i)

		tKey := reflect.New(tKeyTyp).Elem()
		tVal := reflect.New(tValTyp).Elem()

		c.PushField(fmt.Sprintf("%d(key)", i))
		keyEnd := copyReflect(c, tKey, sKey)
		c.PopField()
		if !keyEnd {
			continue
		}

		c.PushField(fmt.Sprintf("%d(val)", i))
		valEnd := copyReflect(c, tVal, sVal)
		c.PopField()
		if !valEnd {
			continue
		}

		target.SetMapIndex(tKey, tVal)
	}

	return true
}

func copyReflectToMapFromSliceWithEmptyStructValue(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	tElemTyp := tTyp.Key()

	sTyp := source.Type()
	sElemTyp := sTyp.Elem()

	if !isCopyTypeConvertible(tElemTyp, sElemTyp) {
		return false
	}

	target.Set(reflect.MakeMap(target.Type()))

	tValue := reflect.ValueOf(struct{}{})
	for i, l := 0, source.Len(); i < l; i++ {
		// sKey := reflect.ValueOf(i)
		sElem := source.Index(i)

		tElem := reflect.New(tElemTyp).Elem()

		c.PushField(fmt.Sprintf("%d", i))
		keyEnd := copyReflect(c, tElem, sElem)
		c.PopField()
		if !keyEnd {
			continue
		}

		target.SetMapIndex(tElem, tValue)
	}

	return true
}
