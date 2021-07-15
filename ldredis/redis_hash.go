/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"go.uber.org/zap"
)

func (that *redisStruct) hmget(ctx Context, key string, fields []string, retry int) ([]interface{}, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.HMGet(key, fields...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}
		if err == Nil {
			return nil, nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis hmget fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) HMGet(ctx Context, key string, fields ...string) ([]interface{}, error) {
	return that.hmget(ctx, key, fields, 1)
}
func (that *redisStruct) HMGetRetry(ctx Context, key string, field string, retry int) ([]interface{}, error) {
	return that.hmget(ctx, key, []string{field}, retry)
}
func (that *redisStruct) HMGetArray(ctx Context, key string, fields []string, retry int) ([]interface{}, error) {
	return that.hmget(ctx, key, fields, retry)
}
