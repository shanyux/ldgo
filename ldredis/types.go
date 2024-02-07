/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type (
	Cmd   = redis.Cmd
	Cmder = redis.Cmder

	PubSub       = redis.PubSub
	Message      = redis.Message
	Subscription = redis.Subscription

	Sort         = redis.Sort
	ScanCmd      = redis.ScanCmd
	ScanIterator = redis.ScanIterator

	BoolCmd   = redis.BoolCmd
	StatusCmd = redis.StatusCmd

	DurationCmd = redis.DurationCmd
	StringCmd   = redis.StringCmd
	IntCmd      = redis.IntCmd
	FloatCmd    = redis.FloatCmd
	SliceCmd    = redis.SliceCmd

	StringsCmd         = redis.StringSliceCmd
	MapStringStringCmd = redis.MapStringStringCmd
	StringSetCmd       = redis.StringStructMapCmd

	ZMember     = redis.Z
	ZAddArgs    = redis.ZAddArgs
	ZStore      = redis.ZStore
	ZRangeBy    = redis.ZRangeBy
	ZSliceCmd   = redis.ZSliceCmd
	ZWithKey    = redis.ZWithKey
	ZWithKeyCmd = redis.ZWithKeyCmd

	XAddArgs         = redis.XAddArgs
	XReadArgs        = redis.XReadArgs
	XReadGroupArgs   = redis.XReadGroupArgs
	XClaimArgs       = redis.XClaimArgs
	XPending         = redis.XPending
	XPendingCmd      = redis.XPendingCmd
	XPendingExt      = redis.XPendingExt
	XPendingExtArgs  = redis.XPendingExtArgs
	XPendingExtCmd   = redis.XPendingExtCmd
	XMessage         = redis.XMessage
	XMessageSliceCmd = redis.XMessageSliceCmd
	XStream          = redis.XStream
	XStreamSliceCmd  = redis.XStreamSliceCmd

	GeoLocation    = redis.GeoLocation
	GeoLocationCmd = redis.GeoLocationCmd
	GeoPos         = redis.GeoPos
	GeoPosCmd      = redis.GeoPosCmd
	GeoRadiusQuery = redis.GeoRadiusQuery
)

type Cmdable interface {
	redis.Cmdable

	Do(ctx context.Context, args ...interface{}) *Cmd
	Process(ctx context.Context, cmd Cmder) error
	Close() error
	// Discard() error
	// Exec() ([]Cmder, error)

	Subscribe(ctx context.Context, channels ...string) *PubSub
	PSubscribe(ctx context.Context, channels ...string) *PubSub
}

type CodecCmdable interface {
	Get(c context.Context, key string) *CodecCmd                                               //
	GetBit(c context.Context, key string, offset int64) *IntCmd                                // same as Cmdable
	GetRange(c context.Context, key string, start, end int64) *StringCmd                       // same as Cmdable
	GetSet(c context.Context, key string, value interface{}) *CodecCmd                         //
	Incr(c context.Context, key string) *IntCmd                                                // same as Cmdable
	IncrBy(c context.Context, key string, value int64) *IntCmd                                 // same as Cmdable
	IncrByFloat(c context.Context, key string, value float64) *FloatCmd                        // same as Cmdable
	MGet(c context.Context, keys ...string) *CodecSliceCmd                                     //
	MSet(c context.Context, pairs ...interface{}) *StatusCmd                                   //
	MSetNX(c context.Context, pairs ...interface{}) *BoolCmd                                   //
	Set(c context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd //
	SetBit(c context.Context, key string, offset int64, value int) *IntCmd                     // same as Cmdable
	SetNX(c context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd //
	SetXX(c context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd //
	SetRange(c context.Context, key string, offset int64, value string) *IntCmd                // same as Cmdable
	StrLen(c context.Context, key string) *IntCmd                                              // same as Cmdable

	HDel(c context.Context, key string, fields ...string) *IntCmd                // same as Cmdable
	HExists(c context.Context, key, field string) *BoolCmd                       // same as Cmdable
	HGet(c context.Context, key, field string) *CodecCmd                         //
	HGetAll(c context.Context, key string) *MapStringCodecCmd                    //
	HIncrBy(c context.Context, key, field string, incr int64) *IntCmd            // same as Cmdable
	HIncrByFloat(c context.Context, key, field string, incr float64) *FloatCmd   // same as Cmdable
	HKeys(c context.Context, key string) *StringsCmd                             // same as Cmdable
	HLen(c context.Context, key string) *IntCmd                                  // same as Cmdable
	HMGet(c context.Context, key string, fields ...string) *CodecSliceCmd        //
	HMSet(c context.Context, key string, fields map[string]interface{}) *BoolCmd //
	HSet(c context.Context, key, field string, value interface{}) *IntCmd        //
	HSetNX(c context.Context, key, field string, value interface{}) *BoolCmd     //
	HVals(c context.Context, key string) *CodecsCmd                              //

	BLPop(c context.Context, timeout time.Duration, keys ...string) *CodecsCmd                 //
	BRPop(c context.Context, timeout time.Duration, keys ...string) *CodecsCmd                 //
	BRPopLPush(c context.Context, source, destination string, timeout time.Duration) *CodecCmd //
	LIndex(c context.Context, key string, index int64) *CodecCmd                               //
	LInsert(c context.Context, key, op string, pivot, value interface{}) *IntCmd               //
	LInsertBefore(c context.Context, key string, pivot, value interface{}) *IntCmd             //
	LInsertAfter(c context.Context, key string, pivot, value interface{}) *IntCmd              //
	LLen(c context.Context, key string) *IntCmd                                                // same as Cmdable
	LPop(c context.Context, key string) *CodecCmd                                              //
	LPush(c context.Context, key string, values ...interface{}) *IntCmd                        //
	LPushX(c context.Context, key string, value interface{}) *IntCmd                           //
	LRange(c context.Context, key string, start, stop int64) *CodecsCmd                        //
	LRem(c context.Context, key string, count int64, value interface{}) *IntCmd                //
	LSet(c context.Context, key string, index int64, value interface{}) *StatusCmd             //
	LTrim(c context.Context, key string, start, stop int64) *StatusCmd                         // same as Cmdable
	RPop(c context.Context, key string) *CodecCmd                                              //
	RPopLPush(c context.Context, source, destination string) *CodecCmd                         //
	RPush(c context.Context, key string, values ...interface{}) *IntCmd                        //
	RPushX(c context.Context, key string, value interface{}) *IntCmd                           //

	SAdd(c context.Context, key string, members ...interface{}) *IntCmd               //
	SCard(c context.Context, key string) *IntCmd                                      // same as Cmdable
	SDiff(c context.Context, keys ...string) *CodecsCmd                               //
	SDiffStore(c context.Context, destination string, keys ...string) *IntCmd         // same as Cmdable
	SInter(c context.Context, keys ...string) *CodecsCmd                              //
	SInterStore(c context.Context, destination string, keys ...string) *IntCmd        // same as Cmdable
	SIsMember(c context.Context, key string, member interface{}) *BoolCmd             //
	SMembers(c context.Context, key string) *CodecsCmd                                //
	SMembersMap(c context.Context, key string) *CodecSetCmd                           //
	SMove(c context.Context, source, destination string, member interface{}) *BoolCmd //
	SPop(c context.Context, key string) *CodecCmd                                     //
	SPopN(c context.Context, key string, count int64) *CodecsCmd                      //
	SRandMember(c context.Context, key string) *CodecCmd                              //
	SRandMemberN(c context.Context, key string, count int64) *CodecsCmd               //
	SRem(c context.Context, key string, members ...interface{}) *IntCmd               //
	SUnion(c context.Context, keys ...string) *CodecsCmd                              //
	SUnionStore(c context.Context, destination string, keys ...string) *IntCmd        // same as Cmdable

	// XAdd(a *XAddArgs) *StringCmd

	ZAdd(c context.Context, key string, members ...ZMember) *IntCmd                          //
	ZAddNX(c context.Context, key string, members ...ZMember) *IntCmd                        //
	ZAddXX(c context.Context, key string, members ...ZMember) *IntCmd                        //
	ZCard(c context.Context, key string) *IntCmd                                             // same as Cmdable
	ZCount(c context.Context, key, min, max string) *IntCmd                                  // same as Cmdable
	ZLexCount(c context.Context, key, min, max string) *IntCmd                               // same as Cmdable
	ZIncrBy(c context.Context, key string, increment float64, member interface{}) *FloatCmd  //
	ZInterStore(c context.Context, destination string, store *ZStore) *IntCmd                // same as Cmdable
	ZPopMax(c context.Context, key string, count ...int64) *ZCodecSliceCmd                   //
	ZPopMin(c context.Context, key string, count ...int64) *ZCodecSliceCmd                   //
	ZRange(c context.Context, key string, start, stop int64) *CodecsCmd                      //
	ZRangeWithScores(c context.Context, key string, start, stop int64) *ZCodecSliceCmd       //
	ZRangeByScore(c context.Context, key string, opt *ZRangeBy) *CodecsCmd                   //
	ZRangeByLex(c context.Context, key string, opt *ZRangeBy) *CodecsCmd                     //
	ZRangeByScoreWithScores(c context.Context, key string, opt *ZRangeBy) *ZCodecSliceCmd    //
	ZRank(c context.Context, key, member interface{}) *IntCmd                                //
	ZRem(c context.Context, key string, members ...interface{}) *IntCmd                      //
	ZRemRangeByRank(c context.Context, key string, start, stop int64) *IntCmd                // same as Cmdable
	ZRemRangeByScore(c context.Context, key, min, max string) *IntCmd                        // same as Cmdable
	ZRemRangeByLex(c context.Context, key, min, max string) *IntCmd                          // same as Cmdable
	ZRevRange(c context.Context, key string, start, stop int64) *CodecsCmd                   //
	ZRevRangeWithScores(c context.Context, key string, start, stop int64) *ZCodecSliceCmd    //
	ZRevRangeByScore(c context.Context, key string, opt *ZRangeBy) *CodecsCmd                //
	ZRevRangeByLex(c context.Context, key string, opt *ZRangeBy) *CodecsCmd                  //
	ZRevRangeByScoreWithScores(c context.Context, key string, opt *ZRangeBy) *ZCodecSliceCmd //
	ZRevRank(c context.Context, key, member interface{}) *IntCmd                             //
	ZScore(c context.Context, key, member interface{}) *FloatCmd                             //
}
