/*
 * Copyright (C) distroy
 */

package cmp

import (
	"strings"
)

func CompareInterface(a, b interface{}) int {
	aa := reflectValueOf(a)
	bb := reflectValueOf(b)
	return CompareReflect(aa, bb)
}

func CompareBool[T ~bool](a, b T) int {
	switch {
	case a == b:
		return 0
	case bool(a):
		return 1
	default:
		return -1
	}
}

func CompareComplex[T ~complex64 | ~complex128](aa, bb T) int {
	a, b := complex128(aa), complex128(bb)
	if r := CompareOrderable(real(a), real(b)); r != 0 {
		return r
	}
	return CompareOrderable(imag(a), imag(b))
}

func CompareString[T ~string](a, b T) int {
	return strings.Compare(string(a), string(b))
	// return bytes.Compare(ldconv.StrToBytesUnsafe(a), ldconv.StrToBytesUnsafe(b))
}

func CompareComparer[T Comparer[T]](a, b T) int { return a.Compare(b) }
