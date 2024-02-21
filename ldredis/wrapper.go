/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"unsafe"

	"github.com/distroy/ldgo/v2/ldredis/internal"
	redis "github.com/redis/go-redis/v9"
)

func getOptions(c *Redis) *internal.Options     { return &c.opts }
func getOptionsPointer(c *Redis) unsafe.Pointer { return unsafe.Pointer(&c.opts) }

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
