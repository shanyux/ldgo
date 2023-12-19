/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"fmt"

	gorm2 "github.com/distroy/ldgo/ldgorm/internal/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type GormDb = gorm2.GormDb

func NewGormDb(db *gorm.DB) *GormDb {
	g := &GormDb{}
	g = g.Set(db)
	return g
}

func NewTestGormDb() (*GormDb, error) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	return NewGormDb(db), nil
}

func MustNewTestGormDb() *GormDb {
	db, err := NewTestGormDb()
	if err != nil {
		panic(fmt.Sprintf("new test gorm fail. err:%v", err))
	}

	return db
}
