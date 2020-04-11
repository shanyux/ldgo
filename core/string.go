/*
 * Copyright (C) distroy
 */

package core

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

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
