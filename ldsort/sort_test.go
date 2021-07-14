/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_Sort(t *testing.T) {
	convey.Convey("sort int array", t, func() {
		array := [6]int{6, 4, 5, 3, 1, 2}
		Sort(array[:], func(a, b int) bool { return a < b })
		convey.So(array, convey.ShouldResemble, [6]int{1, 2, 3, 4, 5, 6})
	})
}
