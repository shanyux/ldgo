/*
 * Copyright (C) distroy
 */

package ldcontext

import (
	"context"
	"reflect"
	"time"

	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldlogger"
	"go.uber.org/zap"
)

type CancelFunc = context.CancelFunc

func Default() Context { return defaultContext }
func Console() Context { return consoleContext }

type Context interface {
	context.Context

	LogD(msg string, fields ...zap.Field)
	LogI(msg string, fields ...zap.Field)
	LogW(msg string, fields ...zap.Field)
	LogE(msg string, fields ...zap.Field)

	LogDf(fmt string, args ...interface{})
	LogIf(fmt string, args ...interface{})
	LogWf(fmt string, args ...interface{})
	LogEf(fmt string, args ...interface{})
}

func NewContext(parent context.Context, fields ...zap.Field) Context {
	if len(fields) == 0 {
		if c, ok := parent.(Context); ok {
			return c
		}

		return Default()
	}

	if parent == nil {
		return newCtx(newLogCtx(Default(), ldlogger.With(defaultLogger(), fields...)))
	}

	return newCtx(newLogCtx(parent, ldlogger.With(GetLogger(parent), fields...)))
}

func ContextName(c context.Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	return reflect.TypeOf(c).String()
}

func GetError(c context.Context) lderr.Error {
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

func WithLogger(parent context.Context, log ldlogger.Logger, fields ...zap.Field) Context {
	if log == nil {
		log = GetLogger(parent)
	}
	log = ldlogger.With(log, fields...)
	return newCtx(newLogCtx(parent, log))
}

func GetLogger(c context.Context) ldlogger.Logger {
	log, ok := c.Value(ctxKeyLogger).(ldlogger.Logger)
	if !ok {
		log = defaultLogger()
	}
	return log
}

func WithValue(parent context.Context, key, val interface{}) Context {
	return newCtx(context.WithValue(parent, key, val))
}

func WithCancel(parent context.Context) Context {
	child, cancel := context.WithCancel(parent)
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
	if cur, ok := parent.Deadline(); ok && cur.Before(deadline) {
		// The current deadline is already sooner than the new one.
		return NewContext(parent)
	}

	child, cancel := context.WithDeadline(parent, deadline)
	return newCtx(newCancelCtx(child, cancel))
}
