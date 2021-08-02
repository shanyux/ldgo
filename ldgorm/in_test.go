/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/distroy/ldgo/ldhook"
	"github.com/smartystreets/goconvey/convey"
)

func Test_In(t *testing.T) {
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

		convey.Convey("In(nil)", func() {
			cond := In(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("In(int)", func() {
			convey.So(func() { In(0) }, convey.ShouldPanic)
		})
		convey.Convey("In([3]int)", func() {
			cond := In([3]int{310, 320, 330})
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " IN (?)", Args: []interface{}{[3]int{310, 320, 330}}},
			})
		})
		convey.Convey("In([]int)", func() {
			cond := In([]int{310, 320, 330})
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " IN (?)", Args: []interface{}{[]int{310, 320, 330}}},
			})
		})
	})
}
