/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/distroy/ldgo/ldhook"
	"github.com/smartystreets/goconvey/convey"
)

type testSelectWhere struct {
	VersionId FieldWhere `gormwhere:"name:version_id;"`
	ChannelId FieldWhere `gormwhere:"name:channel_id;index:2;notempty"`
	ProjectId FieldWhere `gormwhere:"name:project_id;index:1"`
	Type      FieldWhere `gormwhere:"name:type;"`
}

func Test_BuildWhere(t *testing.T) {
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

		cond := &testSelectWhere{}
		convey.Convey("not empty fields are nil", func() {
			convey.So(func() { BuildWhere(nil, cond) }, convey.ShouldPanic)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("channel_id == 20", func() {
			cond.ChannelId = Equal(20)

			BuildWhere(nil, cond)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: "channel_id = ?", Args: []interface{}{20}},
			})
		})
		convey.Convey("project_id == 10 && channel_id == 20", func() {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)

			BuildWhere(nil, cond)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: "project_id = ?", Args: []interface{}{10}},
				{Query: "channel_id = ?", Args: []interface{}{20}},
			})
		})
		convey.Convey("project_id == 10 && channel_id == 20 && versionId > 30", func() {
			cond.ProjectId = Equal(10)
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)

			BuildWhere(nil, cond)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: "project_id = ?", Args: []interface{}{10}},
				{Query: "channel_id = ?", Args: []interface{}{20}},
				{Query: "version_id > ?", Args: []interface{}{30}},
			})
		})

		convey.Convey("channel_id == 20 && versionId > 30 && type in {1,2,3}", func() {
			cond.ChannelId = Equal(20)
			cond.VersionId = Gt(30)
			cond.Type = In([]int{1, 2, 3})

			BuildWhere(nil, cond)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: "channel_id = ?", Args: []interface{}{20}},
				{Query: "version_id > ?", Args: []interface{}{30}},
				{Query: "type IN (?)", Args: []interface{}{[]int{1, 2, 3}}},
			})
		})
	})
}
