/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"go.uber.org/zap"
)

// exists
func (that *redisStruct) exists(ctx Context, keys []string, retry int) (int, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.Exists(keys...)
		err := cmd.Err()
		if err == nil {
			return int(cmd.Val()), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis exists fail", zap.Strings("keys", keys), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) Exists(ctx Context, keys ...string) (int, error) {
	return that.exists(ctx, keys, 1)
}
func (that *redisStruct) ExistsRetry(ctx Context, key string, retry int) (bool, error) {
	v, err := that.exists(ctx, []string{key}, retry)
	return v != 0, err
}
func (that *redisStruct) ExistsArray(ctx Context, keys []string, retry int) (int, error) {
	return that.exists(ctx, keys, retry)
}

// del
func (that *redisStruct) del(ctx Context, keys []string, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.Del(keys...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis del fail", zap.Strings("keys", keys), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) Del(ctx Context, keys ...string) (int64, error) {
	return that.del(ctx, keys, 1)
}
func (that *redisStruct) DelRetry(ctx Context, key string, retry int) (bool, error) {
	n, err := that.del(ctx, []string{key}, retry)
	return n > 0, err
}
func (that *redisStruct) DelArray(ctx Context, keys []string, retry int) (int64, error) {
	return that.del(ctx, keys, 1)
}
