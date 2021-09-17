/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"

	"github.com/go-redis/redis"
)

type cmdable interface {
	Cmdable

	withContext(ctx context.Context) cmdable
	Context() context.Context

	WrapProcess(fn func(oldProcess func(Cmder) error) func(Cmder) error)

	Do(args ...interface{}) *Cmd
	Process(cmd Cmder) error
	Close() error
	// Discard() error
	// Exec() ([]Cmder, error)

	Subscribe(channels ...string) *PubSub
	PSubscribe(channels ...string) *PubSub
}

type redisClientWrapper struct {
	*redis.Client
}

func (r redisClientWrapper) withContext(ctx context.Context) cmdable {
	r.Client = r.Client.WithContext(ctx)
	return r
}

type redisClusterWrapper struct {
	*redis.ClusterClient
}

func (r redisClusterWrapper) withContext(ctx context.Context) cmdable {
	r.ClusterClient = r.ClusterClient.WithContext(ctx)
	return r
}

func newRedisClient(cfg *Config) redisClientWrapper {
	return redisClientWrapper{
		Client: redis.NewClient(cfg.toClient()),
	}
}

func newRedisCluster(cfg *Config) redisClusterWrapper {
	return redisClusterWrapper{
		ClusterClient: redis.NewClusterClient(cfg.toCluster()),
	}
}
