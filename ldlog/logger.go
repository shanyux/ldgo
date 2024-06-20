/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
)

func newLogger(log core) *Logger {
	return &Logger{
		core: log,
	}
}

type Logger struct {
	core
}

func (l *Logger) Wrapper() *Wrapper { return (*Wrapper)(l) }

func (l *Logger) clone() *Logger {
	cp := *l
	return &cp
}

func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	log := l.Core().WithOptions(opts...)
	l = l.clone()
	l.log = log
	l.sugar = log.Sugar()
	return l
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return l
	}
	log := l.Core().With(fields...)
	l = l.clone()
	l.log = log
	l.sugar = log.Sugar()
	return l
}

func (l *Logger) GetSequence() string { return l.sequence }
func (l *Logger) WithSequence(seq string) *Logger {
	if seq == "" {
		return l
	}
	log := l.Core().With(zap.String(sequenceKey, seq))
	l = l.clone()
	l.log = log
	l.sugar = log.Sugar()
	l.sequence = seq
	return l
}

func (l *Logger) Debug(msg string, fields ...zap.Field) { l.zCore(1).Debug(msg, fields...) }
func (l *Logger) Info(msg string, fields ...zap.Field)  { l.zCore(1).Info(msg, fields...) }
func (l *Logger) Warn(msg string, fields ...zap.Field)  { l.zCore(1).Warn(msg, fields...) }
func (l *Logger) Error(msg string, fields ...zap.Field) { l.Core().Error(msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...zap.Field) { l.Core().Fatal(msg, fields...) }

func (l *Logger) Debugf(fmt string, args ...interface{}) { l.zSugar(1).Debugf(fmt, args...) }
func (l *Logger) Infof(fmt string, args ...interface{})  { l.zSugar(1).Infof(fmt, args...) }
func (l *Logger) Warnf(fmt string, args ...interface{})  { l.zSugar(1).Warnf(fmt, args...) }
func (l *Logger) Errorf(fmt string, args ...interface{}) { l.Sugar().Errorf(fmt, args...) }
func (l *Logger) Fatalf(fmt string, args ...interface{}) { l.Sugar().Fatalf(fmt, args...) }

func (l *Logger) Debugln(args ...interface{}) { l.zSugar(1).Debug(pw(args)) }
func (l *Logger) Infoln(args ...interface{})  { l.zSugar(1).Info(pw(args)) }
func (l *Logger) Warnln(args ...interface{})  { l.zSugar(1).Warn(pw(args)) }
func (l *Logger) Errorln(args ...interface{}) { l.Sugar().Error(pw(args)) }
func (l *Logger) Fatalln(args ...interface{}) { l.Sugar().Fatal(pw(args)) }
