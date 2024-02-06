/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

type testOrderStruct struct {
	ProjectId FieldOrderer `ldgormorder:"column:project_id"`
	ChannelId FieldOrderer `ldgormorder:"column:channel_id"`
	VersionId FieldOrderer `ldgormorder:"column:version_id"`
	Type      FieldOrderer `ldgormorder:"column:type"`
}

func testGetOrderFromSql(sql string) string {
	const key = " ORDER BY "
	idx := strings.Index(sql, key)
	// log.Printf("=== sql:%s, key:`%s`, idx:%d", sql, key, idx)
	if idx < 0 {
		return ""
	}
	return sql[idx+len(key):]
}

func TestOrder(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		db := testGetGorm()
		defer db.Close()

		var rows []*testTable
		c.Convey("channel_id, version_id DESC", func(c convey.C) {
			order := Order(&testOrderStruct{
				ChannelId: FieldOrder(1).Asc(),
				VersionId: FieldOrder(2).Desc(),
			})

			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				db := New(tx)
				db = ApplyOptions(db, order)
				db = db.Find(&rows)
				return db.Get()
			})
			res := testGetOrderFromSql(sql)
			c.So(res, convey.ShouldEqual, "`channel_id`,`version_id` DESC")
		})

		c.Convey("channel_id DESC, type", func(c convey.C) {
			order := Order(&testOrderStruct{
				ChannelId: FieldOrder(1).Desc(),
				Type:      FieldOrder(2),
			})

			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				db := New(tx)
				db = ApplyOptions(db, order)
				db = db.Find(&rows)
				return db.Get()
			})
			res := testGetOrderFromSql(sql)
			c.So(res, convey.ShouldEqual, "`channel_id` DESC,`type`")
		})

		c.Convey("FIELD type", func(c convey.C) {
			order := Order(&testOrderStruct{
				Type: FieldOrder(1).Field([]int{2, 4, 3}),
			})

			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				db := New(tx)
				db = ApplyOptions(db, order)
				db = db.Find(&rows)
				return db.Get()
			})
			res := testGetOrderFromSql(sql)
			c.So(res, convey.ShouldEqual, "FIELD(`type`, 2, 4, 3)")
		})

		c.Convey("FIELD type DESC", func(c convey.C) {
			order := Order(&testOrderStruct{
				Type: FieldOrder(1).Field([]int{2, 4, 3}).Desc(),
			})

			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				db := New(tx)
				db = ApplyOptions(db, order)
				db = db.Find(&rows)
				return db.Get()
			})
			res := testGetOrderFromSql(sql)
			c.So(res, convey.ShouldEqual, "FIELD(`type`, 2, 4, 3) DESC")
		})
	})
}
