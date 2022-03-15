/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestCodecRedis_String(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-codec-redis-string"
		expiration := time.Duration(0)

		convey.Convey("set/setnx/setxx/get", func() {
			s0 := rds.WithCodec(JsonCodec()).SetXX(key, "100", expiration)
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeFalse)

			s1 := rds.WithCodec(JsonCodec()).SetNX(key, "100", expiration)
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeTrue)

			g := rds.WithCodec(JsonCodec()).Get(key)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, "100")

			g1 := rds.Get(key)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, `"100"`)

			s2 := rds.WithCodec(JsonCodec()).SetNX(key, "123", expiration)
			convey.So(s2.Err(), convey.ShouldBeNil)
			convey.So(s2.Val(), convey.ShouldBeFalse)

			s3 := rds.WithCodec(JsonCodec()).SetXX(key, "123", expiration)
			convey.So(s3.Err(), convey.ShouldBeNil)
			convey.So(s3.Val(), convey.ShouldBeTrue)

			g3 := rds.Get(key)
			convey.So(g3.Err(), convey.ShouldBeNil)
			convey.So(g3.Val(), convey.ShouldResemble, `"123"`)

			s4 := rds.WithCodec(JsonCodec()).Set(key, "234", expiration)
			convey.So(s4.Err(), convey.ShouldBeNil)

			g4 := rds.Get(key)
			convey.So(g4.Err(), convey.ShouldBeNil)
			convey.So(g4.Val(), convey.ShouldResemble, `"234"`)
		})
		convey.Convey("mset/msetnx/mget", func() {
			key0 := key + "-0"
			key1 := key + "-1"
			key2 := key + "-2"

			s0 := rds.WithCodec(JsonCodec()).MSetNX(key0, "234", key1, "abc")
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeTrue)

			g := rds.WithCodec(JsonCodec()).MGet(key0, key1, key2)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, []interface{}{"234", "abc", nil})

			g0 := rds.MGet(key0, key1, key2)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, []interface{}{`"234"`, `"abc"`, nil})

			s1 := rds.WithCodec(JsonCodec()).MSetNX(key0, "100", key2, "xyz")
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeFalse)

			g1 := rds.MGet(key0, key1, key2)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, []interface{}{`"234"`, `"abc"`, nil})

			s2 := rds.WithCodec(JsonCodec()).MSet(key0, "100", key2, "xyz")
			convey.So(s2.Err(), convey.ShouldBeNil)

			g2 := rds.MGet(key0, key1, key2)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldResemble, []interface{}{`"100"`, `"abc"`, `"xyz"`})
		})
	})
}

func TestCodecRedis_Hash(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-codec-redis-hash"
		field := "field"

		convey.Convey("hset/hsetnx/hget", func() {
			s0 := rds.WithCodec(JsonCodec()).HSetNX(key, field, "100")
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeTrue)

			g0 := rds.HGet(key, field)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, `"100"`)

			g := rds.WithCodec(JsonCodec()).HGet(key, field)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, "100")

			s1 := rds.WithCodec(JsonCodec()).HSetNX(key, field, "100")
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeFalse)

			s2 := rds.WithCodec(JsonCodec()).HSet(key, field, "abc")
			convey.So(s2.Err(), convey.ShouldBeNil)

			g2 := rds.HGet(key, field)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldResemble, `"abc"`)
		})
		convey.Convey("hmset/hmget/hgetall/hvals", func() {
			field0 := field + "-0"
			field1 := field + "-1"
			field2 := field + "-2"

			s := rds.WithCodec(JsonCodec()).HMSet(key, map[string]interface{}{
				field0: "100",
				field1: 100,
			})
			convey.So(s.Err(), convey.ShouldBeNil)

			g := rds.HMGet(key, field0, field1, field2)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, []interface{}{`"100"`, "100", nil})

			g0 := rds.WithCodec(JsonCodec()).HMGet(key, field0, field1, field2)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, []interface{}{"100", float64(100), nil})

			g1 := rds.WithCodec(JsonCodec()).HGetAll(key)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, map[string]interface{}{field0: "100", field1: float64(100)})

			// d1 := rds.WithCodec(JsonCodec()).HDel(key, field0)
			// convey.So(d1.Err(), convey.ShouldBeNil)

			g2 := rds.WithCodec(JsonCodec()).HVals(key)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldContain, float64(100))
			convey.So(g2.Val(), convey.ShouldContain, "100")
		})
	})
}

func TestCodecRedis_ZSet(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-codec-redis-zset"

		members := []ZMember{
			{Score: 1, Member: "100"},
			{Score: 10, Member: "abc"},
			{Score: 100, Member: "xyz"},
		}

		zadd := rds.WithCodec(JsonCodec()).ZAdd(key, members...)
		convey.So(zadd.Err(), convey.ShouldBeNil)

		convey.Convey("zrange", func() {
			zrange0 := rds.ZRange(key, 0, -1)
			convey.So(zrange0.Err(), convey.ShouldBeNil)
			convey.So(zrange0.Val(), convey.ShouldResemble, []string{`"100"`, `"abc"`, `"xyz"`})

			zrange1 := rds.WithCodec(JsonCodec()).ZRange(key, 0, -1)
			convey.So(zrange1.Err(), convey.ShouldBeNil)
			convey.So(zrange1.Val(), convey.ShouldResemble, []interface{}{"100", "abc", "xyz"})
		})

		convey.Convey("zincrby/zscore", func() {
			zincrby := rds.WithCodec(JsonCodec()).ZIncrBy(key, 10, "100")
			convey.So(zincrby.Err(), convey.ShouldBeNil)
			convey.So(zincrby.Val(), convey.ShouldResemble, float64(11))

			zscore := rds.WithCodec(JsonCodec()).ZScore(key, "100")
			convey.So(zscore.Err(), convey.ShouldBeNil)
			convey.So(zscore.Val(), convey.ShouldResemble, float64(11))
		})

		convey.Convey("zrank/zrevrank", func() {
			zrank0 := rds.WithCodec(JsonCodec()).ZRank(key, "100")
			convey.So(zrank0.Err(), convey.ShouldBeNil)
			convey.So(zrank0.Val(), convey.ShouldResemble, int64(0))

			zrank1 := rds.WithCodec(JsonCodec()).ZRank(key, "aaa")
			convey.So(zrank1.Err(), convey.ShouldEqual, Nil)

			zrevrank := rds.WithCodec(JsonCodec()).ZRevRank(key, "100")
			convey.So(zrevrank.Err(), convey.ShouldBeNil)
			convey.So(zrevrank.Val(), convey.ShouldResemble, int64(2))
		})

		convey.Convey("zrem", func() {
			zrem := rds.WithCodec(JsonCodec()).ZRem(key, "abc")
			convey.So(zrem.Err(), convey.ShouldBeNil)
			convey.So(zrem.Val(), convey.ShouldResemble, int64(1))

			zrange := rds.ZRange(key, 0, -1)
			convey.So(zrange.Err(), convey.ShouldBeNil)
			convey.So(zrange.Val(), convey.ShouldResemble, []string{`"100"`, `"xyz"`})
		})
	})
}
