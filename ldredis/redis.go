/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldlogger"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var _ Cmdable = &Redis{}

func NewRedis(cli redis.Cmdable) *Redis {
	if rds, ok := cli.(*Redis); ok {
		return rds
	}

	switch v := cli.(type) {
	case *Redis:
		return v

	case *redis.Client:
		return newRedis(redisClientWrapper{Client: v})

	case *redis.ClusterClient:
		return newRedis(redisClusterWrapper{ClusterClient: v})
	}

	panic("redis client type is invalid")
}

func NewRedisByConfig(cfg *Config) *Redis {
	if cfg.Cluster {
		return newRedis(newRedisCluster(cfg))
	}

	return newRedis(newRedisClient(cfg))
}

func newRedis(cli cmdable) *Redis {
	cli = cli.withContext(cli.Context())

	rds := &Redis{
		cmdable:  cli,
		logger:   ldlogger.Discard(),
		reporter: discardReporter{},
	}

	rds.cmdable.WrapProcess(func(oldProcess func(Cmder) error) func(Cmder) error {
		rds.oldProcess = oldProcess
		return rds.process
	})

	return rds
}

// Redis struct
type Redis struct {
	cmdable

	logger     ldlogger.Logger
	reporter   Reporter
	oldProcess func(Cmder) error
	retry      int
}

func (c *Redis) clone() *Redis {
	cp := *c
	return &cp
}

func (c *Redis) WithContext(ctx Context) *Redis {
	c = c.clone()
	c.cmdable = c.cmdable.withContext(ctx)

	log := ldcontext.GetLogger(ctx)
	log = ldlogger.WithOptions(log, zap.AddCallerSkip(1))
	c.logger = log

	c.cmdable.WrapProcess(func(oldProcess func(Cmder) error) func(Cmder) error {
		return c.process
	})

	return c
}

func (c *Redis) WithLogger(log ldlogger.Logger) *Redis {
	if log != ldlogger.Discard() {
		log = ldlogger.WithOptions(log, zap.AddCallerSkip(1))
	}

	c = c.clone()
	c.logger = log
	return c
}

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

func (c *Redis) Client() Cmdable {
	return c.cmdable
}
