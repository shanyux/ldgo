/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/ldlog"
	"github.com/go-redis/redis"
)

const (
	Nil = redis.Nil
)

func isErrNil(err error) bool {
	return err == nil || err == Nil
}

var _ Cmdable = (*Redis)(nil)

const (
	redisRetryIntervalMin = 10 * time.Millisecond
	redisRetryIntervalDef = 1 * time.Second
)

func New(cli redis.Cmdable) *Redis {
	if rds, ok := cli.(*Redis); ok {
		return rds
	}

	switch v := cli.(type) {
	case *Redis:
		return v

	case cmdable:
		return newRedis(v)

	case *redis.Client:
		return newRedis(redisClientWrapper{Client: v})

	case *redis.ClusterClient:
		return newRedis(redisClusterWrapper{ClusterClient: v})
	}

	panic("redis client type must be `*ldredis.Redis` or `*redis.Client` or `*redis.ClusterClient`")
}

func NewByConfig(cfg *Config) *Redis {
	if cfg.Cluster {
		return newRedis(newRedisCluster(cfg))
	}

	return newRedis(newRedisClient(cfg))
}

func newRedis(cli cmdable) *Redis {
	ctx := ldctx.NewContext(cli.Context())
	c := &Redis{
		cmdable:  cli,
		origin:   cli,
		ctx:      ctx,
		reporter: discardReporter{},
		caller:   true,
	}

	c.cmdable = c.cmdable.withContext(ctx)
	c.cmdable.WrapProcess(func(oldProcess func(cmd Cmder) error) func(cmd Cmder) error {
		c.oldProcess = oldProcess
		return c.defaultProcess
	})
	c.cmdable.WrapProcessPipeline(func(oldProcess func(cmds []Cmder) error) func(cmds []Cmder) error {
		c.oldProcessPipeline = oldProcess
		return c.defaultProcessPipeline
	})

	return c
}

// Redis struct
type Redis struct {
	cmdable

	origin             cmdable
	oldProcess         func(cmd Cmder) error
	oldProcessPipeline func(cmds []Cmder) error

	ctx           Context
	log           *ldlog.Logger
	reporter      Reporter
	retry         int
	retryInterval time.Duration
	caller        bool
}

func (c *Redis) Client() Cmdable  { return c.origin }
func (c *Redis) context() Context { return c.ctx }

func (c *Redis) logger() *ldlog.Logger {
	if c.log != nil {
		return c.log
	}
	return ldctx.GetLogger(c.context())
}

func (c *Redis) clone(ctx ...Context) *Redis {
	cp := *c
	c = &cp

	if len(ctx) != 0 {
		c.origin = c.origin.withContext(ctx[0])
	}

	c.cmdable = c.origin.withContext(c.origin.Context())
	c.cmdable.WrapProcess(func(oldProcess func(cmd Cmder) error) func(cmd Cmder) error {
		return c.defaultProcess
	})
	c.cmdable.WrapProcessPipeline(func(oldProcess func(cmds []Cmder) error) func(cmds []Cmder) error {
		return c.defaultProcessPipeline
	})

	return c
}

func (c *Redis) WithCodec(codec Codec) *CodecRedis {
	return &CodecRedis{
		client: c,
		codec:  codec,
	}
}

func (c *Redis) WithContext(ctx Context) *Redis {
	if ctx == nil {
		ctx = ldctx.Default()
	}

	c = c.clone(ctx)
	c.ctx = ctx
	c.log = nil
	return c
}

// func (c *Redis) WithLogger(l *ldlog.Logger) *Redis {
// 	c = c.clone()
// 	c.log = l
// 	return c
// }

// WithRetry should be called like these:
//
//	WithRetry(3)
//	WithRetry(3, {interval})
func (c *Redis) WithRetry(retry int, interval ...time.Duration) *Redis {
	c = c.clone()
	c.retry = retry

	c.retryInterval = redisRetryIntervalDef
	if len(interval) > 0 {
		d := interval[0]
		if d < redisRetryIntervalMin {
			d = redisRetryIntervalMin
		}
		c.retryInterval = d
	}

	return c
}

func (c *Redis) WithReport(reporter Reporter) *Redis {
	if reporter == nil {
		reporter = discardReporter{}
	}

	c = c.clone()
	c.reporter = reporter
	return c
}

func (c *Redis) WithCaller(enable bool) *Redis {
	c = c.clone()
	c.caller = enable
	return c
}
