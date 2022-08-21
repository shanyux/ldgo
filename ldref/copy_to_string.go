/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"runtime"
	"strconv"
)

func typeNameOfReflect(v reflect.Value) string {
	if v.Kind() == reflect.Invalid {
		return "nil"
	}

	return v.Type().String()
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

	// case reflect.Complex64, reflect.Complex128:
	// 	n := source.Complex()
	// 	target.SetString(strconv.FormatComplex(n, 'f', -1, 128))

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
