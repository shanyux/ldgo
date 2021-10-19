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

func newLoggerWrapper(log *zap.Logger) wrapper {
	return wrapper{
		log:   log,
		sugar: log.Sugar(),
	}
}

type wrapper struct {
	log   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l wrapper) Sync() { l.log.Sync() }

func (l wrapper) Enabled(lvl zapcore.Level) bool { return l.enabled(lvl) }

func (l wrapper) Core() *zap.Logger         { return l.log }
func (l wrapper) Sugar() *zap.SugaredLogger { return l.sugar }

func (l wrapper) enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }

func (l wrapper) Debugf(fmt string, args ...interface{}) { l.sugar.Debugf(fmt, args...) }
func (l wrapper) Debug(args ...interface{})              { l.sugar.Debug(pw(args)) }
func (l wrapper) Debugln(args ...interface{})            { l.sugar.Debug(pw(args)) }
func (l wrapper) Debugz(fmt string, fields ...zap.Field) { l.log.Debug(fmt, fields...) }

func (l wrapper) Infof(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l wrapper) Info(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l wrapper) Infoln(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l wrapper) Infoz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l wrapper) Printf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l wrapper) Print(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l wrapper) Println(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l wrapper) Printz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l wrapper) Logf(fmt string, args ...interface{}) { l.sugar.Infof(fmt, args...) }
func (l wrapper) Log(args ...interface{})              { l.sugar.Info(pw(args)) }
func (l wrapper) Logln(args ...interface{})            { l.sugar.Info(pw(args)) }
func (l wrapper) Logz(fmt string, fields ...zap.Field) { l.log.Info(fmt, fields...) }

func (l wrapper) Warnf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l wrapper) Warn(args ...interface{})              { l.sugar.Warn(pw(args)) }
func (l wrapper) Warnln(args ...interface{})            { l.sugar.Warn(pw(args)) }
func (l wrapper) Warnz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l wrapper) Warningf(fmt string, args ...interface{}) { l.sugar.Warnf(fmt, args...) }
func (l wrapper) Warning(args ...interface{})              { l.sugar.Warn(pw(args)) }
func (l wrapper) Warningln(args ...interface{})            { l.sugar.Warn(pw(args)) }
func (l wrapper) Warningz(fmt string, fields ...zap.Field) { l.log.Warn(fmt, fields...) }

func (l wrapper) Errorf(fmt string, args ...interface{}) { l.sugar.Errorf(fmt, args...) }
func (l wrapper) Error(args ...interface{})              { l.sugar.Error(pw(args)) }
func (l wrapper) Errorln(args ...interface{})            { l.sugar.Error(pw(args)) }
func (l wrapper) Errorz(fmt string, fields ...zap.Field) { l.log.Error(fmt, fields...) }

func (l wrapper) Fatalf(fmt string, args ...interface{}) { l.sugar.Fatalf(fmt, args...) }
func (l wrapper) Fatal(args ...interface{})              { l.sugar.Fatal(pw(args)) }
func (l wrapper) Fatalln(args ...interface{})            { l.sugar.Fatal(pw(args)) }
func (l wrapper) Fatalz(fmt string, fields ...zap.Field) { l.log.Fatal(fmt, fields...) }

func (l wrapper) V(v int) bool {
	if v <= 0 {
		return !l.Enabled(zapcore.DebugLevel)
	}
	return true
}
