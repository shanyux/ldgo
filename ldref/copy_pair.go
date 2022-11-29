/*
 * Copyright (C) distroy
 */

package ldref

import "reflect"

type copyPair struct {
	To   reflect.Kind
	From reflect.Kind
}

type copyFuncType = func(c *context, target, source reflect.Value) (end bool)

var copyFuncMap = map[copyPair]copyFuncType{}

func registerCopyFunc(m map[copyPair]copyFuncType) {
	for pair, fnCopy := range m {
		copyFuncMap[pair] = fnCopy
	}
}

func init() {
	registerCopyFunc(map[copyPair]copyFuncType{
		// {To: reflect.Ptr, From: reflect.Invalid}: copyReflectToPtrFromInvalid,
	})
}
