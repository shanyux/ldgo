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

func zCore(c context.Context) *zap.Logger         { return GetLogger(c).Core() }
func zSugar(c context.Context) *zap.SugaredLogger { return GetLogger(c).Sugar() }

func LogD(c context.Context, msg string, fields ...zap.Field) { zCore(c).Debug(msg, fields...) }
func LogI(c context.Context, msg string, fields ...zap.Field) { zCore(c).Info(msg, fields...) }
func LogW(c context.Context, msg string, fields ...zap.Field) { zCore(c).Warn(msg, fields...) }
func LogE(c context.Context, msg string, fields ...zap.Field) { zCore(c).Error(msg, fields...) }
func LogF(c context.Context, msg string, fields ...zap.Field) { zCore(c).Fatal(msg, fields...) }

func LogDf(c context.Context, fmt string, args ...interface{}) { zSugar(c).Debugf(fmt, args...) }
func LogIf(c context.Context, fmt string, args ...interface{}) { zSugar(c).Infof(fmt, args...) }
func LogWf(c context.Context, fmt string, args ...interface{}) { zSugar(c).Warnf(fmt, args...) }
func LogEf(c context.Context, fmt string, args ...interface{}) { zSugar(c).Errorf(fmt, args...) }
func LogFf(c context.Context, fmt string, args ...interface{}) { zSugar(c).Fatalf(fmt, args...) }
