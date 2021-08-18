/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
	"testing"

	"github.com/distroy/ldgo/ldlogger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

type testFilter struct {
	VersionId FieldWhere `gormwhere:"name:version_id;"`
	ChannelId FieldWhere `gormwhere:"name:channel_id;order:2;notempty"`
	ProjectId FieldWhere `gormwhere:"name:project_id;order:1"`
	Type      FieldWhere `gormwhere:"name:type;"`
}

type testTable struct {
	ProjectId int64 `gorm:"column:project_id"`
	ChannelId int64 `gorm:"column:channel_id"`
	Type      int64 `gorm:"column:type"`
	VersionId int64 `gorm:"column:version_id"`
}

func (_ *testTable) TableName() string { return "test_table" }

func testGetGorm() *gorm.DB {
	db, _ := gorm.Open("sqlite3", ":memory:")
	// convey.So(err, convey.ShouldBeNil)
	db.LogMode(false)
	db.CreateTable(&testTable{})
	db.SetLogger(ldlogger.Console().WithOptions(zap.IncreaseLevel(zap.ErrorLevel)).Wrap())
	return db
}

func testGetWhereFromSql(scope *gorm.Scope) string {
	const key = " WHERE "
	sql := scope.SQL
	idx := strings.Index(sql, key)
	return sql[idx:]
}

func Test_BuildWhere(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"
		gormDb := testGetGorm()
		defer gormDb.Close()

		var res whereResult
		gormDb.Callback().Query().After("gorm:query").Register("ldgorm:after_query", func(scope *gorm.Scope) {
			res.Query = testGetWhereFromSql(scope)
			res.Args = scope.SQLVars
		})

		var rows []*testTable

		cond := &testFilter{}
		convey.Convey("not empty fields are nil", func() {
			convey.So(func() { BuildWhere(gormDb, cond) }, convey.ShouldPanic)
		})

		convey.Convey("channel_id == 20", func() {
			cond.ChannelId = Equal(20)

			BuildWhere(gormDb, cond).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: " WHERE (channel_id = ?)",
				Args:  []interface{}{20},
			})
		})
		convey.Convey("project_id == 10 && channel_id == 20", func() {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)

			BuildWhere(gormDb, cond).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: " WHERE (project_id = ? AND channel_id = ?)",
				Args:  []interface{}{10, 20},
			})
		})
		convey.Convey("project_id == 10 && channel_id == 20 && version_id > 30", func() {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)

			BuildWhere(gormDb, cond).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: " WHERE (project_id = ? AND channel_id = ? AND version_id > ?)",
				Args:  []interface{}{10, 20, 30},
			})
		})

		convey.Convey("channel_id == 20 && version_id > 30 && type in {1,2,3}", func() {
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)
			cond.Type = In([]int{1, 2, 3})

			BuildWhere(gormDb, cond).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: " WHERE (channel_id = ? AND version_id > ? AND type IN (?,?,?))",
				Args:  []interface{}{20, 30, 1, 2, 3},
			})
		})

		convey.Convey("channel_id == 20 && (version_id > 30 || version_id < 100)", func() {
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30).Or(Lt(100))

			BuildWhere(gormDb, cond).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: " WHERE (channel_id = ? AND (version_id > ? OR version_id < ?))",
				Args:  []interface{}{20, 30, 100},
			})
		})
	})
}
