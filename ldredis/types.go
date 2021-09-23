/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"github.com/distroy/ldgo/ldcontext"
	"github.com/go-redis/redis"
)

type (
	Context = ldcontext.Context
)

type (
	Cmd   = redis.Cmd
	Cmder = redis.Cmder

	PubSub = redis.PubSub

	BoolCmd   = redis.BoolCmd
	StatusCmd = redis.StatusCmd

	SliceCmd  = redis.SliceCmd
	StringCmd = redis.StringCmd
	IntCmd    = redis.IntCmd
	FloatCmd  = redis.FloatCmd

	StringSliceCmd     = redis.StringSliceCmd
	StringStringMapCmd = redis.StringStringMapCmd
	StringStructMapCmd = redis.StringStructMapCmd
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
