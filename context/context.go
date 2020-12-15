/*
 * Copyright (C) distroy
 */

package context

import (
	"context"
	"time"

	"github.com/distroy/ldgo/logger"

	"go.uber.org/zap"
)

var (
	defaultContext = ctx{Context: context.Background()}
	consoleContext = Default().WithLogger(logger.Console())
)

func defaultLogger() logger.Logger { return logger.Default }

func Default() Context { return defaultContext }
func Console() Context { return consoleContext }

type CancelFunc = context.CancelFunc

type Context interface {
	context.Context

	WithValue(key, val interface{}) Context
	WithCancel() (Context, CancelFunc)
	WithDeadline(deadline time.Time) (Context, CancelFunc)

	WithLogger(l logger.Logger) Context
	GetLogger() logger.Logger

	With(field ...zap.Field) Context
	LogSync()

	LogD(msg string, fields ...zap.Field)
	LogI(msg string, fields ...zap.Field)
	LogW(msg string, fields ...zap.Field)
	LogE(msg string, fields ...zap.Field)

	LogDf(fmt string, args ...interface{})
	LogIf(fmt string, args ...interface{})
	LogWf(fmt string, args ...interface{})
	LogEf(fmt string, args ...interface{})
}

func NewContext(c context.Context, fields ...zap.Field) Context {
	if c == nil {
		return Default().With(fields...)
	}
	switch _c := c.(type) {
	case ctx:
		return _c.With(fields...)
	case *ctx:
		return _c.With(fields...)
	}
	return newContext(c).With(fields...)
}

type ctx struct {
	context.Context
}

func newContext(c context.Context) ctx { return ctx{Context: c} }

func (c ctx) WithValue(key, val interface{}) Context {
	return newContext(context.WithValue(c.Context, key, val))
}

func (c ctx) With(fields ...zap.Field) Context {
	return newContext(logger.NewContext(c.Context, c.logger().With(fields...)))
}

func (c ctx) WithDeadline(deadline time.Time) (Context, CancelFunc) {
	ctx, cancel := context.WithDeadline(c.Context, deadline)
	return newContext(ctx), cancel
}

func (c ctx) WithCancel() (Context, CancelFunc) {
	ctx, cancel := context.WithCancel(c.Context)
	return newContext(ctx), cancel
}

func (c ctx) WithLogger(l logger.Logger) Context {
	return newContext(logger.NewContext(c.Context, l))
}

func (c ctx) GetLogger() logger.Logger {
	return c.logger()
}

func (c ctx) logger() logger.Logger {
	l := logger.FromContext(c)
	if l == nil || l == logger.Default {
		l = defaultLogger()
	}
	return l
}

func (c ctx) zCore() *zap.Logger         { return c.logger().Core() }
func (c ctx) zSugar() *zap.SugaredLogger { return c.logger().Sugar() }

func (c ctx) LogSync() { c.logger().Sync() }

func (c ctx) LogD(msg string, fields ...zap.Field) { c.zCore().Debug(msg, fields...) }
func (c ctx) LogI(msg string, fields ...zap.Field) { c.zCore().Info(msg, fields...) }
func (c ctx) LogW(msg string, fields ...zap.Field) { c.zCore().Warn(msg, fields...) }
func (c ctx) LogE(msg string, fields ...zap.Field) { c.zCore().Error(msg, fields...) }
func (c ctx) LogF(msg string, fields ...zap.Field) { c.zCore().Fatal(msg, fields...) }

func (c ctx) LogDf(fmt string, args ...interface{}) { c.zSugar().Debugf(fmt, args...) }
func (c ctx) LogIf(fmt string, args ...interface{}) { c.zSugar().Infof(fmt, args...) }
func (c ctx) LogWf(fmt string, args ...interface{}) { c.zSugar().Warnf(fmt, args...) }
func (c ctx) LogEf(fmt string, args ...interface{}) { c.zSugar().Errorf(fmt, args...) }
func (c ctx) LogFf(fmt string, args ...interface{}) { c.zSugar().Fatalf(fmt, args...) }
