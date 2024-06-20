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
		wrapper: newWrapper(client),
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
	wrapper

	closer closer
}

func (r testRedisWrapper) Close() error {
	err := r.wrapper.Close()
	r.closer.Close()
	return err
}

func (r testRedisWrapper) Clone() cmdable {
	r.wrapper = r.wrapper.Clone()
	return r
}
