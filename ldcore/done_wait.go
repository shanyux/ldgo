/*
 * Copyright (C) distroy
 */

package ldcore

import (
	"sync"
)

type DoneWait interface {
	Stop()
	Chan() <-chan none

	Add(n int)
	Done()
	Wait()
}

func NewDoneWait() DoneWait {
	return newDoneWait()
}

func newDoneWait() *doneWait {
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
