/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/smartystreets/goconvey/convey"
)

type testOrderStruct struct {
	ProjectId FieldOrder `gormorder:"column:project_id"`
	ChannelId FieldOrder `gormorder:"column:channel_id"`
	VersionId FieldOrder `gormorder:"column:version_id"`
	Type      FieldOrder `gormorder:"column:type"`
}

func testGetOrderFromSql(scope *gorm.Scope) string {
	const key = " ORDER BY "
	sql := scope.SQL
	idx := strings.Index(sql, key)
	if idx < 0 {
		return ""
	}
	return sql[idx+len(key):]
}

func TestOrder(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		gormDb := testGetGorm()
		defer gormDb.Close()

		var res string
		gormDb.Callback().Query().After("gorm:query").Register("ldgorm:after_query", func(scope *gorm.Scope) {
			res = testGetOrderFromSql(scope)
		})

		var rows []*testTable

		convey.Convey("channel_id, version_id DESC", func() {
			order := Order(&testOrderStruct{
				ChannelId: NewFieldOrder(1).Asc(),
				VersionId: NewFieldOrder(2).Desc(),
			})

			ApplyOptions(gormDb, order).Find(&rows)
			convey.So(res, convey.ShouldResemble, "`channel_id`,`version_id` DESC")
		})

		convey.Convey("channel_id DESC, type", func() {
			order := Order(&testOrderStruct{
				ChannelId: NewFieldOrder(1).Desc(),
				Type:      NewFieldOrder(2),
			})

			ApplyOptions(gormDb, order).Find(&rows)
			convey.So(res, convey.ShouldResemble, "`channel_id` DESC,`type`")
		})
	})
}
