/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"

	"github.com/distroy/ldgo/v2/ldredis/internal"
	redis "github.com/redis/go-redis/v9"
)

type (
	Reporter = internal.Reporter
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
