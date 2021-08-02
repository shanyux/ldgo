/*
 * Copyright (C) distroy
 */

package ldgorm

func Equal(value interface{}) FieldWhere {
	return buildWhereSingleValue(" = ?", value)
}

type whereSingleValue struct {
	whereBase

	Query string      `json:"query"`
	Value interface{} `json:"value"`
}

func (that whereSingleValue) buildGorm(db *DB, field string) *DB {
	return db.Where(field+that.Query, that.Value)
}

func buildWhereSingleValue(query string, value interface{}) FieldWhere {
	val, isSet := getWhereValue(value)
	if !isSet {
		return whereEmpty{}
	}
	return whereSingleValue{
		Value: val,
		Query: query,
	}
}
