/*
 * Copyright (C) distroy
 */

package ldredis

import "go.uber.org/zap"

func (that *redisStruct) Incr(ctx Context, key string) (int64, error) {
	return that.incrby(ctx, key, 1, 1)
}

// incrby
func (that *redisStruct) incrby(ctx Context, key string, n int64, retry int) (int64, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.IncrBy(key, n)
		err := cmd.Err()
		if err == nil {
			v := cmd.Val()
			return v, nil
		}

		if i++; i >= retry {
			ctx.LogE("redis incrby fail", zap.String("key", key), zap.Int64("n", n),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) IncrBy(ctx Context, key string, n int64) (int64, error) {
	return that.incrby(ctx, key, n, 1)
}
func (that *redisStruct) IncrByRetry(ctx Context, key string, n int64, retry int) (int64, error) {
	return that.incrby(ctx, key, n, retry)
}

func (that *redisStruct) incrBy2(ctx Context, key string, n int64, retry int) (int64, int64, error) {
	v, err := that.incrby(ctx, key, n, retry)
	if err != nil {
		return 0, 0, err
	}
	return (v - n + 1), v, nil
}
func (that *redisStruct) IncrBy2(ctx Context, key string, n int64) (int64, int64, error) {
	return that.incrBy2(ctx, key, n, 1)

}
func (that *redisStruct) IncrBy2Retry(ctx Context, key string, n int64, retry int) (int64, int64, error) {
	return that.incrBy2(ctx, key, n, retry)
}
