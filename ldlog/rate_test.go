/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"io"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newTestZapLog(writer io.Writer) *zap.Logger {
	testTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("2024-06-13T10:50:01.011+0800")
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		ConsoleSeparator: "|",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       testTimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
	}

	// enc := zapcore.NewJSONEncoder(cfg)
	enc := zapcore.NewConsoleEncoder(cfg)

	handle := zapcore.AddSync(writer)
	zCore := zapcore.NewCore(enc, handle, zap.DebugLevel)

	core := zapcore.NewTee(zCore)
	return zap.New(core)
}

func Test_core_enable(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		test_core_enable_rate(c)
		test_core_enable_interval(c)
	})
}

func test_core_enable_rate(c convey.C) {
	l := &core{log: newTestZapLog(io.Discard)}
	lvl := zapcore.InfoLevel

	c.Convey("rate", func(c convey.C) {
		l.enabler = IntervalEnabler(time.Second)
		c.Convey("1", func(c convey.C) {
			l.enabler = RateEnabler(1)
			for i := 0; i < 100; i++ {
				c.So(l.enable(lvl, 0), convey.ShouldBeTrue)
			}
		})
		c.Convey("0", func(c convey.C) {
			l.enabler = RateEnabler(0)
			for i := 0; i < 100; i++ {
				c.So(l.enable(lvl, 0), convey.ShouldBeFalse)
			}
		})
		c.Convey("0.5", func(c convey.C) {
			l.enabler = RateEnabler(0.5)
			var (
				totalCnt = 20000
				diff     = 1000
			)
			trueCnt := 0
			for i := 0; i < totalCnt; i++ {
				if l.enable(lvl, 0) {
					trueCnt++
				}
			}
			half := totalCnt / 2
			c.So(trueCnt, convey.ShouldBeBetweenOrEqual, half-diff, half+diff)
		})
	})
}

func test_core_enable_interval(c convey.C) {
	l := &core{log: newTestZapLog(io.Discard)}
	lvl := zapcore.InfoLevel

	c.Convey("interval", func(c convey.C) {
		l.enabler = RateEnabler(0)
		c.Convey("0", func(c convey.C) {
			l.enabler = IntervalEnabler(0)
			for i := 0; i < 100; i++ {
				c.So(l.enable(lvl, 0), convey.ShouldBeTrue)
			}
		})
		c.Convey("1s", func(c convey.C) {
			interval := time.Millisecond * 50
			l.enabler = IntervalEnabler(interval)

			time.Sleep(interval)
			c.So(l.enable(lvl, 1), convey.ShouldBeTrue)
			for i := 0; i < 100; i++ {
				c.So(l.enable(lvl, 1), convey.ShouldBeFalse)
			}
			time.Sleep(interval)
			c.So(l.enable(lvl, 1), convey.ShouldBeTrue)
			c.So(l.enable(lvl, 1), convey.ShouldBeFalse)
		})
	})
}
