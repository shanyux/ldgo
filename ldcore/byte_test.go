/*
 * Copyright (C) distroy
 */

package ldcore

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_ByteToUpper(t *testing.T) {
	convey.Convey("", t, func() {
		b0 := byte('a')
		b1 := byte('A')
		for i := byte(0); i < 'z'-'a'; i++ {
			c0 := b0 + i
			c1 := b1 + i
			r0 := ToUpper(c0)
			convey.So(r0, convey.ShouldEqual, c1)
		}
	})
}

func Test_ByteToLower(t *testing.T) {
	convey.Convey("", t, func() {
		b0 := byte('A')
		b1 := byte('a')
		for i := byte(0); i < 'z'-'a'; i++ {
			c0 := b0 + i
			c1 := b1 + i
			r0 := ToLower(c0)
			convey.So(r0, convey.ShouldEqual, c1)
		}
	})
}
