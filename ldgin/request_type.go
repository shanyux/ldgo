/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"
	"sync"
)

var (
	requestTypes = &sync.Map{}
)

type requestType struct {
	Type         reflect.Type
	ReqPool      sync.Pool
	ReqZero      reflect.Value
	FormFields   []reflect.StructField
	JsonFields   []reflect.StructField
	UriFields    []reflect.StructField
	HeaderFields []reflect.StructField
}

func getRequestType(t reflect.Type) *requestType {
	if v, _ := requestTypes.Load(t); v != nil {
		reqT, _ := v.(*requestType)
		if reqT != nil {
			return reqT
		}
	}

	reqT := &requestType{
		Type:         t,
		ReqZero:      reflect.Zero(t),
		FormFields:   getStructFieldsByTag(t, "form"),
		JsonFields:   getStructFieldsByTag(t, "json"),
		UriFields:    getStructFieldsByTag(t, "uri"),
		HeaderFields: getStructFieldsByTag(t, "header"),
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
	reqType.ReqPool.Put(val)
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
		if valSrc.IsZero() {
			continue
		}
		valDst := dst.FieldByIndex(field.Index)
		valDst.Set(valSrc)
	}
}
