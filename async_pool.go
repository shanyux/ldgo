/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

type AsyncPoolConfig = ldcore.AsyncPoolConfig
type AsyncPool = ldcore.AsyncPool

func NewAsyncPool(cfg *AsyncPoolConfig) AsyncPool {
	return ldcore.NewAsyncPool(cfg)
}
