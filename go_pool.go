/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

type GoPool = ldcore.GoPool

func NewGoPool() GoPool {
	return ldcore.NewGoPool()
}

func GoPoolGo(fn func(done <-chan struct{})) GoPool { return ldcore.GoPoolGo(fn) }
func GoPoolRun(fn func()) GoPool                    { return ldcore.GoPoolRun(fn) }

func GoPoolGoN(n int, fn func(done <-chan struct{})) GoPool { return ldcore.GoPoolGoN(n, fn) }

func GoPoolRunN(n int, fn func()) GoPool { return ldcore.GoPoolRunN(n, fn) }
