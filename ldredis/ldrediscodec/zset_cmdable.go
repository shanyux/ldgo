/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldredis"
	"github.com/redis/go-redis/v9"
)

type SortedSetCmdable[T any] interface {
	BZPopMax(cc context.Context, timeout time.Duration, keys ...string) *ldredis.ZWithKeyCmd
	BZPopMin(cc context.Context, timeout time.Duration, keys ...string) *ldredis.ZWithKeyCmd
	BZMPop(cc context.Context, timeout time.Duration, order string, count int64, keys ...string) *ldredis.ZSliceWithKeyCmd
	ZAdd(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd //
	ZAddLT(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd
	ZAddGT(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd
	ZAddNX(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd //
	ZAddXX(cc context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd //
	ZAddArgs(cc context.Context, key string, args ldredis.ZAddArgs) *ldredis.IntCmd
	ZAddArgsIncr(cc context.Context, key string, args ldredis.ZAddArgs) *ldredis.FloatCmd
	ZCard(cc context.Context, key string) *ldredis.IntCmd                                  // same as ldredis.Cmdable
	ZCount(cc context.Context, key, min, max string) *ldredis.IntCmd                       // same as ldredis.Cmdable
	ZLexCount(cc context.Context, key, min, max string) *ldredis.IntCmd                    // same as ldredis.Cmdable
	ZIncrBy(cc context.Context, key string, increment float64, member T) *ldredis.FloatCmd //
	ZInter(cc context.Context, store *ldredis.ZStore) *ldredis.StringSliceCmd
	ZInterWithScores(cc context.Context, store *ldredis.ZStore) *ldredis.ZSliceCmd
	ZInterCard(cc context.Context, limit int64, keys ...string) *ldredis.IntCmd
	ZInterStore(cc context.Context, destination string, store *ldredis.ZStore) *ldredis.IntCmd // same as ldredis.Cmdable
	ZMPop(cc context.Context, order string, count int64, keys ...string) *ldredis.ZSliceWithKeyCmd
	ZMScore(cc context.Context, key string, members ...T) *ldredis.FloatSliceCmd
	ZPopMax(cc context.Context, key string, count ...int64) *ZMemberSliceCmd[T]                        //
	ZPopMin(cc context.Context, key string, count ...int64) *ZMemberSliceCmd[T]                        //
	ZRange(cc context.Context, key string, start, stop int64) *TypeSliceCmd[T]                         //
	ZRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T]            //
	ZRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]              //
	ZRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]                //
	ZRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] //
	ZRangeArgs(cc context.Context, z ldredis.ZRangeArgs) *ldredis.StringSliceCmd
	ZRangeArgsWithScores(cc context.Context, z ldredis.ZRangeArgs) *ldredis.ZSliceCmd
	ZRangeStore(cc context.Context, dst string, z ldredis.ZRangeArgs) *ldredis.IntCmd
	ZRank(cc context.Context, key string, member T) *ldredis.IntCmd //
	ZRankWithScore(cc context.Context, key, member string) *ldredis.RankWithScoreCmd
	ZRem(cc context.Context, key string, members ...T) *ldredis.IntCmd                                    //
	ZRemRangeByRank(cc context.Context, key string, start, stop int64) *ldredis.IntCmd                    // same as ldredis.Cmdable
	ZRemRangeByScore(cc context.Context, key, min, max string) *ldredis.IntCmd                            // same as ldredis.Cmdable
	ZRemRangeByLex(cc context.Context, key, min, max string) *ldredis.IntCmd                              // same as ldredis.Cmdable
	ZRevRange(cc context.Context, key string, start, stop int64) *TypeSliceCmd[T]                         //
	ZRevRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T]            //
	ZRevRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]              //
	ZRevRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]                //
	ZRevRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] //
	ZRevRank(cc context.Context, key string, member T) *ldredis.IntCmd                                    //
	ZRevRankWithScore(cc context.Context, key, member string) *ldredis.RankWithScoreCmd
	ZScore(cc context.Context, key string, member T) *ldredis.FloatCmd //
	ZUnionStore(cc context.Context, dest string, store *ldredis.ZStore) *ldredis.IntCmd
	ZRandMember(cc context.Context, key string, count int) *ldredis.StringSliceCmd
	ZRandMemberWithScores(cc context.Context, key string, count int) *ldredis.ZSliceCmd
	ZUnion(cc context.Context, store ldredis.ZStore) *ldredis.StringSliceCmd
	ZUnionWithScores(cc context.Context, store ldredis.ZStore) *ldredis.ZSliceCmd
	ZDiff(cc context.Context, keys ...string) *ldredis.StringSliceCmd
	ZDiffWithScores(cc context.Context, keys ...string) *ldredis.ZSliceCmd
	ZDiffStore(cc context.Context, destination string, keys ...string) *ldredis.IntCmd
	ZScan(cc context.Context, key string, cursor uint64, match string, count int64) *ldredis.ScanCmd // same as ldredis.Cmdable
}

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
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZPopMax(cc, key, count...))
}
func (c *Redis[T]) ZIncrBy(cc context.Context, key string, increment float64, member T) *ldredis.FloatCmd {
	// return c.client.ZIncrBy(cc, key, increment, c.mustMarshal(member))
	cmd := redis.NewFloatCmd(cc, "zincrby", key, increment, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *Redis[T]) ZPopMin(cc context.Context, key string, count ...int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZPopMin(cc, key, count...))
}
func (c *Redis[T]) ZRange(cc context.Context, key string, start, stop int64) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRange(cc, key, start, stop))
}
func (c *Redis[T]) ZRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZRangeWithScores(cc, key, start, stop))
}
func (c *Redis[T]) ZRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRangeByScore(cc, key, opt))
}
func (c *Redis[T]) ZRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRangeByLex(cc, key, opt))
}
func (c *Redis[T]) ZRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZRangeByScoreWithScores(cc, key, opt))
}
func (c *Redis[T]) ZRank(cc context.Context, key, member T) *ldredis.IntCmd {
	cmd := redis.NewIntCmd(cc, "zrank", key, c.mustMarshal(member))
	c.client.Process(cc, cmd)
	return cmd
}
func (c *Redis[T]) ZRem(cc context.Context, key string, members ...T) *ldredis.IntCmd {
	return c.client.ZRem(cc, key, c.marshalSlice(members)...)
}
func (c *Redis[T]) ZRevRange(cc context.Context, key string, start, stop int64) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRevRange(cc, key, start, stop))
}
func (c *Redis[T]) ZRevRangeWithScores(cc context.Context, key string, start, stop int64) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZRevRangeWithScores(cc, key, start, stop))
}
func (c *Redis[T]) ZRevRangeByScore(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRevRangeByScore(cc, key, opt))
}
func (c *Redis[T]) ZRevRangeByLex(cc context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T] {
	return newTypeSliceCmd[T](cc, c.base, c.client.ZRevRangeByLex(cc, key, opt))
}
func (c *Redis[T]) ZRevRangeByScoreWithScores(cc context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] {
	return newZMemberSliceCmd[T](cc, c.base, c.client.ZRevRangeByScoreWithScores(cc, key, opt))
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
