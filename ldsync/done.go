/*
 * Copyright (C) distroy
 */

package ldsync

import (
	"sync"
)

type none = struct{}

type Done interface {
	Stop()
	Chan() <-chan none
}

func NewDone() Done {
	return &done{
		c: make(chan struct{}),
	}
}

type done struct {
	c chan none
	o sync.Once
}

func (that *done) Stop() {
	that.o.Do(func() { close(that.c) })
}
func (that *done) Chan() <-chan none {
	return that.c
}
