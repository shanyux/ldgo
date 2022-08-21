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
	target, isTargetPtr := prepareCopyTargetReflect(c, target)

	var copyFunc func(c *context, target, source reflect.Value) (end bool)
	switch target.Kind() {
	default:
		copyFunc = func(c *context, target, source reflect.Value) (end bool) {
			// c.AddErrorf("")
			return
		}

	case reflect.Interface:
		copyFunc = copyReflectToIface

	case reflect.Ptr:
		copyFunc = copyReflectToPtr

	case reflect.UnsafePointer:
		copyFunc = copyReflectToUnsafePointer

	case reflect.Func:
		copyFunc = copyReflectToFunc

	case reflect.Bool:
		copyFunc = copyReflectToBool

	case reflect.Complex64, reflect.Complex128:
		copyFunc = copyReflectToComplex

	case reflect.Float32, reflect.Float64:
		copyFunc = copyReflectToFloat

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		copyFunc = copyReflectToInt

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		copyFunc = copyReflectToUint

	case reflect.String:
		copyFunc = copyReflectToString

	case reflect.Struct:
		copyFunc = copyReflectToStruct

	case reflect.Slice:
		copyFunc = copyReflectToSlice

	case reflect.Array:
		copyFunc = copyReflectToArray

	case reflect.Map:
		copyFunc = copyReflectToMap
	}

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
	if fromType.ConvertibleTo(toType) {
		return true
	}

	switch toType.Kind() {
	case reflect.Complex64, reflect.Complex128:
		switch fromType.Kind() {
		default:
			return false

		case reflect.Bool:
		case reflect.Float32, reflect.Float64:
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		}
		return true

	case reflect.String:
		switch fromType.Kind() {
		default:
			return false

		case reflect.Func:
		case reflect.Bool:
		case reflect.Float32, reflect.Float64:
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		}
		return true
	}

	switch fromType.Kind() {
	case reflect.Complex64, reflect.Complex128:
		switch fromType.Kind() {
		default:
			return false

		case reflect.Bool:
		case reflect.Float32, reflect.Float64:
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		}
		return true

	case reflect.String:
		switch toType.Kind() {
		default:
			return false

		case reflect.Bool:
		case reflect.Float32, reflect.Float64:
		case reflect.Complex64, reflect.Complex128:
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		}
		return true
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

	case reflect.Map, reflect.Slice, reflect.Chan:
		fallthrough

	case reflect.Ptr, reflect.Func:
		target.SetPointer(unsafe.Pointer(source.Pointer()))

		// case reflect.Uintptr:
		// 	target.SetPointer(unsafe.Pointer(uintptr(source.Uint())))
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
