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

	StringSliceCmd     = redis.StringSliceCmd
	StringStringMapCmd = redis.StringStringMapCmd
	StringStructMapCmd = redis.StringStructMapCmd

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
	Get(key string) *CodecCmd
	GetSet(key string, value interface{}) *CodecCmd
	Set(key string, value interface{}, expiration time.Duration) *StatusCmd
	SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd

	HGet(key, field string) *CodecCmd
	HGetAll(key string) *StringCodecMapCmd
	// HMGet(key string, fields ...string) *SliceCmd
	HMSet(key string, fields map[string]interface{}) *StatusCmd
	HSet(key, field string, value interface{}) *BoolCmd
	HSetNX(key, field string, value interface{}) *BoolCmd
	HVals(key string) *CodecSliceCmd

	// LIndex(key string, index int64) *CodecCmd
	// LPush(key string, values ...interface{}) *IntCmd
	// LPushX(key string, value interface{}) *IntCmd
	// LRange(key string, start, stop int64) *CodecSliceCmd
	// LSet(key string, index int64, value interface{}) *StatusCmd
	//
	// SMembers(key string) *CodecSliceCmd
	// // SMembersMap(key string) *StringStructMapCmd
	// SPop(key string) *CodecCmd
	// SPopN(key string, count int64) *CodecSliceCmd
	// SRandMember(key string) *CodecCmd
	// SRandMemberN(key string, count int64) *CodecSliceCmd

	// XAdd(a *XAddArgs) *StringCmd
}
