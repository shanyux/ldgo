/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"fmt"
)

type FieldOrderer interface {
	buildGorm(db *GormDb, field string) *GormDb

	Order(i int) FieldOrderer
	Desc() FieldOrderer
	Asc() FieldOrderer

	getOrder() int
}

type fieldOrder struct {
	OrderNum int
	IsDesc   bool
}

func FieldOrder(i int) FieldOrderer {
	return fieldOrder{
		OrderNum: i,
	}
}

func (that fieldOrder) getOrder() int { return that.OrderNum }

func (that fieldOrder) buildGorm(db *GormDb, field string) *GormDb {
	exp := fmt.Sprintf("`%s`", field)
	if that.IsDesc {
		exp = fmt.Sprintf("`%s` DESC", field)
	}

	db = db.Order(exp)
	return db
}

func (that fieldOrder) Order(i int) FieldOrderer {
	that.OrderNum = i
	return that
}

func (that fieldOrder) Desc() FieldOrderer {
	that.IsDesc = true
	return that
}

func (that fieldOrder) Asc() FieldOrderer {
	that.IsDesc = false
	return that
}
