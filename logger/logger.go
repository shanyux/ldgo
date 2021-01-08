/*
 * Copyright (C) distroy
 */

package logger

import (
	"go.uber.org/zap"
)

func newLogger(log loggerWrap) logger {
	return logger{
		loggerWrap: log,
	}
}

type Logger interface {
	Sync()

	Wrap() LoggerWrap
	Core() *zap.Logger
	Sugar() *zap.SugaredLogger

	With(fields ...zap.Field) Logger
	WithOptions(opts ...zap.Option) Logger

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

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	if len(fields) > 0 {
		log = log.With(fields...)
	}
	return newLogger(newLoggerWrap(log))
}

type logger struct {
	loggerWrap
}

func (l logger) With(fields ...zap.Field) Logger {
	log := l.log.With(fields...)
	return newLogger(newLoggerWrap(log))
}

func (l logger) WithOptions(opts ...zap.Option) Logger {
	log := l.log.WithOptions(opts...)
	return newLogger(newLoggerWrap(log))
}

func (l logger) Wrap() LoggerWrap { return l.loggerWrap }

func (l logger) Debug(msg string, fields ...zap.Field) { l.log.Debug(msg, fields...) }
func (l logger) Info(msg string, fields ...zap.Field)  { l.log.Info(msg, fields...) }
func (l logger) Warn(msg string, fields ...zap.Field)  { l.log.Warn(msg, fields...) }
func (l logger) Error(msg string, fields ...zap.Field) { l.log.Error(msg, fields...) }
func (l logger) Fatal(msg string, fields ...zap.Field) { l.log.Fatal(msg, fields...) }
