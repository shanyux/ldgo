/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"github.com/distroy/ldgo/ldmath"
	"github.com/jinzhu/gorm"
)

type Option interface {
	buildGorm(db *gorm.DB) *gorm.DB
}

func ApplyOptions(db *gorm.DB, opts ...Option) *gorm.DB {
	for _, opt := range opts {
		db = opt.buildGorm(db)
	}
	return db
}

type pagingOption struct {
	Page     int64 // the first page is 1
	Pagesize int64
}

func (that pagingOption) buildGorm(db *gorm.DB) *gorm.DB {
	if that.Pagesize > 0 {
		that.Page = ldmath.MaxInt64(1, that.Page)
		offset := (that.Page - 1) * that.Pagesize
		db = db.Offset(offset).Limit(that.Pagesize)
	}

	return db
}

// Paging return the paging option
// the first page is 1
// if pagesize <= 0, it will query all rows
func Paging(page int64, pagesize int64) Option {
	return pagingOption{
		Page:     page,
		Pagesize: pagesize,
	}
}
