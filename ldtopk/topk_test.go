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

func testTopk(n, k int) {
	less := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}

	name := fmt.Sprintf("topk-n:%d-k:%d", n, k)
	convey.Convey(name, func() {
		origin := make([]interface{}, 0, n)

		topk := &TopK{
			Size: k,
			Less: less,
		}

		for i := 0; i < n; i++ {
			x := ldrand.Intn(100)
			origin = append(origin, x)
			topk.Add(x)
		}

		ldsort.Sort(origin, less)
		ldsort.Sort(topk.Data(), less)

		size := ldmath.MinInt(n, k)
		origin = origin[:size]
		convey.So(topk.Data(), convey.ShouldResemble, origin)
	})
}

func TestTopK(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		testTopk(10, 20)
		testTopk(100, 5)
		testTopk(100, 10)
	})
}
