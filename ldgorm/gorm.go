/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"fmt"

	gorm2 "github.com/distroy/ldgo/v2/ldgorm/internal/jinzhu/gorm"
	"gorm.io/gorm"
)

type GormDb = gorm2.GormDb

func New(db *gorm.DB) *GormDb {
	return gorm2.New(db)
}

func NewGormDb(db *gorm.DB) *GormDb {
	return gorm2.New(db)
}

func NewTestGormDb() (*GormDb, error) {
	return gorm2.NewTestDb()
}

func MustNewTestGormDb() *GormDb {
	db, err := NewTestGormDb()
	if err != nil {
		panic(fmt.Sprintf("new test gorm fail. err:%v", err))
	}

	return db
}
