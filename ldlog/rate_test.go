/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func Test_core_checkRateOrInterval(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		test_core_checkRateOrInterval_rate(c)
		test_core_checkRateOrInterval_interval(c)
	})
}

func test_core_checkRateOrInterval_rate(c convey.C) {
	l := &core{}

	c.Convey("rate", func(c convey.C) {
		l.setInterval(time.Second)
		c.Convey("1", func(c convey.C) {
			l.setRate(1)
			for i := 0; i < 100; i++ {
				c.So(l.checkRateOrInterval(0), convey.ShouldBeTrue)
			}
		})
		c.Convey("0", func(c convey.C) {
			l.setRate(0)
			for i := 0; i < 100; i++ {
				c.So(l.checkRateOrInterval(0), convey.ShouldBeFalse)
			}
		})
		c.Convey("0.5", func(c convey.C) {
			l.setRate(0.5)
			var (
				totalCnt = 20000
				diff     = 1000
			)
			trueCnt := 0
			for i := 0; i < totalCnt; i++ {
				if l.checkRateOrInterval(0) {
					trueCnt++
				}
			}
			half := totalCnt / 2
			c.So(trueCnt, convey.ShouldBeBetweenOrEqual, half-diff, half+diff)
		})
	})
}

func test_core_checkRateOrInterval_interval(c convey.C) {
	l := &core{}

	c.Convey("interval", func(c convey.C) {
		l.setRate(0)
		c.Convey("0", func(c convey.C) {
			l.setInterval(0)
			for i := 0; i < 100; i++ {
				c.So(l.checkRateOrInterval(0), convey.ShouldBeTrue)
			}
		})
		c.Convey("1s", func(c convey.C) {
			interval := time.Second
			l.setInterval(interval)

			c.So(l.checkRateOrInterval(1), convey.ShouldBeTrue)
			for i := 0; i < 100; i++ {
				c.So(l.checkRateOrInterval(1), convey.ShouldBeFalse)
			}
			time.Sleep(interval)
			c.So(l.checkRateOrInterval(1), convey.ShouldBeTrue)
			c.So(l.checkRateOrInterval(1), convey.ShouldBeFalse)
		})
	})
}
