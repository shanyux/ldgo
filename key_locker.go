/*
 * Copyright (C) distroy
 */

package ldgo

import "sync"

var _KeyLockerDataPool = &sync.Pool{New: func() interface{} { return &keyLockerData{} }}

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
	d := kl.add(key)
	kl.locker.Unlock()

	d.locker.Lock()
}

func (kl *keyLocker) Unlock(key interface{}) {
	kl.locker.Lock()
	d := kl.sub(key)
	kl.locker.Unlock()

	if d != nil {
		d.locker.Unlock()
	}
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
	d := _KeyLockerDataPool.Get().(*keyLockerData)
	d.count = 0
	kl.keys[key] = d
	return d
}

func (kl *keyLocker) add(key interface{}) *keyLockerData {
	d := kl.get(key)
	if d == nil {
		d = kl.new(key)
	}

	d.count++
	return d
}

func (kl *keyLocker) sub(key interface{}) *keyLockerData {
	d := kl.get(key)
	if d == nil {
		return nil
	}

	d.count--
	if d.count <= 0 {
		delete(kl.keys, key)
		_KeyLockerDataPool.Put(d)
	}

	return d
}
