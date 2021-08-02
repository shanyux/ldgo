/*
 * Copyright (C) distroy
 */

package ldgorm

func Like(value string) FieldWhere {
	return buildWhereSingleValue(" LIKE ?", value)
}

func LikePrefix(value string) FieldWhere {
	value = value + "%"
	return buildWhereSingleValue(" LIKE ?", value)
}

func LikeSuffix(value string) FieldWhere {
	value = "%" + value
	return buildWhereSingleValue(" LIKE ?", value)
}

func LikeContain(value string) FieldWhere {
	value = "%" + value + "%"
	return buildWhereSingleValue(" LIKE ?", value)
}
