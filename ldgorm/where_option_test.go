/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
	"testing"

	"github.com/distroy/ldgo/ldlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/smartystreets/goconvey/convey"
)

type testFilter struct {
	VersionId FieldWherer `gormwhere:"column:version_id;"`
	ChannelId FieldWherer `gormwhere:"column:channel_id;order:2;notempty"`
	ProjectId FieldWherer `gormwhere:"column:project_id;order:1"`
	Type      FieldWherer `gormwhere:"column:type;"`
}

type testTable struct {
	ProjectId int64 `gorm:"column:project_id"`
	ChannelId int64 `gorm:"column:channel_id"`
	Type      int64 `gorm:"column:type"`
	VersionId int64 `gorm:"column:version_id"`
}

func (_ *testTable) TableName() string { return "test_table" }

func testGetGorm() *GormDb {
	v, _ := gorm.Open("sqlite3", ":memory:")

	db := NewGormDb(v)
	// convey.So(err, convey.ShouldBeNil)
	db.SetLogger(ldlog.Discard().Wrapper())
	db.CreateTable(&testTable{})
	return db
}

func testGetWhereFromSql(scope *gorm.Scope) string {
	const key = " WHERE "
	sql := scope.SQL
	idx := strings.Index(sql, key)
	if idx < 0 {
		return ""
	}
	return sql[idx+len(key):]
}

func TestWhereOption(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		gormDb := testGetGorm()
		defer gormDb.Close()

		var res whereResult
		gormDb.Callback().Query().After("gorm:query").Register("ldgorm:after_query", func(scope *gorm.Scope) {
			res.Query = testGetWhereFromSql(scope)
			res.Args = scope.SQLVars
		})

		var rows []*testTable

		convey.Convey("(project_id = 100 && channel_id > 100) || (project_id = 123 && channel_id < 234)", func() {
			where := Where(&testFilter{
				ProjectId: Equal(10),
				ChannelId: Gt(100),
			}).Or(&testFilter{
				ProjectId: Equal(123),
				ChannelId: Lt(234),
			})

			where.buildGorm(gormDb).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: "((`project_id` = ? AND `channel_id` > ?) OR (`project_id` = ? AND `channel_id` < ?))",
				Args:  []interface{}{10, 100, 123, 234},
			})
		})

		convey.Convey("(project_id = 10 && channel_id >= 0) && ((channel_id < 100 && version_id > 220) || channel_id > 200 && version_id < 110)", func() {
			where1 := Where(&testFilter{
				ProjectId: Equal(10),
				ChannelId: Between(0, nil),
			})
			where2 := Where(&testFilter{
				ChannelId: Lt(100),
				VersionId: Gt(220),
			}).Or(&testFilter{
				ChannelId: Gt(200),
				VersionId: Lt(110),
			})
			where := where1.And(where2)

			ApplyOptions(gormDb, where).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: "((`project_id` = ? AND `channel_id` >= ?) AND ((`channel_id` < ? AND `version_id` > ?) OR (`channel_id` > ? AND `version_id` < ?)))",
				Args:  []interface{}{10, 0, 100, 220, 200, 110},
			})
		})

		convey.Convey("(((channel_id < 100 AND version_id > 220) OR (channel_id > 200 AND version_id < 110)) AND (project_id = 10 AND channel_id >= 0))", func() {
			where1 := Where(&testFilter{
				ChannelId: Lt(100),
				VersionId: Gt(220),
			}).Or(&testFilter{
				ChannelId: Gt(200),
				VersionId: Lt(110),
			})
			where2 := Where(&testFilter{
				ProjectId: Equal(10),
				ChannelId: Between(0, nil),
			})
			where := Where(where1).And(where2)

			ApplyOptions(gormDb, where).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: "(((`channel_id` < ? AND `version_id` > ?) OR (`channel_id` > ? AND `version_id` < ?)) AND (`project_id` = ? AND `channel_id` >= ?))",
				Args:  []interface{}{100, 220, 200, 110, 10, 0},
			})
		})

		convey.Convey("right(channel_id, 1) = 1 && channel_id = 10 && b(channel_id)", func() {
			where := Where(&testFilter{
				ChannelId: Expr(`right({{column}}, ?) = ?`, 1, "%").And(Equal(10)).And(Expr(`b({{column}})`)),
			})

			where.buildGorm(gormDb).Find(&rows)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: "(right(`channel_id`, ?) = ? AND `channel_id` = ? AND b(`channel_id`))",
				Args:  []interface{}{1, "%", 10},
			})
		})
	})
}
