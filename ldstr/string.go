/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"reflect"
	"strings"

	"github.com/distroy/ldgo/ldconv"
)

func getReplaceSplits(splits []string) (l, r string) {
	switch len(splits) {
	case 0:
		return "{", "}"
	case 1:
		return splits[0], splits[0]
	default:
		return splits[0], splits[1]
	}
}

// StrMapReplace returns new string.
//
//	StrMapReplace("user: {user}", map[string]string{"user", "x"}) returns "user: x"
//	StrMapReplace("user: #user#", map[string]string{"user", "x"}, "#") returns "user: x"
//	StrMapReplace("user: <user>", map[string]string{"user", "<", ">"}, "#") returns "user: x"
func StrMapReplace(tmpl string, m map[string]string, splits ...string) string {
	if len(m) == 0 {
		return tmpl
	}

	l, r := getReplaceSplits(splits)

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
	return replacer.Replace(tmpl)
}

// StrIMapReplace returns new string.
//
//	StrIMapReplace("user: {user}", map[interface{}]interface{}{"user", "x"}) returns "user: x"
//	StrIMapReplace("user: #user#", map[interface{}]interface{}{"user", "x"}, "#") returns "user: x"
//	StrIMapReplace("user: <user>", map[interface{}]interface{}{"user", "<", ">"}, "#") returns "user: x"
func StrIMapReplace(tmpl string, m interface{}, splits ...string) string {
	if m, ok := m.(map[string]string); ok {
		return StrMapReplace(tmpl, m, splits...)
	}

	l, r := getReplaceSplits(splits)

	val := reflect.ValueOf(m)
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
		return tmpl
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
	return replacer.Replace(tmpl)
}
