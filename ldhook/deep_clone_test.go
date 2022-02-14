/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDeepClone(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("*int", func() {
			v0 := new(int)
			*v0 = 12345
			v1 := Clone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("[]*int", func() {
			ptrToInt := func(n int) *int { return &n }
			v0 := []*int{ptrToInt(1), ptrToInt(2), ptrToInt(3)}
			v1 := DeepClone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
			d1, _ := v1.([]*int)
			for i := range v0 {
				convey.So(d1[i], convey.ShouldNotEqual, v0[i])
			}
		})
		convey.Convey("[3]int", func() {
			v0 := [3]int{1, 2, 3}
			v1 := DeepClone(v0)
			convey.So(v1, convey.ShouldEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
		})
		convey.Convey("map[string]*int", func() {
			ptrToInt := func(n int) *int { return &n }
			v0 := map[string]*int{
				"a": ptrToInt(1),
				"b": ptrToInt(2),
				"c": ptrToInt(3),
				// "d": nil,
			}
			v1 := DeepClone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
			d1, _ := v1.(map[string]*int)
			for i := range v0 {
				if d1[i] == nil {
					continue
				}
				convey.So(d1[i], convey.ShouldNotEqual, v0[i])
			}
		})
		convey.Convey("struct", func() {
			v0 := testCloneStruct{
				String: "abc",
				Int:    123,
				Struct: &testCloneStruct{
					String: "xyz",
					Int:    234,
				},
			}
			v1 := DeepClone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
			d1, _ := v1.(testCloneStruct)
			convey.So(d1.Struct, convey.ShouldNotEqual, v0.Struct)
		})
		convey.Convey("*struct", func() {
			v0 := &testCloneStruct{
				String: "abc",
				Int:    123,
				Struct: &testCloneStruct{
					String: "xyz",
					Int:    234,
				},
			}
			v1 := DeepClone(v0)
			convey.So(v1, convey.ShouldNotEqual, v0)
			convey.So(v1, convey.ShouldResemble, v0)
			d1, _ := v1.(*testCloneStruct)
			convey.So(d1.Struct, convey.ShouldNotEqual, v0.Struct)
		})
	})
}
