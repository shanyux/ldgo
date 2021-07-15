/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"go.uber.org/zap"
)

func (that *redisStruct) sadd(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SAdd(key, members...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis sadd fail", zap.String("key", key), zap.Reflect("members", members),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) SAdd(ctx Context, key string, members ...interface{}) (int64, error) {
	return that.sadd(ctx, key, members, 1)
}
func (that *redisStruct) SAddArray(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	return that.sadd(ctx, key, members, retry)
}
func (that *redisStruct) SAddRetry(ctx Context, key string, member interface{}, retry int) (int64, error) {
	return that.sadd(ctx, key, []interface{}{member}, retry)
}

func (that *redisStruct) smembers(ctx Context, key string, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SMembers(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis smembers fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SMembers(ctx Context, key string) ([]string, error) {
	return that.smembers(ctx, key, 1)
}
func (that *redisStruct) SMembersRetry(ctx Context, key string, retry int) ([]string, error) {
	return that.smembers(ctx, key, retry)
}

func (that *redisStruct) smembersmap(ctx Context, key string, retry int) (map[string]struct{}, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SMembersMap(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis smembers(map) fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SMembersMap(ctx Context, key string) (map[string]struct{}, error) {
	return that.smembersmap(ctx, key, 1)
}
func (that *redisStruct) SMembersMapRetry(ctx Context, key string, retry int) (map[string]struct{}, error) {
	return that.smembersmap(ctx, key, retry)
}

func (that *redisStruct) srandmembersn(ctx Context, key string, count int64, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SRandMemberN(key, count)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis srandmembers fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SRandMembers(ctx Context, key string) ([]string, error) {
	return that.srandmembersn(ctx, key, 1, 1)
}
func (that *redisStruct) SRandMembersRetry(ctx Context, key string, retry int) ([]string, error) {
	return that.srandmembersn(ctx, key, 1, retry)
}
func (that *redisStruct) SRandMembersN(ctx Context, key string, count int64) ([]string, error) {
	return that.srandmembersn(ctx, key, count, 1)
}
func (that *redisStruct) SRandMembersNRetry(ctx Context, key string, count int64, retry int) ([]string, error) {
	return that.srandmembersn(ctx, key, count, retry)
}

func (that *redisStruct) sismember(ctx Context, key string, member interface{}, retry int) (bool, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SIsMember(key, member)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis sismember fail", zap.String("key", key), zap.Reflect("member", member),
				zap.Int("retry", i), zap.Error(err))
			return false, err
		}
	}
}
func (that *redisStruct) SIsMember(ctx Context, key string, member interface{}) (bool, error) {
	return that.sismember(ctx, key, member, 1)
}
func (that *redisStruct) SIsMemberRetry(ctx Context, key string, member interface{}, retry int) (bool, error) {
	return that.sismember(ctx, key, member, retry)
}

func (that *redisStruct) sdiff(ctx Context, keys []string, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SDiff(keys...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis sdiff fail", zap.Strings("keys", keys), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SDiff(ctx Context, keys ...string) ([]string, error) {
	return that.sdiff(ctx, keys, 1)
}
func (that *redisStruct) SDiffArray(ctx Context, keys []string, retry int) ([]string, error) {
	return that.sdiff(ctx, keys, retry)
}
func (that *redisStruct) SDiffRetry(ctx Context, key0, key1 string, retry int) ([]string, error) {
	return that.sdiff(ctx, []string{key0, key1}, retry)
}

func (that *redisStruct) sunion(ctx Context, keys []string, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SUnion(keys...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis sunion fail", zap.Strings("keys", keys), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SUnion(ctx Context, keys ...string) ([]string, error) {
	return that.sunion(ctx, keys, 1)
}
func (that *redisStruct) SUnionRetry(ctx Context, key0, key1 string, retry int) ([]string, error) {
	return that.sunion(ctx, []string{key0, key1}, retry)
}
func (that *redisStruct) SUnionArray(ctx Context, keys []string, retry int) ([]string, error) {
	return that.sunion(ctx, keys, retry)
}

func (that *redisStruct) sinter(ctx Context, keys []string, retry int) ([]string, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SInter(keys...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis sinter fail", zap.Strings("keys", keys), zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) SInter(ctx Context, keys ...string) ([]string, error) {
	return that.sinter(ctx, keys, 1)
}
func (that *redisStruct) SInterRetry(ctx Context, key0, key1 string, retry int) ([]string, error) {
	return that.sinter(ctx, []string{key0, key1}, retry)
}
func (that *redisStruct) SInterArray(ctx Context, keys []string, retry int) ([]string, error) {
	return that.sinter(ctx, keys, retry)
}

func (that *redisStruct) srem(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SRem(key, members...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis srem fail", zap.String("key", key), zap.Reflect("members", members),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) SRem(ctx Context, key string, members ...interface{}) (int64, error) {
	return that.srem(ctx, key, members, 1)
}
func (that *redisStruct) SRemRetry(ctx Context, key string, member interface{}) (int64, error) {
	return that.srem(ctx, key, []interface{}{member}, 1)
}
func (that *redisStruct) SRemArray(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	return that.srem(ctx, key, members, retry)
}

func (that *redisStruct) scard(ctx Context, key string, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.SCard(key)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis scard fail", zap.String("key", key), zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) SCard(ctx Context, key string) (int64, error) {
	return that.scard(ctx, key, 1)
}
func (that *redisStruct) SCardRetry(ctx Context, key string, retry int) (int64, error) {
	return that.scard(ctx, key, retry)
}
