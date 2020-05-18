/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"log"
	"runtime/debug"
	"sync"
)

type GoPool interface {
	Go(fn func(done <-chan struct{}))
	GoN(n int, fn func(done <-chan struct{}))
	Stop()
	Wait()
}

func NewGoPool() GoPool {
	return &goPool{
		done: make(chan struct{}),
	}
}

type goPool struct {
	wg   sync.WaitGroup
	once sync.Once
	done chan struct{}
}

func (that *goPool) Wait() { that.wg.Wait() }

func (that *goPool) Stop() {
	that.once.Do(func() {
		close(that.done)
	})
}

func (that *goPool) Go(fn func(done <-chan struct{})) {
	fnGo := that.withRecover(fn)
	that.wg.Add(1)
	go fnGo()
}

func (that *goPool) GoN(n int, fn func(done <-chan struct{})) {
	fnGo := that.withRecover(fn)
	that.wg.Add(n)
	for i := 0; i < n; i++ {
		go fnGo()
	}
}

func (that *goPool) withRecover(fn func(done <-chan none)) func() {
	return func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err, BytesToStrUnsafe(debug.Stack()))
			}
			that.wg.Done()
		}()

		fn(that.done)
	}
}
