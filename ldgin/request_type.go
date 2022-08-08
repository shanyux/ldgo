/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"
	"sync"

	"github.com/distroy/ldgo/ldref"
	"github.com/gin-gonic/gin"
)

var (
	requestTypes = &sync.Map{}
)

var requestBinds = []*requestBind{
	{Tag: "form", Func: (*gin.Context).ShouldBindQuery},
	{Tag: "json", Func: (*gin.Context).ShouldBindJSON},
	{Tag: "uri", Func: (*gin.Context).ShouldBindUri},
	{Tag: "header", Func: (*gin.Context).ShouldBindHeader},
}

type requestBind struct {
	Tag    string
	Func   func(g *gin.Context, o interface{}) error
	Fields []reflect.StructField
}

type requestType struct {
	Type    reflect.Type
	ReqPool sync.Pool
	ReqZero reflect.Value
	Binds   []*requestBind
}

func getRequestType(t reflect.Type) *requestType {
	if v, _ := requestTypes.Load(t); v != nil {
		reqT, ok := v.(*requestType)
		if ok {
			return reqT
		}
	}

	reqT := &requestType{
		Type:    t,
		ReqZero: reflect.Zero(t),
		Binds:   make([]*requestBind, 0, len(requestBinds)),
	}

	for _, v := range requestBinds {
		fields := getStructFieldsByTag(t, v.Tag)
		if len(fields) == 0 {
			continue
		}

		reqT.Binds = append(reqT.Binds, &requestBind{
			Tag:    v.Tag,
			Func:   v.Func,
			Fields: fields,
		})
	}

	if v, ok := requestTypes.LoadOrStore(t, reqT); ok {
		reqT, _ = v.(*requestType)
	}
	return reqT
}

func newRequest(reqType *requestType) reflect.Value {
	if i := reqType.ReqPool.Get(); i != nil {
		if v, ok := i.(reflect.Value); !ok {
			return v
		}
	}

	return reflect.New(reqType.Type)
}

func delRequest(reqType *requestType, val reflect.Value) {
	val.Elem().Set(reqType.ReqZero)
	reqType.ReqPool.Put(val) // nolint
}

func getStructFieldsByTag(objT reflect.Type, tag string) []reflect.StructField {
	fields := make([]reflect.StructField, 0, objT.NumField())
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i)
		tagStr, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		if tagStr == "" || tagStr == "-" {
			continue
		}

		fields = append(fields, field)
	}

	return fields
}

func fillHttpRequestByFeilds(dst, src reflect.Value, fields []reflect.StructField) {
	for _, field := range fields {
		dst.Type().NumField()
		valSrc := src.FieldByIndex(field.Index)
		if ldref.IsValZero(valSrc) {
			continue
		}
		valDst := dst.FieldByIndex(field.Index)
		valDst.Set(valSrc)
	}
}
