/*
 * Copyright (C) distroy
 */

package ldslice

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("part elements have been filtered", func() {
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			n := Filter(s, func(d int) bool { return d%2 == 0 })
			convey.So(s[:n], convey.ShouldResemble, []int{0, 2, 8, 4, 6})
		})
		convey.Convey("no element has been filtered", func() {
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			n := Filter(s, func(d int) bool { return true })
			convey.So(s[:n], convey.ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		})
		convey.Convey("all elements have been filtered", func() {
			s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			n := Filter(s, func(d int) bool { return false })
			convey.So(s[:n], convey.ShouldResemble, []int{})
		})
	})
}
