/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldgopool"

type GoPool = ldgopool.GoPool

func NewGoPool() GoPool {
	return ldgopool.NewGoPool()
}

func GoPoolGo(fn func(done <-chan struct{})) GoPool { return ldgopool.GoPoolGo(fn) }
func GoPoolRun(fn func()) GoPool                    { return ldgopool.GoPoolRun(fn) }

func GoPoolGoN(n int, fn func(done <-chan struct{})) GoPool { return ldgopool.GoPoolGoN(n, fn) }

func GoPoolRunN(n int, fn func()) GoPool { return ldgopool.GoPoolRunN(n, fn) }
