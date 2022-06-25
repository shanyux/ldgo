/*
 * Copyright (C) distroy
 */

package ldcontext

import (
	"context"
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldlog"
	"go.uber.org/zap"
)

type (
	StdContext = context.Context
	CancelFunc = ldctx.CancelFunc
)

func Default() Context { return ldctx.Default() }
func Console() Context { return ldctx.Console() }
func Discard() Context { return ldctx.Discard() }

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
	return ldctx.NewContext(parent, fields...)
}

func ContextName(c context.Context) string { return ldctx.ContextName(c) }
func GetError(c StdContext) lderr.Error    { return ldctx.GetError(c) }

func WithLogger(parent context.Context, log ldlog.LoggerInterface, fields ...zap.Field) Context {
	if log != nil {
		return ldctx.WithLogger(parent, ldlog.GetLogger(log.Core()), fields...)
	}
	return ldctx.WithLogger(parent, ldctx.GetLogger(parent), fields...)
}

func GetLogger(c context.Context) ldlog.LoggerInterface { return ldctx.GetLogger(c) }

func WithValue(parent context.Context, key, val interface{}) Context {
	return ldctx.WithValue(parent, key, val)
}

func WithCancel(parent context.Context) Context  { return ldctx.WithCancel(parent) }
func GetCancelFunc(c context.Context) CancelFunc { return ldctx.GetCancelFunc(c) }
func TryCancel(c context.Context) bool           { return ldctx.TryCancel(c) }

func WithTimeout(parent context.Context, timeout time.Duration) Context {
	return ldctx.WithTimeout(parent, timeout)
}

func WithDeadline(parent context.Context, deadline time.Time) Context {
	return ldctx.WithDeadline(parent, deadline)
}
