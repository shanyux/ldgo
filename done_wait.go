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

type DoneWait interface {
	Stop()
	Chan() <-chan none

	Add(n int)
	Done()
	Wait()
}

func NewDoneWait() DoneWait {
	return &doneWait{
		ch: make(chan struct{}),
	}
}

type doneWait struct {
	sync.WaitGroup

	ch   chan none
	once sync.Once
}

func (that *doneWait) Stop() {
	that.once.Do(func() { close(that.ch) })
}
func (that *doneWait) Chan() <-chan none {
	return that.ch
}
