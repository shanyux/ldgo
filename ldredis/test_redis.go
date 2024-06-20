/*
 * Copyright (C) Renchong Tan
 */

package ldredis

import (
	"fmt"

	miniredis "github.com/alicebob/miniredis/v2"
	redis "github.com/redis/go-redis/v9"
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
		cmdable: newWrapper(client),
		closer:  server,
	}), nil
}

func MustNewTestRedis() *Redis {
	r, err := NewTestRedis()
	if err != nil {
		panic(fmt.Errorf("new test redis fail. err:%w", err))
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

func (r testRedisWrapper) Close() error {
	err := r.cmdable.Close()
	r.closer.Close()
	return err
}
