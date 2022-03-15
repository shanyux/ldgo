/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/distroy/ldgo/ldctx"
	"github.com/go-redis/redis"
	"github.com/smartystreets/goconvey/convey"
)

func testMemoryRedis() *Redis {
	server, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})
	return New(client)
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
