/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	redis "github.com/redis/go-redis/v9"
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
	switch v := cli.(type) {
	case *Redis:
		return v

	case cmdable:
		return newRedis(v)

	case *redis.Client:
		return newRedis(v)

	case *redis.ClusterClient:
		return newRedis(v)
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
	c := &Redis{
		cmdable:  cli,
		origin:   cli,
		reporter: discardReporter{},
		caller:   true,
	}

	c.cmdable.AddHook(newHook(c))
	return c
}

// Redis struct
type Redis struct {
	cmdable

	origin             cmdable
	oldProcess         func(cmd Cmder) error
	oldProcessPipeline func(cmds []Cmder) error

	reporter      Reporter
	retry         int
	retryInterval time.Duration
	caller        bool
}

func (c *Redis) Client() Cmdable { return c.origin }

func (c *Redis) clone() *Redis {
	cp := *c
	c = &cp

	return c
}

func (c *Redis) WithCodec(codec Codec) *CodecRedis {
	return &CodecRedis{
		client: c,
		codec:  codec,
	}
}

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
