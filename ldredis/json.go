/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	if bytes.Compare(b, []byte("null")) == 0 {
		return nil, nil
	}
	return b, nil
}

func (c jsonCodec) Unmarshal(b []byte) (interface{}, error) {
	if c.typ == nil {
		var v interface{}
		if err := json.Unmarshal(b, &v); err != nil {
			return nil, err
		}
		return v, nil
	}

	var p reflect.Value
	var v reflect.Value

	switch c.typ.Kind() {
	default:
		return nil, fmt.Errorf("the type is not supported. %s", c.typ.String())

	case reflect.Ptr:
		p = reflect.New(c.typ.Elem())
		v = p

	case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
		p = reflect.New(c.typ)
		v = p.Elem()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.UnsafePointer,
		reflect.Float32, reflect.Float64,
		reflect.Bool, reflect.String:

		p = reflect.New(c.typ)
		v = p.Elem()
	}

	if err := json.Unmarshal(b, p.Interface()); err != nil {
		return nil, err
	}

	return v.Interface(), nil
}
