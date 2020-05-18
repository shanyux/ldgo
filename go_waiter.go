/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"log"
	"runtime/debug"
	"sync"
)

type Waiter interface {
	Wait()
}

type goWaiter struct {
	wg sync.WaitGroup
}

func Go(fn func()) Waiter {
	return GoN(1, fn)
}

func GoN(n int, fn func()) Waiter {
	pool := &goWaiter{}
	fnGo := pool.withRecover(fn)
	for i := 0; i < n; i++ {
		fnGo()
	}
	return pool
}

func (that *goWaiter) Wait() { that.wg.Wait() }

func (that *goWaiter) withRecover(fn func()) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, BytesToStrUnsafe(debug.Stack()))
			}
			that.wg.Done()
		}()

		fn()
	}
}
