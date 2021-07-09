/*
 * Copyright (C) distroy
 */

package ldcore

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_SortInt64s(t *testing.T) {
	convey.Convey("", t, func() {
		l := []int64{223, 562, 424, 642, 223, 123, 496, 623, 845, 375}

		SortInt64s(l)

		convey.So(IsSortedInt64s(l), convey.ShouldBeTrue)
		convey.So(l, convey.ShouldResemble, []int64{
			123, 223, 223, 375, 424, 496, 562, 623, 642, 845,
		})

		convey.So(SearchInt64s(l, 0), convey.ShouldEqual, 0)
		convey.So(SearchInt64s(l, 123), convey.ShouldEqual, 0)
		convey.So(SearchInt64s(l, 223), convey.ShouldEqual, 1)
		convey.So(SearchInt64s(l, 300), convey.ShouldEqual, 3)
		convey.So(SearchInt64s(l, 10000), convey.ShouldEqual, 10)
	})
}
