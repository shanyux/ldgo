/*
 * Copyright (C) distroy
 */

package ldgorm

import "github.com/jinzhu/gorm"

type FieldOrder interface {
	buildGorm(db *gorm.DB, field string) *gorm.DB

	Order(i int) FieldOrder
	Desc() FieldOrder
	Asc() FieldOrder

	getOrder() int
}

type fieldOrder struct {
	OrderNum int
	IsDesc   bool
}

func NewFieldOrder(i int) FieldOrder {
	return fieldOrder{
		OrderNum: i,
	}
}

func (that fieldOrder) getOrder() int { return that.OrderNum }

func (that fieldOrder) buildGorm(db *gorm.DB, field string) *gorm.DB {
	exp := field
	if that.IsDesc {
		exp = field + " DESC"
	}

	db = db.Order(exp)
	return db
}

func (that fieldOrder) Order(i int) FieldOrder {
	that.OrderNum = i
	return that
}

func (that fieldOrder) Desc() FieldOrder {
	that.IsDesc = true
	return that
}

func (that fieldOrder) Asc() FieldOrder {
	that.IsDesc = false
	return that
}
