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

type AsyncPool interface {
	Async() chan<- func()
	Stop()
	Wait()
}

func NewAsyncPool(num int) AsyncPool {
	ap := &asyncPool{
		ch: make(chan func(), 256),
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

func (that *asyncPool) Stop() {
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
		that.withRecover(fn)
	}
}

func (that *asyncPool) withRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
		}
		that.wg.Done()
	}()

	fn()
}