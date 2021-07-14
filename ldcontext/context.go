/*
 * Copyright (C) distroy
 */

package ldcontext

import (
	"context"
	"time"

	"github.com/distroy/ldgo/ldlogger"
	"go.uber.org/zap"
)

var (
	defaultContext = ctx{Context: context.Background()}
	consoleContext = Default().WithLogger(ldlogger.Console())
)

func defaultLogger() ldlogger.Logger { return ldlogger.Default() }

func Default() Context { return defaultContext }
func Console() Context { return consoleContext }

type cancelFunc = context.CancelFunc

type Context interface {
	context.Context

	WithValue(key, val interface{}) Context

	TryCancel() bool
	WithCancel() Context
	WithDeadline(deadline time.Time) Context
	WithTimeout(timeout time.Duration) Context

	WithLogger(l ldlogger.Logger) Context
	GetLogger() ldlogger.Logger

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

	cancel cancelFunc
}

func newContext(c context.Context) ctx { return ctx{Context: c} }

func (c ctx) WithValue(key, val interface{}) Context {
	c.Context = context.WithValue(c.Context, key, val)
	return c
}

func (c ctx) With(fields ...zap.Field) Context {
	c.Context = ldlogger.NewContext(c.Context, c.logger().With(fields...))
	return c
}

func (c ctx) TryCancel() bool {
	if c.cancel == nil {
		return false
	}

	c.cancel()
	return true
}

func (c ctx) WithDeadline(d time.Time) Context {
	if c.cancel != nil {
		if cur, ok := c.Deadline(); ok && cur.Before(d) {
			// The current deadline is already sooner than the new one.
			return c
		}
	}

	ctx, cancel := context.WithDeadline(c.Context, d)
	c.Context = ctx
	c.cancel = cancel
	return c
}

func (c ctx) WithTimeout(t time.Duration) Context {
	return c.WithDeadline(time.Now().Add(t))
}

func (c ctx) WithCancel() Context {
	if c.cancel != nil {
		return c
	}

	ctx, cancel := context.WithCancel(c.Context)
	c.Context = ctx
	c.cancel = cancel
	return c
}

func (c ctx) WithLogger(l ldlogger.Logger) Context {
	c.Context = ldlogger.NewContext(c.Context, l)
	return c
}

func (c ctx) GetLogger() ldlogger.Logger {
	return c.logger()
}

func (c ctx) logger() ldlogger.Logger {
	l := ldlogger.FromContext(c)
	if l == nil || l == ldlogger.Default() {
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
