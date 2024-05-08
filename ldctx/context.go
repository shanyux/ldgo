/*
 * Copyright (C) distroy
 */

package ldctx

import (
	"context"
	"reflect"
	"time"

	"github.com/distroy/ldgo/v2/lderr"
	"github.com/distroy/ldgo/v2/ldlog"
	"go.uber.org/zap"
)

type (
	StdContext = context.Context
	CancelFunc = context.CancelFunc
)

func Default() Context { return defaultContext }
func Console() Context { return consoleContext }
func Discard() Context { return discardContext }

type Context interface {
	context.Context

	LogD(msg string, fields ...zap.Field)
	LogI(msg string, fields ...zap.Field)
	LogW(msg string, fields ...zap.Field)
	LogE(msg string, fields ...zap.Field)
	LogF(msg string, fields ...zap.Field)

	LogDf(fmt string, args ...interface{})
	LogIf(fmt string, args ...interface{})
	LogWf(fmt string, args ...interface{})
	LogEf(fmt string, args ...interface{})
	LogFf(fmt string, args ...interface{})
}

func New(parent context.Context, fields ...zap.Field) Context {
	if len(fields) > 0 {
		l := GetLogger(parent).With(fields...)
		parent = unwrap(parent)
		return newCtx(newLogCtx(parent, l))
	}

	if parent == nil {
		return Default()
	}
	if cc, _ := parent.(Context); cc != nil {
		return cc
	}
	return newCtx(parent)
}

func ContextName(c context.Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	return reflect.TypeOf(c).String()
}

func GetError(c StdContext) lderr.Error {
	e := c.Err()
	switch e {
	case nil:
		return nil

	case context.Canceled:
		return lderr.ErrCtxCanceled

	case context.DeadlineExceeded:
		return lderr.ErrCtxDeadlineExceeded
	}

	if err, ok := e.(lderr.Error); ok {
		return err
	}

	return lderr.Wrap(e)
}

func WithLogger(parent context.Context, log *ldlog.Logger, fields ...zap.Field) Context {
	if log == nil {
		return New(parent, fields...)
	}
	log = log.With(fields...)
	return newCtx(newLogCtx(unwrap(parent), log))
}

func GetLogger(c context.Context) *ldlog.Logger {
	log, ok := c.Value(ctxKeyLogger).(*ldlog.Logger)
	if !ok {
		log = defaultLogger()
	}
	return log
}

func WithValue(parent context.Context, key, val interface{}) Context {
	return newCtx(context.WithValue(unwrap(parent), key, val))
}

func WithCancel(parent context.Context) Context {
	child, cancel := context.WithCancel(unwrap(parent))
	return newCtx(newCancelCtx(child, cancel))
}

func GetCancelFunc(c context.Context) CancelFunc {
	cancel, _ := c.Value(ctxKeyCancel).(CancelFunc)
	return cancel
}

func TryCancel(c context.Context) bool {
	cancel := GetCancelFunc(c)
	if cancel == nil {
		return false
	}
	cancel()
	return true
}

func WithTimeout(parent context.Context, timeout time.Duration) Context {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent context.Context, deadline time.Time) Context {
	child, cancel := context.WithDeadline(unwrap(parent), deadline)
	return newCtx(newCancelCtx(child, cancel))
}
