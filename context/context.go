/*
 * Copyright (C) distroy
 */

package context

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/distroy/ldgo/logger"

	"go.uber.org/zap"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = context.Canceled
var _ = fmt.Sprintf
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

type CancelFunc = func()

type Context interface {
	context.Context

	Root() Context

	WithCancel() (Context, CancelFunc)
	WithDeadline(deadline time.Time) (Context, CancelFunc)

	Clone() Context

	Logger(log ...logger.Logger) logger.Logger

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
	if c != nil {
		switch _c := c.(type) {
		case ctx:
			return _c.With(fields...)
		case *ctx:
			return _c.With(fields...)
		}
	}
	if c == nil {
		c = context.Background()
	}

	return &ctx{
		Context: c,
		log:     newLogger(DefaultLogger.With(fields...)),
	}
}

func newContext(c context.Context, _c *ctx) *ctx {
	return &ctx{
		Context: c,
		log:     _c.log,
		root:    _c.first(),
	}
}

type ctx struct {
	context.Context
	log

	root *ctx
}

func (that *ctx) Root() Context { return that.first() }
func (that *ctx) first() *ctx {
	if that.root != nil {
		return that.root
	}
	return that
}

func (that *ctx) Clone() Context { return that.clone() }
func (that *ctx) clone() *ctx {
	ctx := *that
	ctx.root = that.first()
	return &ctx
}

func (that *ctx) With(fields ...zap.Field) Context {
	ctx := that.clone()
	ctx.log = ctx.log.With(fields...)
	return ctx
}

func (that *ctx) WithDeadline(deadline time.Time) (Context, CancelFunc) {
	ctx, cancel := context.WithDeadline(that.Context, deadline)
	return newContext(ctx, that), cancel
}

func (that *ctx) WithCancel() (Context, CancelFunc) {
	ctx, cancel := context.WithCancel(that.Context)
	return newContext(ctx, that), cancel
}

func (that *ctx) Logger(log ...logger.Logger) logger.Logger {
	if len(log) != 0 && log[0] != nil {
		that.log = newLogger(log[0])
	}
	return that.log.Logger
}
