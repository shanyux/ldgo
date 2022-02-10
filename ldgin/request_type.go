/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"math"
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
		reqT, ok := v.(*requestType)
		if ok {
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
		if isReflectValueZero(valSrc) {
			continue
		}
		valDst := dst.FieldByIndex(field.Index)
		valDst.Set(valSrc)
	}
}

func isReflectValueZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isReflectValueZero(v.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()

	case reflect.String:
		return v.Len() == 0

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isReflectValueZero(v.Field(i)) {
				return false
			}
		}
		return true

	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{
			Method: "reflect.Value.IsZero",
			Kind:   v.Kind(),
		})
	}
}
