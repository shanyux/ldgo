/*
 * Copyright (C) distroy
 */

package gorm

import "github.com/distroy/ldgo/v2/ldlog"

type testTable struct {
	ProjectId int64 `gorm:"column:project_id"`
	ChannelId int64 `gorm:"column:channel_id"`
	Type      int64 `gorm:"column:type"`
	VersionId int64 `gorm:"column:version_id"`
}

func (_ *testTable) TableName() string { return "test_table" }

func testGetGorm() *GormDb {
	db, _ := NewTestDb()
	db = db.WithLogger(ldlog.Default().Wrapper())
	return db
}
