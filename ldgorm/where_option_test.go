/*
 * Copyright (C) Shopee
 */

package ldgorm

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/smartystreets/goconvey/convey"
)

func TestWhereOption(t *testing.T) {
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
				Query: " WHERE ((project_id = ? AND channel_id > ?) OR (project_id = ? AND channel_id < ?))",
				Args:  []interface{}{10, 100, 123, 234},
			})
		})
	})
}
