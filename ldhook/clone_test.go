/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testCloneStruct struct {
	String string
	Int    int
	Struct *testCloneStruct
}

func TestClone(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("*int", func() {
			v0 := new(int)
			*v0 = 12345
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("[]int", func() {
			v0 := []int{1, 2, 3}
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("[3]int", func() {
			v0 := [3]int{1, 2, 3}
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("map[string]int", func() {
			v0 := map[string]int{
				"a": 1,
				"b": 2,
			}
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("struct", func() {
			v0 := testCloneStruct{
				String: "abc",
				Int:    123,
			}
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("*struct", func() {
			v0 := &testCloneStruct{
				String: "abc",
				Int:    123,
			}
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
	})
}
