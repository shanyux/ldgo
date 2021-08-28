/*
 * Copyright (C) distroy
 */

package ldlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerWrapper interface {
	Sync()
	Enabled(lvl zapcore.Level) bool

	Core() *zap.Logger
	Sugar() *zap.SugaredLogger

	// Debugf formats according to a format specifier and print the logger
	Debugf(fmt string, args ...interface{})
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugz(fmt string, fields ...zap.Field)

	// Infof formats according to a format specifier and print the logger
	Infof(fmt string, args ...interface{})
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infoz(fmt string, fields ...zap.Field)

	// Printf formats according to a format specifier and print the logger
	Warningf(fmt string, args ...interface{})
	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningz(fmt string, fields ...zap.Field)

	// Warnf formats according to a format specifier and print the logger
	Warnf(fmt string, args ...interface{})
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnz(fmt string, fields ...zap.Field)

	// Errorf formats according to a format specifier and print the logger
	Errorf(fmt string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorz(fmt string, fields ...zap.Field)

	// Fatalf formats according to a format specifier and print the logger
	Fatalf(fmt string, args ...interface{})
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalz(fmt string, fields ...zap.Field)

	// Logf formats according to a format specifier and print the logger
	Logf(fmt string, args ...interface{})
	Log(args ...interface{})
	Logln(args ...interface{})
	Logz(fmt string, fields ...zap.Field)

	// Printf formats according to a format specifier and print the logger
	Printf(fmt string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printz(fmt string, fields ...zap.Field)

	V(v int) bool
}

func newLoggerWrapper(log *zap.Logger) loggerWrap {
	return loggerWrap{
		log:   log,
		sugar: log.Sugar(),
	}
}

type loggerWrap struct {
	log   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l loggerWrap) Sync() { l.log.Sync() }

func (l loggerWrap) Enabled(lvl zapcore.Level) bool { return l.enabled(lvl) }

func (l loggerWrap) Core() *zap.Logger         { return l.log }
func (l loggerWrap) Sugar() *zap.SugaredLogger { return l.sugar }

func (l loggerWrap) enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }

func (l loggerWrap) Debugf(fmt string, args ...interface{}) { l.sugar.Debugf(fmt, args...) }
func (l loggerWrap) Debug(args ...interface{})              { l.sugar.Debug(pw(args)) }
func (l loggerWrap) Debugln(args ...interface{})            { l.sugar.Debug(pw(args)) }
func (l loggerWrap) Debugz(fmt string, fields ...zap.Field) { l.log.Debug(fmt, fields...) }

func (l loggerWrap) Infof(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Info(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l loggerWrap) Infoln(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l loggerWrap) Infoz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Printf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Print(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l loggerWrap) Println(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l loggerWrap) Printz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Logf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l loggerWrap) Log(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l loggerWrap) Logln(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l loggerWrap) Logz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l loggerWrap) Warnf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l loggerWrap) Warn(args ...interface{})              { l.sugar.Warn(pw(args)) }
func (l loggerWrap) Warnln(args ...interface{})            { l.sugar.Warn(pw(args)) }
func (l loggerWrap) Warnz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l loggerWrap) Warningf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l loggerWrap) Warning(args ...interface{})              { l.sugar.Warn(pw(args)) }
func (l loggerWrap) Warningln(args ...interface{})            { l.sugar.Warn(pw(args)) }
func (l loggerWrap) Warningz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l loggerWrap) Errorf(fmt string, args ...interface{}) { l.sugar.Errorf(fmt, args...) }
func (l loggerWrap) Error(args ...interface{})              { l.sugar.Error(pw(args)) }
func (l loggerWrap) Errorln(args ...interface{})            { l.sugar.Error(pw(args)) }
func (l loggerWrap) Errorz(fmt string, fields ...zap.Field) { l.log.Error(fmt, fields...) }

func (l loggerWrap) Fatalf(fmt string, args ...interface{}) { l.sugar.Fatalf(fmt, args...) }
func (l loggerWrap) Fatal(args ...interface{})              { l.sugar.Fatal(pw(args)) }
func (l loggerWrap) Fatalln(args ...interface{})            { l.sugar.Fatal(pw(args)) }
func (l loggerWrap) Fatalz(fmt string, fields ...zap.Field) { l.log.Fatal(fmt, fields...) }

func (l loggerWrap) V(v int) bool {
	if v <= 0 {
		return !l.Enabled(zapcore.DebugLevel)
	}
	return true
}
