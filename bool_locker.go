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

func NewBoolLocker(b bool) sync.Locker {
	if b {
		return &sync.Mutex{}
	}
	return nullLocker{}
}

type nullLocker struct{}

func (_ nullLocker) Lock()   {}
func (_ nullLocker) Unlock() {}
