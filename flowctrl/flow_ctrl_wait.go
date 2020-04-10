/*
 * Copyright (C) distroy
 */

package flowctrl

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

func newFlowCtrlWait() *flowCtrlWait {
	return &flowCtrlWait{
		ch: make(chan *flowCtrlCall, 32),
		n:  0,
	}
}

type flowCtrlWait struct {
	ch chan *flowCtrlCall
	n  int32
}

func (that *flowCtrlWait) sleep(c *flowCtrlCall) {
	c.n = 0
	atomic.AddInt32(&that.n, 1)
	that.ch <- c
	<-c.ch
}

func (that *flowCtrlWait) awake() {
	n := atomic.SwapInt32(&that.n, 0)
	for i := int32(0); i < n; i++ {
		c := <-that.ch
		select {
		case c.ch <- struct{}{}:
		default:
		}
	}
}
