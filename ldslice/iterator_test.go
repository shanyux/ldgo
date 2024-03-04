/*
 * Copyright (C) distroy
 */

package ldslice

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testRangeIface[Data any] interface {
	Get() Data
	HasNext() bool
	Next()
}

func testRangeToSlice[Data any](r testRangeIface[Data]) []Data {
	l := make([]Data, 0, 16)
	for ; r.HasNext(); r.Next() {
		l = append(l, r.Get())
	}
	return l
}

func TestIterator(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		l := []int{1, 100, 50, 200, 70, 80, 30, 40, 20, 90, 150}

		c.Convey("forward", func(c convey.C) {
			r := testRangeToSlice[int](&Range[int]{
				Begin: Begin(l),
				End:   End(l),
			})
			c.So(r, convey.ShouldResemble, []int{
				1, 100, 50, 200, 70, 80, 30, 40, 20, 90, 150,
			})
		})
		c.Convey("backward", func(c convey.C) {
			r := testRangeToSlice[int](&ReverseRange[int]{
				Begin: RBegin(l),
				End:   REnd(l),
			})
			c.So(r, convey.ShouldResemble, []int{
				150, 90, 20, 40, 30, 80, 70, 200, 50, 100, 1,
			})
		})
	})
}
