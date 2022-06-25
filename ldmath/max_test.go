/*
 * Copyright (C) distroy
 */

package ldmath

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMax(t *testing.T) {
	convey.Convey("", t, func() {
		convey.So(MaxInt(3, 4), convey.ShouldEqual, 4)
		convey.So(MaxInt8(3, 4), convey.ShouldEqual, int8(4))
		convey.So(MaxInt16(3, 4), convey.ShouldEqual, int16(4))
		convey.So(MaxInt32(3, 4), convey.ShouldEqual, int32(4))
		convey.So(MaxInt64(3, 4), convey.ShouldEqual, int64(4))
	})
}

func TestMin(t *testing.T) {
	convey.Convey("", t, func() {
		convey.So(MinInt(4, 3), convey.ShouldEqual, 3)
		convey.So(MinInt8(4, 3), convey.ShouldEqual, int8(3))
		convey.So(MinInt16(4, 3), convey.ShouldEqual, int16(3))
		convey.So(MinInt32(4, 3), convey.ShouldEqual, int32(3))
		convey.So(MinInt64(4, 3), convey.ShouldEqual, int64(3))
	})
}
