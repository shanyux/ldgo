/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/distroy/ldgo/ldconv"
	"github.com/jinzhu/gorm"
)

const _WHERE_TAG = "gormwhere"

var (
	_WHERE_FIELD_TYPE = reflect.TypeOf((*FieldWhere)(nil)).Elem()

	whereCache = &sync.Map{}
)

func BuildWhere(db *gorm.DB, cond interface{}) *gorm.DB {
	if cond == nil {
		return db
	}

	val := reflect.ValueOf(cond)
	w := getWhereInfo(val.Type())

	return w.buildWhere(db, val)
}

type whereReflect struct {
	Fields []*fieldWhereReflect
}

type fieldWhereReflect struct {
	Tags       map[string]string
	Name       string
	Order      int32
	FieldOrder int
	NotEmpty   bool
}

func (w *whereReflect) buildWhere(db *gorm.DB, val reflect.Value) *gorm.DB {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	root := db

	for _, f := range w.Fields {
		fw, _ := val.Field(f.FieldOrder).Interface().(FieldWhere)
		if fw == nil || fw.isEmpty() {
			if f.NotEmpty {
				panic(fmt.Sprintf("the where field must not be empty. %s", f.Name))
			}
			continue
		}

		db = fw.buildGorm(root, f.Name)
	}
	return db
}

type sortSliceFieldInfo []*fieldWhereReflect

func (s sortSliceFieldInfo) Len() int      { return len(s) }
func (s sortSliceFieldInfo) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortSliceFieldInfo) Less(i, j int) bool {
	if s[i].Order != s[j].Order {
		return s[i].Order < s[j].Order
	}
	return s[i].FieldOrder < s[j].FieldOrder
}

func getWhereInfo(typ reflect.Type) *whereReflect {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("the where type must be struct or pointer to struct. %s", typ))
	}

	cache := whereCache
	if v, _ := cache.Load(typ); v != nil {
		reqT, _ := v.(*whereReflect)
		if reqT != nil {
			return reqT
		}
	}

	fields := make([]*fieldWhereReflect, 0, typ.NumField())
	for i, n := 0, typ.NumField(); i < n; i++ {
		f := getWhereFieldInfo(typ, i)
		if f == nil {
			continue
		}

		fields = append(fields, f)
	}

	if len(fields) == 0 {
		panic("struct must have at least one where field")
	}

	sort.Sort(sortSliceFieldInfo(fields))
	w := &whereReflect{
		Fields: fields,
	}

	cache.Store(typ, w)
	return w
}

func getWhereFieldInfo(typ reflect.Type, i int) *fieldWhereReflect {
	field := typ.Field(i)
	tag, ok := field.Tag.Lookup(_WHERE_TAG)
	if !ok {
		return nil
	}
	if len(tag) == 0 {
		return nil
	}

	tags := parseWhereTagString(tag)
	if _, ok := tags["-"]; ok {
		return nil
	}
	name, _ := tags["name"]
	if len(name) == 0 {
		return nil
	}

	if !field.Type.Implements(_WHERE_FIELD_TYPE) {
		panic("where field type must be `ldgorm.FieldWhere`")
	}

	order, _ := tags["order"]
	_, notEmpty := tags["notempty"]

	return &fieldWhereReflect{
		Tags:       tags,
		Name:       name,
		Order:      ldconv.AsInt32(order, math.MaxInt32),
		FieldOrder: i,
		NotEmpty:   notEmpty,
	}
}

func parseWhereTagString(tag string) map[string]string {
	tagList := strings.Split(tag, ";")
	m := make(map[string]string)
	for _, v := range tagList {
		if len(v) == 0 {
			continue
		}

		l := strings.SplitN(v, ":", 2)
		k := strings.TrimSpace(strings.ToLower(l[0]))
		if len(k) == 0 {
			continue
		}

		v := k
		if len(l) >= 2 {
			v = l[1]
		}
		m[k] = v
	}
	return m
}
