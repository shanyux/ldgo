/*
 * Copyright (C) distroy
 */

package ldrand

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestRandString(t *testing.T) {
	convey.Convey("", t, func() {
		a := String(16)
		b := String(16)
		convey.So(a, convey.ShouldHaveLength, 16)
		convey.So(b, convey.ShouldHaveLength, 16)
		convey.So(a, convey.ShouldNotEqual, b)
	})
}
