/*
 * Copyright (C) distroy
 */

package ldtopk

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/distroy/ldgo/ldmath"
	"github.com/distroy/ldgo/ldsort"
	"github.com/smartystreets/goconvey/convey"
)

func testTopkInt(n, k int) {
	name := fmt.Sprintf("n:%d k:%d", n, k)
	convey.Convey(name, func() {
		rand.Seed(time.Now().UnixNano())
		origin := make([]int, 0, n)
		topk := make([]int, 0, k)

		for i := 0; i < n; i++ {
			x := rand.Intn(100)
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

func TestTopkInt(t *testing.T) {
	convey.Convey("", t, func() {
		testTopkInt(10, 20)
		testTopkInt(100, 5)
		testTopkInt(100, 10)
	})
}
