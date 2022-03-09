/*
 * Copyright (C) distroy
 */

package ldtopk

import (
	"fmt"
	"testing"

	"github.com/distroy/ldgo/ldmath"
	"github.com/distroy/ldgo/ldrand"
	"github.com/distroy/ldgo/ldsort"
	"github.com/smartystreets/goconvey/convey"
)

func testTopkInts(n, k int) {
	name := fmt.Sprintf("topkints-n:%d-k:%d", n, k)
	convey.Convey(name, func() {
		origin := make([]int, 0, n)
		topk := make([]int, 0, k)

		for i := 0; i < n; i++ {
			x := ldrand.Intn(100)
			origin = append(origin, x)
			topk, _ = TopkIntsAdd(topk, k, x)
		}

		ldsort.SortInts(origin)
		ldsort.SortInts(topk)

		size := ldmath.MinInt(n, k)
		origin = origin[:size]
		convey.So(topk, convey.ShouldResemble, origin)
	})
}

func TestTopkInts(t *testing.T) {
	convey.Convey("", t, func() {
		testTopkInts(10, 20)
		testTopkInts(100, 5)
		testTopkInts(100, 10)
	})
}
