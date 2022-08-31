/*
 * Copyright (C) distroy
 */

package ldref

import "reflect"

var (
	typeOfByteSlice = reflect.TypeOf((*[]byte)(nil)).Elem()
	typeOfRuneSlice = reflect.TypeOf((*[]rune)(nil)).Elem()
)
