/*
 * Copyright (C) Renchong Tan
 */

package ldredis

import (
	"context"
	"fmt"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

func NewTestRedis() (*Redis, error) {
	server, err := miniredis.Run()
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})

	return New(&testRedisWrapper{
		cmdable: redisClientWrapper{Client: client},
		closer:  server,
	}), nil
}

func MustNewTestRedis() *Redis {
	r, err := NewTestRedis()
	if err != nil {
		panic(fmt.Sprintf("new test redis fail. err:%v", err))
	}
	return r
}

type closer interface {
	Close()
}

type testRedisWrapper struct {
	cmdable

	closer closer
}

func (r testRedisWrapper) withContext(ctx context.Context) cmdable {
	r.cmdable = r.cmdable.withContext(ctx)
	return r
}

func (r testRedisWrapper) Close() error {
	err := r.cmdable.Close()
	r.closer.Close()
	return err
}
