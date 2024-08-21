/*
 * Copyright (C) distroy
 */

package ldctx

import (
	"context"

	"github.com/distroy/ldgo/v2/ldlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKeyType int

const (
	ctxKeyLogger ctxKeyType = iota
	ctxKeyMap
)

var (
	defaultContext = context.Background()
	consoleContext = WithLogger(context.Background(), ldlog.Console())
	discardContext = WithLogger(context.Background(), ldlog.Discard())
)

func defaultLogger() *ldlog.Logger { return ldlog.Default() }

type stringer interface {
	String() string
}

func zCore(c context.Context, lvl zapcore.Level, skip int) *zap.Logger {
	l := GetLogger(c)
	if !l.Enabler().Enable(lvl, skip+1) {
		l = ldlog.Discard()
	}
	return l.Core()
}
func zSugar(c context.Context, lvl zapcore.Level, skip int) *zap.SugaredLogger {
	l := GetLogger(c)
	if !l.Enabler().Enable(lvl, skip+1) {
		l = ldlog.Discard()
	}
	return l.Sugar()
}

const (
	lvlD = zapcore.DebugLevel
	lvlI = zapcore.InfoLevel
	lvlW = zapcore.WarnLevel
	lvlE = zapcore.ErrorLevel
	lvlF = zapcore.FatalLevel
)

func LogD(c Context, msg string, fields ...zap.Field) { zCore(c, lvlD, 1).Debug(msg, fields...) }
func LogI(c Context, msg string, fields ...zap.Field) { zCore(c, lvlI, 1).Info(msg, fields...) }
func LogW(c Context, msg string, fields ...zap.Field) { zCore(c, lvlW, 1).Warn(msg, fields...) }
func LogE(c Context, msg string, fields ...zap.Field) { zCore(c, lvlE, 1).Error(msg, fields...) }
func LogF(c Context, msg string, fields ...zap.Field) { zCore(c, lvlF, 1).Fatal(msg, fields...) }

func LogDf(c Context, fmt string, args ...interface{}) { zSugar(c, lvlD, 1).Debugf(fmt, args...) }
func LogIf(c Context, fmt string, args ...interface{}) { zSugar(c, lvlI, 1).Infof(fmt, args...) }
func LogWf(c Context, fmt string, args ...interface{}) { zSugar(c, lvlW, 1).Warnf(fmt, args...) }
func LogEf(c Context, fmt string, args ...interface{}) { zSugar(c, lvlE, 1).Errorf(fmt, args...) }
func LogFf(c Context, fmt string, args ...interface{}) { zSugar(c, lvlF, 1).Fatalf(fmt, args...) }
