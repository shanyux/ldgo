/*
 * Copyright (C) distroy
 */

package ldtime

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

var _TZ_TEST = fixedTimezone("Asia/Bejing", +(8 * time.Hour))

func paresTimeStr(s string) time.Time {
	const TIME_FORMAT = "2006-01-02T15:04:05-0700"
	tm, _ := time.Parse(TIME_FORMAT, s)
	return tm
}

func fixedTimezone(name string, tm time.Duration) *time.Location {
	sec := int(tm / time.Second)
	return time.FixedZone(name, sec)
}

func Test_DateBegin(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		tm := paresTimeStr("2020-11-10T06:42:23+0800")
		convey.So(tm.Nanosecond(), convey.ShouldEqual, 0)
		convey.So(tm.Unix(), convey.ShouldEqual, 1604961743)

		convey.Convey("+0800", func() {
			t := tm
			t = DateBegin(t)
			convey.So(t.Nanosecond(), convey.ShouldEqual, 0)
			convey.So(t.Unix(), convey.ShouldEqual, 1604937600)
		})
		convey.Convey("+0000", func() {
			t := tm.In(time.FixedZone("test", 0))
			t = DateBegin(t)
			convey.So(t.Nanosecond(), convey.ShouldEqual, 0)
			convey.So(t.Unix(), convey.ShouldEqual, 1604880000)
		})
	})
}

func Test_TimeToDateNum(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("2020-12-24 16:59:54", func() {
			t := paresTimeStr("2020-12-24T16:59:54+0800")
			n := TimeToDateNum(t)
			convey.So(n, convey.ShouldEqual, int64(20201224))
		})
	})
}

func Test_DateNumToTime(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("20201224 +0800", func() {
			n := int64(20201224)
			t := DateNumToTime(n, _TZ_TEST)
			convey.So(t, convey.ShouldResemble, paresTimeStr("2020-12-24T00:00:00+0800").In(_TZ_TEST))
		})
	})
}

func Test_TimeToNum(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("2020-12-24T16:59:54+0800", func() {
			t := paresTimeStr("2020-12-24T16:59:54+0800")
			n := TimeToNum(t)
			convey.So(n, convey.ShouldEqual, int64(20201224165954))
		})
	})
}
