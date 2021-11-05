/*
 * Copyright (C) distroy
 */

package ldctx

import (
	"context"

	"github.com/distroy/ldgo/ldlog"
	"go.uber.org/zap"
)

type ctxKeyType int

const (
	ctxKeyLogger ctxKeyType = iota
	ctxKeyCancel
)

var (
	defaultContext Context = newCtx(context.Background())
	consoleContext Context = newCtx(newLogCtx(context.Background(), ldlog.Console()))
	discardContext Context = newCtx(newLogCtx(context.Background(), ldlog.Discard()))
)

func defaultLogger() *ldlog.Logger { return ldlog.Default() }

type stringer interface {
	String() string
}

type ctx struct {
	context.Context
}

func newCtx(c context.Context) ctx { return ctx{Context: c} }

func (c ctx) String() string { return ContextName(c.Context) }

func (c ctx) logger() *ldlog.Logger      { return GetLogger(c.Context) }
func (c ctx) zCore() *zap.Logger         { return c.logger().Core() }
func (c ctx) zSugar() *zap.SugaredLogger { return c.logger().Sugar() }

// func (c ctx) GetLogger() ldlog.Logger { return c.logger() }
// func (c ctx) LogSync()                   { c.logger().Sync() }

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

type logCtx struct {
	context.Context

	logger *ldlog.Logger
}

func newLogCtx(parent context.Context, log *ldlog.Logger) *logCtx {
	return &logCtx{
		Context: parent,
		logger:  log,
	}
}

func (c *logCtx) String() string { return ContextName(c.Context) + ".WithLogger" }

func (c *logCtx) Value(key interface{}) interface{} {
	if key == ctxKeyLogger {
		return c.logger
	}
	return c.Context.Value(key)
}

type cancelCtx struct {
	context.Context

	cancel CancelFunc
}

func newCancelCtx(parent context.Context, cancel CancelFunc) *cancelCtx {
	return &cancelCtx{
		Context: parent,
		cancel:  cancel,
	}
}

func (c *cancelCtx) String() string { return ContextName(c.Context) }

func (c *cancelCtx) Value(key interface{}) interface{} {
	if key == ctxKeyCancel {
		return c.cancel
	}
	return c.Context.Value(key)
}
