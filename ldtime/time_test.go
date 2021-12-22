/*
 * Copyright (C) distroy
 */

package ldtime

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

const (
	testTimeFormat = "2006-01-02T15:04:05-0700"
)

var testTz = fixedTimezone("Asia/Bejing", +(8 * time.Hour))

func paresTimeStr(s string) time.Time {
	const timeFormat = "2006-01-02T15:04:05-0700"
	tm, _ := time.Parse(timeFormat, s)
	return tm
}

func fixedTimezone(name string, tm time.Duration) *time.Location {
	sec := int(tm / time.Second)
	return time.FixedZone(name, sec)
}

func TestTimeToDateNum(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("2020-12-24 16:59:54", func() {
			t := paresTimeStr("2020-12-24T16:59:54+0800")
			n := TimeToDateNum(t)
			convey.So(n, convey.ShouldEqual, int64(20201224))
		})
	})
}

func TestDateNumToTime(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("20201224 +0800", func() {
			n := int64(20201224)
			t := DateNumToTime(n, testTz)
			convey.So(t, convey.ShouldResemble, paresTimeStr("2020-12-24T00:00:00+0800").In(testTz))
		})
	})
}

func TestTimeToNum(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("2020-12-24T16:59:54+0800", func() {
			t := paresTimeStr("2020-12-24T16:59:54+0800")
			n := TimeToNum(t)
			convey.So(n, convey.ShouldEqual, int64(20201224165954))
		})
	})
}
