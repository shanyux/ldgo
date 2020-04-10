/*
 * Copyright (C) distroy
 */

package logger

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

type LoggerWrap interface {
	Sync()
	Enabled(lvl zapcore.Level) bool

	Core() *zap.Logger
	Sugar() *zap.SugaredLogger

	With(fields ...zap.Field) LoggerWrap
	WithOptions(opts ...zap.Option) LoggerWrap

	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(fmt string, args ...interface{})
	Debugz(fmt string, fields ...zap.Field)

	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(fmt string, args ...interface{})
	Infoz(fmt string, fields ...zap.Field)

	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(fmt string, args ...interface{})
	Warningz(fmt string, fields ...zap.Field)

	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(fmt string, args ...interface{})
	Warnz(fmt string, fields ...zap.Field)

	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(fmt string, args ...interface{})
	Errorz(fmt string, fields ...zap.Field)

	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(fmt string, args ...interface{})
	Fatalz(fmt string, fields ...zap.Field)

	Log(args ...interface{})
	Logln(args ...interface{})
	Logf(fmt string, args ...interface{})
	Logz(fmt string, fields ...zap.Field)

	Print(args ...interface{})
	Printf(fmt string, args ...interface{})
	Println(args ...interface{})
	Printz(fmt string, fields ...zap.Field)

	V(v int) bool
}

func newLoggerWrap(log *zap.Logger) loggerWrap {
	return loggerWrap{
		log:   log,
		sugar: log.Sugar(),
	}
}

type loggerWrap struct {
	log   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l loggerWrap) With(fields ...zap.Field) LoggerWrap {
	log := l.log.With(fields...)
	return newLoggerWrap(log)
}

func (l loggerWrap) WithOptions(opts ...zap.Option) LoggerWrap {
	log := l.log.WithOptions(opts...)
	return newLoggerWrap(log)
}

func (l loggerWrap) Sync() { l.log.Sync() }

func (l loggerWrap) Enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }

func (l loggerWrap) Core() *zap.Logger         { return l.log }
func (l loggerWrap) Sugar() *zap.SugaredLogger { return l.sugar }

func (l loggerWrap) emptyFunc(args ...interface{})  {}
func (l loggerWrap) enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }
func (l loggerWrap) f(lvl zapcore.Level) func(args ...interface{}) {
	if l.enabled(lvl) {
		switch lvl {
		case _DEBUG:
			return l.sugar.Debug
		case _INFO:
			return l.sugar.Info
		case _WARN:
			return l.sugar.Warn
		case _ERROR:
			return l.sugar.Error
		case _DPANIC:
			return l.sugar.DPanic
		case _PANIC:
			return l.sugar.Panic
		case _FATAL:
			return l.sugar.Fatal
		}
	}
	return l.emptyFunc
}

func (l loggerWrap) Debug(args ...interface{})              { l.f(_DEBUG)(sprintln(args)) }
func (l loggerWrap) Debugln(args ...interface{})            { l.f(_DEBUG)(sprintln(args)) }
func (l loggerWrap) Debugf(fmt string, args ...interface{}) { l.sugar.Debugf(fmt, args...) }
func (l loggerWrap) Debugz(fmt string, fields ...zap.Field) { l.log.Debug(fmt, fields...) }

func (l loggerWrap) Info(args ...interface{})              { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Infoln(args ...interface{})            { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Infof(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Infoz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Print(args ...interface{})              { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Println(args ...interface{})            { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Printf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Printz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Log(args ...interface{})              { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Logln(args ...interface{})            { l.f(_INFO)(sprintln(args)) }
func (l loggerWrap) Logf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Logz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Warn(args ...interface{})              { l.f(_WARN)(sprintln(args)) }
func (l loggerWrap) Warnln(args ...interface{})            { l.f(_WARN)(sprintln(args)) }
func (l loggerWrap) Warnf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l loggerWrap) Warnz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l loggerWrap) Warning(args ...interface{})              { l.f(_WARN)(sprintln(args)) }
func (l loggerWrap) Warningln(args ...interface{})            { l.f(_WARN)(sprintln(args)) }
func (l loggerWrap) Warningf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l loggerWrap) Warningz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l loggerWrap) Error(args ...interface{})              { l.f(_ERROR)(sprintln(args)) }
func (l loggerWrap) Errorln(args ...interface{})            { l.f(_ERROR)(sprintln(args)) }
func (l loggerWrap) Errorf(fmt string, args ...interface{}) { l.sugar.Errorf(fmt, args...) }
func (l loggerWrap) Errorz(fmt string, fields ...zap.Field) { l.log.Error(fmt, fields...) }

func (l loggerWrap) Fatal(args ...interface{})              { l.f(_FATAL)(sprintln(args)) }
func (l loggerWrap) Fatalln(args ...interface{})            { l.f(_FATAL)(sprintln(args)) }
func (l loggerWrap) Fatalf(fmt string, args ...interface{}) { l.sugar.Fatalf(fmt, args...) }
func (l loggerWrap) Fatalz(fmt string, fields ...zap.Field) { l.log.Fatal(fmt, fields...) }

func (l loggerWrap) V(v int) bool {
	if v <= 0 {
		return !l.Enabled(zapcore.DebugLevel)
	}
	return true
}
