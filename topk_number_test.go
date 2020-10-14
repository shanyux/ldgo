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

func testTopkInt(t testing.TB, n, k int) {
	name := fmt.Sprintf("n:%d k:%d", n, k)
	convey.Convey(name, t, func() {
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

func TestTopkInt(t *testing.T) {
	testTopkInt(t, 10, 20)
	testTopkInt(t, 100, 5)
	testTopkInt(t, 100, 10)
}
