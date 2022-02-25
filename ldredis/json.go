/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"encoding/json"
	"reflect"
)

func JsonCodec(i ...interface{}) Codec {
	if len(i) == 0 || i[0] == nil {
		return &jsonCodec{}
	}

	return jsonCodec{
		typ: reflect.TypeOf(i[0]),
	}
}

type jsonCodec struct {
	typ reflect.Type
}

func (c jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c jsonCodec) Unmarshal(b []byte) (interface{}, error) {
	var p interface{}
	var v interface{}

	if c.typ != nil {
		switch c.typ.Kind() {
		case reflect.Map:
			v = reflect.MakeMap(c.typ).Interface()

		case reflect.Slice, reflect.Array:
			v = reflect.MakeSlice(reflect.SliceOf(c.typ.Elem()), 0, 0).Interface()

		case reflect.Struct:
			v = reflect.New(c.typ).Elem().Interface()

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Uintptr, reflect.UnsafePointer,
			reflect.Float32, reflect.Float64,
			reflect.Bool, reflect.String:
			v = reflect.New(c.typ).Elem().Interface()

		case reflect.Ptr:
			p = reflect.New(c.typ.Elem()).Interface()

		default:
		}
	}

	if p == nil {
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return v, nil
	}

	if err := json.Unmarshal(b, p); err != nil {
		return nil, err
	}
	return p, nil
}
