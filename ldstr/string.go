/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"reflect"
	"strings"

	"github.com/distroy/ldgo/ldconv"
)

func StrMapReplace(s string, m map[string]string, l string, r string) string {
	if len(m) == 0 {
		return s
	}

	replaceKeyValues := make([]string, 0, len(m)*2)
	for k, v := range m {
		kk := k
		vv := v

		kk = strings.Join([]string{l, kk, r}, "")
		if kk == "" {
			continue
		}
		replaceKeyValues = append(replaceKeyValues, kk, vv)
	}

	replacer := strings.NewReplacer(replaceKeyValues...)
	return replacer.Replace(s)
}

func StrIMapReplace(s string, _m interface{}, l string, r string) string {
	if m, ok := _m.(map[string]string); ok {
		return StrMapReplace(s, m, l, r)
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
	if val.Len() == 0 {
		return s
	}

	replaceKeyValues := make([]string, 0, val.Len()*2)
	for it := val.MapRange(); it.Next(); {
		k := it.Key().Interface()
		v := it.Value().Interface()

		kk := ldconv.AsString(k)
		vv := ldconv.AsString(v)

		kk = strings.Join([]string{l, kk, r}, "")
		if kk == "" {
			continue
		}
		replaceKeyValues = append(replaceKeyValues, kk, vv)
	}

	replacer := strings.NewReplacer(replaceKeyValues...)
	return replacer.Replace(s)
}
