/*
 * Copyright (C) distroy
 */

package internal

import (
	"context"

	"github.com/distroy/ldgo/v2/ldctx"
)

type ctxKey int

const (
	ctxKeyInProcess ctxKey = iota
)

func InProcess(c context.Context) bool {
	b, _ := c.Value(ctxKeyInProcess).(bool)
	return b
}

func NewContext(c context.Context) ldctx.Context {
	ctx := ldctx.New(c)
	ctx = ldctx.WithValue(ctx, ctxKeyInProcess, true)
	return ctx
}
