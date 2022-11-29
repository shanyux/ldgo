/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"

	"github.com/distroy/ldgo/lderr"
)

func Copy(target, source interface{}) error {
	c := getContext(false)
	defer putContext(c)

	return copyWithContext(c, target, source)
}

func DeepCopy(target, source interface{}) error {
	c := getContext(true)
	defer putContext(c)

	return copyWithContext(c, target, source)
}

func copyWithContext(c *context, target, source interface{}) lderr.Error {
	sVal := valueOf(source)
	// sVal, _ = valueElment(sVal)

	if tVal, ok := target.(reflect.Value); ok {
		if !tVal.CanAddr() {
			if tVal.Kind() != reflect.Ptr {
				return lderr.ErrReflectTargetNotPtr

			} else if tVal.IsNil() {
				return lderr.ErrReflectTargetNilPtr
			}
		}

		copyReflect(c, tVal, sVal)
		return c.Error()
	}

	tVal := reflect.ValueOf(target)
	if tVal.Kind() != reflect.Ptr {
		return lderr.ErrReflectTargetNotPtr
	}
	if tVal.IsNil() {
		return lderr.ErrReflectTargetNilPtr
	}

	// tVal = tVal.Elem()

	copyReflect(c, tVal, sVal)
	return c.Error()
}

func valueOf(v interface{}) reflect.Value {
	if vv, ok := v.(reflect.Value); ok {
		return vv
	}

	return reflect.ValueOf(v)
}

func elementOfReflect(v reflect.Value) (reflect.Value, bool) {
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		return v.Elem(), true
	}
	return v, false
}

func indirectType(_type reflect.Type) (typ reflect.Type, isTypePtr bool) {
	typ = _type
	for typ.Kind() == reflect.Ptr {
		isTypePtr = true
		typ = typ.Elem()
	}
	return
}

func indirectTypeV2(_type reflect.Type) (typ reflect.Type, lvl int) {
	typ = _type
	lvl = 0
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		lvl++
	}
	return
}

func indirectCopySourceV2(_source reflect.Value) (source reflect.Value, lvl int) {
	source = _source
	lvl = 0
	for source.Kind() == reflect.Ptr && !source.IsNil() {
		source = source.Elem()
		lvl++
	}
	return
}

func indirectCopyTargetV2(_target reflect.Value) (target reflect.Value, lvl int) {
	target = _target
	lvl = 0
	for target.Kind() == reflect.Ptr {
		if target.IsNil() {
			target.Set(reflect.New(target.Type().Elem()))
		}

		target = target.Elem()
		lvl++
	}

	return
}

func indirectSourceReflect(_source reflect.Value) (source reflect.Value, isSourcePtr bool) {
	source = _source
	for source.Kind() == reflect.Interface || (source.Kind() == reflect.Ptr && !source.IsNil()) {
		if source.Kind() == reflect.Interface {
			source = reflect.ValueOf(source.Interface())
			continue
		}

		isSourcePtr = true
		source = source.Elem()
	}

	return
}

func indirectTargetReflect(_target reflect.Value) (target reflect.Value, isTargetPtr bool) {
	target = _target
	for target.Kind() == reflect.Ptr {
		isTargetPtr = true
		if target.IsNil() {
			target.Set(reflect.New(target.Type().Elem()))
		}

		target = target.Elem()
	}

	return
}

func newSourceForDeepCopy(c *context, value reflect.Value, isPtr bool) reflect.Value {
	if !isPtr {
		return value
	}

	cp := reflect.New(value.Type())
	if !c.IsDeep {
		cp.Elem().Set(value)
	} else {
		copyReflect(c, cp.Elem(), value)
	}
	return cp
}

func prepareCopyTargetReflect(c *context, _target reflect.Value) (target reflect.Value, isTargetPtr bool) {
	target = _target
	if !c.IsDeep {
		return elementOfReflect(target)
	}

	return indirectTargetReflect(target)
}

func prepareCopySourceReflect(c *context, _source reflect.Value) (source reflect.Value, isSourcePtr bool) {
	source = _source
	if !c.IsDeep {
		if source.Kind() == reflect.Interface {
			source = reflect.ValueOf(source.Interface())
		}
		return elementOfReflect(source)
	}

	return indirectSourceReflect(source)
}

func copyReflect(c *context, target, source reflect.Value) bool {
	_target := target
	_source := source

	if !target.CanAddr() {
		target = target.Elem()
	}
	// if !source.CanAddr() && source.Kind() == reflect.Ptr {
	// 	source = source.Elem()
	// }

	pair := copyPair{To: target.Kind(), From: source.Kind()}
	copyFunc := copyFuncMap[pair]
	if copyFunc == nil {
		if target.Kind() != reflect.Ptr && source.Kind() == reflect.Ptr {
			source, _ = indirectCopySourceV2(source)

		} else if source.Kind() == reflect.Interface {
			source = reflect.ValueOf(source.Interface())

		}

		pair := copyPair{To: target.Kind(), From: source.Kind()}
		copyFunc = copyFuncMap[pair]
	}

	if copyFunc != nil {
		end := copyFunc(c, target, source)
		if end {
			return true
		}
	}

	// clear target
	if !_target.CanAddr() {
		_target.Elem().Set(reflect.Zero(_target.Elem().Type()))
	} else {
		_target.Set(reflect.Zero(_target.Type()))
	}

	c.AddErrorf("%s can not copy to %s", typeNameOfReflect(_source), typeNameOfReflect(_target))
	return false
}

func isCopyTypeConvertible(toType, fromType reflect.Type) bool {
	toType, _ = indirectTypeV2(toType)
	fromType, _ = indirectTypeV2(fromType)

	pair := copyPair{To: toType.Kind(), From: fromType.Kind()}
	_, ok := copyFuncMap[pair]
	return ok
}
