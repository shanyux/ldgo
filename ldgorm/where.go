/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"reflect"
)

type FieldWhere interface {
	isEmpty() bool
	buildGorm(db *DB, field string) *DB
}

func getWhereValue(v interface{}) (interface{}, bool) {
	if v == nil {
		return nil, false
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, false
		}
		return val.Elem().Interface(), true
	}

	return v, true
}

type whereEmpty struct{}

func (_ whereEmpty) isEmpty() bool                      { return true }
func (_ whereEmpty) buildGorm(db *DB, field string) *DB { return db }

type whereBase struct{}

func (_ whereBase) isEmpty() bool { return false }
