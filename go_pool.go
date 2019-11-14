/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

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
			}
			that.wg.Done()
		}()

		fn(that.done)
	}
}
