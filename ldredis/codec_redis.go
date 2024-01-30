/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	redis "github.com/redis/go-redis/v9"
)

type client = Redis

var _ CodecCmdable = (*CodecRedis)(nil)

type CodecRedis struct {
	*client

	codec Codec
}

func (c *CodecRedis) Client() *Redis { return c.client }

func (c *CodecRedis) clone(cli ...*Redis) *CodecRedis {
	cp := *c
	c = &cp
	if len(cli) > 0 && cli[0] != nil {
		cp.client = cli[0]
	}
	return c
}

func (c *CodecRedis) WithContext(ctx Context) *CodecRedis { return c.clone(c.client.WithContext(ctx)) }
func (c *CodecRedis) WithRetry(retry int) *CodecRedis     { return c.clone(c.client.WithRetry(retry)) }
func (c *CodecRedis) WithReport(r Reporter) *CodecRedis   { return c.clone(c.client.WithReport(r)) }
func (c *CodecRedis) WithCaller(enable bool) *CodecRedis  { return c.clone(c.client.WithCaller(enable)) }

// ***** redis string begin *****

func (c *CodecRedis) MGet(keys ...string) *CodecSliceCmd {
	return newCodecSliceCmd(c, c.client.MGet(keys...))
}
func (c *CodecRedis) MSet(pairs ...interface{}) *StatusCmd {
	return c.client.MSet(c.marshalPairs(pairs)...)
}
func (c *CodecRedis) MSetNX(pairs ...interface{}) *BoolCmd {
	return c.client.MSetNX(c.marshalPairs(pairs)...)
}

func (c *CodecRedis) Get(key string) *CodecCmd {
	return newCodecCmd(c, c.client.Get(key))
}
func (c *CodecRedis) GetSet(key string, value interface{}) *CodecCmd {
	return newCodecCmd(c, c.client.GetSet(key, c.mustMarshal(value)))
}
func (c *CodecRedis) Set(key string, value interface{}, expiration time.Duration) *StatusCmd {
	return c.client.Set(key, c.mustMarshal(value), expiration)
}
func (c *CodecRedis) SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd {
	return c.client.SetNX(key, c.mustMarshal(value), expiration)
}
func (c *CodecRedis) SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd {
	return c.client.SetXX(key, c.mustMarshal(value), expiration)
}

// ***** redis string end *****

// ***** redis hash begin *****

func (c *CodecRedis) HGet(key, field string) *CodecCmd {
	return newCodecCmd(c, c.client.HGet(key, field))
}
func (c *CodecRedis) HGetAll(key string) *StringCodecMapCmd {
	return newStringCodecMapCmd(c, c.client.HGetAll(key))
}
func (c *CodecRedis) HMGet(key string, fields ...string) *CodecSliceCmd {
	return newCodecSliceCmd(c, c.client.HMGet(key, fields...))
}
func (c *CodecRedis) HMSet(key string, fields map[string]interface{}) *StatusCmd {
	return c.client.HMSet(key, c.marshalStringInterfaceMap(fields))
}
func (c *CodecRedis) HSet(key, field string, value interface{}) *BoolCmd {
	return c.client.HSet(key, field, c.mustMarshal(value))
}
func (c *CodecRedis) HSetNX(key, field string, value interface{}) *BoolCmd {
	return c.client.HSetNX(key, field, c.mustMarshal(value))
}
func (c *CodecRedis) HVals(key string) *CodecsCmd {
	return newCodecsCmd(c, c.client.HVals(key))
}

// ***** redis hash end *****

// ***** redis list begin *****

func (c *CodecRedis) BLPop(timeout time.Duration, keys ...string) *CodecsCmd {
	return newCodecsCmd(c, c.client.BLPop(timeout, keys...))
}
func (c *CodecRedis) BRPop(timeout time.Duration, keys ...string) *CodecsCmd {
	return newCodecsCmd(c, c.client.BRPop(timeout, keys...))
}
func (c *CodecRedis) BRPopLPush(source, destination string, timeout time.Duration) *CodecCmd {
	return newCodecCmd(c, c.client.BRPopLPush(source, destination, timeout))
}
func (c *CodecRedis) LIndex(key string, index int64) *CodecCmd {
	return newCodecCmd(c, c.client.LIndex(key, index))
}
func (c *CodecRedis) LInsert(key, op string, pivot, value interface{}) *IntCmd {
	return c.client.LInsert(key, op, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LInsertBefore(key string, pivot, value interface{}) *IntCmd {
	return c.client.LInsertBefore(key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LInsertAfter(key string, pivot, value interface{}) *IntCmd {
	return c.client.LInsertAfter(key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LPop(key string) *CodecCmd {
	return newCodecCmd(c, c.client.LPop(key))
}
func (c *CodecRedis) LPush(key string, values ...interface{}) *IntCmd {
	return c.client.LPush(key, c.marshalSlice(values)...)
}
func (c *CodecRedis) LPushX(key string, value interface{}) *IntCmd {
	return c.client.LPushX(key, c.mustMarshal(value))
}
func (c *CodecRedis) LRange(key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(c, c.client.LRange(key, start, stop))
}
func (c *CodecRedis) LRem(key string, count int64, value interface{}) *IntCmd {
	return c.client.LRem(key, count, c.mustMarshal(value))
}
func (c *CodecRedis) LSet(key string, index int64, value interface{}) *StatusCmd {
	return c.client.LSet(key, index, c.mustMarshal(value))
}
func (c *CodecRedis) RPop(key string) *CodecCmd {
	return newCodecCmd(c, c.client.RPop(key))
}
func (c *CodecRedis) RPopLPush(source, destination string) *CodecCmd {
	return newCodecCmd(c, c.client.RPopLPush(source, destination))
}
func (c *CodecRedis) RPush(key string, values ...interface{}) *IntCmd {
	return c.client.RPush(key, c.marshalSlice(values)...)
}
func (c *CodecRedis) RPushX(key string, value interface{}) *IntCmd {
	return c.client.RPushX(key, c.mustMarshal(value))
}

// ***** redis list begin *****

// ***** redis set begin *****

func (c *CodecRedis) SAdd(key string, members ...interface{}) *IntCmd {
	return c.client.SAdd(key, c.marshalSlice(members)...)
}
func (c *CodecRedis) SDiff(keys ...string) *CodecsCmd {
	return newCodecsCmd(c, c.client.SDiff(keys...))
}

func (c *CodecRedis) SInter(keys ...string) *CodecsCmd {
	return newCodecsCmd(c, c.client.SInter(keys...))
}

func (c *CodecRedis) SIsMember(key string, member interface{}) *BoolCmd {
	return c.client.SIsMember(key, c.mustMarshal(member))
}
func (c *CodecRedis) SMembers(key string) *CodecsCmd {
	return newCodecsCmd(c, c.client.SMembers(key))
}
func (c *CodecRedis) SMembersMap(key string) *CodecSetCmd {
	return newCodecSetCmd(c, c.client.SMembersMap(key))
}
func (c *CodecRedis) SMove(src, dest string, member interface{}) *BoolCmd {
	return c.client.SMove(src, dest, c.mustMarshal(member))
}
func (c *CodecRedis) SPop(key string) *CodecCmd {
	return newCodecCmd(c, c.client.SPop(key))
}
func (c *CodecRedis) SPopN(key string, count int64) *CodecsCmd {
	return newCodecsCmd(c, c.client.SPopN(key, count))
}
func (c *CodecRedis) SRandMember(key string) *CodecCmd {
	return newCodecCmd(c, c.client.SRandMember(key))
}
func (c *CodecRedis) SRandMemberN(key string, count int64) *CodecsCmd {
	return newCodecsCmd(c, c.client.SRandMemberN(key, count))
}
func (c *CodecRedis) SRem(key string, members ...interface{}) *IntCmd {
	return c.client.SRem(key, c.marshalSlice(members)...)
}
func (c *CodecRedis) SUnion(keys ...string) *CodecsCmd {
	return newCodecsCmd(c, c.client.SUnion(keys...))
}

// func (c *CodecRedis)SUnionStore(destination string, keys ...string) *IntCmd // same as Cmdable

// ***** redis set end *****

// ***** redis zset begin *****

func (c *CodecRedis) ZAdd(key string, members ...ZMember) *IntCmd {
	return c.client.ZAdd(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddNX(key string, members ...ZMember) *IntCmd {
	return c.client.ZAddNX(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddXX(key string, members ...ZMember) *IntCmd {
	return c.client.ZAddXX(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddCh(key string, members ...ZMember) *IntCmd {
	return c.client.ZAddCh(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddNXCh(key string, members ...ZMember) *IntCmd {
	return c.client.ZAddNXCh(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddXXCh(key string, members ...ZMember) *IntCmd {
	return c.client.ZAddXXCh(key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZIncr(key string, member ZMember) *FloatCmd {
	return c.client.ZIncr(key, c.marshalZMember(member))
}
func (c *CodecRedis) ZIncrNX(key string, member ZMember) *FloatCmd {
	return c.client.ZIncrNX(key, c.marshalZMember(member))
}
func (c *CodecRedis) ZIncrXX(key string, member ZMember) *FloatCmd {
	return c.client.ZIncrXX(key, c.marshalZMember(member))
}
func (c *CodecRedis) ZPopMax(key string, count ...int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZPopMax(key, count...))
}
func (c *CodecRedis) ZIncrBy(key string, increment float64, member interface{}) *FloatCmd {
	cmd := redis.NewFloatCmd("zincrby", key, increment, c.mustMarshal(member))
	c.client.Process(cmd)
	return cmd
}
func (c *CodecRedis) ZPopMin(key string, count ...int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZPopMin(key, count...))
}
func (c *CodecRedis) ZRange(key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRange(key, start, stop))
}
func (c *CodecRedis) ZRangeWithScores(key string, start, stop int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZRangeWithScores(key, start, stop))
}
func (c *CodecRedis) ZRangeByScore(key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRangeByScore(key, opt))
}
func (c *CodecRedis) ZRangeByLex(key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRangeByLex(key, opt))
}
func (c *CodecRedis) ZRangeByScoreWithScores(key string, opt ZRangeBy) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZRangeByScoreWithScores(key, opt))
}
func (c *CodecRedis) ZRank(key, member interface{}) *IntCmd {
	cmd := redis.NewIntCmd("zrank", key, c.mustMarshal(member))
	c.client.Process(cmd)
	return cmd
}
func (c *CodecRedis) ZRem(key string, members ...interface{}) *IntCmd {
	return c.client.ZRem(key, c.marshalSlice(members)...)
}
func (c *CodecRedis) ZRevRange(key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRevRange(key, start, stop))
}
func (c *CodecRedis) ZRevRangeWithScores(key string, start, stop int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZRevRangeWithScores(key, start, stop))
}
func (c *CodecRedis) ZRevRangeByScore(key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRevRangeByScore(key, opt))
}
func (c *CodecRedis) ZRevRangeByLex(key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(c, c.client.ZRevRangeByLex(key, opt))
}
func (c *CodecRedis) ZRevRangeByScoreWithScores(key string, opt ZRangeBy) *ZCodecSliceCmd {
	return newZCodecSliceCmd(c, c.client.ZRevRangeByScoreWithScores(key, opt))
}
func (c *CodecRedis) ZRevRank(key, member interface{}) *IntCmd {
	cmd := redis.NewIntCmd("zrevrank", key, c.mustMarshal(member))
	c.client.Process(cmd)
	return cmd
}
func (c *CodecRedis) ZScore(key, member interface{}) *FloatCmd {
	cmd := redis.NewFloatCmd("zscore", key, c.mustMarshal(member))
	c.client.Process(cmd)
	return cmd
}

// ***** redis zset end *****
