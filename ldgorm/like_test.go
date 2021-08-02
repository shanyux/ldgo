/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/distroy/ldgo/ldhook"
	"github.com/smartystreets/goconvey/convey"
)

func Test_Like(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*DB).Where,
				Double: func(db *DB, query interface{}, args ...interface{}) *DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Like(abc)", func() {
			cond := Like("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " LIKE ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Like(%abc%)", func() {
			cond := Like("%abc%")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " LIKE ?", Args: []interface{}{"%abc%"}},
			})
		})
	})
}
