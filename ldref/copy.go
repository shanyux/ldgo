/*
 * Copyright (C) distroy
 */

package ldref

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
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

func elementOfValue(v reflect.Value) (reflect.Value, bool) {
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		return v.Elem(), true
	}
	return v, false
}

func isType(target, source reflect.Type) bool {
	if target == source {
		return true
	}

	if target.Kind() == reflect.Interface && source.Implements(target) {
		return true
	}

	return false
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
	if !c.IsDeep {
		return elementOfValue(target)
	}

	return indirectTargetReflect(target)
}

func prepareCopySourceReflect(c *context, _source reflect.Value) (source reflect.Value, isSourcePtr bool) {
	if !c.IsDeep {
		if source.Kind() == reflect.Interface {
			source = reflect.ValueOf(source.Interface())
		}
		return elementOfValue(source)
	}

	return indirectSourceReflect(source)
}

func copyReflect(c *context, target, source reflect.Value) {
	_target := target
	target, isTargetPtr := prepareCopyTargetReflect(c, target)

	var copyFunc func(c *context, target, source reflect.Value) (end bool)
	switch target.Kind() {
	default:
		// c.AddErrorf("")

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

	case reflect.Float32, reflect.Float64:
		copyFunc = copyReflectToFloat

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		copyFunc = copyReflectToInt

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		copyFunc = copyReflectToUint

	case reflect.Struct:
		copyFunc = copyReflectToStruct

	case reflect.Slice:
		copyFunc = copyReflectToSlice

	case reflect.Array:
		copyFunc = copyReflectToArray

	case reflect.Map:
		copyFunc = copyReflectToMap
	}

	if copyFunc != nil {
		end := copyFunc(c, target, source)
		if end {
			return
		}
	}

	// clear target
	if isTargetPtr {
		_target.Elem().Set(reflect.ValueOf(nil))
	} else {
		_target.Set(reflect.Zero(_target.Type()))
	}

	c.AddErrorf("%s can not convert to %s", source.Type().String(), target.Type().String())
}

func copyReflectToIface(c *context, target, source reflect.Value) bool {
	// source, isSourcePtr := prepareCopySourceReflect(c, source)
	tType := target.Type()

	if source.Kind() == reflect.Invalid {
		target.Set(reflect.ValueOf(nil))
		return true
	}

	if !c.IsDeep {
		if source.Type().Implements(tType) {
			target.Set(source)
			return true
		}

		return false
	}

	source, isSourcePtr := prepareCopySourceReflect(c, source)
	if source.Kind() == reflect.Invalid {
		target.Set(reflect.ValueOf(nil))
		return true
	}

	// source is nil
	if source.Kind() == reflect.Ptr {
		if source.Type().Implements(tType) {
			target.Set(source)
		}

		c.AddErrorf("%s is nil, can not convert to %s", source.Type().String(), target.Type().String())
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

	return false
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

		} else if strings.EqualFold(s, "true") {
			target.SetBool(false)
		}

		n, err := strconv.ParseFloat(s, 64)
		target.SetBool(n != 0)
		if err != nil {
			c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
		}
	}

	return true
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func copyReflectToFloat(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		// c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())
		return false

	case reflect.Invalid:
		target.SetFloat(0)

	case reflect.Bool:
		b := source.Bool()
		target.SetFloat(float64(bool2int(b)))

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetFloat(n)
		if target.OverflowFloat(n) {
			c.AddErrorf("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		r := real(n)
		target.SetFloat(r)
		if target.OverflowFloat(r) {
			c.AddErrorf("%s(%v) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetFloat(float64(n))
		if target.OverflowFloat(float64(n)) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetFloat(float64(n))
		if target.OverflowFloat(float64(n)) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseFloat(s, 64)
		target.SetFloat(n)
		if err != nil {
			c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
		}
	}

	return true
}

func copyReflectToInt(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		// c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())
		return false

	case reflect.Invalid:
		target.SetInt(0)

	case reflect.Bool:
		b := source.Bool()
		target.SetInt(int64(bool2int(b)))

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetInt(int64(n))
		if target.OverflowInt(int64(n)) {
			c.AddErrorf("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		r := real(n)
		target.SetInt(int64(r))
		if r > math.MaxInt64 || target.OverflowInt(int64(r)) {
			c.AddErrorf("%s(%v) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetInt(n)
		if target.OverflowInt(n) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetInt(int64(n))
		if n > math.MaxInt64 || target.OverflowInt(int64(n)) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseInt(s, 10, 64)
		target.SetInt(n)
		if err != nil {
			c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
		}
	}

	return true
}

func copyReflectToUint(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		// c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())
		return false

	case reflect.Invalid:
		target.SetUint(0)

	case reflect.Bool:
		b := source.Bool()
		target.SetUint(uint64(bool2int(b)))

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetUint(uint64(n))
		if target.OverflowInt(int64(n)) {
			c.AddErrorf("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		r := real(n)
		target.SetUint(uint64(r))
		if r > math.MaxUint64 || r < 0 || target.OverflowInt(int64(r)) {
			c.AddErrorf("%s(%v) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetUint(uint64(n))
		if n < 0 || target.OverflowUint(uint64(n)) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetUint(n)
		if target.OverflowUint(n) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseUint(s, 10, 64)
		target.SetUint(n)
		if err != nil {
			c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
		}
	}

	return true
}

func copyReflectToString(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		return false

	case reflect.Invalid:
		target.SetString("")

	case reflect.Bool:
		b := source.Bool()
		if b {
			target.SetString("true")
		} else {
			target.SetString("false")
		}

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetString(strconv.FormatFloat(n, 'f', -1, 64))

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		target.SetString(strconv.FormatComplex(n, 'f', -1, 128))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetString(strconv.FormatInt(n, 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetString(strconv.FormatUint(n, 10))

	case reflect.String:
		s := source.String()
		target.SetString(s)

	case reflect.Array:
		source = source.Slice(0, source.Len())
		fallthrough

	case reflect.Slice:
		switch ss := source.Interface().(type) {
		default:
			return false

		case []byte:
			target.SetString(string(ss))

		case []rune:
			target.SetString(string(ss))
		}

	case reflect.Func:
		funcName := runtime.FuncForPC(source.Pointer()).Name()
		target.SetString(funcName)
	}

	return true
}

func copyReflectToStruct(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	tTyp := target.Type()
	sTyp := source.Type()
	if !c.IsDeep && tTyp == sTyp {
		target.Set(source)
		return true
	}

	tInfo := getCopyTypeInfo(target.Type())
	switch source.Kind() {
	default:
		return false

	case reflect.Struct:
		sInfo := getCopyTypeInfo(source.Type())
		for _, sFieldInfo := range sInfo.Fields {
			tFieldInfo := tInfo.Fields[sFieldInfo.Name]
			if tFieldInfo == nil {
				continue
			}

			tField := target.Field(tFieldInfo.Index)
			sField := source.Field(sFieldInfo.Index)

			c.PushField(tFieldInfo.Name)
			copyReflect(c, tField, sField)
			c.PopField()
		}

	case reflect.Map:
		if source.Type().Key().Kind() != reflect.String {
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
	}

	return true
}

func copyReflectToArray(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	tTyp := target.Type()
	sTyp := source.Type()
	if !c.IsDeep && tTyp == sTyp {
		target.Set(source)
		return true
	}

	if source.Kind() == reflect.Invalid {
		target.Set(source)
		return true
	}

	if source.Kind() != reflect.Array {
		source = source.Slice(0, source.Len())
	}

	if source.Kind() != reflect.Slice || !sTyp.Elem().ConvertibleTo(tTyp.Elem()) {
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

func copyReflectToSlice(c *context, target, source reflect.Value) bool {
	// source, _ = prepareCopySourceReflect(c, source)
	source, _ = indirectSourceReflect(source)

	tTyp := target.Type()
	sTyp := source.Type()
	if !c.IsDeep && tTyp == sTyp {
		target.Set(source)
		return true
	}

	if source.Kind() == reflect.Invalid {
		target.Set(source)
		return true
	}

	if source.Kind() != reflect.Array {
		source = source.Slice(0, source.Len())
	}

	if source.Kind() != reflect.Slice || !sTyp.Elem().ConvertibleTo(tTyp.Elem()) {
		return false
	}

	l := source.Len()
	if l > target.Len() {
		target.Set(reflect.MakeSlice(tTyp.Elem(), l, l))
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

func isMapElemTypeConvertible(toType, fromType reflect.Type) bool {
	if fromType.ConvertibleTo(toType) {
		return true
	}

	switch toType.Kind() {
	case reflect.String:
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
	}

	return false
}

func copyReflectToMap(c *context, target, source reflect.Value) bool {
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		return false

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
		key := reflect.ValueOf(sFieldInfo.Name)
		value := reflect.Zero(valTyp)
		sField := source.Field(sFieldInfo.Index)

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

	if !isMapElemTypeConvertible(tKeyTyp, sKeyTyp) {
		return false
	}

	if !isMapElemTypeConvertible(tValTyp, sValTyp) {
		return false
	}

	// if target.IsNil() {
	target.Set(reflect.MakeMap(target.Type()))
	// }

	for sIt := source.MapRange(); sIt.Next(); {
		sKey := sIt.Key()
		sVal := sIt.Value()

		tKey := reflect.Zero(tKeyTyp)
		tVal := reflect.Zero(tValTyp)

		c.PushField(fmt.Sprintf("%v(key)", sKey.Interface()))
		copyReflect(c, tKey, sKey)
		c.PopField()

		c.PushField(fmt.Sprintf("%v(val)", sKey.Interface()))
		copyReflect(c, tVal, sVal)
		c.PopField()

		target.SetMapIndex(sKey, tVal)
	}

	return true
}

func copyReflectToMapFromSlice(c *context, target, source reflect.Value) bool {
	tTyp := target.Type()
	tKeyTyp := tTyp.Key()
	tValTyp := tTyp.Elem()

	sTyp := source.Type()
	sKeyTyp := reflect.TypeOf(int(0))
	sValTyp := sTyp.Elem()

	if !isMapElemTypeConvertible(tKeyTyp, sKeyTyp) {
		return false
	}

	if !isMapElemTypeConvertible(tValTyp, sValTyp) {
		return false
	}

	// if target.IsNil() {
	target.Set(reflect.MakeMap(target.Type()))
	// }

	for i := 0; i < source.Len(); i++ {
		sKey := reflect.ValueOf(i)
		sVal := source.Index(i)

		tKey := reflect.Zero(tKeyTyp)
		tVal := reflect.Zero(tValTyp)

		c.PushField(fmt.Sprintf("%d(key)", i))
		copyReflect(c, tKey, sKey)
		c.PopField()

		c.PushField(fmt.Sprintf("%d(val)", i))
		copyReflect(c, tVal, sVal)
		c.PopField()

		target.SetMapIndex(sKey, tVal)
	}

	return true
}
