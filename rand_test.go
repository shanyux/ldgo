/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_RandString(t *testing.T) {
	convey.Convey("", t, func() {
		a := RandString(16)
		b := RandString(16)
		convey.So(a, convey.ShouldHaveLength, 16)
		convey.So(b, convey.ShouldHaveLength, 16)
		convey.So(a, convey.ShouldNotEqual, b)
	})
}
