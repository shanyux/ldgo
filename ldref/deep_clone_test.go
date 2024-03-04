/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/distroy/ldgo/v2/lderr"
	"github.com/smartystreets/goconvey/convey"
)

func TestDeepClone(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("*int", func(c convey.C) {
			v0 := new(int)
			*v0 = 12345
			v1 := Clone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("[]*int", func(c convey.C) {
			ptrToInt := func(n int) *int { return &n }
			v0 := []*int{ptrToInt(1), ptrToInt(2), ptrToInt(3)}
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
			d1 := v1
			for i := range v0 {
				c.So(d1[i], convey.ShouldNotEqual, v0[i])
			}
		})
		c.Convey("[3]int", func(c convey.C) {
			v0 := [3]int{1, 2, 3}
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
		c.Convey("map[string]*int", func(c convey.C) {
			ptrToInt := func(n int) *int { return &n }
			v0 := map[string]*int{
				"a": ptrToInt(1),
				"b": ptrToInt(2),
				"c": ptrToInt(3),
				// "d": nil,
			}
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
			d1 := v1
			for i := range v0 {
				if d1[i] == nil {
					continue
				}
				c.So(d1[i], convey.ShouldNotEqual, v0[i])
				c.So(d1[i], convey.ShouldResemble, v0[i])
			}
		})
		c.Convey("struct", func(c convey.C) {
			v0 := testCloneStruct{
				String: "abc",
				Int:    123,
				Struct: &testCloneStruct{
					String: "xyz",
					Int:    234,
				},
			}
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
			d1 := v1
			c.So(d1.Struct, convey.ShouldNotEqual, v0.Struct)
		})
		c.Convey("*struct", func(c convey.C) {
			v0 := &testCloneStruct{
				String: "abc",
				Int:    123,
				Struct: &testCloneStruct{
					String: "xyz",
					Int:    234,
				},
			}
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
			d1 := v1
			c.So(d1.Struct, convey.ShouldNotEqual, v0.Struct)
		})

		c.Convey("error", func(c convey.C) {
			v0 := lderr.ErrUnkown
			v1 := DeepClone(v0)
			c.So(v1, convey.ShouldNotEqual, v0)
			c.So(v1, convey.ShouldResemble, v0)
		})
	})
}
