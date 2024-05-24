/*
 * Copyright (C) distroy
 */

package ldctx

import (
	"context"

	"github.com/distroy/ldgo/v2/ldlog"
	"go.uber.org/zap"
)

type ctxKeyType int

const (
	ctxKeyLogger ctxKeyType = iota
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

func zCore(c context.Context, skip int) *zap.Logger {
	l := GetLogger(c)
	if !l.CheckRateOrInterval(skip + 1) {
		l = ldlog.Discard()
	}
	return l.Core()
}
func zSugar(c context.Context, skip int) *zap.SugaredLogger {
	l := GetLogger(c)
	if !l.CheckRateOrInterval(skip + 1) {
		l = ldlog.Discard()
	}
	return l.Sugar()
}

func lCore(c context.Context) *zap.Logger         { return GetLogger(c).Core() }
func lSugar(c context.Context) *zap.SugaredLogger { return GetLogger(c).Sugar() }

func LogD(c context.Context, msg string, fields ...zap.Field) { zCore(c, 1).Debug(msg, fields...) }
func LogI(c context.Context, msg string, fields ...zap.Field) { zCore(c, 1).Info(msg, fields...) }
func LogW(c context.Context, msg string, fields ...zap.Field) { zCore(c, 1).Warn(msg, fields...) }
func LogE(c context.Context, msg string, fields ...zap.Field) { lCore(c).Error(msg, fields...) }
func LogF(c context.Context, msg string, fields ...zap.Field) { lCore(c).Fatal(msg, fields...) }

func LogDf(c context.Context, fmt string, args ...interface{}) { zSugar(c, 1).Debugf(fmt, args...) }
func LogIf(c context.Context, fmt string, args ...interface{}) { zSugar(c, 1).Infof(fmt, args...) }
func LogWf(c context.Context, fmt string, args ...interface{}) { zSugar(c, 1).Warnf(fmt, args...) }
func LogEf(c context.Context, fmt string, args ...interface{}) { lSugar(c).Errorf(fmt, args...) }
func LogFf(c context.Context, fmt string, args ...interface{}) { lSugar(c).Fatalf(fmt, args...) }
