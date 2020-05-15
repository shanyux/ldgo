/*
 * Copyright (C) distroy
 */

package flowctrl

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	_Pool = &sync.Pool{New: func() interface{} { return newFlowCtrlCall() }}
)

type FlowCtrl interface {
	Close()

	Apply(n int64) int64
	TryApply(n int64) int64
}

func New(limit int64, interval ...time.Duration) FlowCtrl {
	if len(interval) == 0 {
		interval = []time.Duration{time.Second}
	} else if interval[0] <= 0 {
		interval[0] = time.Second
	}

	fc := &flowCtrl{}
	fc.limit = limit
	fc.wait = newFlowCtrlWait()
	fc.done = make(chan struct{})

	go fc.tickerGoroutine(interval[0])
	return fc
}

type flowCtrl struct {
	once  sync.Once
	done  chan struct{}
	limit int64
	flow  int64
	wait  *flowCtrlWait
}

func (that *flowCtrl) Close() {
	that.once.Do(func() {
		close(that.done)
	})
}

func (that *flowCtrl) tickerGoroutine(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-that.done:
			return

		case <-ticker.C:
			atomic.StoreInt64(&that.flow, 0)
			that.wait.awake()
		}
	}
}

func (that *flowCtrl) popCall() *flowCtrlCall {
	c := _Pool.Get().(*flowCtrlCall)
	return c
}

func (that *flowCtrl) pushCall(c *flowCtrlCall) {
	c.n = 0
	_Pool.Put(c)
}

func (that *flowCtrl) Apply(n int64) int64 {
	if n <= 0 || that.limit <= 0 {
		return n
	}

	call := that.popCall()
	r := that.apply(call, n)
	that.pushCall(call)
	return r
}

func (that *flowCtrl) TryApply(n int64) int64 {
	if n <= 0 || that.limit <= 0 {
		return n
	}

	return that.tryApply(n)
}

func (that *flowCtrl) apply(call *flowCtrlCall, n int64) int64 {
	for {
		if r := that.tryApply(n); r > 0 {
			call.n += r
			return r
		}

		that.wait.sleep(call)
	}
}

func (that *flowCtrl) tryApply(n int64) int64 {
	limit := that.limit
	if atomic.LoadInt64(&that.flow) >= limit {
		return 0
	}

	flow := atomic.AddInt64(&that.flow, n)
	if flow-n >= limit {
		return 0
	}

	if flow > limit {
		n = flow - limit
	}

	return n
}
