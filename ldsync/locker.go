/*
 * Copyright (C) distroy
 */

package ldsync

import "sync"

var (
	_ Locker = (*AutoLocker)(nil)
)

type Locker interface {
	sync.Locker

	TryLock() bool
}

type AutoLocker struct {
	Locker sync.Locker

	locked bool
}

func (l *AutoLocker) Lock() {
	if l.locked {
		return
	}
	l.Locker.Lock()
	l.locked = true
}

func (l *AutoLocker) Unlock() {
	if !l.locked {
		return
	}
	l.Locker.Unlock()
	l.locked = false
}

func (l *AutoLocker) TryLock() bool {
	if l.locked {
		return true
	}
	locker, _ := l.Locker.(Locker)
	if locker == nil {
		return false
	}
	ok := locker.TryLock()
	l.locked = ok
	return ok
}
