/*
 * Copyright (C) distroy
 */

package ldsync

import (
	"sync"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestKeyLocker(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		kl := NewKeyLocker()
		key1 := "123"
		key2 := 123

		convey.So(kl.TryLock(key1), convey.ShouldBeTrue)
		convey.So(kl.TryLock(key2), convey.ShouldBeTrue)
		convey.So(kl.TryLock(key1), convey.ShouldBeFalse)
		convey.So(kl.TryLock(key2), convey.ShouldBeFalse)

		var key1UnlockTime, key1LockTime time.Time
		var key2UnlockTime, key2LockTime time.Time

		wg := &sync.WaitGroup{}
		wg.Add(4)
		go func() {
			defer wg.Done()
			kl.Lock(key1)
			key1LockTime = time.Now()
		}()
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Millisecond)
			kl.Unlock(key1)
			key1UnlockTime = time.Now()
		}()
		go func() {
			defer wg.Done()
			kl.Lock(key2)
			key2LockTime = time.Now()
		}()
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			kl.Unlock(key2)
			key2UnlockTime = time.Now()
		}()
		wg.Wait()

		convey.So(key1UnlockTime, convey.ShouldHappenBefore, key1LockTime)
		convey.So(key2UnlockTime, convey.ShouldHappenBefore, key2LockTime)
		convey.So(key1LockTime, convey.ShouldHappenBefore, key2LockTime)
		convey.So(key1UnlockTime, convey.ShouldHappenBefore, key2UnlockTime)

		kl.Unlock(key1)
		kl.Unlock(key2)
		kl.Unlock(key2) // double unlock

		convey.So(kl.TryLock(key1), convey.ShouldBeTrue)
		convey.So(kl.TryLock(key2), convey.ShouldBeTrue)
		convey.So(kl.TryLock(key1), convey.ShouldBeFalse)
		convey.So(kl.TryLock(key2), convey.ShouldBeFalse)
	})
}
