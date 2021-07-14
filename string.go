/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldconv"
	"github.com/distroy/ldgo/ldstr"
)

func BytesToStr(b []byte) string {
	return ldconv.BytesToStr(b)
}

// BytesToStrUnsafe the source bytes cannot free
func BytesToStrUnsafe(b []byte) string {
	return ldconv.BytesToStrUnsafe(b)
}

func StrToBytes(s string) []byte {
	return ldconv.StrToBytes(s)
}

// StrToBytesUnsafe the result bytes cannot be modified and the source string cannot free
func StrToBytesUnsafe(s string) []byte {
	return ldconv.StrToBytesUnsafe(s)
}

func StrMapReplace(s string, m map[string]string, l string, r string) string {
	return ldstr.StrMapReplace(s, m, l, r)
}

func StrIMapReplace(s string, m interface{}, l string, r string) string {
	return ldstr.StrIMapReplace(s, m, l, r)
}
