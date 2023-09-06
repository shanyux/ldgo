/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestAsyncPool(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("nomarl", func(c convey.C) {
			n := int32(0)
			fn := func() {
				time.Sleep(time.Millisecond)
				atomic.AddInt32(&n, 1)
			}

			p := NewAsyncPool(10)
			for i := 0; i < 10; i++ {
				p.Async() <- fn
			}

			c.So(atomic.LoadInt32(&n), convey.ShouldEqual, 0)

			p.Close()
			p.Wait()
			c.So(atomic.LoadInt32(&n), convey.ShouldEqual, 10)
		})

		c.Convey("panic", func(c convey.C) {
			fn := func() {
				panic(11)
			}

			p := NewAsyncPool(10)
			p.Async() <- fn

			p.Close()
			p.Wait()
		})
	})
}
