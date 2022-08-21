/*
 * Copyright (C) distroy
 */

package ldref

import (
	"math"
	"reflect"
	"strconv"
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func copyReflectToComplex(c *context, target, source reflect.Value) bool {
	source, _ = indirectSourceReflect(source)

	switch source.Kind() {
	default:
		return false

	case reflect.Invalid:
		target.SetComplex(0)

	case reflect.Bool:
		b := source.Bool()
		target.SetComplex(complex(float64(bool2int(b)), 0))

	case reflect.Float32, reflect.Float64:
		n := source.Float()
		x := complex(n, 0)
		target.SetComplex(x)
		if target.OverflowComplex(x) {
			c.AddErrorf("%s(%f) overflow", target.Type().String(), n)
		}

	case reflect.Complex64, reflect.Complex128:
		n := source.Complex()
		target.SetComplex(n)
		if target.OverflowComplex(n) {
			c.AddErrorf("%s(%v) overflow", target.Type().String(), n)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := source.Int()
		x := complex(float64(n), 0)
		target.SetComplex(x)
		if target.OverflowComplex(x) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n := source.Uint()
		x := complex(float64(n), 0)
		target.SetComplex(x)
		if target.OverflowComplex(x) {
			c.AddErrorf("%s(%d) overflow", target.Type().String(), n)
		}

	case reflect.String:
		return copyReflectToComplexFromString(c, target, source)
	}

	return true
}

func copyReflectToComplexFromString(c *context, target, source reflect.Value) bool {
	s := source.String()
	if x, err := strconv.ParseComplex(s, 128); err == nil {
		target.SetComplex(x)
		if target.OverflowComplex(x) {
			c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
		}
		return true
	}

	if n, err := strconv.ParseFloat(s, 64); err == nil {
		x := complex(n, 0)
		target.SetComplex(x)
		if target.OverflowComplex(x) {
			c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
		}
		return true
	}

	c.AddErrorf("can not convert to %s, %q", target.Type().String(), s)
	return true
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

		} else if target.OverflowFloat(n) {
			c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
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

		} else if target.OverflowInt(n) {
			c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
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

		} else if target.OverflowUint(n) {
			c.AddErrorf("%s(%s) overflow", target.Type().String(), s)
		}
	}

	return true
}
