/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"
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

func (c *CodecRedis) WithRetry(retry int) *CodecRedis    { return c.clone(c.client.WithRetry(retry)) }
func (c *CodecRedis) WithReport(r Reporter) *CodecRedis  { return c.clone(c.client.WithReport(r)) }
func (c *CodecRedis) WithCaller(enable bool) *CodecRedis { return c.clone(c.client.WithCaller(enable)) }

// ***** redis string begin *****

func (c *CodecRedis) MGet(cc context.Context, keys ...string) *CodecSliceCmd {
	return newCodecSliceCmd(cc, c, c.client.MGet(cc, keys...))
}
func (c *CodecRedis) MSet(cc context.Context, pairs ...interface{}) *StatusCmd {
	return c.client.MSet(cc, c.marshalPairs(pairs)...)
}
func (c *CodecRedis) MSetNX(cc context.Context, pairs ...interface{}) *BoolCmd {
	return c.client.MSetNX(cc, c.marshalPairs(pairs)...)
}

func (c *CodecRedis) Get(cc context.Context, key string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.Get(cc, key))
}
func (c *CodecRedis) GetSet(cc context.Context, key string, value interface{}) *CodecCmd {
	return newCodecCmd(cc, c, c.client.GetSet(cc, key, c.mustMarshal(value)))
}
func (c *CodecRedis) Set(cc context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd {
	return c.client.Set(cc, key, c.mustMarshal(value), expiration)
}
func (c *CodecRedis) SetNX(cc context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	return c.client.SetNX(cc, key, c.mustMarshal(value), expiration)
}
func (c *CodecRedis) SetXX(cc context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd {
	return c.client.SetXX(cc, key, c.mustMarshal(value), expiration)
}

// ***** redis string end *****

// ***** redis hash begin *****

func (c *CodecRedis) HGet(cc context.Context, key, field string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.HGet(cc, key, field))
}
func (c *CodecRedis) HGetAll(cc context.Context, key string) *StringCodecMapCmd {
	return newStringCodecMapCmd(cc, c, c.client.HGetAll(cc, key))
}
func (c *CodecRedis) HMGet(cc context.Context, key string, fields ...string) *CodecSliceCmd {
	return newCodecSliceCmd(cc, c, c.client.HMGet(cc, key, fields...))
}
func (c *CodecRedis) HMSet(cc context.Context, key string, fields map[string]interface{}) *StatusCmd {
	return c.client.HMSet(cc, key, c.marshalStringInterfaceMap(fields))
}
func (c *CodecRedis) HSet(cc context.Context, key, field string, value interface{}) *BoolCmd {
	return c.client.HSet(cc, key, field, c.mustMarshal(value))
}
func (c *CodecRedis) HSetNX(cc context.Context, key, field string, value interface{}) *BoolCmd {
	return c.client.HSetNX(cc, key, field, c.mustMarshal(value))
}
func (c *CodecRedis) HVals(cc context.Context, key string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.HVals(cc, key))
}

// ***** redis hash end *****

// ***** redis list begin *****

func (c *CodecRedis) BLPop(cc context.Context, timeout time.Duration, keys ...string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.BLPop(cc, timeout, keys...))
}
func (c *CodecRedis) BRPop(cc context.Context, timeout time.Duration, keys ...string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.BRPop(cc, timeout, keys...))
}
func (c *CodecRedis) BRPopLPush(cc context.Context, source, destination string, timeout time.Duration) *CodecCmd {
	return newCodecCmd(cc, c, c.client.BRPopLPush(cc, source, destination, timeout))
}
func (c *CodecRedis) LIndex(cc context.Context, key string, index int64) *CodecCmd {
	return newCodecCmd(cc, c, c.client.LIndex(cc, key, index))
}
func (c *CodecRedis) LInsert(cc context.Context, key, op string, pivot, value interface{}) *IntCmd {
	return c.client.LInsert(cc, key, op, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LInsertBefore(cc context.Context, key string, pivot, value interface{}) *IntCmd {
	return c.client.LInsertBefore(cc, key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LInsertAfter(cc context.Context, key string, pivot, value interface{}) *IntCmd {
	return c.client.LInsertAfter(cc, key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *CodecRedis) LPop(cc context.Context, key string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.LPop(cc, key))
}
func (c *CodecRedis) LPush(cc context.Context, key string, values ...interface{}) *IntCmd {
	return c.client.LPush(cc, key, c.marshalSlice(values)...)
}
func (c *CodecRedis) LPushX(cc context.Context, key string, value interface{}) *IntCmd {
	return c.client.LPushX(cc, key, c.mustMarshal(value))
}
func (c *CodecRedis) LRange(cc context.Context, key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.LRange(cc, key, start, stop))
}
func (c *CodecRedis) LRem(cc context.Context, key string, count int64, value interface{}) *IntCmd {
	return c.client.LRem(cc, key, count, c.mustMarshal(value))
}
func (c *CodecRedis) LSet(cc context.Context, key string, index int64, value interface{}) *StatusCmd {
	return c.client.LSet(cc, key, index, c.mustMarshal(value))
}
func (c *CodecRedis) RPop(cc context.Context, key string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.RPop(cc, key))
}
func (c *CodecRedis) RPopLPush(cc context.Context, source, destination string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.RPopLPush(cc, source, destination))
}
func (c *CodecRedis) RPush(cc context.Context, key string, values ...interface{}) *IntCmd {
	return c.client.RPush(cc, key, c.marshalSlice(values)...)
}
func (c *CodecRedis) RPushX(cc context.Context, key string, value interface{}) *IntCmd {
	return c.client.RPushX(cc, key, c.mustMarshal(value))
}

// ***** redis list begin *****

// ***** redis set begin *****

func (c *CodecRedis) SAdd(cc context.Context, key string, members ...interface{}) *IntCmd {
	return c.client.SAdd(cc, key, c.marshalSlice(members)...)
}
func (c *CodecRedis) SDiff(cc context.Context, keys ...string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SDiff(cc, keys...))
}

func (c *CodecRedis) SInter(cc context.Context, keys ...string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SInter(cc, keys...))
}

func (c *CodecRedis) SIsMember(cc context.Context, key string, member interface{}) *BoolCmd {
	return c.client.SIsMember(cc, key, c.mustMarshal(member))
}
func (c *CodecRedis) SMembers(cc context.Context, key string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SMembers(cc, key))
}
func (c *CodecRedis) SMembersMap(cc context.Context, key string) *CodecSetCmd {
	return newCodecSetCmd(cc, c, c.client.SMembersMap(cc, key))
}
func (c *CodecRedis) SMove(cc context.Context, src, dest string, member interface{}) *BoolCmd {
	return c.client.SMove(cc, src, dest, c.mustMarshal(member))
}
func (c *CodecRedis) SPop(cc context.Context, key string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.SPop(cc, key))
}
func (c *CodecRedis) SPopN(cc context.Context, key string, count int64) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SPopN(cc, key, count))
}
func (c *CodecRedis) SRandMember(cc context.Context, key string) *CodecCmd {
	return newCodecCmd(cc, c, c.client.SRandMember(cc, key))
}
func (c *CodecRedis) SRandMemberN(cc context.Context, key string, count int64) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SRandMemberN(cc, key, count))
}
func (c *CodecRedis) SRem(cc context.Context, key string, members ...interface{}) *IntCmd {
	return c.client.SRem(cc, key, c.marshalSlice(members)...)
}
func (c *CodecRedis) SUnion(cc context.Context, keys ...string) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.SUnion(cc, keys...))
}

// func (c *CodecRedis)SUnionStore(cc context.Context, destination string, keys ...string) *IntCmd // same as Cmdable

// ***** redis set end *****

// ***** redis zset begin *****

func (c *CodecRedis) ZAdd(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAdd(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddNX(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAddNX(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddXX(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAddXX(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddCh(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAddCh(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddNXCh(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAddNXCh(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZAddXXCh(cc context.Context, key string, members ...ZMember) *IntCmd {
	return c.client.ZAddXXCh(cc, key, c.marshalZMembers(members)...)
}
func (c *CodecRedis) ZIncr(cc context.Context, key string, member ZMember) *FloatCmd {
	return c.client.ZIncr(cc, key, c.marshalZMember(member))
}
func (c *CodecRedis) ZIncrNX(cc context.Context, key string, member ZMember) *FloatCmd {
	return c.client.ZIncrNX(cc, key, c.marshalZMember(member))
}
func (c *CodecRedis) ZIncrXX(cc context.Context, key string, member ZMember) *FloatCmd {
	return c.client.ZIncrXX(cc, key, c.marshalZMember(member))
}
func (c *CodecRedis) ZPopMax(cc context.Context, key string, count ...int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZPopMax(cc, key, count...))
}
func (c *CodecRedis) ZIncrBy(cc context.Context, key string, increment float64, member interface{}) *FloatCmd {
	cmd := redis.NewFloatCmd(cc, cc, "zincrby", key, increment, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *CodecRedis) ZPopMin(cc context.Context, key string, count ...int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZPopMin(cc, key, count...))
}
func (c *CodecRedis) ZRange(cc context.Context, key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRange(cc, key, start, stop))
}
func (c *CodecRedis) ZRangeWithScores(cc context.Context, key string, start, stop int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZRangeWithScores(cc, key, start, stop))
}
func (c *CodecRedis) ZRangeByScore(cc context.Context, key string, opt *ZRangeBy) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRangeByScore(cc, key, opt))
}
func (c *CodecRedis) ZRangeByLex(cc context.Context, key string, opt *ZRangeBy) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRangeByLex(cc, key, opt))
}
func (c *CodecRedis) ZRangeByScoreWithScores(cc context.Context, key string, opt *ZRangeBy) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZRangeByScoreWithScores(cc, key, opt))
}
func (c *CodecRedis) ZRank(cc context.Context, key, member interface{}) *IntCmd {
	cmd := redis.NewIntCmd(cc, "zrank", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *CodecRedis) ZRem(cc context.Context, key string, members ...interface{}) *IntCmd {
	return c.client.ZRem(cc, key, c.marshalSlice(members)...)
}
func (c *CodecRedis) ZRevRange(cc context.Context, key string, start, stop int64) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRevRange(cc, key, start, stop))
}
func (c *CodecRedis) ZRevRangeWithScores(cc context.Context, key string, start, stop int64) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZRevRangeWithScores(cc, key, start, stop))
}
func (c *CodecRedis) ZRevRangeByScore(cc context.Context, key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRevRangeByScore(cc, key, opt))
}
func (c *CodecRedis) ZRevRangeByLex(cc context.Context, key string, opt ZRangeBy) *CodecsCmd {
	return newCodecsCmd(cc, c, c.client.ZRevRangeByLex(cc, key, opt))
}
func (c *CodecRedis) ZRevRangeByScoreWithScores(cc context.Context, key string, opt ZRangeBy) *ZCodecSliceCmd {
	return newZCodecSliceCmd(cc, c, c.client.ZRevRangeByScoreWithScores(cc, key, opt))
}
func (c *CodecRedis) ZRevRank(cc context.Context, key, member interface{}) *IntCmd {
	cmd := redis.NewIntCmd(cc, "zrevrank", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *CodecRedis) ZScore(cc context.Context, key, member interface{}) *FloatCmd {
	cmd := redis.NewFloatCmd(cc, "zscore", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}

// ***** redis zset end *****
