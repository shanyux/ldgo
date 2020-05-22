/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"sync"
)

type none = struct{}

type Done interface {
	Stop()
	Done() <-chan none
}

func NewDone() Done {
	return &done{
		Chan: make(chan struct{}),
	}
}

type done struct {
	Chan chan none
	Once sync.Once
}

func (that *done) Stop() {
	that.Once.Do(func() { close(that.Chan) })
}
func (that *done) Done() <-chan none {
	return that.Chan
}
