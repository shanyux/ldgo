/*
 * Copyright (C) distroy
 */

package ldref

import "reflect"

func resizeSliceReflect(slice reflect.Value, size int) {
	l := slice.Len()
	if l >= size {
		return
	}

	tmp := reflect.MakeSlice(slice.Type(), size, size)
	for i := 0; i < l; i++ {
		tmp.Index(i).Set(slice.Index(i))
	}
	slice.Set(tmp)
}
