/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"go.uber.org/zap"
)

// expire
func (that *redisStruct) expire(ctx Context, key string, expiration time.Duration, retry int) (bool, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.Expire(key, expiration)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis expire fail", zap.String("key", key), zap.Duration("expiration", expiration),
				zap.Int("retry", i), zap.Error(err))
			return false, err
		}
	}
}
func (that *redisStruct) Expire(ctx Context, key string, expiration time.Duration) (bool, error) {
	return that.expire(ctx, key, expiration, 1)
}

func (that *redisStruct) ExpireRetry(ctx Context, key string, expiration time.Duration, retry int) (bool, error) {
	return that.expire(ctx, key, expiration, retry)
}

// expireat
func (that *redisStruct) expireAt(ctx Context, key string, tm time.Time, retry int) (bool, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.ExpireAt(key, tm)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis expireat fail", zap.String("key", key), zap.Time("time", tm),
				zap.Int("retry", i), zap.Error(err))
			return false, err
		}
	}
}
func (that *redisStruct) ExpireAt(ctx Context, key string, tm time.Time) (bool, error) {
	return that.expireAt(ctx, key, tm, 1)
}

// ttl
func (that *redisStruct) ttl(ctx Context, key string, retry int) (time.Duration, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.TTL(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis ttl fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) Ttl(ctx Context, key string) (time.Duration, error) {
	return that.ttl(ctx, key, 1)
}
func (that *redisStruct) TtlRetry(ctx Context, key string, retry int) (time.Duration, error) {
	return that.ttl(ctx, key, retry)
}
