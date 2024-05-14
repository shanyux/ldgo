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
	Context    = context.Context
	CancelFunc = context.CancelFunc
)

func Default() context.Context { return defaultContext }
func Console() context.Context { return consoleContext }
func Discard() context.Context { return discardContext }

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

func WithLogger(parent context.Context, log *ldlog.Logger, fields ...zap.Field) context.Context {
	if log == nil && len(fields) == 0 {
		return parent
	}
	if log == nil {
		log = GetLogger(parent)
	}
	log = log.With(fields...)
	return WithValue(parent, ctxKeyLogger, log)
}

func GetLogger(c context.Context) *ldlog.Logger {
	log, ok := c.Value(ctxKeyLogger).(*ldlog.Logger)
	if !ok {
		log = defaultLogger()
	}
	return log
}

func WithValue(parent context.Context, key, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

func WithCancel(parent context.Context) (context.Context, CancelFunc) {
	return context.WithCancel(parent)
}

func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent context.Context, deadline time.Time) (context.Context, CancelFunc) {
	return context.WithDeadline(parent, deadline)
}
