/*
 * Copyright (C) distroy
 */

package ldsheet

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/distroy/ldgo/ldtagmap"
)

type model struct {
	Inited  bool
	Headed  bool
	Type    reflect.Type
	Headers map[string]int
	Fields  []*modelField
}

type modelField struct {
	Field       reflect.StructField
	Tags        ldtagmap.Tags
	NotEmpty    bool
	Name        string
	Prefix      string
	FieldIndex  int
	HeaderIndex int
}

func parseModelField(field reflect.StructField) *modelField {
	tag, ok := field.Tag.Lookup(_TAG_NAME)
	if !ok {
		return &modelField{
			Field:       field,
			HeaderIndex: -1,
		}
	}

	tags := ldtagmap.Parse(tag)
	if _, ok := tags["-"]; ok {
		return nil
	}

	name := tags.Get("name")
	prefix := tags.Get("prefix")
	notEmpty := tags.Has("notempty")

	return &modelField{
		Field:       field,
		Name:        name,
		Prefix:      prefix,
		NotEmpty:    notEmpty,
		HeaderIndex: -1,
	}
}

func getFieldIndex(headers map[string]int, field *modelField) int {
	if name := field.Name; len(name) > 0 {
		if idx, ok := headers[name]; ok {
			return idx
		}
	}

	if prefix := field.Prefix; len(prefix) > 0 {
		for key, idx := range headers {
			if strings.HasPrefix(key, prefix) {
				return idx
			}
		}
	}

	if len(field.Name) > 0 || len(field.Prefix) > 0 {
		return -1
	}

	name := field.Field.Name
	if idx, ok := headers[name]; ok {
		return idx
	}

	name = splitStringWord(field.Field.Name)
	if idx, ok := headers[name]; ok {
		return idx
	}

	name = strings.ToLower(name)
	if idx, ok := headers[name]; ok {
		return idx
	}

	return -1
}

func checkFieldEmpty(f *modelField) error {
	if f.HeaderIndex >= 0 {
		return nil
	}

	if !f.NotEmpty {
		return nil
	}

	field := f.Field.Name
	if len(f.Name) > 0 {
		return fmt.Errorf("[ldsheet] the field is missed in header. field:%s, name:%s", field, f.Name)
	}

	if len(f.Prefix) > 0 {
		return fmt.Errorf("[ldsheet] the field is missed in header. field:%s, prefix:%s", field, f.Prefix)
	}

	return fmt.Errorf("[ldsheet] the field is missed in header. field:%s", field)
}

func parseFieldValue(field *modelField, obj reflect.Value, line []string) error {
	if field.HeaderIndex < 0 {
		return nil
	}

	if field.HeaderIndex >= len(line) || len(line[field.HeaderIndex]) == 0 {
		if field.NotEmpty {
			return fmt.Errorf("[ldsheet] the field must not be empty. field:%s", field.Field.Name)
		}
		return nil
	}

	str := line[field.HeaderIndex]
	fVal := obj.Field(field.FieldIndex)
	err := parseStringValue(fVal, str)
	if err != nil {
		return fmt.Errorf("[ldsheet] parse field value fail. type:%s, err:%s",
			fVal.Type().String(), err)
	}

	return nil
}
