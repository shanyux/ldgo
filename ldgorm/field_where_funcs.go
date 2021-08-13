/*
 * Copyright (C) distroy
 */

package ldgorm

import "reflect"

func Equal(value interface{}) FieldWhere    { return newFieldWhereWithCheck(" = ?", value) }
func NotEqual(value interface{}) FieldWhere { return newFieldWhereWithCheck(" <> ?", value) }

func IsNull() FieldWhere    { return newFieldWhere(" IS NULL") }
func IsNotNull() FieldWhere { return newFieldWhere(" IS NOT NULL") }

func Between(min, max interface{}) FieldWhere {
	minVal, isMinSet := getWhereValue(min)
	maxVal, isMaxSet := getWhereValue(max)
	if !isMinSet && !isMaxSet {
		return fieldWhereEmpty{}
	}
	if !isMinSet {
		return newFieldWhere(" <= ?", maxVal)
	}
	if !isMaxSet {
		return newFieldWhere(" >= ?", minVal)
	}
	if minVal == maxVal {
		return newFieldWhere(" = ?", minVal)
	}

	return newFieldWhere(" BETWEEN ? AND ?", minVal, maxVal)
}

func Gt(value interface{}) FieldWhere {
	return newFieldWhereWithCheck(" > ?", value)
}
func Lt(value interface{}) FieldWhere {
	return newFieldWhereWithCheck(" < ?", value)
}
func Gte(value interface{}) FieldWhere {
	return newFieldWhereWithCheck(" >= ?", value)
}
func Lte(value interface{}) FieldWhere {
	return newFieldWhereWithCheck(" <= ?", value)
}

func In(value interface{}) FieldWhere {
	if value == nil {
		return fieldWhereEmpty{}
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	default:
		panic("the condition value for in must be slice or array type")

	case reflect.Slice, reflect.Array:
		break
	}

	if val.Len() == 0 {
		return fieldWhereEmpty{}
	}

	return newFieldWhere(" IN (?)", value)
}

func NotIn(value interface{}) FieldWhere {
	if value == nil {
		return fieldWhereEmpty{}
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	default:
		panic("the condition value for in must be slice or array type")

	case reflect.Slice, reflect.Array:
		break
	}

	if val.Len() == 0 {
		return fieldWhereEmpty{}
	}

	return newFieldWhere(" NOT IN (?)", value)
}

func Like(value string) FieldWhere {
	return newFieldWhereWithCheck(" LIKE ?", value)
}

func LikePrefix(value string) FieldWhere {
	value = value + "%"
	return newFieldWhereWithCheck(" LIKE ?", value)
}

func LikeSuffix(value string) FieldWhere {
	value = "%" + value
	return newFieldWhereWithCheck(" LIKE ?", value)
}

func LikeContain(value string) FieldWhere {
	value = "%" + value + "%"
	return newFieldWhereWithCheck(" LIKE ?", value)
}

func NotLike(value string) FieldWhere {
	return newFieldWhereWithCheck(" NOT LIKE ?", value)
}

func NotLikePrefix(value string) FieldWhere {
	value = value + "%"
	return newFieldWhereWithCheck(" NOT LIKE ?", value)
}

func NotLikeSuffix(value string) FieldWhere {
	value = "%" + value
	return newFieldWhereWithCheck(" NOT LIKE ?", value)
}

func NotLikeContain(value string) FieldWhere {
	value = "%" + value + "%"
	return newFieldWhereWithCheck(" NOT LIKE ?", value)
}
