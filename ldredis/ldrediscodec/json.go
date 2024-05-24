/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"bytes"
	"encoding/json"
	"reflect"
)

type codecBase[T any] struct{}

func (c codecBase[T]) fillIfPointer(v *T) (ptrOfT interface{}) {
	p := reflect.ValueOf(v).Elem()
	if p.Kind() != reflect.Ptr {
		return v
	}

	tmp := reflect.New(p.Type().Elem())
	p.Set(tmp)
	return p.Interface()
}

func Json[T any](d T) JsonCodec[T] { return JsonCodec[T]{} }

type JsonCodec[T any] struct {
	codecBase[T]
}

func (c JsonCodec[T]) Marshal(v T) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	b := buf.Bytes()
	// trim tail line end
	if i := len(b) - 1; i >= 0 && b[i] == '\n' {
		b = b[:i]
	}
	if bytes.Compare(b, []byte("null")) == 0 {
		return nil, nil
	}
	return b, nil
}

func (c JsonCodec[T]) Unmarshal(b []byte) (T, error) {
	var v T
	p := c.fillIfPointer(&v)
	return v, json.Unmarshal(b, p)
}
