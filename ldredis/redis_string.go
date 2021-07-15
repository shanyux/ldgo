/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"go.uber.org/zap"
)

// get
func (that *redisStruct) get(ctx Context, key string, retry int) (string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.Get(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if err == Nil {
			ctx.LogW("redis get nil", zap.String("key", key))
			return "", nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis get fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return "", err
		}
	}
}
func (that *redisStruct) Get(ctx Context, key string) (string, error) {
	return that.get(ctx, key, 1)
}
func (that *redisStruct) GetRetry(ctx Context, key string, retry int) (string, error) {
	return that.get(ctx, key, retry)
}

// set
// Zero expiration means the key has no expiration time.
func (that *redisStruct) set(ctx Context, key string, val interface{}, expiration time.Duration, retry int) error {
	cli := that.client
	for i := 0; ; {
		cmd := cli.Set(key, val, expiration)
		err := cmd.Err()
		if err == nil {
			return nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis set fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return err
		}
	}
}

func (that *redisStruct) Set(ctx Context, key string, val interface{}, expiration time.Duration) error {
	return that.set(ctx, key, val, expiration, 1)
}
func (that *redisStruct) SetRetry(ctx Context, key string, val interface{}, expiration time.Duration, retry int) error {
	return that.set(ctx, key, val, expiration, retry)
}

// SetXx `SET key value [expiration] XX` command.
// Zero expiration means the key has no expiration time.
func (that *redisStruct) setxx(ctx Context, key string, val interface{}, expiration time.Duration, retry int) (bool, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.SetXX(key, val, expiration)
		err := cmd.Err()
		if err == nil {
			if !cmd.Val() {
				ctx.LogW("redis setxx existed", zap.String("key", key))
			}
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis setxx fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return false, err
		}
	}
}
func (that *redisStruct) SetXx(ctx Context, key string, val interface{}, expiration time.Duration) (bool, error) {
	return that.setxx(ctx, key, val, expiration, 1)
}
func (that *redisStruct) SetXxRetry(ctx Context, key string, val interface{}, expiration time.Duration, retry int) (bool, error) {
	return that.setxx(ctx, key, val, expiration, retry)
}

// SetNx `SET key value [expiration] NX` command.
// Zero expiration means the key has no expiration time.
func (that *redisStruct) setnx(ctx Context, key string, val interface{}, expiration time.Duration, retry int) (bool, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.SetNX(key, val, expiration)
		err := cmd.Err()
		if err == nil {
			if !cmd.Val() {
				ctx.LogW("redis setnx existed", zap.String("key", key))
			}
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis setnx fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return false, err
		}
	}
}
func (that *redisStruct) SetNx(ctx Context, key string, val interface{}, expiration time.Duration) (bool, error) {
	return that.setnx(ctx, key, val, expiration, 1)
}
func (that *redisStruct) SetNxRetry(ctx Context, key string, val interface{}, expiration time.Duration, retry int) (bool, error) {
	return that.setnx(ctx, key, val, expiration, retry)
}
