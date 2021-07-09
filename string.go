/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

func BytesToStr(b []byte) string {
	return ldcore.BytesToStr(b)
}

// BytesToStrUnsafe the source bytes cannot free
func BytesToStrUnsafe(b []byte) string {
	return ldcore.BytesToStrUnsafe(b)
}

func StrToBytes(s string) []byte {
	return ldcore.StrToBytes(s)
}

// StrToBytesUnsafe the result bytes cannot be modified and the source string cannot free
func StrToBytesUnsafe(s string) []byte {
	return ldcore.StrToBytesUnsafe(s)
}

func StrMapReplace(s string, m map[string]string, l string, r string) string {
	return StrMapReplace(s, m, l, r)
}

func StrIMapReplace(s string, _m interface{}, l string, r string) string {
	return StrIMapReplace(s, m, l, r)
}
