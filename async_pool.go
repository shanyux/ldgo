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

type AsyncPoolConfig struct {
	GoroutineNum int
	ChannelSize  int
}

type AsyncPool interface {
	Async() chan<- func()
	Close()
	Wait()
}

func NewAsyncPool(cfg *AsyncPoolConfig) AsyncPool {
	size := cfg.ChannelSize
	num := cfg.GoroutineNum
	ap := &asyncPool{
		ch: make(chan func(), size),
	}

	ap.wg.Add(num)
	for i := 0; i < num; i++ {
		go ap.main()
	}

	return ap
}

type asyncPool struct {
	once sync.Once
	wg   sync.WaitGroup
	ch   chan func()
}

func (that *asyncPool) Wait() { that.wg.Wait() }

func (that *asyncPool) Close() {
	that.once.Do(func() {
		close(that.ch)
	})
}

func (that *asyncPool) Async() chan<- func() {
	return that.ch
}

func (that *asyncPool) main() {
	defer that.wg.Done()

	for fn := range that.ch {
		that.doWithRecover(fn)
	}
}

func (that *asyncPool) doWithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	fn()
}
