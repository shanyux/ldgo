/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSort(t *testing.T) {
	convey.Convey("sort int array", t, func(c convey.C) {
		compare := func(a, b int) int { return a - b }
		array := []int{6, 4, 5, 3, 1, 2}
		c.So(IsSorted[int](array, compare), convey.ShouldEqual, false)

		Sort[int](array, compare)
		c.So(array, convey.ShouldResemble, []int{1, 2, 3, 4, 5, 6})
		c.So(IsSorted[int](array, compare), convey.ShouldEqual, true)
	})
}
