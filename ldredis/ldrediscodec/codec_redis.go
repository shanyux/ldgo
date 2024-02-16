/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldredis"
	redis "github.com/redis/go-redis/v9"
)

func New[T comparable](rds *ldredis.Redis, codec Codec[T]) *Redis[T] {
	return &Redis[T]{
		base: base[T]{
			client: rds,
			codec:  codec,
		},
	}
}

var _ Cmdable[any] = (*Redis[any])(nil)

type Redis[T comparable] struct {
	base[T]
}

func (c *Redis[T]) Client() *ldredis.Redis { return c.client }

func (c *Redis[T]) clone(cli ...*ldredis.Redis) *Redis[T] {
	cp := *c
	c = &cp
	if len(cli) > 0 && cli[0] != nil {
		cp.client = cli[0]
	}
	return c
}

func (c *Redis[T]) WithRetry(retry int) *Redis[T] {
	return c.clone(c.client.WithRetry(retry))
}
func (c *Redis[T]) WithReport(r Reporter) *Redis[T] {
	return c.clone(c.client.WithReport(r))
}
func (c *Redis[T]) WithCaller(enable bool) *Redis[T] {
	return c.clone(c.client.WithCaller(enable))
}

// ***** redis string begin *****

func (c *Redis[T]) MGet(cc context.Context, keys ...string) *SliceCmd[T] {
	return newSliceCmd(cc, c, c.client.MGet(cc, keys...))
}
func (c *Redis[T]) MSet(cc context.Context, pairs ...interface{}) *ldredis.StatusCmd {
	return c.client.MSet(cc, c.marshalPairs(pairs)...)
}
func (c *Redis[T]) MSetNX(cc context.Context, pairs ...interface{}) *ldredis.BoolCmd {
	return c.client.MSetNX(cc, c.marshalPairs(pairs)...)
}

func (c *Redis[T]) Get(cc context.Context, key string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.Get(cc, key))
}
func (c *Redis[T]) GetSet(cc context.Context, key string, value T) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.GetSet(cc, key, c.mustMarshal(value)))
}
func (c *Redis[T]) Set(cc context.Context, key string, value T, expiration time.Duration) *ldredis.StatusCmd {
	return c.client.Set(cc, key, c.mustMarshal(value), expiration)
}
func (c *Redis[T]) SetNX(cc context.Context, key string, value T, expiration time.Duration) *ldredis.BoolCmd {
	return c.client.SetNX(cc, key, c.mustMarshal(value), expiration)
}
func (c *Redis[T]) SetXX(cc context.Context, key string, value T, expiration time.Duration) *ldredis.BoolCmd {
	return c.client.SetXX(cc, key, c.mustMarshal(value), expiration)
}

// ***** redis string end *****

// ***** redis hash begin *****

func (c *Redis[T]) HGet(cc context.Context, key, field string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.HGet(cc, key, field))
}
func (c *Redis[T]) HGetAll(cc context.Context, key string) *MapStringAnyCmd[T] {
	return newMapStringAnyCmd[T](cc, c, c.client.HGetAll(cc, key))
}
func (c *Redis[T]) HMGet(cc context.Context, key string, fields ...string) *SliceCmd[T] {
	return newSliceCmd(cc, c, c.client.HMGet(cc, key, fields...))
}
func (c *Redis[T]) HMSet(cc context.Context, key string, fields map[string]T) *ldredis.BoolCmd {
	return c.client.HMSet(cc, key, c.marshalMap(fields)...)
}
func (c *Redis[T]) HSet(cc context.Context, key, field string, value T) *ldredis.IntCmd {
	return c.client.HSet(cc, key, field, c.mustMarshal(value))
}
func (c *Redis[T]) HSetNX(cc context.Context, key, field string, value T) *ldredis.BoolCmd {
	return c.client.HSetNX(cc, key, field, c.mustMarshal(value))
}
func (c *Redis[T]) HVals(cc context.Context, key string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.HVals(cc, key))
}

// ***** redis hash end *****

// ***** redis list begin *****

func (c *Redis[T]) BLPop(cc context.Context, timeout time.Duration, keys ...string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.BLPop(cc, timeout, keys...))
}
func (c *Redis[T]) BRPop(cc context.Context, timeout time.Duration, keys ...string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.BRPop(cc, timeout, keys...))
}
func (c *Redis[T]) BRPopLPush(cc context.Context, source, destination string, timeout time.Duration) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.BRPopLPush(cc, source, destination, timeout))
}
func (c *Redis[T]) LIndex(cc context.Context, key string, index int64) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.LIndex(cc, key, index))
}
func (c *Redis[T]) LInsert(cc context.Context, key, op string, pivot, value T) *ldredis.IntCmd {
	return c.client.LInsert(cc, key, op, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *Redis[T]) LInsertBefore(cc context.Context, key string, pivot, value T) *ldredis.IntCmd {
	return c.client.LInsertBefore(cc, key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *Redis[T]) LInsertAfter(cc context.Context, key string, pivot, value T) *ldredis.IntCmd {
	return c.client.LInsertAfter(cc, key, c.mustMarshal(pivot), c.mustMarshal(value))
}
func (c *Redis[T]) LPop(cc context.Context, key string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.LPop(cc, key))
}
func (c *Redis[T]) LPush(cc context.Context, key string, values ...T) *ldredis.IntCmd {
	return c.client.LPush(cc, key, c.marshalSlice(values)...)
}
func (c *Redis[T]) LPushX(cc context.Context, key string, value T) *ldredis.IntCmd {
	return c.client.LPushX(cc, key, c.mustMarshal(value))
}
func (c *Redis[T]) LRange(cc context.Context, key string, start, stop int64) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.LRange(cc, key, start, stop))
}
func (c *Redis[T]) LRem(cc context.Context, key string, count int64, value T) *ldredis.IntCmd {
	return c.client.LRem(cc, key, count, c.mustMarshal(value))
}
func (c *Redis[T]) LSet(cc context.Context, key string, index int64, value T) *ldredis.StatusCmd {
	return c.client.LSet(cc, key, index, c.mustMarshal(value))
}
func (c *Redis[T]) RPop(cc context.Context, key string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.RPop(cc, key))
}
func (c *Redis[T]) RPopLPush(cc context.Context, source, destination string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.RPopLPush(cc, source, destination))
}
func (c *Redis[T]) RPush(cc context.Context, key string, values ...T) *ldredis.IntCmd {
	return c.client.RPush(cc, key, c.marshalSlice(values)...)
}
func (c *Redis[T]) RPushX(cc context.Context, key string, value T) *ldredis.IntCmd {
	return c.client.RPushX(cc, key, c.mustMarshal(value))
}

// ***** redis list begin *****

// ***** redis set begin *****

func (c *Redis[T]) SAdd(cc context.Context, key string, members ...T) *ldredis.IntCmd {
	return c.client.SAdd(cc, key, c.marshalSlice(members)...)
}
func (c *Redis[T]) SDiff(cc context.Context, keys ...string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SDiff(cc, keys...))
}

func (c *Redis[T]) SInter(cc context.Context, keys ...string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SInter(cc, keys...))
}

func (c *Redis[T]) SIsMember(cc context.Context, key string, member T) *ldredis.BoolCmd {
	return c.client.SIsMember(cc, key, c.mustMarshal(member))
}
func (c *Redis[T]) SMembers(cc context.Context, key string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SMembers(cc, key))
}
func (c *Redis[T]) SMembersMap(cc context.Context, key string) *AnySetCmd[T] {
	return newAnySetCmd[T](cc, c, c.client.SMembersMap(cc, key))
}
func (c *Redis[T]) SMove(cc context.Context, src, dest string, member T) *ldredis.BoolCmd {
	return c.client.SMove(cc, src, dest, c.mustMarshal(member))
}
func (c *Redis[T]) SPop(cc context.Context, key string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.SPop(cc, key))
}
func (c *Redis[T]) SPopN(cc context.Context, key string, count int64) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SPopN(cc, key, count))
}
func (c *Redis[T]) SRandMember(cc context.Context, key string) *AnyCmd[T] {
	return newAnyCmd[T](cc, c, c.client.SRandMember(cc, key))
}
func (c *Redis[T]) SRandMemberN(cc context.Context, key string, count int64) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SRandMemberN(cc, key, count))
}
func (c *Redis[T]) SRem(cc context.Context, key string, members ...T) *ldredis.IntCmd {
	return c.client.SRem(cc, key, c.marshalSlice(members)...)
}
func (c *Redis[T]) SUnion(cc context.Context, keys ...string) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.SUnion(cc, keys...))
}

// func (c *CodecRedis[T])SUnionStore(cc context.Context, destination string, keys ...string) *IntCmd // same as Cmdable

// ***** redis set end *****

// ***** redis zset begin *****

func (c *Redis[T]) ZAdd(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd {
	return c.client.ZAdd(cc, key, c.marshalZMembers(members)...)
}
func (c *Redis[T]) ZAddNX(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd {
	return c.client.ZAddNX(cc, key, c.marshalZMembers(members)...)
}
func (c *Redis[T]) ZAddXX(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd {
	return c.client.ZAddXX(cc, key, c.marshalZMembers(members)...)
}
func (c *Redis[T]) ZPopMax(cc context.Context, key string, count ...int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZPopMax(cc, key, count...))
}
func (c *Redis[T]) ZIncrBy(cc context.Context, key string, increment float64, member T) *ldredis.FloatCmd {
	// return c.client.ZIncrBy(cc, key, increment, c.mustMarshal(member))
	cmd := redis.NewFloatCmd(cc, "zincrby", key, increment, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *Redis[T]) ZPopMin(cc context.Context, key string, count ...int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZPopMin(cc, key, count...))
}
func (c *Redis[T]) ZRange(cc context.Context, key string, start, stop int64) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRange(cc, key, start, stop))
}
func (c *Redis[T]) ZRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZRangeWithScores(cc, key, start, stop))
}
func (c *Redis[T]) ZRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRangeByScore(cc, key, opt))
}
func (c *Redis[T]) ZRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRangeByLex(cc, key, opt))
}
func (c *Redis[T]) ZRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZRangeByScoreWithScores(cc, key, opt))
}
func (c *Redis[T]) ZRank(cc context.Context, key, member T) *ldredis.IntCmd {
	cmd := redis.NewIntCmd(cc, "zrank", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *Redis[T]) ZRem(cc context.Context, key string, members ...T) *ldredis.IntCmd {
	return c.client.ZRem(cc, key, c.marshalSlice(members)...)
}
func (c *Redis[T]) ZRevRange(cc context.Context, key string, start, stop int64) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRevRange(cc, key, start, stop))
}
func (c *Redis[T]) ZRevRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZRevRangeWithScores(cc, key, start, stop))
}
func (c *Redis[T]) ZRevRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRevRangeByScore(cc, key, opt))
}
func (c *Redis[T]) ZRevRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *AnySliceCmd[T] {
	return newAnySliceCmd[T](cc, c, c.client.ZRevRangeByLex(cc, key, opt))
}
func (c *Redis[T]) ZRevRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c, c.client.ZRevRangeByScoreWithScores(cc, key, opt))
}
func (c *Redis[T]) ZRevRank(cc context.Context, key, member T) *ldredis.IntCmd {
	cmd := redis.NewIntCmd(cc, "zrevrank", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *Redis[T]) ZScore(cc context.Context, key, member T) *ldredis.FloatCmd {
	cmd := redis.NewFloatCmd(cc, "zscore", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}

// ***** redis zset end *****
