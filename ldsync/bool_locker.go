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
	return NullLocker{}
}

type NullLocker struct{}

func (_ NullLocker) Lock()   {}
func (_ NullLocker) Unlock() {}
