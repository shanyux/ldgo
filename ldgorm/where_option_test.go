/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testFilter struct {
	Table     string
	VersionId FieldWherer `ldgormwhere:"column:version_id;"`
	ChannelId FieldWherer `ldgormwhere:"column:channel_id;order:2;notempty"`
	ProjectId FieldWherer `ldgormwhere:"column:project_id;order:1"`
	Type      FieldWherer `ldgormwhere:"column:type;"`
}

type testFilterWithTableName testFilter

func (p testFilterWithTableName) TableName() string {
	if p.Table != "" {
		return p.Table
	}
	return "test_table"
}

type testTable struct {
	ProjectId int64 `gorm:"column:project_id"`
	ChannelId int64 `gorm:"column:channel_id"`
	Type      int64 `gorm:"column:type"`
	VersionId int64 `gorm:"column:version_id"`
}

func (_ *testTable) TableName() string { return "test_table" }

func testGetGorm() *GormDb {
	db := MustNewTestGormDb()
	// db = db.WithLogger(ldlog.Discard().Wrapper())
	db.CreateTable(&testTable{})
	return db
}

func testGetWhereFromSql(sql string) string {
	const key = " WHERE "
	idx := strings.Index(sql, key)
	if idx < 0 {
		return ""
	}
	res := sql[idx+len(key):]
	return res
}

func TestWhereOption(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		db := testGetGorm()
		defer db.Close()

		var rows []*testTable
		c.Convey("(project_id = 100 && channel_id > 100) || (project_id = 123 && channel_id < 234)", func(c convey.C) {
			where := Where(&testFilter{
				ProjectId: Equal(10),
				ChannelId: Gt(100),
			}).Or(&testFilter{
				ProjectId: Equal(123),
				ChannelId: Lt(234),
			})

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = where.buildGorm(db)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual,
				"(`project_id` = 10 AND `channel_id` > 100) OR (`project_id` = 123 AND `channel_id` < 234)")
		})

		c.Convey("(project_id = 10 && channel_id >= 0) && ((channel_id < 100 && version_id > 220) || channel_id > 200 && version_id < 110)", func(c convey.C) {
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

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = where.buildGorm(db)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual,
				"(`project_id` = 10 AND `channel_id` >= 0) AND ((`channel_id` < 100 AND `version_id` > 220) OR (`channel_id` > 200 AND `version_id` < 110))")
		})

		c.Convey("(((channel_id < 100 AND version_id > 220) OR (channel_id > 200 AND version_id < 110)) AND (project_id = 10 AND channel_id >= 0))", func(c convey.C) {
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

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = where.buildGorm(db)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual,
				"((`channel_id` < 100 AND `version_id` > 220) OR (`channel_id` > 200 AND `version_id` < 110)) AND (`project_id` = 10 AND `channel_id` >= 0)")
		})

		c.Convey("right(channel_id, 1) = % && channel_id = 10 && b(channel_id)", func(c convey.C) {
			where := Where(&testFilter{
				ChannelId: Expr(`right({{column}}, ?) = ?`, 1, "%").And(Equal(10)).And(Expr(`b({{column}})`)),
			})

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = where.buildGorm(db)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual,
				"right(`channel_id`, 1) = \"%\" AND `channel_id` = 10 AND b(`channel_id`)")
		})
	})
}
