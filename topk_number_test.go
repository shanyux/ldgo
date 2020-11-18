/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func test_TopkInt(n, k int) {
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

		SortInts(origin)
		SortInts(topk)

		size := MinInt(n, k)
		origin = origin[:size]
		convey.So(topk, convey.ShouldResemble, origin)
	})
}

func Test_TopkInt(t *testing.T) {
	convey.Convey("", t, func() {
		test_TopkInt(10, 20)
		test_TopkInt(100, 5)
		test_TopkInt(100, 10)
	})
}
