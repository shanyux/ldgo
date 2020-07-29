/*
 * Copyright (C) distroy
 */

package core

import (
	"reflect"
	"unsafe"
)

type bytesFromString struct {
	Slice reflect.SliceHeader
	Data  unsafe.Pointer
}

func BytesToStr(b []byte) string {
	return string(b)
}

func StrToBytes(s string) []byte {
	return []byte(s)
}

// BytesToStrUnsafe the source bytes cannot free
func BytesToStrUnsafe(b []byte) string {
	if b == nil {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// StrToBytesUnsafe the result bytes cannot be modified and the source string cannot free
func StrToBytesUnsafe(s string) []byte {
	// return *(*[]byte)(unsafe.Pointer(&s))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &bytesFromString{
		Slice: reflect.SliceHeader{
			Data: sh.Data,
			Len:  sh.Len,
			Cap:  sh.Len,
		},
		Data: unsafe.Pointer(sh.Data),
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}
