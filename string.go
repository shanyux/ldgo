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
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/distroy/ldgo/core"
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

func BytesToStr(b []byte) string {
	return core.BytesToStr(b)
}

func BytesToStrUnsafe(b []byte) string {
	return core.BytesToStrUnsafe(b)
}

func StrToBytes(s string) []byte {
	return core.StrToBytes(s)
}

// StrToBytesUnsafe the result may not be modified
func StrToBytesUnsafe(s string) []byte {
	return core.StrToBytesUnsafe(s)
}

var _RAND_STRING_LETTERS = StrToBytesUnsafe("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	letters := _RAND_STRING_LETTERS
	b := make([]byte, n)
	for i, _ := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return BytesToStrUnsafe(b)
}

func StrMapReplace(s string, m map[string]string, l string, r string) string {
	if len(m) == 0 {
		return s
	}

	builder := &strings.Builder{}
	builder.Grow(len(s))
	for len(s) > 0 {
		bpos := strings.Index(s, l)
		if bpos < 0 {
			builder.WriteString(s)
			break
		}
		builder.WriteString(s[:bpos])
		s = s[bpos+len(l):]

		epos := strings.Index(s, r)
		if epos < 0 {
			builder.WriteString(l)
			builder.WriteString(s)
			break
		}

		key := s[:epos]
		val, ok := m[key]
		if !ok {
			builder.WriteString(l)
			continue
		}

		builder.WriteString(val)
		s = s[epos+len(r):]
	}
	return builder.String()
}

func StrIMapReplace(s string, _m interface{}, l string, r string) string {
	if _m, ok := _m.(map[string]string); ok {
		return StrMapReplace(s, _m, l, r)
	}

	val := reflect.ValueOf(_m)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return ""
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Map {
		return ""
	}
	m := make(map[string]string, val.Len())

	for it := val.MapRange(); it.Next(); {
		k := it.Key().Interface()
		v := it.Value().Interface()
		m[AsString(k)] = AsString(v)
	}

	return StrMapReplace(s, m, l, r)
}
