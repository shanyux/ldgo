/*
 * Copyright (C) distroy
 */

package ldgorm

import "reflect"

func In(value interface{}) FieldWhere {
	if value == nil {
		return whereEmpty{}
	}
	val := reflect.ValueOf(value)
	switch val.Kind() {
	default:
		panic("the condition value for in must be slice or array type")

	case reflect.Slice, reflect.Array:
		break
	}

	if val.Len() == 0 {
		return whereEmpty{}
	}

	return whereIn{Value: value}
}

type whereIn struct {
	whereBase

	Value interface{} `json:"in"`
}

func (that whereIn) buildGorm(db *DB, field string) *DB {
	db = db.Where(field+" IN (?)", that.Value)
	return db
}
