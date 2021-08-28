/*
 * Copyright (C) distroy
 */

package ldlogger

import (
	"go.uber.org/zap"
)

func newLogger(log wrapper) logger {
	return logger{
		wrapper: log,
	}
}

type Logger interface {
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

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	if len(fields) > 0 {
		log = log.With(fields...)
	}
	return newLogger(newLoggerWrapper(log))
}

func GetWrapper(l Logger) LoggerWrapper {
	return wrapper{
		log:   l.Core(),
		sugar: l.Sugar(),
	}
}

func With(l Logger, fields ...zap.Field) Logger {
	if len(fields) == 0 {
		return l
	}
	log := l.Core().With(fields...)
	return newLogger(newLoggerWrapper(log))
}

func WithOptions(l Logger, opts ...zap.Option) Logger {
	log := l.Core().WithOptions(opts...)
	return newLogger(newLoggerWrapper(log))
}

type logger struct {
	wrapper
}

func (l logger) Debug(msg string, fields ...zap.Field) { l.log.Debug(msg, fields...) }
func (l logger) Info(msg string, fields ...zap.Field)  { l.log.Info(msg, fields...) }
func (l logger) Warn(msg string, fields ...zap.Field)  { l.log.Warn(msg, fields...) }
func (l logger) Error(msg string, fields ...zap.Field) { l.log.Error(msg, fields...) }
func (l logger) Fatal(msg string, fields ...zap.Field) { l.log.Fatal(msg, fields...) }
