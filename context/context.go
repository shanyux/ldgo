/*
 * Copyright (C) distroy
 */

package context

import (
	"context"

	"github.com/distroy/ldgo/ldctx"
	"go.uber.org/zap"
)

func Default() Context { return ldctx.Default() }
func Console() Context { return ldctx.Console() }

type Context = ldctx.Context

func NewContext(c context.Context, fields ...zap.Field) Context {
	return ldctx.NewContext(c, fields...)
}
