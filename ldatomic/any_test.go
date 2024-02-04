/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAny(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("int", func(c convey.C) {
			v := NewAny[int](0)
			c.So(v.Load(), convey.ShouldEqual, 0)

			c.So(v.CompareAndSwap(0, 1), convey.ShouldBeTrue)
			c.So(v.Load(), convey.ShouldEqual, 1)

			c.So(v.Swap(20), convey.ShouldEqual, 1)
			c.So(v.Load(), convey.ShouldEqual, 20)

			c.So(v.CompareAndSwap(1, 2), convey.ShouldBeFalse)
			c.So(v.Load(), convey.ShouldEqual, 20)
		})
		c.Convey("uint", func(c convey.C) {
			v := &Any[uint]{}
			c.So(v.Load(), convey.ShouldEqual, 0)

			v.Store(1)
			c.So(v.Load(), convey.ShouldEqual, 1)

			c.So(v.Swap(20), convey.ShouldEqual, 1)
			c.So(v.Load(), convey.ShouldEqual, 20)

			c.So(v.CompareAndSwap(0, 1), convey.ShouldBeFalse)
			c.So(v.Load(), convey.ShouldEqual, 20)

			c.So(v.CompareAndSwap(1, 2), convey.ShouldBeFalse)
			c.So(v.Load(), convey.ShouldEqual, 20)
		})
	})
}
