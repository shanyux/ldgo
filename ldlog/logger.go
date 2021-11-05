/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
)

func newLogger(log Wrapper) *Logger {
	return &Logger{
		wrapper: log,
	}
}

func GetLogger(log *zap.Logger, fields ...zap.Field) *Logger {
	if len(fields) > 0 {
		log = log.With(fields...)
	}
	return newLogger(newWrapper(log, log.Sugar()))
}

type Logger struct {
	wrapper Wrapper
}

func (l *Logger) Sync() { l.wrapper.Sync() }

func (l *Logger) Core() *zap.Logger         { return l.wrapper.Core() }
func (l *Logger) Sugar() *zap.SugaredLogger { return l.wrapper.Sugar() }

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	log := l.Core().WithOptions(opts...)
	return newLogger(newWrapper(log, log.Sugar()))
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return l
	}
	log := l.Core().With(fields...)
	return newLogger(newWrapper(log, log.Sugar()))
}

func (l *Logger) Wrapper() *Wrapper {
	c := l.wrapper
	return &c
}

func (l *Logger) Debug(msg string, fields ...zap.Field) { l.Core().Debug(msg, fields...) }
func (l *Logger) Info(msg string, fields ...zap.Field)  { l.Core().Info(msg, fields...) }
func (l *Logger) Warn(msg string, fields ...zap.Field)  { l.Core().Warn(msg, fields...) }
func (l *Logger) Error(msg string, fields ...zap.Field) { l.Core().Error(msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...zap.Field) { l.Core().Fatal(msg, fields...) }

func (l *Logger) Debugf(fmt string, args ...interface{}) { l.Sugar().Debugf(fmt, args...) }
func (l *Logger) Infof(fmt string, args ...interface{})  { l.Sugar().Infof(fmt, args...) }
func (l *Logger) Warnf(fmt string, args ...interface{})  { l.Sugar().Warnf(fmt, args...) }
func (l *Logger) Errorf(fmt string, args ...interface{}) { l.Sugar().Errorf(fmt, args...) }
func (l *Logger) Fatalf(fmt string, args ...interface{}) { l.Sugar().Fatalf(fmt, args...) }

func (l *Logger) Debugln(args ...interface{}) { l.Sugar().Debug(pw(args)) }
func (l *Logger) Infoln(args ...interface{})  { l.Sugar().Info(pw(args)) }
func (l *Logger) Warnln(args ...interface{})  { l.Sugar().Warn(pw(args)) }
func (l *Logger) Errorln(args ...interface{}) { l.Sugar().Error(pw(args)) }
func (l *Logger) Fatalln(args ...interface{}) { l.Sugar().Fatal(pw(args)) }
