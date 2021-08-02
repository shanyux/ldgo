/*
 * Copyright (C) distroy
 */

package ldgorm

type testGormWhere struct {
	Query interface{}
	Args  []interface{}
}

type testGormOrder struct {
	Value   interface{}
	Reorder []bool
}
