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

func TestOnce(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		times := 0

		once := &Once{}
		c.So(once.Done(), convey.ShouldBeFalse)

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				once.Do(func() {
					time.Sleep(time.Microsecond)
					times++
				})
			}()
		}
		wg.Wait()

		c.So(once.Done(), convey.ShouldBeTrue)
		c.So(times, convey.ShouldEqual, 1)
	})
}
