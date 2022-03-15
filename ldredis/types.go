/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/go-redis/redis"
)

type (
	Context = ldctx.Context
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
	StringStringMapCmd = redis.StringStringMapCmd
	StringSetCmd       = redis.StringStructMapCmd

	ZMember     = redis.Z
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

	Do(args ...interface{}) *Cmd
	Process(cmd Cmder) error
	Close() error
	// Discard() error
	// Exec() ([]Cmder, error)

	Subscribe(channels ...string) *PubSub
	PSubscribe(channels ...string) *PubSub
}

type CodecCmdable interface {
	Get(key string) *CodecCmd                                               //
	GetBit(key string, offset int64) *IntCmd                                // same as Cmdable
	GetRange(key string, start, end int64) *StringCmd                       // same as Cmdable
	GetSet(key string, value interface{}) *CodecCmd                         //
	Incr(key string) *IntCmd                                                // same as Cmdable
	IncrBy(key string, value int64) *IntCmd                                 // same as Cmdable
	IncrByFloat(key string, value float64) *FloatCmd                        // same as Cmdable
	MGet(keys ...string) *CodecSliceCmd                                     //
	MSet(pairs ...interface{}) *StatusCmd                                   //
	MSetNX(pairs ...interface{}) *BoolCmd                                   //
	Set(key string, value interface{}, expiration time.Duration) *StatusCmd //
	SetBit(key string, offset int64, value int) *IntCmd                     // same as Cmdable
	SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd //
	SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd //
	SetRange(key string, offset int64, value string) *IntCmd                // same as Cmdable
	StrLen(key string) *IntCmd                                              // same as Cmdable

	HDel(key string, fields ...string) *IntCmd                  // same as Cmdable
	HExists(key, field string) *BoolCmd                         // same as Cmdable
	HGet(key, field string) *CodecCmd                           //
	HGetAll(key string) *StringCodecMapCmd                      //
	HIncrBy(key, field string, incr int64) *IntCmd              // same as Cmdable
	HIncrByFloat(key, field string, incr float64) *FloatCmd     // same as Cmdable
	HKeys(key string) *StringsCmd                               // same as Cmdable
	HLen(key string) *IntCmd                                    // same as Cmdable
	HMGet(key string, fields ...string) *CodecSliceCmd          //
	HMSet(key string, fields map[string]interface{}) *StatusCmd //
	HSet(key, field string, value interface{}) *BoolCmd         //
	HSetNX(key, field string, value interface{}) *BoolCmd       //
	HVals(key string) *CodecsCmd                                //

	BLPop(timeout time.Duration, keys ...string) *CodecsCmd                 //
	BRPop(timeout time.Duration, keys ...string) *CodecsCmd                 //
	BRPopLPush(source, destination string, timeout time.Duration) *CodecCmd //
	LIndex(key string, index int64) *CodecCmd                               //
	LInsert(key, op string, pivot, value interface{}) *IntCmd               //
	LInsertBefore(key string, pivot, value interface{}) *IntCmd             //
	LInsertAfter(key string, pivot, value interface{}) *IntCmd              //
	LLen(key string) *IntCmd                                                // same as Cmdable
	LPop(key string) *CodecCmd                                              //
	LPush(key string, values ...interface{}) *IntCmd                        //
	LPushX(key string, value interface{}) *IntCmd                           //
	LRange(key string, start, stop int64) *CodecsCmd                        //
	LRem(key string, count int64, value interface{}) *IntCmd                //
	LSet(key string, index int64, value interface{}) *StatusCmd             //
	LTrim(key string, start, stop int64) *StatusCmd                         // same as Cmdable
	RPop(key string) *CodecCmd                                              //
	RPopLPush(source, destination string) *CodecCmd                         //
	RPush(key string, values ...interface{}) *IntCmd                        //
	RPushX(key string, value interface{}) *IntCmd                           //

	SAdd(key string, members ...interface{}) *IntCmd               //
	SCard(key string) *IntCmd                                      // same as Cmdable
	SDiff(keys ...string) *CodecsCmd                               //
	SDiffStore(destination string, keys ...string) *IntCmd         // same as Cmdable
	SInter(keys ...string) *CodecsCmd                              //
	SInterStore(destination string, keys ...string) *IntCmd        // same as Cmdable
	SIsMember(key string, member interface{}) *BoolCmd             //
	SMembers(key string) *CodecsCmd                                //
	SMembersMap(key string) *StringSetCmd                          // same as Cmdable
	SMove(source, destination string, member interface{}) *BoolCmd //
	SPop(key string) *CodecCmd                                     //
	SPopN(key string, count int64) *CodecsCmd                      //
	SRandMember(key string) *CodecCmd                              //
	SRandMemberN(key string, count int64) *CodecsCmd               //
	SRem(key string, members ...interface{}) *IntCmd               //
	SUnion(keys ...string) *CodecsCmd                              //
	SUnionStore(destination string, keys ...string) *IntCmd        // same as Cmdable

	// XAdd(a *XAddArgs) *StringCmd

	ZAdd(key string, members ...ZMember) *IntCmd                          //
	ZAddNX(key string, members ...ZMember) *IntCmd                        //
	ZAddXX(key string, members ...ZMember) *IntCmd                        //
	ZAddCh(key string, members ...ZMember) *IntCmd                        //
	ZAddNXCh(key string, members ...ZMember) *IntCmd                      //
	ZAddXXCh(key string, members ...ZMember) *IntCmd                      //
	ZIncr(key string, member ZMember) *FloatCmd                           //
	ZIncrNX(key string, member ZMember) *FloatCmd                         //
	ZIncrXX(key string, member ZMember) *FloatCmd                         //
	ZCard(key string) *IntCmd                                             // same as Cmdable
	ZCount(key, min, max string) *IntCmd                                  // same as Cmdable
	ZLexCount(key, min, max string) *IntCmd                               // same as Cmdable
	ZIncrBy(key string, increment float64, member interface{}) *FloatCmd  //
	ZInterStore(destination string, store ZStore, keys ...string) *IntCmd // same as Cmdable
	ZPopMax(key string, count ...int64) *ZCodecSliceCmd                   //
	ZPopMin(key string, count ...int64) *ZCodecSliceCmd                   //
	ZRange(key string, start, stop int64) *CodecsCmd                      //
	ZRangeWithScores(key string, start, stop int64) *ZCodecSliceCmd       //
	ZRangeByScore(key string, opt ZRangeBy) *CodecsCmd                    //
	ZRangeByLex(key string, opt ZRangeBy) *CodecsCmd                      //
	ZRangeByScoreWithScores(key string, opt ZRangeBy) *ZCodecSliceCmd     //
	ZRank(key, member interface{}) *IntCmd                                //
	ZRem(key string, members ...interface{}) *IntCmd                      //
	ZRemRangeByRank(key string, start, stop int64) *IntCmd                // same as Cmdable
	ZRemRangeByScore(key, min, max string) *IntCmd                        // same as Cmdable
	ZRemRangeByLex(key, min, max string) *IntCmd                          // same as Cmdable
	ZRevRange(key string, start, stop int64) *CodecsCmd                   //
	ZRevRangeWithScores(key string, start, stop int64) *ZCodecSliceCmd    //
	ZRevRangeByScore(key string, opt ZRangeBy) *CodecsCmd                 //
	ZRevRangeByLex(key string, opt ZRangeBy) *CodecsCmd                   //
	ZRevRangeByScoreWithScores(key string, opt ZRangeBy) *ZCodecSliceCmd  //
	ZRevRank(key, member interface{}) *IntCmd                             //
	ZScore(key, member interface{}) *FloatCmd                             //
}
