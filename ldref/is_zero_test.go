/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestIsZero(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("invalid", func() {
			convey.So(IsZero(nil), convey.ShouldEqual, true)
		})

		convey.Convey("bool", func() {
			convey.So(IsZero(false), convey.ShouldEqual, true)
			convey.So(IsZero(true), convey.ShouldEqual, false)
		})

		convey.Convey("int", func() {
			convey.So(IsZero(int(0)), convey.ShouldEqual, true)
			convey.So(IsZero(int(1)), convey.ShouldEqual, false)
		})

		convey.Convey("uint", func() {
			convey.So(IsZero(uint(0)), convey.ShouldEqual, true)
			convey.So(IsZero(uint(1)), convey.ShouldEqual, false)
		})

		convey.Convey("float64", func() {
			convey.So(IsZero(float64(0)), convey.ShouldEqual, true)
			convey.So(IsZero(float64(1)), convey.ShouldEqual, false)
		})

		convey.Convey("*int", func() {
			n := 1
			convey.So(IsZero((*int)(nil)), convey.ShouldEqual, true)
			convey.So(IsZero((*int)(&n)), convey.ShouldEqual, false)
		})

		convey.Convey("[]int", func() {
			convey.So(IsZero(([]int)(nil)), convey.ShouldEqual, true)
			convey.So(IsZero([]int{}), convey.ShouldEqual, false)
		})

		convey.Convey("[...]int", func() {
			convey.So(IsZero([3]int{}), convey.ShouldEqual, true)
			convey.So(IsZero([3]int{0, 0, 1}), convey.ShouldEqual, false)
		})

		convey.Convey("string", func() {
			convey.So(IsZero(""), convey.ShouldEqual, true)
			convey.So(IsZero("1"), convey.ShouldEqual, false)
		})
	})
}
