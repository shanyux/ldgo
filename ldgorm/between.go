/*
 * Copyright (C) distroy
 */

package ldgorm

func Between(min, max interface{}) FieldWhere {
	minVal, isMinSet := getWhereValue(min)
	maxVal, isMaxSet := getWhereValue(max)
	if !isMinSet && !isMaxSet {
		return whereEmpty{}
	}
	if !isMinSet {
		return whereSingleValue{
			Value: maxVal,
			Query: " <= ?",
		}
	}
	if !isMaxSet {
		return whereSingleValue{
			Value: minVal,
			Query: " >= ?",
		}
	}
	if minVal == maxVal {
		return Equal(minVal)
	}
	return whereBetween{
		Min: minVal,
		Max: maxVal,
	}
}

func Gt(value interface{}) FieldWhere {
	return buildWhereSingleValue(" > ?", value)
}
func Lt(value interface{}) FieldWhere {
	return buildWhereSingleValue(" < ?", value)
}
func Gte(value interface{}) FieldWhere {
	return buildWhereSingleValue(" >= ?", value)
}
func Lte(value interface{}) FieldWhere {
	return buildWhereSingleValue(" <= ?", value)
}

type whereBetween struct {
	whereBase

	Min interface{} `json:"min"`
	Max interface{} `json:"max"`
}

func (that whereBetween) buildGorm(db *DB, field string) *DB {
	return db.Where(field+" BETWEEN ? AND ?", that.Min, that.Max)
}
