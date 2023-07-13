/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/ldgopool"
	"github.com/distroy/ldgo/ldlog"
	"github.com/smartystreets/goconvey/convey"
)

func testMemoryRedis() *Redis {
	return MustNewTestRedis()
}

func TestMutex_Lock(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		lockKey := "test-key"
		timeout := 1 * time.Second

		ctx := ldctx.WithTimeout(ldctx.Discard(), timeout)
		rds = rds.WithContext(ctx)

		m0 := NewMutex(rds)
		m1 := NewMutex(rds)

		convey.So(m0.Lock(lockKey), convey.ShouldBeNil)
		convey.So(m1.Lock(lockKey), convey.ShouldNotBeNil)
		convey.So(m0.Unlock(), convey.ShouldBeNil)
	})
}

func TestMutex_Unlock(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		lockKey := "test-key"
		timeout := 1 * time.Second
		ctx := ldctx.WithTimeout(ldctx.Discard(), timeout)

		convey.Convey("unlock after context not timeout", func() {
			m := NewMutex(rds).WithContext(ctx)

			convey.So(m.Lock(lockKey), convey.ShouldBeNil)

			// time.Sleep(timeout - 1*time.Second)

			convey.So(m.Unlock(), convey.ShouldBeNil)
		})

		convey.Convey("unlock after context timeout", func() {
			m := NewMutex(rds).WithContext(ctx)

			convey.So(m.Lock(lockKey), convey.ShouldBeNil)

			<-m.Events()
			// time.Sleep(timeout + 1*time.Second)

			convey.So(m.Unlock(), convey.ShouldBeNil)
		})
	})
}

func TestMutex_WithLockForce(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		r := testMemoryRedis()
		defer r.Close()

		ctx := ldctx.Default()
		ctx = ldctx.WithLogger(ctx, ldlog.Discard())
		r = r.WithContext(ctx)

		lockKey := "test-key"
		interval := 100 * time.Millisecond
		timeout0 := 2 * time.Second
		timeout1 := 1 * time.Second

		convey.Convey("lock force without timeout", func(c convey.C) {
			var t0, t1 time.Time
			m0 := NewMutex(r)
			m0 = m0.WithInterval(time.Second)

			m1 := NewMutex(r)
			m1 = m1.WithLockForce(true, interval)

			gos := ldgopool.NewGoPool()

			c.So(m0.Lock(lockKey), convey.ShouldBeNil)
			gos.Run(func() {
				time.Sleep(timeout0)
				c.So(m0.Unlock(), convey.ShouldBeNil)

				t0 = time.Now()
			})

			gos.Run(func() {
				c.So(m1.Lock(lockKey), convey.ShouldBeNil)
				t1 = time.Now()

				// time.Sleep(timeout)
				c.So(m1.Unlock(), convey.ShouldBeNil)
			})

			gos.Wait()
			c.So(t0, convey.ShouldHappenBefore, t1)
		})

		convey.Convey("lock force with timeout", func(c convey.C) {
			convey.Convey("lock succ", func(c convey.C) {
				var t0, t1 time.Time
				m0 := NewMutex(r)

				m1 := NewMutex(r)
				m1 = m1.WithLockForce(true, interval, timeout1)

				gos := ldgopool.NewGoPool()

				c.So(m0.Lock(lockKey), convey.ShouldBeNil)
				gos.Run(func() {

					time.Sleep(timeout0)
					c.So(m0.Unlock(), convey.ShouldBeNil)

					t0 = time.Now()
				})

				gos.Run(func() {
					m := m1
					c.So(m.Lock(lockKey), convey.ShouldNotBeNil)
					t1 = time.Now()

					// c.So(m.Unlock(), convey.ShouldNotBeNil)
				})

				gos.Wait()
				c.So(t0, convey.ShouldHappenAfter, t1)
			})
			convey.Convey("lock timeout", func(c convey.C) {
				var t0, t1 time.Time
				m0 := NewMutex(r)

				m1 := NewMutex(r)
				m1 = m1.WithLockForce(true, interval, timeout0-time.Second)

				gos := ldgopool.NewGoPool()

				c.So(m0.Lock(lockKey), convey.ShouldBeNil)
				gos.Run(func() {

					time.Sleep(timeout0)
					c.So(m0.Unlock(), convey.ShouldBeNil)

					t0 = time.Now()
				})

				gos.Run(func() {
					m := m1
					c.So(m.Lock(lockKey), convey.ShouldNotBeNil)
					t1 = time.Now()

					// c.So(m.Unlock(), convey.ShouldNotBeNil)
				})

				gos.Wait()
				c.So(t0, convey.ShouldHappenAfter, t1)
			})
		})
	})
}
