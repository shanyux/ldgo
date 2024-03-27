/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"strings"
	"sync"
)

const defaultCopyTagName = "json"

var copyTypePool = &sync.Map{}
var copyNameKeysInTag = []string{
	"name",
	"column",
}

type copyFieldInfo struct {
	reflect.StructField

	Name   string
	Index  int
	Ignore bool
}

type copyStructKey struct {
	Type    reflect.Type
	TagName string
}

type copyStructValue struct {
	Type    reflect.Type
	TagName string
	Fields  map[string]*copyFieldInfo
	Ignores []*copyFieldInfo
}

func getCopyTypeInfo(typ reflect.Type, tagName string) *copyStructValue {
	if tagName == "" {
		tagName = defaultCopyTagName
	}
	key := copyStructKey{
		Type:    typ,
		TagName: tagName,
	}
	i, _ := copyTypePool.Load(key)
	if i != nil {
		return i.(*copyStructValue)
	}

	p := parseCopyStructInfo(typ, tagName)
	copyTypePool.LoadOrStore(key, p)

	return p
}

func parseCopyStructInfo(typ reflect.Type, tagName string) *copyStructValue {
	num := typ.NumField()
	res := &copyStructValue{
		Type:   typ,
		Fields: make(map[string]*copyFieldInfo, num),
	}

	for i := 0; i < num; i++ {
		f := parseCopyFieldInfo(i, typ.Field(i), tagName)
		if f.Ignore {
			res.Ignores = append(res.Ignores, f)
			continue
		}
		res.Fields[f.Name] = f
	}

	return res
}

func parseCopyFieldInfo(index int, field reflect.StructField, tagName string) *copyFieldInfo {
	f := &copyFieldInfo{
		StructField: field,
		Name:        field.Name,
		Index:       index,
	}

	tagStr := field.Tag.Get(tagName)
	if tagStr == "" {
		return f
	}

	if tagStr == "-" {
		f.Ignore = true
		return f
	}

	tagList := strings.FieldsFunc(tagStr, func(r rune) bool { return r == ';' || r == ',' })

	name := tagList[0]
	if name == "" {
		return f
	}

	idx := strings.Index(name, ":")
	if idx < 0 {
		f.Name = name
		return f
	}

	tagMap := make(map[string]string, len(tagList))
	tagMap[strings.ToLower(name[:idx])] = name[idx+1:]
	for _, str := range tagList[1:] {
		idx := strings.Index(str, ":")
		if idx < 0 {
			// tagMap[str] = ""
			continue
		}
		tagMap[strings.ToLower(str[:idx])] = str[idx+1:]
	}

	for _, key := range copyNameKeysInTag {
		val := tagMap[key]
		if val != "" {
			f.Name = val
			return f
		}
	}

	return f
}
