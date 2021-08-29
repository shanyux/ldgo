/*
 * Copyright (C) distroy
 */

package ldlogger

import (
	"bytes"
	"testing"
	"time"

	"github.com/distroy/ldgo/ldhook"
	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		patches := ldhook.NewPatches()
		defer patches.Reset()
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: time.Now,
				Double: ldhook.Values{time.Unix(1629610258, 0)},
			},
		})

		type LoggerValue struct {
			Name string
		}

		writer := bytes.NewBuffer(nil)
		l := NewLogger(Writer(writer))
		l = With(l, zap.String("abc", "xxx"))

		convey.Convey("error", func() {
			l.Error("error message")
			convey.So(writer.String(), convey.ShouldEqual,
				"2021-08-22T13:30:58.000+0800|ERROR|-|ldlogger/logger_test.go:37|error message,abc=xxx\n")
		})

		convey.Convey("warn", func() {
			l.Warn("warn message")
			convey.So(writer.String(), convey.ShouldEqual,
				"2021-08-22T13:30:58.000+0800|WARN|-|ldlogger/logger_test.go:43|warn message,abc=xxx\n")
		})

		convey.Convey("info", func() {
			l.Info("info message")
			convey.So(writer.String(), convey.ShouldEqual,
				"2021-08-22T13:30:58.000+0800|INFO|-|ldlogger/logger_test.go:49|info message,abc=xxx\n")
		})

		convey.Convey("infoln", func() {
			l.Infoln("infoln message", &LoggerValue{Name: "abc"})
			convey.So(writer.String(), convey.ShouldEqual,
				"2021-08-22T13:30:58.000+0800|INFO|-|ldlogger/logger_test.go:55|infoln message &{abc},abc=xxx\n")
		})
	})
}
