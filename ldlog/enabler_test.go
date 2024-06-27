/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap/zapcore"
)

func TestRateEnabler_Enable(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		l := &rateEnabler{}
		info := zapcore.InfoLevel
		err := zapcore.ErrorLevel

		c.Convey("1", func(c convey.C) {
			l.rate = 1
			for i := 0; i < 100; i++ {
				c.So(l.Enable(info), convey.ShouldBeTrue)
			}
		})
		c.Convey("0", func(c convey.C) {
			l.rate = 0
			for i := 0; i < 100; i++ {
				c.So(l.Enable(info), convey.ShouldBeFalse)
			}
			c.So(l.Enable(err), convey.ShouldBeTrue)
		})
		c.Convey("0.5", func(c convey.C) {
			l.rate = 0.5
			var (
				totalCnt = 20000
				diff     = 1000
			)
			trueCnt := 0
			for i := 0; i < totalCnt; i++ {
				if l.Enable(info) {
					trueCnt++
				}
				c.So(l.Enable(err), convey.ShouldBeTrue)
			}
			half := totalCnt / 2
			c.So(trueCnt, convey.ShouldBeBetweenOrEqual, half-diff, half+diff)
		})
	})
}

func TestIntervalEnabler_Enable(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		l := &intervalEnabler{}
		info := zapcore.InfoLevel
		err := zapcore.ErrorLevel

		c.Convey("0", func(c convey.C) {
			l.interval = 0
			for i := 0; i < 100; i++ {
				c.So(l.Enable(info, 0), convey.ShouldBeTrue)
				c.So(l.Enable(err, 0), convey.ShouldBeTrue)
			}
		})
		c.Convey("1s", func(c convey.C) {
			interval := time.Millisecond * 50
			l.interval = interval

			time.Sleep(interval)
			c.So(l.Enable(info, 1), convey.ShouldBeTrue)
			c.So(l.Enable(err, 1), convey.ShouldBeTrue)
			for i := 0; i < 100; i++ {
				c.So(l.Enable(info, 1), convey.ShouldBeFalse)
				c.So(l.Enable(err, 1), convey.ShouldBeTrue)
			}
			time.Sleep(interval)
			c.So(l.Enable(err, 1), convey.ShouldBeTrue)
			c.So(l.Enable(info, 1), convey.ShouldBeTrue)
			c.So(l.Enable(info, 1), convey.ShouldBeFalse)
		})
	})
}
