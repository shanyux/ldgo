/*
 * Copyright (C) distroy
 */

package ldredis

import (
	redis "github.com/redis/go-redis/v9"
)

type cmdable interface {
	Cmdable

	AddHook(hook redis.Hook)
}

func newRedisClient(cfg *Config) cmdable {
	return redis.NewClient(cfg.toClient())
}

func newRedisCluster(cfg *Config) cmdable {
	return redis.NewClusterClient(cfg.toCluster())
}
