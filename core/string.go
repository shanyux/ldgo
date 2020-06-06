/*
 * Copyright (C) distroy
 */

package core

import (
	"reflect"
	"unsafe"
)

func BytesToStr(b []byte) string {
	return BytesToStrUnsafe(b)
}

func StrToBytes(s string) []byte {
	return []byte(s)
}

func BytesToStrUnsafe(b []byte) string {
	if b == nil {
		return ""
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

// StrToBytesUnsafe the result may not be modified
func StrToBytesUnsafe(s string) []byte {
	// return *(*[]byte)(unsafe.Pointer(&s))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
