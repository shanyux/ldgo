/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldredis"
)

type (
	client   = ldredis.Redis
	Reporter = ldredis.Reporter
)

type ZMember[T any] struct {
	Score  float64
	Member T
}

type Cmdable[T comparable] interface {
	Get(c context.Context, key string) *TypeCmd[T]                                                     //
	GetBit(c context.Context, key string, offset int64) *ldredis.IntCmd                                // same as ldredis.Cmdable
	GetRange(c context.Context, key string, start, end int64) *ldredis.StringCmd                       // same as ldredis.Cmdable
	GetSet(c context.Context, key string, value T) *TypeCmd[T]                                         //
	Incr(c context.Context, key string) *ldredis.IntCmd                                                // same as ldredis.Cmdable
	IncrBy(c context.Context, key string, value int64) *ldredis.IntCmd                                 // same as ldredis.Cmdable
	IncrByFloat(c context.Context, key string, value float64) *ldredis.FloatCmd                        // same as ldredis.Cmdable
	MGet(c context.Context, keys ...string) *SliceCmd[T]                                               //
	MSet(c context.Context, pairs ...interface{}) *ldredis.StatusCmd                                   //
	MSetNX(c context.Context, pairs ...interface{}) *ldredis.BoolCmd                                   //
	Set(c context.Context, key string, value interface{}, expiration time.Duration) *ldredis.StatusCmd //
	SetBit(c context.Context, key string, offset int64, value int) *ldredis.IntCmd                     // same as ldredis.Cmdable
	SetNX(c context.Context, key string, value interface{}, expiration time.Duration) *ldredis.BoolCmd //
	SetXX(c context.Context, key string, value interface{}, expiration time.Duration) *ldredis.BoolCmd //
	SetRange(c context.Context, key string, offset int64, value string) *ldredis.IntCmd                // same as ldredis.Cmdable
	StrLen(c context.Context, key string) *ldredis.IntCmd                                              // same as ldredis.Cmdable

	HDel(c context.Context, key string, fields ...string) *ldredis.IntCmd                // same as ldredis.Cmdable
	HExists(c context.Context, key, field string) *ldredis.BoolCmd                       // same as ldredis.Cmdable
	HGet(c context.Context, key, field string) *TypeCmd[T]                               //
	HGetAll(c context.Context, key string) *MapStringTypeCmd[T]                          //
	HIncrBy(c context.Context, key, field string, incr int64) *ldredis.IntCmd            // same as ldredis.Cmdable
	HIncrByFloat(c context.Context, key, field string, incr float64) *ldredis.FloatCmd   // same as ldredis.Cmdable
	HKeys(c context.Context, key string) *ldredis.StringSliceCmd                         // same as ldredis.Cmdable
	HLen(c context.Context, key string) *ldredis.IntCmd                                  // same as ldredis.Cmdable
	HMGet(c context.Context, key string, fields ...string) *SliceCmd[T]                  //
	HMSet(c context.Context, key string, fields map[string]interface{}) *ldredis.BoolCmd //
	HSet(c context.Context, key, field string, value interface{}) *ldredis.IntCmd        //
	HSetNX(c context.Context, key, field string, value interface{}) *ldredis.BoolCmd     //
	HVals(c context.Context, key string) *TypeSliceCmd[T]                                //

	BLPop(c context.Context, timeout time.Duration, keys ...string) *TypeSliceCmd[T]             //
	BRPop(c context.Context, timeout time.Duration, keys ...string) *TypeSliceCmd[T]             //
	BRPopLPush(c context.Context, source, destination string, timeout time.Duration) *TypeCmd[T] //
	LIndex(c context.Context, key string, index int64) *TypeCmd[T]                               //
	LInsert(c context.Context, key, op string, pivot, value interface{}) *ldredis.IntCmd         //
	LInsertBefore(c context.Context, key string, pivot, value interface{}) *ldredis.IntCmd       //
	LInsertAfter(c context.Context, key string, pivot, value interface{}) *ldredis.IntCmd        //
	LLen(c context.Context, key string) *ldredis.IntCmd                                          // same as ldredis.Cmdable
	LPop(c context.Context, key string) *TypeCmd[T]                                              //
	LPush(c context.Context, key string, values ...interface{}) *ldredis.IntCmd                  //
	LPushX(c context.Context, key string, value interface{}) *ldredis.IntCmd                     //
	LRange(c context.Context, key string, start, stop int64) *TypeSliceCmd[T]                    //
	LRem(c context.Context, key string, count int64, value interface{}) *ldredis.IntCmd          //
	LSet(c context.Context, key string, index int64, value interface{}) *ldredis.StatusCmd       //
	LTrim(c context.Context, key string, start, stop int64) *ldredis.StatusCmd                   // same as ldredis.Cmdable
	RPop(c context.Context, key string) *TypeCmd[T]                                              //
	RPopLPush(c context.Context, source, destination string) *TypeCmd[T]                         //
	RPush(c context.Context, key string, values ...interface{}) *ldredis.IntCmd                  //
	RPushX(c context.Context, key string, value interface{}) *ldredis.IntCmd                     //

	SAdd(c context.Context, key string, members ...interface{}) *ldredis.IntCmd               //
	SCard(c context.Context, key string) *ldredis.IntCmd                                      // same as ldredis.Cmdable
	SDiff(c context.Context, keys ...string) *TypeSliceCmd[T]                                 //
	SDiffStore(c context.Context, destination string, keys ...string) *ldredis.IntCmd         // same as ldredis.Cmdable
	SInter(c context.Context, keys ...string) *TypeSliceCmd[T]                                //
	SInterStore(c context.Context, destination string, keys ...string) *ldredis.IntCmd        // same as ldredis.Cmdable
	SIsMember(c context.Context, key string, member interface{}) *ldredis.BoolCmd             //
	SMembers(c context.Context, key string) *TypeSliceCmd[T]                                  //
	SMembersMap(c context.Context, key string) *TypeSetCmd[T]                                 //
	SMove(c context.Context, source, destination string, member interface{}) *ldredis.BoolCmd //
	SPop(c context.Context, key string) *TypeCmd[T]                                           //
	SPopN(c context.Context, key string, count int64) *TypeSliceCmd[T]                        //
	SRandMember(c context.Context, key string) *TypeCmd[T]                                    //
	SRandMemberN(c context.Context, key string, count int64) *TypeSliceCmd[T]                 //
	SRem(c context.Context, key string, members ...interface{}) *ldredis.IntCmd               //
	SUnion(c context.Context, keys ...string) *TypeSliceCmd[T]                                //
	SUnionStore(c context.Context, destination string, keys ...string) *ldredis.IntCmd        // same as ldredis.Cmdable

	// XAdd(a *XAddArgs) *StringCmd

	ZAdd(c context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd                           //
	ZAddNX(c context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd                         //
	ZAddXX(c context.Context, key string, members ...ZMember[T]) *ldredis.IntCmd                         //
	ZCard(c context.Context, key string) *ldredis.IntCmd                                                 // same as ldredis.Cmdable
	ZCount(c context.Context, key, min, max string) *ldredis.IntCmd                                      // same as ldredis.Cmdable
	ZLexCount(c context.Context, key, min, max string) *ldredis.IntCmd                                   // same as ldredis.Cmdable
	ZIncrBy(c context.Context, key string, increment float64, member interface{}) *ldredis.FloatCmd      //
	ZInterStore(c context.Context, destination string, store *ldredis.ZStore) *ldredis.IntCmd            // same as ldredis.Cmdable
	ZPopMax(c context.Context, key string, count ...int64) *ZMemberSliceCmd[T]                           //
	ZPopMin(c context.Context, key string, count ...int64) *ZMemberSliceCmd[T]                           //
	ZRange(c context.Context, key string, start, stop int64) *TypeSliceCmd[T]                            //
	ZRangeWithScores(c context.Context, key string, start, stop int64) *ZMemberSliceCmd[T]               //
	ZRangeByScore(c context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]                 //
	ZRangeByLex(c context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]                   //
	ZRangeByScoreWithScores(c context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T]    //
	ZRank(c context.Context, key, member interface{}) *ldredis.IntCmd                                    //
	ZRem(c context.Context, key string, members ...interface{}) *ldredis.IntCmd                          //
	ZRemRangeByRank(c context.Context, key string, start, stop int64) *ldredis.IntCmd                    // same as ldredis.Cmdable
	ZRemRangeByScore(c context.Context, key, min, max string) *ldredis.IntCmd                            // same as ldredis.Cmdable
	ZRemRangeByLex(c context.Context, key, min, max string) *ldredis.IntCmd                              // same as ldredis.Cmdable
	ZRevRange(c context.Context, key string, start, stop int64) *TypeSliceCmd[T]                         //
	ZRevRangeWithScores(c context.Context, key string, start, stop int64) *ZMemberSliceCmd[T]            //
	ZRevRangeByScore(c context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]              //
	ZRevRangeByLex(c context.Context, key string, opt *ldredis.ZRangeBy) *TypeSliceCmd[T]                //
	ZRevRangeByScoreWithScores(c context.Context, key string, opt *ldredis.ZRangeBy) *ZMemberSliceCmd[T] //
	ZRevRank(c context.Context, key, member interface{}) *ldredis.IntCmd                                 //
	ZScore(c context.Context, key, member interface{}) *ldredis.FloatCmd                                 //
}
