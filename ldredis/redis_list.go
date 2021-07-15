/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"go.uber.org/zap"
)

func (that *redisStruct) llen(ctx Context, key string, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.LLen(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis llen fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) LLen(ctx Context, key string) (int64, error) {
	return that.llen(ctx, key, 1)
}
func (that *redisStruct) LLenRetry(ctx Context, key string, retry int) (int64, error) {
	return that.llen(ctx, key, retry)
}

func (that *redisStruct) lrange(ctx Context, key string, start, stop int64, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.LRange(key, start, stop)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis lrange fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) LRange(ctx Context, key string, start, stop int64) ([]string, error) {
	return that.lrange(ctx, key, start, stop, 1)
}
func (that *redisStruct) LRangeRetry(ctx Context, key string, start, stop int64, retry int) ([]string, error) {
	return that.lrange(ctx, key, start, stop, retry)
}

func (that *redisStruct) lpop(ctx Context, key string, retry int) (string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.LPop(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis lpop fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return "", err
		}
	}
}
func (that *redisStruct) LPop(ctx Context, key string) (string, error) {
	return that.lpop(ctx, key, 1)
}
func (that *redisStruct) LPopRetry(ctx Context, key string, retry int) (string, error) {
	return that.lpop(ctx, key, retry)
}

func (that *redisStruct) rpop(ctx Context, key string, retry int) (string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.RPop(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis rpop fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return "", err
		}
	}
}
func (that *redisStruct) RPop(ctx Context, key string) (string, error) {
	return that.rpop(ctx, key, 1)
}
func (that *redisStruct) RPopRetry(ctx Context, key string, retry int) (string, error) {
	return that.rpop(ctx, key, retry)
}

func (that *redisStruct) lpush(ctx Context, key string, vals []interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.LPush(key, vals...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis lpush fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) LPush(ctx Context, key string, vals ...interface{}) (int64, error) {
	return that.lpush(ctx, key, vals, 1)
}
func (that *redisStruct) LPushArray(ctx Context, key string, vals []interface{}) (int64, error) {
	return that.lpush(ctx, key, vals, 1)
}
func (that *redisStruct) LPushRetry(ctx Context, key string, vals []interface{}, retry int) (int64, error) {
	return that.lpush(ctx, key, vals, retry)
}

func (that *redisStruct) lpushx(ctx Context, key string, val interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.LPushX(key, val)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis lpushx fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) LPushX(ctx Context, key string, val interface{}) (int64, error) {
	return that.lpushx(ctx, key, val, 1)
}
func (that *redisStruct) LPushXRetry(ctx Context, key string, val interface{}, retry int) (int64, error) {
	return that.lpushx(ctx, key, val, retry)
}

func (that *redisStruct) rpush(ctx Context, key string, vals []interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.RPush(key, vals...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis rpush fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) RPush(ctx Context, key string, vals ...interface{}) (int64, error) {
	return that.rpush(ctx, key, vals, 1)
}
func (that *redisStruct) RPushArray(ctx Context, key string, vals []interface{}) (int64, error) {
	return that.rpush(ctx, key, vals, 1)
}
func (that *redisStruct) RPushRetry(ctx Context, key string, vals []interface{}, retry int) (int64, error) {
	return that.rpush(ctx, key, vals, retry)
}

func (that *redisStruct) rpushx(ctx Context, key string, val interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.RPushX(key, val)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis rpushx fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) RPushX(ctx Context, key string, val interface{}) (int64, error) {
	return that.rpushx(ctx, key, val, 1)
}
func (that *redisStruct) RPushXRetry(ctx Context, key string, val interface{}, retry int) (int64, error) {
	return that.rpushx(ctx, key, val, retry)
}
