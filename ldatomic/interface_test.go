/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestInterface(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewInterface(nil).Load(), convey.ShouldBeNil)
		convey.So(NewInterface("123").Load(), convey.ShouldEqual, "123")

		p := Interface{}
		convey.So(p.Load(), convey.ShouldBeNil)

		p.Store("123")
		convey.So(p.Load(), convey.ShouldEqual, "123")
		p.Store(123)
		convey.So(p.Load(), convey.ShouldEqual, 123)
		p.Store(nil)
		convey.So(p.Load(), convey.ShouldEqual, nil)

		convey.So(p.Swap("123"), convey.ShouldEqual, nil)
		convey.So(p.Swap(123), convey.ShouldEqual, "123")
		convey.So(p.Swap(nil), convey.ShouldEqual, 123)
		convey.So(p.Load(), convey.ShouldEqual, nil)
	})
}
