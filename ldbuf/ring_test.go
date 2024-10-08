/*
 * Copyright (C) distroy
 */

package ldbuf

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewRing(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		b := NewRing[int](4)
		c.So(b.Cap(), convey.ShouldEqual, 4)
		c.So(b.Size(), convey.ShouldEqual, 0)
		c.So(b.Pop(), convey.ShouldEqual, 0)

		b.Put(1)
		b.Put(2)
		b.Put(3)
		b.Put(4)
		b.Put(5)

		c.So(b.Pop(), convey.ShouldEqual, 1)
		c.So(b.Pop(), convey.ShouldEqual, 2)
		c.So(b.Size(), convey.ShouldEqual, 2)

		b.Put(6)
		b.Put(7)
		c.So(b.Pop(), convey.ShouldEqual, 3)
		c.So(b.Pop(), convey.ShouldEqual, 4)
		c.So(b.Pop(), convey.ShouldEqual, 6)
		c.So(b.Size(), convey.ShouldEqual, 1)
	})
}
