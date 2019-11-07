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
