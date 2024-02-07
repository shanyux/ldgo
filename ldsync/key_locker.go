/*
 * Copyright (C) distroy
 */

package ldsync

import "sync"

var keyLockerDataPool = &Pool[*keyLockerData]{
	New: func() *keyLockerData { return &keyLockerData{} },
}

type KeyLocker interface {
	Lock(key interface{})
	TryLock(key interface{}) bool
	Unlock(key interface{})
}

func NewKeyLocker() KeyLocker {
	return &keyLocker{
		keys: make(map[interface{}]*keyLockerData),
	}
}

type keyLockerData struct {
	locker sync.Mutex
	count  int64
}

type keyLocker struct {
	locker sync.Mutex
	keys   map[interface{}]*keyLockerData
}

func (kl *keyLocker) Lock(key interface{}) {
	kl.locker.Lock()
	d := kl.getAndAdd(key)
	kl.locker.Unlock()

	d.locker.Lock()
}

func (kl *keyLocker) Unlock(key interface{}) {
	kl.locker.Lock()
	d := kl.getAndSub(key)
	kl.locker.Unlock()

	if d == nil {
		return
	}

	d.locker.Unlock()
}

func (kl *keyLocker) TryLock(key interface{}) bool {
	kl.locker.Lock()

	if d := kl.get(key); d != nil {
		kl.locker.Unlock()
		return false
	}

	d := kl.new(key)
	d.count++
	kl.locker.Unlock()

	d.locker.Lock()
	return true
}

func (kl *keyLocker) get(key interface{}) *keyLockerData {
	if kl.keys == nil {
		kl.keys = make(map[interface{}]*keyLockerData, 1)
	}

	return kl.keys[key]
}

func (kl *keyLocker) new(key interface{}) *keyLockerData {
	d := keyLockerDataPool.Get()
	d.count = 0
	kl.keys[key] = d
	return d
}

func (kl *keyLocker) getAndAdd(key interface{}) *keyLockerData {
	d := kl.get(key)
	if d == nil {
		d = kl.new(key)
	}

	d.count++
	return d
}

func (kl *keyLocker) getAndSub(key interface{}) *keyLockerData {
	d := kl.get(key)
	if d == nil {
		return nil
	}

	d.count--
	if d.count <= 0 {
		delete(kl.keys, key)
		keyLockerDataPool.Put(d)
	}

	return d
}
