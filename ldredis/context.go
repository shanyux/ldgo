/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"

	"github.com/distroy/ldgo/v2/ldctx"
)

type ctxKey int

const (
	ctxKeyInProcess ctxKey = iota
)

func isInProcess(c context.Context) bool {
	b, _ := c.Value(ctxKeyInProcess).(bool)
	return b
}

func newContext(c context.Context) ldctx.Context {
	ctx := ldctx.New(c)
	ctx = ldctx.WithValue(ctx, ctxKeyInProcess, true)
	return ctx
}
