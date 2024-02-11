/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldredis/internal"
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
		cmdable: cli,
		opts: internal.Options{
			Reporter: internal.DiscardReporter{},
			Caller:   true,
		},
	}

	c.cmdable.AddHook(newHook(c))
	return c
}

// Redis struct
type Redis struct {
	cmdable

	opts internal.Options
}

func (c *Redis) Client() Cmdable { return c.cmdable }

func (c *Redis) clone() *Redis {
	cp := *c
	c = &cp

	return c
}

// WithRetry should be called like these:
//
//	WithRetry(3)
//	WithRetry(3, {interval})
func (c *Redis) WithRetry(retry int, interval ...time.Duration) *Redis {
	c = c.clone()
	c.opts.Retry = retry

	c.opts.RetryInterval = redisRetryIntervalDef
	if len(interval) > 0 {
		d := interval[0]
		if d < redisRetryIntervalMin {
			d = redisRetryIntervalMin
		}
		c.opts.RetryInterval = d
	}

	return c
}

func (c *Redis) WithReport(reporter Reporter) *Redis {
	if reporter == nil {
		reporter = internal.DiscardReporter{}
	}

	c = c.clone()
	c.opts.Reporter = reporter
	return c
}

func (c *Redis) WithCaller(enable bool) *Redis {
	c = c.clone()
	c.opts.Caller = enable
	return c
}

func (c *Redis) Process(ctx context.Context, cmd Cmder) error {
	if err := c.cmdable.Process(ctx, cmd); err != nil {
		return err
	}

	v, _ := cmd.(internal.CmderWithParse)
	if v != nil {
		return v.Parse(ctx)
	}

	return nil
}
