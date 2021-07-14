/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldgopool"

type AsyncPoolConfig = ldgopool.AsyncPoolConfig
type AsyncPool = ldgopool.AsyncPool

func NewAsyncPool(cfg *AsyncPoolConfig) AsyncPool {
	return ldgopool.NewAsyncPool(cfg)
}
