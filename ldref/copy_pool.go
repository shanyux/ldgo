/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"sync"
)

var copyTypePool = &sync.Map{}

type copyFieldInfo struct {
	reflect.StructField

	Name  string
	Index int
}

type copyStructInfo struct {
	Fields map[string]*copyFieldInfo
}

func getCopyTypeInfo(typ reflect.Type) *copyStructInfo {
	i, _ := copyTypePool.Load(typ)
	p, _ := i.(*copyStructInfo)
	if p != nil {
		return p
	}

	return p
}

func parseCopyStructInfo(typ reflect.Type) *copyStructInfo {
	num := typ.NumField()
	res := &copyStructInfo{
		Fields: make(map[string]*copyFieldInfo, num),
	}

	for i := 0; i < num; i++ {
		field := parseCopyFieldInfo(i, typ.Field(i))
		if field == nil {
			continue
		}
		res.Fields[field.Name] = field
	}

	return res
}

func parseCopyFieldInfo(index int, field reflect.StructField) *copyFieldInfo {
	f := &copyFieldInfo{
		StructField: field,
		Name:        field.Name,
		Index:       index,
	}

	tagStr := field.Tag.Get("copy")
	if tagStr == "" {
		return f
	}

	if tagStr == "-" {
		return nil
	}

	if name := tagStr; name != "" {
		f.Name = name
	}

	return f
}
