/*
 * Copyright (C) distroy
 */

package gorm

import (
	"github.com/distroy/ldgo/ldlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type testTable struct {
	ProjectId int64 `gorm:"column:project_id"`
	ChannelId int64 `gorm:"column:channel_id"`
	Type      int64 `gorm:"column:type"`
	VersionId int64 `gorm:"column:version_id"`
}

func (_ *testTable) TableName() string { return "test_table" }

func testGetGorm() *GormDb {
	v, _ := gorm.Open("sqlite3", ":memory:")

	db := &GormDb{gormDb: v}
	db.SetLogger(ldlog.Discard().Wrapper())
	// db = db.WithLogger(ldlog.Console())
	db.CreateTable(&testTable{})

	return db
}
