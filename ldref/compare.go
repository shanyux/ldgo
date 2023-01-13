/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"

	"github.com/distroy/ldgo/internal/cmp"
)

func Compare(a, b interface{}) int {
	return cmp.CompareInterface(a, b)
}

func CompareReflect(a, b reflect.Value) int {
	return cmp.CompareReflect(a, b)
}
