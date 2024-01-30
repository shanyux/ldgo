/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/distroy/ldgo/v2/lderr"
	"github.com/smartystreets/goconvey/convey"
)

type testErrorStruct struct {
	value interface{}
}

func (testErrorStruct) Error() string { return "" }

type testErrorStruct2 struct {
	value interface{}
}

func (*testErrorStruct2) Error() string { return "" }

func TestMerge(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("fail", func() {
			convey.Convey("to invalid type", func() {
				err := Merge(1, 2)
				convey.So(err, convey.ShouldEqual, lderr.ErrReflectTargetNotPtr)
			})
			convey.Convey("to nil ptr", func() {
				err := Merge((*int)(nil), 2)
				convey.So(err, convey.ShouldEqual, lderr.ErrReflectTargetNilPtr)
			})
		})

		convey.Convey("succ", func() {
			convey.Convey("to interface", func() {
				convey.Convey("from struct", func() {
					var target error
					source := testErrorStruct{value: "abcde"}

					err := Merge(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, testErrorStruct{value: "abcde"})
				})
				convey.Convey("from ptr to struct 1", func() {
					var target error
					source := &testErrorStruct{value: "abcde"}

					err := Merge(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, testErrorStruct{value: "abcde"})
				})
				convey.Convey("from ptr to struct 2", func() {
					var target error
					source := &testErrorStruct2{value: "abcde"}

					err := Merge(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, &testErrorStruct2{value: "abcde"})
				})
			})

			convey.Convey("from nil ptr", func() {
				var (
					target int = 1
				)

				err := Merge(&target, (*int)(nil))
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 1)
			})

			convey.Convey("normal type 1", func() {
				var (
					target int
					source = 1234
				)

				err := Merge(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 1234)
			})
			convey.Convey("normal type 2", func() {
				var (
					target int
					source = 1234
				)

				err := Merge(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 1234)
			})

			convey.Convey("ptr", func() {
				convey.Convey("no clone", func() {
					var (
						target (*int)
						source = 1234
					)

					err := Merge(&target, &source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, &source)
				})
				convey.Convey("clone", func() {
					var (
						target (*int)
						source = 1234
					)

					err := Merge(&target, &source, &MergeConfig{Clone: true})
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldNotEqual, &source)
					convey.So(target, convey.ShouldResemble, &source)
				})
			})

			convey.Convey("struct", func() {
				var (
					target = &testCloneStruct{
						String: "abc",
					}
					source = &testCloneStruct{
						Int:    1234,
						String: "xyz",
					}
				)

				err := Merge(target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, &testCloneStruct{
					Int:    1234,
					String: "abc",
				})
			})

			convey.Convey("map", func() {
				var (
					target = map[string]any{
						"a": 1234,
						"c": &testCloneStruct{
							String: "abc",
						},
					}
					source = map[string]any{
						"a": 2345,
						"b": "abc",
						"c": &testCloneStruct{
							Int:    1234,
							String: "xyz",
						},
					}
				)

				err := Merge(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[string]any{
					"a": 1234,
					"b": "abc",
					"c": &testCloneStruct{
						Int:    1234,
						String: "abc",
					},
				})
			})

			convey.Convey("slice", func() {
				convey.Convey("no merge elem", func() {
					var (
						target = map[string]any{
							"a": []any(nil),
						}
						source = map[string]any{
							"a": []any{1, 2, 4},
						}
					)

					err := Merge(&target, &source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, map[string]any{
						"a": []any{1, 2, 4},
					})
				})

				convey.Convey("merge elem", func() {
					var (
						target = map[string]any{
							"a": []any{0, 3, 0},
						}
						source = map[string]any{
							"a": []any{1, 2, 4, 7},
						}
					)

					err := Merge(&target, &source, &MergeConfig{SliceElem: true})
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, map[string]any{
						"a": []any{1, 3, 4, 7},
					})
				})
			})

			convey.Convey("array", func() {
				convey.Convey("no merge elem", func() {
					var (
						target = [4]any{}
						source = [4]any{1, 2, 4}
					)

					err := Merge(&target, &source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, [4]any{1, 2, 4})
				})
				convey.Convey("merge elem", func() {
					var (
						target = [4]any{0, 0, 5}
						source = [4]any{0, 2, 0, 14}
					)

					err := Merge(&target, &source, &MergeConfig{ArrayElem: true})
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldResemble, [4]any{0, 2, 5, 14})
				})
			})
		})
	})
}
