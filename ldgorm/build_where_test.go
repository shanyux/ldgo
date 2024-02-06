/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestBuildWhere(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		db := testGetGorm()
		defer db.Close()

		var rows []*testTable
		cond := &testFilter{}
		c.Convey("not empty fields are nil", func(c convey.C) {
			c.So(func() { BuildWhere(db, cond) }, convey.ShouldPanic)
		})

		c.Convey("channel_id == 20", func(c convey.C) {
			cond.ChannelId = Equal(20)

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = BuildWhere(db, cond)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual, "`channel_id` = 20")
		})

		c.Convey("project_id == 10 && channel_id == 20", func(c convey.C) {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = BuildWhere(db, cond)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual, "`project_id` = 10 AND `channel_id` = 20")
		})

		c.Convey("project_id == 10 && channel_id == 20 && version_id > 30", func(c convey.C) {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = BuildWhere(db, cond)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual, "`project_id` = 10 AND `channel_id` = 20 AND `version_id` > 30")
		})

		c.Convey("channel_id == 20 && version_id > 30 && type in {1,2,3}", func(c convey.C) {
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)
			cond.Type = In([]int{1, 2, 3})

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = BuildWhere(db, cond)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual, "`channel_id` = 20 AND `version_id` > 30 AND `type` IN (1,2,3)")
		})

		c.Convey("channel_id == 20 && (version_id > 30 || version_id < 100)", func(c convey.C) {
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30).Or(Lt(100))

			sql := db.ToSQL(func(tx *GormDb) *GormDb {
				db := tx
				db = BuildWhere(db, cond)
				db = db.Find(&rows)
				return db
			})
			res := testGetWhereFromSql(sql)
			c.So(res, convey.ShouldEqual, "`channel_id` = 20 AND (`version_id` > 30 OR `version_id` < 100)")
		})

		c.Convey("with table name", func(c convey.C) {
			c.Convey("customer table name", func(c convey.C) {
				cond := &testFilterWithTableName{}
				cond.Table = "a"
				cond.ChannelId = Equal(20)
				cond.VersionId = Gt(30).Or(Lt(100))

				sql := db.ToSQL(func(tx *GormDb) *GormDb {
					db := tx
					db = BuildWhere(db, cond)
					db = db.Find(&rows)
					return db
				})
				res := testGetWhereFromSql(sql)
				c.So(res, convey.ShouldEqual,
					"`a`.`channel_id` = 20 AND (`a`.`version_id` > 30 OR `a`.`version_id` < 100)")
			})
			c.Convey("fixed table name", func(c convey.C) {
				cond := &testFilterWithTableName{}
				cond.ChannelId = Equal(20)
				cond.VersionId = Gt(30).Or(Lt(100))

				sql := db.ToSQL(func(tx *GormDb) *GormDb {
					db := tx
					db = BuildWhere(db, cond)
					db = db.Find(&rows)
					return db
				})
				res := testGetWhereFromSql(sql)
				c.So(res, convey.ShouldEqual,
					"`test_table`.`channel_id` = 20 AND (`test_table`.`version_id` > 30 OR `test_table`.`version_id` < 100)")
			})
		})
	})
}
