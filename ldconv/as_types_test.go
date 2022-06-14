/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAsBool(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AsBool("abc"), convey.ShouldEqual, false)
		convey.So(AsBool("abc", true), convey.ShouldEqual, true)

		convey.So(AsBool("true", false), convey.ShouldEqual, true)
		convey.So(AsBool("false", true), convey.ShouldEqual, false)
		convey.So(AsBool([]byte("null"), true), convey.ShouldEqual, false)

		convey.So(AsBool(0, true), convey.ShouldEqual, false)

		convey.So(AsBool(int(-1)), convey.ShouldEqual, true)
		convey.So(AsBool(uint(1)), convey.ShouldEqual, true)
	})
}

func TestAsInt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AsInt("abc"), convey.ShouldEqual, 0)
		convey.So(AsInt("abc", 1), convey.ShouldEqual, 1)

		convey.So(AsInt("true", 0), convey.ShouldEqual, 1)
		convey.So(AsInt("false", 1), convey.ShouldEqual, 0)
		convey.So(AsInt([]byte("null"), 1), convey.ShouldEqual, 0)

		convey.So(AsInt(0, 100), convey.ShouldEqual, 0)

		convey.So(AsInt(int(-1)), convey.ShouldEqual, -1)
		convey.So(AsInt(uint(1)), convey.ShouldEqual, 1)
	})
}
