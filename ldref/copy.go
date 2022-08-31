/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"

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

func getCopyReflectFunc(kind reflect.Kind) func(c *context, target, source reflect.Value) (end bool) {
	switch kind {
	case reflect.Interface:
		return copyReflectToIface
	case reflect.Ptr:
		return copyReflectToPtr
	case reflect.UnsafePointer:
		return copyReflectToUnsafePointer
	case reflect.Func:
		return copyReflectToFunc
	case reflect.Bool:
		return copyReflectToBool
	case reflect.Complex64, reflect.Complex128:
		return copyReflectToComplex
	case reflect.Float32, reflect.Float64:
		return copyReflectToFloat
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return copyReflectToInt
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return copyReflectToUint
	case reflect.String:
		return copyReflectToString
	case reflect.Struct:
		return copyReflectToStruct
	case reflect.Slice:
		return copyReflectToSlice
	case reflect.Array:
		return copyReflectToArray
	case reflect.Map:
		return copyReflectToMap
	}

	return func(c *context, target, source reflect.Value) (end bool) {
		// c.AddErrorf("")
		return
	}
}

func copyReflect(c *context, target, source reflect.Value) bool {
	_target := target
	target, isTargetPtr := prepareCopyTargetReflect(c, target)

	copyFunc := getCopyReflectFunc(target.Kind())
	if end := copyFunc(c, target, source); end {
		return true
	}

	// clear target
	if isTargetPtr {
		_target.Elem().Set(reflect.Zero(_target.Elem().Type()))
	} else {
		_target.Set(reflect.Zero(_target.Type()))
	}

	c.AddErrorf("%s can not convert to %s", typeNameOfReflect(source), typeNameOfReflect(target))
	return false
}

func isCopyTypeConvertible(toType, fromType reflect.Type) bool {
	toType, _ = indirectType(toType)
	if toType.Kind() == reflect.UnsafePointer && fromType.Kind() == reflect.Ptr {
		return true
	}

	fromType, _ = indirectType(fromType)

	if fromType.ConvertibleTo(toType) {
		return true
	}

	if isCopyTypeConvertibleOneWay(toType, fromType) {
		return true
	}

	if isCopyTypeConvertibleOneWay(fromType, toType) {
		return true
	}

	return false
}

func isCopyTypeConvertibleOneWay(toType, fromType reflect.Type) bool {
	switch toType.Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Complex64, reflect.Complex128:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fallthrough
	case reflect.String:
		switch fromType.Kind() {
		default:
			return false

		case reflect.Bool:
		case reflect.Float32, reflect.Float64:
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		case reflect.String:
		}
		return true

	case reflect.UnsafePointer:
		switch fromType.Kind() {
		default:
			return false

		case reflect.Func:
		}
	}

	return false
}

func copyReflectToIface(c *context, target, source reflect.Value) bool {
	// source, isSourcePtr := prepareCopySourceReflect(c, source)
	tTyp := target.Type()

	if source.Kind() == reflect.Invalid {
		target.Set(reflect.Zero(tTyp))
		return true
	}

	if !c.IsDeep {
		if source.Type().Implements(tTyp) {
			target.Set(source)
			return true
		}

		return false
	}

	source, isSourcePtr := indirectSourceReflect(source)
	if source.Kind() == reflect.Invalid {
		target.Set(reflect.Zero(tTyp))
		return true
	}

	// source is nil
	if source.Kind() == reflect.Ptr {
		if source.Type().Implements(tTyp) {
			target.Set(source)

		} else {
			c.AddErrorf("%s is nil, can not convert to %s", source.Type().String(), target.Type().String())
		}

		return true
	}

	if source.Type().Implements(target.Type()) {
		target.Set(newSourceForDeepCopy(c, source, isSourcePtr))
		return true

	} else if isSourcePtr && source.Addr().Type().Implements(target.Type()) {
		target.Set(newSourceForDeepCopy(c, source, isSourcePtr))
		return true
	}

	return false
}

func copyReflectToPtr(c *context, target, source reflect.Value) bool {
	if c.IsDeep {
		// unreachable code
		return false
	}

	tTyp := target.Type()
	if source.Kind() == reflect.Invalid {
		target.Set(reflect.Zero(target.Type()))
		return true
	}

	// _source := source
	for source.Kind() == reflect.Ptr || source.Kind() == reflect.Interface {
		if source.Kind() == reflect.Interface {
			source = reflect.ValueOf(source.Interface())
			continue

		} else if source.Type() == tTyp {
			target.Set(source)
			return true
		}

		source = source.Elem()
	}

	_target := target
	target, isTargetPtr := indirectTargetReflect(target)
	if end := copyReflect(c, target, source); end {
		return true
	}

	// clear target
	if isTargetPtr {
		_target.Elem().Set(reflect.Zero(_target.Elem().Type()))
	} else {
		_target.Set(reflect.Zero(_target.Type()))
	}

	return true
}

func copyReflectToFunc(c *context, target, source reflect.Value) bool {
	source, _ = indirectSourceReflect(source)
	if source.Type() == target.Type() {
		target.Set(source)
		return true
	}

	return false
}

func copyReflectToUnsafePointer(c *context, target, source reflect.Value) bool {
	switch source.Kind() {
	case reflect.UnsafePointer:
		target.Set(source)

	case reflect.Ptr, reflect.Func:
		target.SetPointer(unsafe.Pointer(source.Pointer()))
	}

	return false
}

func copyReflectToBool(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		return false

	case reflect.Invalid:
		target.SetBool(false)

	case reflect.Bool:
		b := source.Bool()
		target.SetBool(b)

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetBool(n != 0)

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		target.SetBool(n != 0)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetBool(n != 0)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetBool(n != 0)

	case reflect.String:
		s := source.String()
		if strings.EqualFold(s, "true") {
			target.SetBool(true)
			break

		} else if strings.EqualFold(s, "false") {
			target.SetBool(false)
			break
		}

		n, err := strconv.ParseFloat(s, 64)
		target.SetBool(n != 0)
		if err != nil {
			c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
		}
	}

	return true
}
