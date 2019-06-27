/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = context.Canceled
var _ = errors.New
var _ = fmt.Sprintf
var _ = log.New
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

func BytesToStr(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }
func StrToBytes(s string) []byte { return *(*[]byte)(unsafe.Pointer(&s)) }

var _RAND_STRING_LETTERS = StrToBytes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	letters := _RAND_STRING_LETTERS
	b := make([]byte, n)
	for i, _ := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return BytesToStr(b)
}

func StrMapReplace(_s string, m map[string]string, _l string, _r string) string {
	if len(m) == 0 {
		return _s
	}

	s := StrToBytes(_s)
	l := StrToBytes(_l)
	r := StrToBytes(_r)

	builder := &strings.Builder{}
	builder.Grow(len(s))
	for len(s) > 0 {
		bpos := bytes.Index(s, l)
		if bpos < 0 {
			builder.Write(s)
			break
		}
		builder.Write(s[:bpos])
		s = s[bpos+len(l):]

		epos := bytes.Index(s, r)
		if epos < 0 {
			builder.Write(l)
			builder.Write(s)
			break
		}

		key := s[:epos]
		val, ok := m[BytesToStr(key)]
		if !ok {
			builder.Write(l)
			continue
		}

		builder.Write(StrToBytes(val))
		s = s[epos+len(r):]
	}
	return builder.String()
}
