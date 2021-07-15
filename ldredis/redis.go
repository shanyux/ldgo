/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/go-redis/redis"
)

type RedisConfig struct {
	Cluster  bool
	Addrs    []string
	Addr     string
	Password string
}

type (
	PubSub = redis.PubSub

	BoolCmd   = redis.BoolCmd
	StatusCmd = redis.StatusCmd
	StringCmd = redis.StringCmd
	IntCmd    = redis.IntCmd
	FloatCmd  = redis.FloatCmd
)

const (
	Nil = redis.Nil
)

type (
	Context = ldcontext.Context
	Redis   = *redisStruct
)

type redisStruct struct {
	client redisCmdable
}

func NewRedis(client redisCmdable) Redis {
	return &redisStruct{
		client: client,
	}
}

func NewRedisByConfig(cfg *RedisConfig) Redis {
	return NewRedis(newClient(cfg))
}

func (that *redisStruct) Client() redisCmdable { return that.client }

type redisCmdable interface {
	redis.Cmdable
	Subscribe(channels ...string) *PubSub
	PSubscribe(channels ...string) *PubSub
}

func newClient(cfg *RedisConfig) redisCmdable {
	if cfg.Cluster {
		cli := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Addrs,
			Password: cfg.Password,
		})
		return redisCluster{ClusterClient: cli}
	}

	cli := redis.NewClient(&redis.Options{
		Addr:        cfg.Addr,
		Password:    cfg.Password,
		IdleTimeout: 3 * time.Second,
	})
	return redisClient{Client: cli}
}

type redisClient struct {
	*redis.Client
}

type redisCluster struct {
	*redis.ClusterClient
}
