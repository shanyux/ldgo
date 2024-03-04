/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/distroy/ldgo/v2/lderr"
	"github.com/smartystreets/goconvey/convey"
)

type testCloneStruct struct {
	String string
	Int    int
	Struct *testCloneStruct
}

func TestClone(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("*int", func(c convey.C) {
			v0 := new(int)
			*v0 = 12345
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("[]int", func(c convey.C) {
			v0 := []int{1, 2, 3}
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("[3]int", func(c convey.C) {
			v0 := [3]int{1, 2, 3}
			v1 := Clone(v0)
			c.So(v1, convey.ShouldEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("map[string]int", func(c convey.C) {
			v0 := map[string]int{
				"a": 1,
				"b": 2,
			}
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("struct", func(c convey.C) {
			v0 := testCloneStruct{
				String: "abc",
				Int:    123,
			}
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("*struct", func(c convey.C) {
			v0 := &testCloneStruct{
				String: "abc",
				Int:    123,
			}
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})

		c.Convey("error", func(c convey.C) {
			v0 := lderr.ErrUnkown
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
	})
}
