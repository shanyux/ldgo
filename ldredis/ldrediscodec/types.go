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

type (
	Cmd   = ldredis.Cmd
	Cmder = ldredis.Cmder

	PubSub       = ldredis.PubSub
	Message      = ldredis.Message
	Subscription = ldredis.Subscription

	Sort         = ldredis.Sort
	ScanCmd      = ldredis.ScanCmd
	ScanIterator = ldredis.ScanIterator

	BoolCmd   = ldredis.BoolCmd
	StatusCmd = ldredis.StatusCmd

	DurationCmd = ldredis.DurationCmd
	StringCmd   = ldredis.StringCmd
	IntCmd      = ldredis.IntCmd
	FloatCmd    = ldredis.FloatCmd
	SliceCmd    = ldredis.SliceCmd

	StringsCmd         = ldredis.StringsCmd
	MapStringStringCmd = ldredis.MapStringStringCmd
	StringSetCmd       = ldredis.StringSetCmd

	ZMember     = ldredis.ZMember
	ZAddArgs    = ldredis.ZAddArgs
	ZStore      = ldredis.ZStore
	ZRangeBy    = ldredis.ZRangeBy
	ZSliceCmd   = ldredis.ZSliceCmd
	ZWithKey    = ldredis.ZWithKey
	ZWithKeyCmd = ldredis.ZWithKeyCmd

	XAddArgs         = ldredis.XAddArgs
	XReadArgs        = ldredis.XReadArgs
	XReadGroupArgs   = ldredis.XReadGroupArgs
	XClaimArgs       = ldredis.XClaimArgs
	XPending         = ldredis.XPending
	XPendingCmd      = ldredis.XPendingCmd
	XPendingExt      = ldredis.XPendingExt
	XPendingExtArgs  = ldredis.XPendingExtArgs
	XPendingExtCmd   = ldredis.XPendingExtCmd
	XMessage         = ldredis.XMessage
	XMessageSliceCmd = ldredis.XMessageSliceCmd
	XStream          = ldredis.XStream
	XStreamSliceCmd  = ldredis.XStreamSliceCmd

	GeoLocation    = ldredis.GeoLocation
	GeoLocationCmd = ldredis.GeoLocationCmd
	GeoPos         = ldredis.GeoPos
	GeoPosCmd      = ldredis.GeoPosCmd
	GeoRadiusQuery = ldredis.GeoRadiusQuery
)

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
