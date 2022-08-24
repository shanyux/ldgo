/*
 * Copyright (C) distroy
 */

package ldref

import "reflect"

var (
	typeOfEmptyStruct = reflect.TypeOf((*struct{})(nil)).Elem()
)
