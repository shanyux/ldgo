/*
 * Copyright (C) distroy
 */

package ldtime

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestBeginAndEnd(t *testing.T) {

	convey.Convey(t.Name(), t, func() {
		t0 := time.Unix(1640125769, 0).In(testTz)
		t1 := t0.In(time.UTC)

		convey.So(t0.Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T06:29:29+0800")
		convey.So(t1.Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T22:29:29+0000")

		convey.Convey("minute", func() {
			convey.So(MinuteBegin(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T06:29:00+0800")
			convey.So(MinuteEnd(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T06:29:59+0800")

			convey.So(MinuteBegin(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T22:29:00+0000")
			convey.So(MinuteEnd(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T22:29:59+0000")
		})

		convey.Convey("hour", func() {
			convey.So(HourBegin(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T06:00:00+0800")
			convey.So(HourEnd(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T06:59:59+0800")

			convey.So(HourBegin(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T22:00:00+0000")
			convey.So(HourEnd(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T22:59:59+0000")
		})

		convey.Convey("date", func() {
			convey.So(DateBegin(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T00:00:00+0800")
			convey.So(DateEnd(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-22T23:59:59+0800")

			convey.So(DateBegin(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T00:00:00+0000")
			convey.So(DateEnd(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-21T23:59:59+0000")
		})

		convey.Convey("month", func() {
			convey.So(MonthBegin(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-01T00:00:00+0800")
			convey.So(MonthEnd(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-31T23:59:59+0800")

			convey.So(MonthBegin(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-01T00:00:00+0000")
			convey.So(MonthEnd(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-31T23:59:59+0000")
		})

		convey.Convey("year", func() {
			convey.So(YearBegin(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-01-01T00:00:00+0800")
			convey.So(YearEnd(t0).Format(testTimeFormat), convey.ShouldResemble, "2021-12-31T23:59:59+0800")

			convey.So(YearBegin(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-01-01T00:00:00+0000")
			convey.So(YearEnd(t1).Format(testTimeFormat), convey.ShouldResemble, "2021-12-31T23:59:59+0000")
		})
	})
}
