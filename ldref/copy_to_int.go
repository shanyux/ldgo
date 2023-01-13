/*
 * Copyright (C) distroy
 */

package ldref

import (
	"math"
	"reflect"
	"strconv"
)

func init() {
	registerCopyFunc(map[copyPair]copyFuncType{
		{To: reflect.Int, From: reflect.Invalid}:    copyReflectToIntFromInvalid,
		{To: reflect.Int, From: reflect.Bool}:       copyReflectToIntFromBool,
		{To: reflect.Int, From: reflect.Complex64}:  copyReflectToIntFromComplex,
		{To: reflect.Int, From: reflect.Complex128}: copyReflectToIntFromComplex,
		{To: reflect.Int, From: reflect.Float32}:    copyReflectToIntFromFloat,
		{To: reflect.Int, From: reflect.Float64}:    copyReflectToIntFromFloat,
		{To: reflect.Int, From: reflect.Int}:        copyReflectToIntFromInt,
		{To: reflect.Int, From: reflect.Int8}:       copyReflectToIntFromInt,
		{To: reflect.Int, From: reflect.Int16}:      copyReflectToIntFromInt,
		{To: reflect.Int, From: reflect.Int32}:      copyReflectToIntFromInt,
		{To: reflect.Int, From: reflect.Int64}:      copyReflectToIntFromInt,
		{To: reflect.Int, From: reflect.Uint}:       copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.Uint8}:      copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.Uint16}:     copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.Uint32}:     copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.Uint64}:     copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.Uintptr}:    copyReflectToIntFromUint,
		{To: reflect.Int, From: reflect.String}:     copyReflectToIntFromString,

		{To: reflect.Int8, From: reflect.Invalid}:    copyReflectToIntFromInvalid,
		{To: reflect.Int8, From: reflect.Bool}:       copyReflectToIntFromBool,
		{To: reflect.Int8, From: reflect.Complex64}:  copyReflectToIntFromComplex,
		{To: reflect.Int8, From: reflect.Complex128}: copyReflectToIntFromComplex,
		{To: reflect.Int8, From: reflect.Float32}:    copyReflectToIntFromFloat,
		{To: reflect.Int8, From: reflect.Float64}:    copyReflectToIntFromFloat,
		{To: reflect.Int8, From: reflect.Int}:        copyReflectToIntFromInt,
		{To: reflect.Int8, From: reflect.Int8}:       copyReflectToIntFromInt,
		{To: reflect.Int8, From: reflect.Int16}:      copyReflectToIntFromInt,
		{To: reflect.Int8, From: reflect.Int32}:      copyReflectToIntFromInt,
		{To: reflect.Int8, From: reflect.Int64}:      copyReflectToIntFromInt,
		{To: reflect.Int8, From: reflect.Uint}:       copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.Uint8}:      copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.Uint16}:     copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.Uint32}:     copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.Uint64}:     copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.Uintptr}:    copyReflectToIntFromUint,
		{To: reflect.Int8, From: reflect.String}:     copyReflectToIntFromString,

		{To: reflect.Int16, From: reflect.Invalid}:    copyReflectToIntFromInvalid,
		{To: reflect.Int16, From: reflect.Bool}:       copyReflectToIntFromBool,
		{To: reflect.Int16, From: reflect.Complex64}:  copyReflectToIntFromComplex,
		{To: reflect.Int16, From: reflect.Complex128}: copyReflectToIntFromComplex,
		{To: reflect.Int16, From: reflect.Float32}:    copyReflectToIntFromFloat,
		{To: reflect.Int16, From: reflect.Float64}:    copyReflectToIntFromFloat,
		{To: reflect.Int16, From: reflect.Int}:        copyReflectToIntFromInt,
		{To: reflect.Int16, From: reflect.Int8}:       copyReflectToIntFromInt,
		{To: reflect.Int16, From: reflect.Int16}:      copyReflectToIntFromInt,
		{To: reflect.Int16, From: reflect.Int32}:      copyReflectToIntFromInt,
		{To: reflect.Int16, From: reflect.Int64}:      copyReflectToIntFromInt,
		{To: reflect.Int16, From: reflect.Uint}:       copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.Uint8}:      copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.Uint16}:     copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.Uint32}:     copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.Uint64}:     copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.Uintptr}:    copyReflectToIntFromUint,
		{To: reflect.Int16, From: reflect.String}:     copyReflectToIntFromString,

		{To: reflect.Int32, From: reflect.Invalid}:    copyReflectToIntFromInvalid,
		{To: reflect.Int32, From: reflect.Bool}:       copyReflectToIntFromBool,
		{To: reflect.Int32, From: reflect.Complex64}:  copyReflectToIntFromComplex,
		{To: reflect.Int32, From: reflect.Complex128}: copyReflectToIntFromComplex,
		{To: reflect.Int32, From: reflect.Float32}:    copyReflectToIntFromFloat,
		{To: reflect.Int32, From: reflect.Float64}:    copyReflectToIntFromFloat,
		{To: reflect.Int32, From: reflect.Int}:        copyReflectToIntFromInt,
		{To: reflect.Int32, From: reflect.Int8}:       copyReflectToIntFromInt,
		{To: reflect.Int32, From: reflect.Int16}:      copyReflectToIntFromInt,
		{To: reflect.Int32, From: reflect.Int32}:      copyReflectToIntFromInt,
		{To: reflect.Int32, From: reflect.Int64}:      copyReflectToIntFromInt,
		{To: reflect.Int32, From: reflect.Uint}:       copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.Uint8}:      copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.Uint16}:     copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.Uint32}:     copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.Uint64}:     copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.Uintptr}:    copyReflectToIntFromUint,
		{To: reflect.Int32, From: reflect.String}:     copyReflectToIntFromString,

		{To: reflect.Int64, From: reflect.Invalid}:    copyReflectToIntFromInvalid,
		{To: reflect.Int64, From: reflect.Bool}:       copyReflectToIntFromBool,
		{To: reflect.Int64, From: reflect.Complex64}:  copyReflectToIntFromComplex,
		{To: reflect.Int64, From: reflect.Complex128}: copyReflectToIntFromComplex,
		{To: reflect.Int64, From: reflect.Float32}:    copyReflectToIntFromFloat,
		{To: reflect.Int64, From: reflect.Float64}:    copyReflectToIntFromFloat,
		{To: reflect.Int64, From: reflect.Int}:        copyReflectToIntFromInt,
		{To: reflect.Int64, From: reflect.Int8}:       copyReflectToIntFromInt,
		{To: reflect.Int64, From: reflect.Int16}:      copyReflectToIntFromInt,
		{To: reflect.Int64, From: reflect.Int32}:      copyReflectToIntFromInt,
		{To: reflect.Int64, From: reflect.Int64}:      copyReflectToIntFromInt,
		{To: reflect.Int64, From: reflect.Uint}:       copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.Uint8}:      copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.Uint16}:     copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.Uint32}:     copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.Uint64}:     copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.Uintptr}:    copyReflectToIntFromUint,
		{To: reflect.Int64, From: reflect.String}:     copyReflectToIntFromString,
	})
}

func copyReflectToIntFromInvalid(c *copyContext, target, source reflect.Value) bool {
	target.SetInt(0)
	return true
}

func copyReflectToIntFromBool(c *copyContext, target, source reflect.Value) bool {
	b := source.Bool()
	target.SetInt(int64(bool2int(b)))
	return true
}

func copyReflectToIntFromFloat(c *copyContext, target, source reflect.Value) bool {
	n := source.Float()
	target.SetInt(int64(n))
	if target.OverflowInt(int64(n)) {
		c.AddErrorf("%s(%f) overflow", target.Type().String(), n)
	}
	return true
}

func copyReflectToIntFromComplex(c *copyContext, target, source reflect.Value) bool {
	n := source.Complex()
	r := real(n)
	target.SetInt(int64(r))
	if r > math.MaxInt64 || target.OverflowInt(int64(r)) {
		c.AddErrorf("%s(%v) overflow", target.Type().String(), n)
	}
	return true
}

func copyReflectToIntFromInt(c *copyContext, target, source reflect.Value) bool {
	n := source.Int()
	target.SetInt(n)
	if target.OverflowInt(n) {
		c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
	}

	return true
}

func copyReflectToIntFromUint(c *copyContext, target, source reflect.Value) bool {
	n := source.Uint()
	target.SetInt(int64(n))
	if n > math.MaxInt64 || target.OverflowInt(int64(n)) {
		c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
	}
	return true
}

func copyReflectToIntFromString(c *copyContext, target, source reflect.Value) bool {
	s := source.String()
	n, err := strconv.ParseInt(s, 10, 64)
	target.SetInt(n)
	if err != nil {
		c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)

	} else if target.OverflowInt(n) {
		c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
	}
	return true
}
