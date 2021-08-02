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

	"github.com/distroy/ldgo/ldconv"
)

const _GORM_WHERE_TAG = "gormwhere"

type fieldWhereInfo struct {
	Condition  FieldWhere
	Value      reflect.Value
	Tags       map[string]string
	Name       string
	Index      int32
	FieldIndex int
	NotEmpty   bool
}

type sortSliceFieldInfo []*fieldWhereInfo

func (s sortSliceFieldInfo) Len() int      { return len(s) }
func (s sortSliceFieldInfo) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortSliceFieldInfo) Less(i, j int) bool {
	if s[i].Index != s[j].Index {
		return s[i].Index < s[j].Index
	}
	return s[i].FieldIndex < s[j].FieldIndex
}

func BuildWhere(db *DB, cond interface{}) *DB {
	if cond == nil {
		return db
	}

	val := reflect.ValueOf(cond)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic(fmt.Sprintf("the condtion must be struct"))
	}

	fields := make([]*fieldWhereInfo, 0, val.NumField())
	for i, n := 0, val.NumField(); i < n; i++ {
		f := getFieldInfo(val, i)
		if f == nil {
			continue
		}
		if cond := f.Condition; cond == nil || cond.isEmpty() {
			if f.NotEmpty {
				panic(fmt.Sprintf("the condtion field must not be empty. %s", f.Name))
			}
			continue
		}
		fields = append(fields, f)
	}
	sort.Sort(sortSliceFieldInfo(fields))

	for _, f := range fields {
		db = f.Condition.buildGorm(db, f.Name)
	}

	return db
}

func getFieldInfo(obj reflect.Value, i int) *fieldWhereInfo {
	field := obj.Type().Field(i)
	tag, ok := field.Tag.Lookup(_GORM_WHERE_TAG)
	if !ok {
		return nil
	}
	if len(tag) == 0 {
		return nil
	}

	tags := parseTagString(tag)
	if _, ok := tags["-"]; ok {
		return nil
	}
	name, _ := tags["NAME"]
	if len(name) == 0 {
		return nil
	}

	index, _ := tags["INDEX"]
	_, notEmpty := tags["NOTEMPTY"]

	value := obj.Field(i)
	cond, _ := value.Interface().(FieldWhere)

	return &fieldWhereInfo{
		Condition:  cond,
		Value:      value,
		Tags:       tags,
		Name:       name,
		Index:      ldconv.AsInt32(index, math.MaxInt32),
		FieldIndex: i,
		NotEmpty:   notEmpty,
	}
}

func parseTagString(tag string) map[string]string {
	tagList := strings.Split(tag, ";")
	m := make(map[string]string)
	for _, v := range tagList {
		if len(v) == 0 {
			continue
		}

		l := strings.SplitN(v, ":", 2)
		k := strings.TrimSpace(strings.ToUpper(l[0]))
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
