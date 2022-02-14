/*
 * Copyright (C) distroy
 */

package ldatomic

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		p := String{}
		convey.So(p.Load(), convey.ShouldEqual, "")

		p.Store("123")
		convey.So(p.Load(), convey.ShouldEqual, "123")
		p.Store("abc")
		convey.So(p.Load(), convey.ShouldEqual, "abc")
		p.Store("")
		convey.So(p.Load(), convey.ShouldEqual, "")

		convey.So(p.Swap("123"), convey.ShouldEqual, "")
		convey.So(p.Swap("abc"), convey.ShouldEqual, "123")
		convey.So(p.Swap(""), convey.ShouldEqual, "abc")
	})
}
