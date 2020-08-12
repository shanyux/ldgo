/*
 * Copyright (C) distroy
 */

package logger

import "context"

type ctxLogKey int

const (
	_ContextKeyLogger ctxLogKey = iota
)

var (
	Default = NewLogger()
	console = NewLogger()
)

func Console() Logger { return console }

func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, _ContextKeyLogger, l)
}

func FromContext(ctx context.Context) Logger {
	l, ok := ctx.Value(_ContextKeyLogger).(Logger)
	if ok {
		return l
	}
	return Default
}
