/*
 * Copyright (C) distroy
 */

package ldref

import (
	"math"
	"reflect"
	"strconv"
	"sync"

	"github.com/distroy/ldgo/lderr"
)

var (
	copyType sync.Map
)

func Copy(target, source interface{}) error {
	sVal := valueOf(source)
	// sVal, _ = valueElment(sVal)

	if tVal, ok := target.(reflect.Value); ok {
		return copyValue(tVal, sVal)
	}

	tVal := reflect.ValueOf(target)
	if tVal.Kind() != reflect.Ptr {
		return lderr.ErrReflectTargetNotPtr
	}

	if tVal.IsNil() {
		return lderr.ErrReflectTargetNilPtr
	}

	return copyValue(tVal, sVal)
}

func valueOf(v interface{}) reflect.Value {
	if vv, ok := v.(reflect.Value); ok {
		return vv
	}

	return reflect.ValueOf(v)
}

func valueElment(v reflect.Value) (reflect.Value, bool) {
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		return v.Elem(), true
	}
	return v, false
}

func typeElement(t reflect.Type) (reflect.Type, bool) {
	if t.Kind() == reflect.Ptr {
		return t.Elem(), true
	}

	return t, false
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

func copyValue(target, source reflect.Value) error {
	c := getContext()
	copyValueWithContext(c, target, source)
	return c.Error()
}

func copyValueWithContext(c *context, target, source reflect.Value) {
	if source.Kind() == reflect.Interface {
		source = reflect.ValueOf(source.Interface())
	}

	target, _ = valueElment(target)
	if isType(target.Type(), source.Type()) {
		target.Set(source)
		return
	}

	source, isSourcePtr := valueElment(source)
	if isSourcePtr && isType(target.Type(), source.Type()) {
		target.Set(source)
		return
	}

	switch target.Kind() {
	default:

	case reflect.Float32, reflect.Float64:
		copyValueToFloat(c, target, source)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		copyValueToInt(c, target, source)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		copyValueToUint(c, target, source)

	}
}

func copyValueToFloat(c *context, target, source reflect.Value) {
	switch source.Kind() {
	default:
		c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())

	case reflect.Bool:
		b := source.Bool()
		if b {
			target.SetFloat(1)
		} else {
			target.SetFloat(0)
		}

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetFloat(n)
		if target.OverflowFloat(n) {
			c.AddError("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetFloat(float64(n))
		if target.OverflowFloat(float64(n)) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetFloat(float64(n))
		if target.OverflowFloat(float64(n)) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseFloat(s, 64)
		target.SetFloat(n)
		if err != nil {
			c.AddError("can not convert to %s, %q", target.Type().String(), s)
		}
	}
}

func copyValueToInt(c *context, target, source reflect.Value) {
	switch source.Kind() {
	default:
		c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())

	case reflect.Bool:
		b := source.Bool()
		if b {
			target.SetInt(1)
		} else {
			target.SetInt(0)
		}

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetInt(int64(n))
		if target.OverflowInt(int64(n)) {
			c.AddError("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetInt(n)
		if target.OverflowInt(n) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetInt(int64(n))
		if n > math.MaxInt64 {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		} else if target.OverflowFloat(float64(n)) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseInt(s, 10, 64)
		target.SetInt(n)
		if err != nil {
			c.AddError("can not convert to %s, %q", target.Type().String(), s)
		}
	}
}

func copyValueToUint(c *context, target, source reflect.Value) {
	switch source.Kind() {
	default:
		c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())

	case reflect.Bool:
		b := source.Bool()
		if b {
			target.SetUint(1)
		} else {
			target.SetUint(0)
		}

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetUint(uint64(n))
		if target.OverflowInt(int64(n)) {
			c.AddError("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetUint(uint64(n))
		if n < 0 {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		} else if target.OverflowInt(n) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetUint(n)
		if target.OverflowFloat(float64(n)) {
			c.AddError("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		s := source.String()
		n, err := strconv.ParseUint(s, 10, 64)
		target.SetUint(n)
		if err != nil {
			c.AddError("can not convert to %s, %q", target.Type().String(), s)
		}
	}
}

func copyValueToString(c *context, target, source reflect.Value) {
	switch source.Kind() {
	default:
		c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		target.SetString(strconv.FormatFloat(n, 'f', -1, 64))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		target.SetString(strconv.FormatInt(n, 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		target.SetString(strconv.FormatUint(n, 10))

	case reflect.String:
		s := source.String()
		target.SetString(s)

	case reflect.Slice:
		switch ss := source.Interface().(type) {
		default:
			c.AddError("%s can not convert to %s", source.Type().String(), target.Type().String())

		case []byte:
			target.SetString(string(ss))

		case []rune:
			target.SetString(string(ss))
		}
	}
}
