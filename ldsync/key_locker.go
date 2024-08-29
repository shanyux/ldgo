/*
 * Copyright (C) distroy
 */

package ldsync

import "sync"

var keyLockerDataPool = &Pool[*keyLockerData]{
	New: func() *keyLockerData { return &keyLockerData{} },
}

func NewKeyLocker() *KeyLocker {
	return &KeyLocker{}
}

type keyLockerData struct {
	locker sync.Mutex
	count  int64
}

type KeyLocker struct {
	locker sync.Mutex
	keys   map[interface{}]*keyLockerData
}

func (kl *KeyLocker) init() {
	if kl.keys == nil {
		kl.keys = make(map[interface{}]*keyLockerData, 1)
	}
}

func (kl *KeyLocker) Lock(key interface{}) {
	kl.locker.Lock()
	d := kl.getAndAdd(key)
	kl.locker.Unlock()

	d.locker.Lock()
}

func (kl *KeyLocker) Unlock(key interface{}) {
	kl.locker.Lock()
	d := kl.getAndSub(key)
	kl.locker.Unlock()

	if d == nil {
		return
	}

	d.locker.Unlock()
}

func (kl *KeyLocker) TryLock(key interface{}) bool {
	kl.locker.Lock()

	if d := kl.get(key); d != nil {
		kl.locker.Unlock()
		return false
	}

	d := kl.new(key)
	d.count++
	kl.locker.Unlock()

	return d.locker.TryLock()
}

func (kl *KeyLocker) get(key interface{}) *keyLockerData {
	kl.init()
	return kl.keys[key]
}

func (kl *KeyLocker) new(key interface{}) *keyLockerData {
	d := keyLockerDataPool.Get()
	d.count = 0
	kl.keys[key] = d
	return d
}

func (kl *KeyLocker) getAndAdd(key interface{}) *keyLockerData {
	d := kl.get(key)
	if d == nil {
		d = kl.new(key)
	}

	d.count++
	return d
}

func (kl *KeyLocker) getAndSub(key interface{}) *keyLockerData {
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
