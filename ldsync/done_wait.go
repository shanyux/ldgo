/*
 * Copyright (C) distroy
 */

package ldsync

import (
	"sync"
)

type DoneWait interface {
	Done

	Add(n int)
	Done()
	Wait()
}

func NewDoneWait() DoneWait {
	return newDoneWait()
}

func newDoneWait() *doneWait {
	return &doneWait{
		done: done{
			c: make(chan struct{}),
		},
	}
}

type doneWait struct {
	sync.WaitGroup
	done
}
