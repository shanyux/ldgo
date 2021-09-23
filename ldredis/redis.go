/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldlogger"
	"github.com/go-redis/redis"
)

const (
	Nil = redis.Nil
)

var _ Cmdable = &Redis{}

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
	ctx := ldcontext.NewContext(cli.Context())
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

	ctx      Context
	log      ldlogger.Logger
	reporter Reporter
	retry    int
	caller   bool
}

func (c *Redis) Client() Cmdable  { return c.origin }
func (c *Redis) context() Context { return c.ctx }

func (c *Redis) logger() ldlogger.Logger {
	if c.log != nil {
		return c.log
	}
	return ldcontext.GetLogger(c.context())
}

func (c *Redis) clone(ctx ...Context) *Redis {
	cp := *c
	c = &cp

	if len(ctx) != 0 {
		c.origin = c.origin.withContext(ctx[0])
	} else {
		c.origin = c.origin.withContext(c.origin.Context())
	}

	c.cmdable = c.origin
	c.cmdable.WrapProcess(func(oldProcess func(cmd Cmder) error) func(cmd Cmder) error {
		return c.defaultProcess
	})
	c.cmdable.WrapProcessPipeline(func(oldProcess func(cmds []Cmder) error) func(cmds []Cmder) error {
		return c.defaultProcessPipeline
	})

	return c
}

func (c *Redis) WithContext(ctx Context) *Redis {
	c = c.clone(ctx)
	c.ctx = ctx
	c.log = nil
	return c
}

// func (c *Redis) WithLogger(l ldlogger.Logger) *Redis {
// 	c = c.clone()
// 	c.log = l
// 	return c
// }

func (c *Redis) WithRetry(retry int) *Redis {
	c = c.clone()
	c.retry = retry
	return c
}

func (c *Redis) WithReport(reporter Reporter) *Redis {
	c = c.clone()
	c.reporter = reporter
	return c
}

func (c *Redis) WithCaller(caller bool) *Redis {
	c = c.clone()
	c.caller = caller
	return c
}
