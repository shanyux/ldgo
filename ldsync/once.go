/*
 * Copyright (C) distroy
 */

package ldsync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done  uint32
	mutex sync.Mutex
}

func (o *Once) Done() bool { return atomic.LoadUint32(&o.done) != 0 }

func (o *Once) Reset() {
	o.mutex.Lock()
	atomic.StoreUint32(&o.done, 0)
	o.mutex.Unlock()
}

func (o *Once) Do(fn func()) {
	if o.Done() {
		return
	}

	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.done != 0 {
		return
	}

	defer atomic.StoreUint32(&o.done, 1)
	fn()
}
