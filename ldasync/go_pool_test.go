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

func TestGoPool(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("nomarl", func(c convey.C) {
			n := int32(0)

			p := GoN(10, func() {
				time.Sleep(time.Millisecond)
				atomic.AddInt32(&n, 1)
			})

			c.So(atomic.LoadInt32(&n), convey.ShouldEqual, 0)
			c.So(p.Count(), convey.ShouldEqual, 10)

			p.Wait()
			c.So(atomic.LoadInt32(&n), convey.ShouldEqual, 10)
			c.So(p.Count(), convey.ShouldEqual, 0)
		})

		c.Convey("panic", func(c convey.C) {
			fn := func() {
				panic(11)
			}

			c.So(func() { Go(fn) }, convey.ShouldNotPanic)
		})
	})
}
