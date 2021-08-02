/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/distroy/ldgo/ldhook"
	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"
)

func Test_Equal(t *testing.T) {
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

		convey.Convey("Equal(nil)", func() {
			cond := Equal(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Equal(int)", func() {
			cond := Equal(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Equal(string)", func() {
			cond := Equal("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Equal(*int)", func() {
			cond := Equal(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Equal(*int(nil))", func() {
			cond := Equal((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Equal(*string)", func() {
			cond := Equal(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Equal(*string(nil))", func() {
			cond := Equal((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}
