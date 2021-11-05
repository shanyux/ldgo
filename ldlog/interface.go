/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_ LoggerInterface  = (*Logger)(nil)
	_ WrapperInterface = (*Wrapper)(nil)
)

type LoggerInterface interface {
	Sync()

	Core() *zap.Logger
	Sugar() *zap.SugaredLogger

	// Debugf formats according to a format specifier and print the logger
	Debugf(fmt string, args ...interface{})
	Debug(msg string, fields ...zap.Field)
	Debugln(args ...interface{})

	// Infof formats according to a format specifier and print the logger
	Infof(fmt string, args ...interface{})
	Info(msg string, fields ...zap.Field)
	Infoln(args ...interface{})

	// Warnf formats according to a format specifier and print the logger
	Warnf(fmt string, args ...interface{})
	Warn(msg string, fields ...zap.Field)
	Warnln(args ...interface{})

	// Errorf formats according to a format specifier and print the logger
	Errorf(fmt string, args ...interface{})
	Error(msg string, fields ...zap.Field)
	Errorln(args ...interface{})

	// Fatalf formats according to a format specifier and print the logger
	Fatalf(fmt string, args ...interface{})
	Fatal(msg string, fields ...zap.Field)
	Fatalln(args ...interface{})
}

type WrapperInterface interface {
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

func GetWrapper(l LoggerInterface) *Wrapper {
	w := newWrapper(l.Core(), l.Sugar())
	return &w
}

func With(l LoggerInterface, fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return newLogger(newWrapper(l.Core(), l.Sugar()))
	}
	return GetLogger(l.Core(), fields...)
}

func WithOptions(l LoggerInterface, opts ...zap.Option) *Logger {
	log := l.Core().WithOptions(opts...)
	return newLogger(newWrapper(log, log.Sugar()))
}
