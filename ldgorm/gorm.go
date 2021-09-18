/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	gorm2 "github.com/distroy/ldgo/ldgorm/jinzhu/gorm"
	"github.com/jinzhu/gorm"
)

type GormDb = gorm2.GormDb

func NewGormDb(db *gorm.DB) *GormDb {
	g := &GormDb{}
	g = g.Set(db)
	return g
}
