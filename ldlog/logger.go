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

type Logger struct {
	wrapper Wrapper
}

func (l *Logger) clone() *Logger {
	cp := *l
	return &cp
}

func (l *Logger) Sync() { l.wrapper.Sync() }

func (l *Logger) Core() *zap.Logger         { return l.wrapper.Core() }
func (l *Logger) Sugar() *zap.SugaredLogger { return l.wrapper.Sugar() }

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	log := l.Core().WithOptions(opts...)
	l = l.clone()
	l.wrapper.log = log
	l.wrapper.sugar = log.Sugar()
	return l
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return l
	}
	log := l.Core().With(fields...)
	l = l.clone()
	l.wrapper.log = log
	l.wrapper.sugar = log.Sugar()
	return l
}

func (l *Logger) GetSequence() string { return l.wrapper.sequence }
func (l *Logger) WithSequence(seq string) *Logger {
	if seq == "" {
		return l
	}
	log := l.Core().With(zap.String(sequenceKey, seq))
	l = l.clone()
	l.wrapper.log = log
	l.wrapper.sugar = log.Sugar()
	l.wrapper.sequence = seq
	return l
}

func (l *Logger) Wrapper() *Wrapper {
	return &l.wrapper
}

func (l *Logger) zCore(skip int) *zap.Logger         { return l.wrapper.zCore(skip + 1) }
func (l *Logger) zSugar(skip int) *zap.SugaredLogger { return l.wrapper.zSugar(skip + 1) }

func (l *Logger) Debug(msg string, fields ...zap.Field) { l.zCore(1).Debug(msg, fields...) }
func (l *Logger) Info(msg string, fields ...zap.Field)  { l.zCore(1).Info(msg, fields...) }
func (l *Logger) Warn(msg string, fields ...zap.Field)  { l.zCore(1).Warn(msg, fields...) }
func (l *Logger) Error(msg string, fields ...zap.Field) { l.zCore(1).Error(msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...zap.Field) { l.zCore(1).Fatal(msg, fields...) }

func (l *Logger) Debugf(fmt string, args ...interface{}) { l.zSugar(1).Debugf(fmt, args...) }
func (l *Logger) Infof(fmt string, args ...interface{})  { l.zSugar(1).Infof(fmt, args...) }
func (l *Logger) Warnf(fmt string, args ...interface{})  { l.zSugar(1).Warnf(fmt, args...) }
func (l *Logger) Errorf(fmt string, args ...interface{}) { l.zSugar(1).Errorf(fmt, args...) }
func (l *Logger) Fatalf(fmt string, args ...interface{}) { l.zSugar(1).Fatalf(fmt, args...) }

func (l *Logger) Debugln(args ...interface{}) { l.zSugar(1).Debug(pw(args)) }
func (l *Logger) Infoln(args ...interface{})  { l.zSugar(1).Info(pw(args)) }
func (l *Logger) Warnln(args ...interface{})  { l.zSugar(1).Warn(pw(args)) }
func (l *Logger) Errorln(args ...interface{}) { l.zSugar(1).Error(pw(args)) }
func (l *Logger) Fatalln(args ...interface{}) { l.zSugar(1).Fatal(pw(args)) }
