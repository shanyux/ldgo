/*
 * Copyright (C) distroy
 */

package ldgo

import "sync"

var _KeyLockerDataPool = &sync.Pool{New: func() interface{} { return &keyLockerData{} }}

type KeyLocker interface {
	Lock(key interface{})
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
	d := kl.add(key)
	d.locker.Lock()
}

func (kl *keyLocker) Unlock(key interface{}) {
	d := kl.sub(key)
	d.locker.Unlock()
}

func (kl *keyLocker) add(key interface{}) *keyLockerData {
	kl.locker.Lock()

	if kl.keys == nil {
		kl.keys = make(map[interface{}]*keyLockerData, 1)
	}

	d := kl.keys[key]
	if d == nil {
		d = _KeyLockerDataPool.Get().(*keyLockerData)
		d.count = 0
	}

	d.count++
	kl.locker.Unlock()

	return d
}

func (kl *keyLocker) sub(key interface{}) *keyLockerData {
	kl.locker.Lock()
	defer kl.locker.Unlock()

	if kl.keys == nil {
		return nil
	}

	d := kl.keys[key]
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
