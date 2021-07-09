/*
 * Copyright (C) distroy
 */

package context

import (
	"context"

	"github.com/distroy/ldgo/ldcontext"
	"go.uber.org/zap"
)

func Default() Context { return ldcontext.Default() }
func Console() Context { return ldcontext.Console() }

type Context = ldcontext.Context

func NewContext(c context.Context, fields ...zap.Field) Context {
	return ldcontext.NewContext(c, fields...)
}
