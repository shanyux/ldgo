/*
 * Copyright (C) distroy
 */

package ldsheet

import (
	"fmt"
	"io"
	"reflect"
)

const (
	_TAG_NAME = "sheet"
)

type Reader struct {
	Reader LineReader
	model  model
}

func NewReader(r LineReader) *Reader {
	return &Reader{
		Reader: r,
	}
}

func (r *Reader) Read(d interface{}) error {
	if err := r.init(d); err != nil {
		return err
	}

	if err := r.readHeader(); err != nil {
		return err
	}

	if err := r.readBody(d); err != nil {
		return err
	}

	return nil
}

func (r *Reader) init(d interface{}) error {
	m := &r.model
	if m.Inited {
		return nil
	}

	typ := reflect.TypeOf(d)
	if typ.Kind() != reflect.Ptr && typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("[sheet] the data must be the pointer to struct. type:%s", typ.String())
	}

	pTyp := typ         // pointer type
	oTyp := pTyp.Elem() // object type

	fields := make([]*modelField, 0, oTyp.NumField())
	for i, l := 0, oTyp.NumField(); i < l; i++ {
		v := parseModelField(oTyp.Field(i))
		if v == nil {
			continue
		}

		v.FieldIndex = i
		fields = append(fields, v)
	}

	m.Type = pTyp
	m.Fields = fields
	m.Inited = true
	return nil
}

func (r *Reader) readHeader() error {
	m := &r.model
	if m.Headed {
		return nil
	}

	header, err := r.Reader.Read()
	if err != nil {
		return fmt.Errorf("[sheet] read header fail. err:%s", err)
	}

	headerMap := make(map[string]int, len(header))
	for i, v := range header {
		headerMap[v] = i
	}

	for _, field := range m.Fields {
		field.HeaderIndex = -1

		if idx := getFieldIndex(headerMap, field); idx >= 0 {
			field.HeaderIndex = idx

		} else if err := checkFieldEmpty(field); err != nil {
			return err
		}
	}

	m.Headers = headerMap
	m.Headed = true
	return nil
}

func (r *Reader) readBody(d interface{}) error {
	m := &r.model

	val := reflect.ValueOf(d)
	typ := val.Type()
	if m.Type != typ {
		return fmt.Errorf("[sheet] read the inconsistently type. expected:%s", m.Type.String())
	}

	obj := val.Elem()
	obj.Set(reflect.Zero(obj.Type()))

	line, err := r.Reader.Read()
	if err != nil {
		if err == io.EOF {
			return io.EOF
		}
		return fmt.Errorf("[sheet] read body fail. err:%s", err)
	}

	for _, field := range m.Fields {
		err := parseFieldValue(field, obj, line)
		if err != nil {
			return err
		}
	}

	return nil
}
