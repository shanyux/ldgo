/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"sync"
)

const (
	copyTagName = "copy"
)

var copyTypePool = &sync.Map{}

type copyFieldInfo struct {
	reflect.StructField

	Name   string
	Index  int
	Ignore bool
}

type copyStructInfo struct {
	Type    reflect.Type
	Fields  map[string]*copyFieldInfo
	Ignores []*copyFieldInfo
}

func getCopyTypeInfo(typ reflect.Type) *copyStructInfo {
	i, _ := copyTypePool.Load(typ)
	p, _ := i.(*copyStructInfo)
	if p != nil {
		return p
	}

	p = parseCopyStructInfo(typ)
	copyTypePool.LoadOrStore(typ, p)

	return p
}

func parseCopyStructInfo(typ reflect.Type) *copyStructInfo {
	num := typ.NumField()
	res := &copyStructInfo{
		Type:   typ,
		Fields: make(map[string]*copyFieldInfo, num),
	}

	for i := 0; i < num; i++ {
		f := parseCopyFieldInfo(i, typ.Field(i))
		if f.Ignore {
			res.Ignores = append(res.Ignores, f)
			continue
		}
		res.Fields[f.Name] = f
	}

	return res
}

func parseCopyFieldInfo(index int, field reflect.StructField) *copyFieldInfo {
	f := &copyFieldInfo{
		StructField: field,
		Name:        field.Name,
		Index:       index,
	}

	tagStr := field.Tag.Get(copyTagName)
	if tagStr == "" {
		return f
	}

	if tagStr == "-" {
		f.Ignore = true
		return f
	}

	if name := tagStr; name != "" {
		f.Name = name
	}

	return f
}
