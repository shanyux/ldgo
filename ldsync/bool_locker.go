/*
 * Copyright (C) distroy
 */

package ldsync

import (
	"sync"
)

func NewBoolLocker(b bool) sync.Locker {
	if b {
		return &sync.Mutex{}
	}
	return nullLocker{}
}

type nullLocker struct{}

func (_ nullLocker) Lock()   {}
func (_ nullLocker) Unlock() {}
