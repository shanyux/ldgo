/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"go.uber.org/zap"
)

// hkeys
func (that *redisStruct) hkeys(ctx Context, key string, retry int) ([]string, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HKeys(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if err == Nil {
			ctx.LogW("redis hkeys nil", zap.String("key", key))
			return nil, nil
		}

		if i++; i >= retry {
			ctx.LogE("redis hkeys fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) HKeys(ctx Context, key string) ([]string, error) {
	return that.hkeys(ctx, key, 1)
}
func (that *redisStruct) HKeysRetry(ctx Context, key string, retry int) ([]string, error) {
	return that.hkeys(ctx, key, retry)
}

// hgetall
func (that *redisStruct) hgetall(ctx Context, key string, retry int) (map[string]string, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HGetAll(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}
		if err == Nil {
			ctx.LogW("redis hgetall nil", zap.String("key", key))
			return nil, nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis hgetall fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) HGetAll(ctx Context, key string) (map[string]string, error) {
	return that.hgetall(ctx, key, 1)
}

func (that *redisStruct) HGetAllRetry(ctx Context, key string, retry int) (map[string]string, error) {
	return that.hgetall(ctx, key, retry)
}

// hdel
func (that *redisStruct) hdel(ctx Context, key string, fields []string, retry int) (int64, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HDel(key, fields...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis hdel fail", zap.String("key", key), zap.Strings("fields", fields),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}

func (that *redisStruct) HDel(ctx Context, key string, fields ...string) (int64, error) {
	return that.hdel(ctx, key, fields, 1)
}
func (that *redisStruct) HDelRetry(ctx Context, key string, field string, retry int) (bool, error) {
	n, err := that.hdel(ctx, key, []string{field}, retry)
	return n > 0, err
}
func (that *redisStruct) HDelArray(ctx Context, key string, fields []string, retry int) (int64, error) {
	return that.hdel(ctx, key, fields, retry)
}

// hset
func (that *redisStruct) hset(ctx Context, key, field string, value interface{}, retry int) error {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HSet(key, field, value)
		err := cmd.Err()
		if err == nil {
			return nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis hset fail", zap.String("key", key), zap.String("field", field),
				zap.Int("retry", i), zap.Error(err))
			return err
		}
	}
}

func (that *redisStruct) HSet(ctx Context, key, field string, value interface{}) error {
	return that.hset(ctx, key, field, value, 1)
}
func (that *redisStruct) HSetRetry(ctx Context, key, field string, value interface{}, retry int) error {
	return that.hset(ctx, key, field, value, retry)
}

func (that *redisStruct) hmset(ctx Context, key string, m map[string]interface{}, retry int) error {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HMSet(key, m)
		err := cmd.Err()
		if err == nil {
			return nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis hmset fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return err
		}
	}
}

func (that *redisStruct) HMSet(ctx Context, key string, m map[string]interface{}) error {
	return that.hmset(ctx, key, m, 1)
}
func (that *redisStruct) HMSetRetry(ctx Context, key string, m map[string]interface{}, retry int) error {
	return that.hmset(ctx, key, m, retry)
}

func (that *redisStruct) hget(ctx Context, key, field string, retry int) (string, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HGet(key, field)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}
		if err == Nil {
			ctx.LogW("redis hget nil", zap.String("key", key), zap.String("field", field))
			return "", nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis hget fail", zap.String("key", key), zap.String("field", field),
				zap.Int("retry", i), zap.Error(err))
			return "", err
		}
	}
}

func (that *redisStruct) HGet(ctx Context, key, field string) (string, error) {
	return that.hget(ctx, key, field, 1)
}
func (that *redisStruct) HGetRetry(ctx Context, key, field string, retry int) (string, error) {
	return that.hget(ctx, key, field, retry)
}
