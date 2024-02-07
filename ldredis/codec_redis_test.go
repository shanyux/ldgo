/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	"github.com/distroy/ldgo/v2/ldctx"
	"github.com/smartystreets/goconvey/convey"
)

func TestCodecRedis_Nil(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		type TestImpl struct{}

		var pTest *TestImpl
		var p interface{} = pTest

		convey.So(pTest == nil, convey.ShouldBeTrue)
		convey.So(p, convey.ShouldBeNil)
		convey.So(p, convey.ShouldNotResemble, nil)

		key := "test-codec-redis-nil"
		expiration := 100 * time.Second

		s := rds.WithCodec(JsonCodec()).SetNX(ctx, key, p, expiration)
		convey.So(s.Err(), convey.ShouldBeNil)
		convey.So(s.Val(), convey.ShouldBeTrue)

		g0 := rds.Get(ctx, key)
		convey.So(g0.Err(), convey.ShouldBeNil)
		convey.So(g0.Val(), convey.ShouldResemble, "")

		g1 := rds.WithCodec(JsonCodec()).Get(ctx, key)
		convey.So(g1.Err(), convey.ShouldBeNil)
		convey.So(g1.Val(), convey.ShouldResemble, nil)
	})
}

func TestCodecRedis_String(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		key := "test-codec-redis-string"
		expiration := time.Duration(0)

		convey.Convey("marshal fail", func(c convey.C) {
			cmd := rds.WithCodec(ProtoV1Codec()).Set(ctx, key, "234", expiration)
			convey.So(cmd.Err().Error(), convey.ShouldStartWith, "the object for marshal must be `proto.Message`.")
		})
		convey.Convey("set/setnx/setxx/get", func() {
			s0 := rds.WithCodec(JsonCodec()).SetXX(ctx, key, "100", expiration)
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeFalse)

			s1 := rds.WithCodec(JsonCodec()).SetNX(ctx, key, "100", expiration)
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeTrue)

			g := rds.WithCodec(JsonCodec()).Get(ctx, key)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, "100")

			g1 := rds.Get(ctx, key)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, `"100"`)

			s2 := rds.WithCodec(JsonCodec()).SetNX(ctx, key, "123", expiration)
			convey.So(s2.Err(), convey.ShouldBeNil)
			convey.So(s2.Val(), convey.ShouldBeFalse)

			s3 := rds.WithCodec(JsonCodec()).SetXX(ctx, key, "123", expiration)
			convey.So(s3.Err(), convey.ShouldBeNil)
			convey.So(s3.Val(), convey.ShouldBeTrue)

			g3 := rds.Get(ctx, key)
			convey.So(g3.Err(), convey.ShouldBeNil)
			convey.So(g3.Val(), convey.ShouldResemble, `"123"`)

			s4 := rds.WithCodec(JsonCodec()).Set(ctx, key, "234", expiration)
			convey.So(s4.Err(), convey.ShouldBeNil)

			g4 := rds.Get(ctx, key)
			convey.So(g4.Err(), convey.ShouldBeNil)
			convey.So(g4.Val(), convey.ShouldResemble, `"234"`)
		})
		convey.Convey("mset/msetnx/mget", func() {
			key0 := key + "-0"
			key1 := key + "-1"
			key2 := key + "-2"

			s0 := rds.WithCodec(JsonCodec()).MSetNX(ctx, key0, "234", key1, "abc")
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeTrue)

			g := rds.WithCodec(JsonCodec()).MGet(ctx, key0, key1, key2)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, []interface{}{"234", "abc", nil})

			g0 := rds.MGet(ctx, key0, key1, key2)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, []interface{}{`"234"`, `"abc"`, nil})

			s1 := rds.WithCodec(JsonCodec()).MSetNX(ctx, key0, "100", key2, "xyz")
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeFalse)

			g1 := rds.MGet(ctx, key0, key1, key2)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, []interface{}{`"234"`, `"abc"`, nil})

			s2 := rds.WithCodec(JsonCodec()).MSet(ctx, key0, "100", key2, "xyz")
			convey.So(s2.Err(), convey.ShouldBeNil)

			g2 := rds.MGet(ctx, key0, key1, key2)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldResemble, []interface{}{`"100"`, `"abc"`, `"xyz"`})
		})
	})
}

func TestCodecRedis_Hash(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		key := "test-codec-redis-hash"
		field := "field"

		convey.Convey("hset/hsetnx/hget", func() {
			s0 := rds.WithCodec(JsonCodec()).HSetNX(ctx, key, field, "100")
			convey.So(s0.Err(), convey.ShouldBeNil)
			convey.So(s0.Val(), convey.ShouldBeTrue)

			g0 := rds.HGet(ctx, key, field)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, `"100"`)

			g := rds.WithCodec(JsonCodec()).HGet(ctx, key, field)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, "100")

			s1 := rds.WithCodec(JsonCodec()).HSetNX(ctx, key, field, "100")
			convey.So(s1.Err(), convey.ShouldBeNil)
			convey.So(s1.Val(), convey.ShouldBeFalse)

			s2 := rds.WithCodec(JsonCodec()).HSet(ctx, key, field, "abc")
			convey.So(s2.Err(), convey.ShouldBeNil)

			g2 := rds.HGet(ctx, key, field)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldResemble, `"abc"`)
		})
		convey.Convey("hmset/hmget/hgetall/hvals", func() {
			field0 := field + "-0"
			field1 := field + "-1"
			field2 := field + "-2"

			s := rds.WithCodec(JsonCodec()).HMSet(ctx, key, map[string]interface{}{
				field0: "100",
				field1: 100,
			})
			convey.So(s.Err(), convey.ShouldBeNil)

			g := rds.HMGet(ctx, key, field0, field1, field2)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, []interface{}{`"100"`, "100", nil})

			g0 := rds.WithCodec(JsonCodec()).HMGet(ctx, key, field0, field1, field2)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldResemble, []interface{}{"100", float64(100), nil})

			g1 := rds.WithCodec(JsonCodec()).HGetAll(ctx, key)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldResemble, map[string]interface{}{field0: "100", field1: float64(100)})

			// d1 := rds.WithCodec(JsonCodec()).HDel(ctx, key, field0)
			// convey.So(d1.Err(), convey.ShouldBeNil)

			g2 := rds.WithCodec(JsonCodec()).HVals(ctx, key)
			convey.So(g2.Err(), convey.ShouldBeNil)
			convey.So(g2.Val(), convey.ShouldHaveLength, 2)
			convey.So(g2.Val(), convey.ShouldContain, float64(100))
			convey.So(g2.Val(), convey.ShouldContain, "100")
		})
	})
}

func TestCodecRedis_ZSet(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		key := "test-codec-redis-zset"

		members := []ZMember{
			{Score: 1, Member: "100"},
			{Score: 10, Member: "abc"},
			{Score: 100, Member: "xyz"},
		}

		zadd := rds.WithCodec(JsonCodec()).ZAdd(ctx, key, members...)
		convey.So(zadd.Err(), convey.ShouldBeNil)

		convey.Convey("zrange", func() {
			zrange0 := rds.ZRange(ctx, key, 0, -1)
			convey.So(zrange0.Err(), convey.ShouldBeNil)
			convey.So(zrange0.Val(), convey.ShouldResemble, []string{`"100"`, `"abc"`, `"xyz"`})

			zrange1 := rds.WithCodec(JsonCodec()).ZRange(ctx, key, 0, -1)
			convey.So(zrange1.Err(), convey.ShouldBeNil)
			convey.So(zrange1.Val(), convey.ShouldResemble, []interface{}{"100", "abc", "xyz"})
		})

		convey.Convey("zincrby/zscore", func() {
			zincrby := rds.WithCodec(JsonCodec()).ZIncrBy(ctx, key, 10, "100")
			convey.So(zincrby.Err(), convey.ShouldBeNil)
			convey.So(zincrby.Val(), convey.ShouldResemble, float64(11))

			zscore := rds.WithCodec(JsonCodec()).ZScore(ctx, key, "100")
			convey.So(zscore.Err(), convey.ShouldBeNil)
			convey.So(zscore.Val(), convey.ShouldResemble, float64(11))
		})

		convey.Convey("zrank/zrevrank", func() {
			zrank0 := rds.WithCodec(JsonCodec()).ZRank(ctx, key, "100")
			convey.So(zrank0.Err(), convey.ShouldBeNil)
			convey.So(zrank0.Val(), convey.ShouldResemble, int64(0))

			zrank1 := rds.WithCodec(JsonCodec()).ZRank(ctx, key, "aaa")
			convey.So(zrank1.Err(), convey.ShouldEqual, Nil)

			zrevrank := rds.WithCodec(JsonCodec()).ZRevRank(ctx, key, "100")
			convey.So(zrevrank.Err(), convey.ShouldBeNil)
			convey.So(zrevrank.Val(), convey.ShouldResemble, int64(2))
		})

		convey.Convey("zrem", func() {
			zrem := rds.WithCodec(JsonCodec()).ZRem(ctx, key, "abc")
			convey.So(zrem.Err(), convey.ShouldBeNil)
			convey.So(zrem.Val(), convey.ShouldResemble, int64(1))

			zrange := rds.ZRange(ctx, key, 0, -1)
			convey.So(zrange.Err(), convey.ShouldBeNil)
			convey.So(zrange.Val(), convey.ShouldResemble, []string{`"100"`, `"xyz"`})
		})
	})
}

func TestCodecRedis_Set(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		key0 := "test-codec-redis-set-0"
		key1 := "test-codec-redis-set-1"

		sadd0 := rds.WithCodec(JsonCodec()).SAdd(ctx, key0, "abc", 100, "100")
		convey.So(sadd0.Err(), convey.ShouldBeNil)

		sadd1 := rds.WithCodec(JsonCodec()).SAdd(ctx, key1, "abc", 234, "xyz")
		convey.So(sadd1.Err(), convey.ShouldBeNil)

		convey.Convey("smembers", func() {
			g0 := rds.SMembers(ctx, key0)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldHaveLength, 3)
			convey.So(g0.Val(), convey.ShouldContain, `"100"`)
			convey.So(g0.Val(), convey.ShouldContain, `100`)
			convey.So(g0.Val(), convey.ShouldContain, `"abc"`)

			g1 := rds.WithCodec(JsonCodec()).SMembers(ctx, key0)
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldHaveLength, 3)
			convey.So(g1.Val(), convey.ShouldContain, "100")
			convey.So(g1.Val(), convey.ShouldContain, float64(100))
			convey.So(g1.Val(), convey.ShouldContain, "abc")
		})
		convey.Convey("smembers map", func() {
			g := rds.WithCodec(JsonCodec()).SMembersMap(ctx, key0)
			convey.So(g.Err(), convey.ShouldBeNil)
			convey.So(g.Val(), convey.ShouldResemble, map[interface{}]struct{}{
				"abc":        {},
				"100":        {},
				float64(100): {},
			})
		})
		convey.Convey("sismember", func() {
			g0 := rds.WithCodec(JsonCodec()).SIsMember(ctx, key0, 100)
			convey.So(g0.Err(), convey.ShouldBeNil)
			convey.So(g0.Val(), convey.ShouldBeTrue)

			g1 := rds.WithCodec(JsonCodec()).SIsMember(ctx, key0, "xyz")
			convey.So(g1.Err(), convey.ShouldBeNil)
			convey.So(g1.Val(), convey.ShouldBeFalse)
		})
		convey.Convey("sdiff/sinter", func() {
			sinter := rds.WithCodec(JsonCodec()).SInter(ctx, key0, key1)
			convey.So(sinter.Err(), convey.ShouldBeNil)
			convey.So(sinter.Val(), convey.ShouldResemble, []interface{}{"abc"})

			sdiff := rds.WithCodec(JsonCodec()).SDiff(ctx, key0, key1)
			convey.So(sdiff.Err(), convey.ShouldBeNil)
			convey.So(sdiff.Val(), convey.ShouldHaveLength, 2)
			convey.So(sdiff.Val(), convey.ShouldContain, float64(100))
			convey.So(sdiff.Val(), convey.ShouldContain, "100")
		})
		convey.Convey("srem", func() {
			srem := rds.WithCodec(JsonCodec()).SRem(ctx, key0, "100")
			convey.So(srem.Err(), convey.ShouldBeNil)
			convey.So(srem.Val(), convey.ShouldEqual, 1)

			smembers := rds.WithCodec(JsonCodec()).SMembers(ctx, key0)
			convey.So(smembers.Err(), convey.ShouldBeNil)
			convey.So(smembers.Val(), convey.ShouldHaveLength, 2)
			convey.So(smembers.Val(), convey.ShouldContain, float64(100))
			convey.So(smembers.Val(), convey.ShouldContain, "abc")
		})
	})
}

func TestCodecRedis_List(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		ctx := ldctx.Discard()

		key := "test-codec-redis-list"

		lpush := rds.WithCodec(JsonCodec()).LPush(ctx, key, "abc", 100, "100")
		convey.So(lpush.Err(), convey.ShouldBeNil)
		convey.So(lpush.Val(), convey.ShouldEqual, 3)

		rpush := rds.WithCodec(JsonCodec()).RPush(ctx, key, "xyz", 234, "100.234")
		convey.So(rpush.Err(), convey.ShouldBeNil)
		convey.So(rpush.Val(), convey.ShouldEqual, 6)

		convey.Convey("llen", func() {
			cmd := rds.WithCodec(JsonCodec()).LLen(ctx, key)
			convey.So(cmd.Err(), convey.ShouldBeNil)
			convey.So(cmd.Val(), convey.ShouldEqual, 6)
		})
		convey.Convey("lrange", func() {
			cmd := rds.WithCodec(JsonCodec()).LRange(ctx, key, 0, -1)
			convey.So(cmd.Err(), convey.ShouldBeNil)
			convey.So(cmd.Val(), convey.ShouldResemble, []interface{}{
				"100", float64(100), "abc", "xyz", float64(234), "100.234",
			})
		})
		convey.Convey("lindex", func() {
			cmd0 := rds.WithCodec(JsonCodec()).LIndex(ctx, key, 0)
			convey.So(cmd0.Err(), convey.ShouldBeNil)
			convey.So(cmd0.Val(), convey.ShouldResemble, "100")

			cmd1 := rds.WithCodec(JsonCodec()).LIndex(ctx, key, 5)
			convey.So(cmd1.Err(), convey.ShouldBeNil)
			convey.So(cmd1.Val(), convey.ShouldResemble, "100.234")
		})
		convey.Convey("lpop", func() {
			cmd := rds.WithCodec(JsonCodec()).LPop(ctx, key)
			convey.So(cmd.Err(), convey.ShouldBeNil)
			convey.So(cmd.Val(), convey.ShouldResemble, "100")
		})
		convey.Convey("rpop", func() {
			cmd := rds.WithCodec(JsonCodec()).RPop(ctx, key)
			convey.So(cmd.Err(), convey.ShouldBeNil)
			convey.So(cmd.Val(), convey.ShouldResemble, "100.234")
		})
		convey.Convey("LRem", func() {
			lpush := rds.WithCodec(JsonCodec()).LPush(ctx, key, "100", "aaa", "100", "bbb")
			convey.So(lpush.Err(), convey.ShouldBeNil)
			convey.So(lpush.Val(), convey.ShouldEqual, 10)

			lrem := rds.WithCodec(JsonCodec()).LRem(ctx, key, 2, "100")
			convey.So(lrem.Err(), convey.ShouldBeNil)
			convey.So(lrem.Val(), convey.ShouldEqual, 2)

			lrange := rds.WithCodec(JsonCodec()).LRange(ctx, key, 0, -1)
			convey.So(lrange.Err(), convey.ShouldBeNil)
			convey.So(lrange.Val(), convey.ShouldResemble, []interface{}{
				"bbb", "aaa", "100", float64(100), "abc", "xyz", float64(234), "100.234",
			})
		})
	})
}
