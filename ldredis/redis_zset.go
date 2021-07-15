/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"strconv"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	ZSET_SCORE_MIN zScore = "-inf"
	ZSET_SCORE_MAX zScore = "+inf"
)

type ZMember = redis.Z
type zScore string

type ZRangeBy struct {
	Min    zScore
	Max    zScore
	Offset int64
	Count  int64
}

func ZScoreExclusive(score float64) zScore {
	return zScore("(" + strconv.FormatFloat(score, 'f', -1, 64))
}
func ZScoreInclusive(score float64) zScore {
	return zScore(strconv.FormatFloat(score, 'f', -1, 64))
}

func (that *redisStruct) zadd(ctx Context, key string, members []ZMember, retry int) (int64, error) {
	cli := that.client
	for i := 0; ; {
		cmd := cli.ZAdd(key, members...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		i++
		if i >= retry {
			ctx.LogE("redis zadd failed", zap.String("key", key), zap.Reflect("members", members),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}
func (that *redisStruct) zaddWithScore(ctx Context, key string, score float64, members []interface{}, retry int) (int64, error) {
	zMembers := make([]ZMember, 0, len(members))
	for _, member := range members {
		zMembers = append(zMembers, ZMember{
			Score:  score,
			Member: member,
		})
	}

	return that.zadd(ctx, key, zMembers, retry)
}
func (that *redisStruct) ZAdd(ctx Context, key string, score float64, members ...interface{}) (int64, error) {
	return that.zaddWithScore(ctx, key, score, members, 1)
}
func (that *redisStruct) ZAddArray(ctx Context, key string, score float64, members []interface{}, retry int) (int64, error) {
	return that.zaddWithScore(ctx, key, score, members, retry)
}
func (that *redisStruct) ZAddRetry(ctx Context, key string, score float64, member interface{}, retry int) (int64, error) {
	zMembers := []ZMember{
		ZMember{
			Score:  score,
			Member: member,
		},
	}
	return that.zadd(ctx, key, zMembers, retry)
}

func (that *redisStruct) ZAddMember(ctx Context, key string, members ...ZMember) (int64, error) {
	return that.zadd(ctx, key, members, 1)
}

func (that *redisStruct) ZAddMemberArray(ctx Context, key string, members []ZMember, retry int) (int64, error) {
	return that.zadd(ctx, key, members, retry)
}

func (that *redisStruct) zrem(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.ZRem(key, members...)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zrem failed", zap.String("key", key), zap.Reflect("members", members),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}

func (that *redisStruct) ZRem(ctx Context, key string, members ...interface{}) (int64, error) {
	return that.zrem(ctx, key, members, 1)
}

func (that *redisStruct) ZRemRetry(ctx Context, key string, member interface{}, retry int) (int64, error) {
	return that.zrem(ctx, key, []interface{}{member}, retry)
}

func (that *redisStruct) ZRemArray(ctx Context, key string, members []interface{}, retry int) (int64, error) {
	return that.zrem(ctx, key, members, retry)
}

// remove members from the min score to max score, min and max score included
func (that *redisStruct) zremrangebyscore(ctx Context, key string, minScore, maxScore zScore, retry int) (int64, error) {
	cli := that.client

	min := string(minScore)
	max := string(maxScore)
	for i := 0; ; {
		cmd := cli.ZRemRangeByScore(key, min, max)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zrem range by score failed", zap.String("key", key),
				zap.String("min", min), zap.String("max", max), zap.Int("retry", i),
				zap.Error(err))
			return 0, err
		}
	}
}

func (that *redisStruct) ZRemScoreRange(ctx Context, key string, min, max zScore) (int64, error) {
	return that.zremrangebyscore(ctx, key, min, max, 1)
}

func (that *redisStruct) ZRemScoreRangeRetry(ctx Context, key string, min, max zScore, retry int) (int64, error) {
	return that.zremrangebyscore(ctx, key, min, max, retry)
}

func (that *redisStruct) zremrangebylex(ctx Context, key string, minScore, maxScore zScore, retry int) (int64, error) {
	cli := that.client

	min := string(minScore)
	max := string(maxScore)
	for i := 0; ; {
		cmd := cli.ZRemRangeByLex(key, min, max)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zrem range bt score failed", zap.String("key", key), zap.String("min", min), zap.String("max", max),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}
}

func (that *redisStruct) zrangebyscore(ctx Context, key string, opt ZRangeBy, retry int) ([]string, error) {
	cli := that.client
	rangeBy := redis.ZRangeBy{
		Min:    string(opt.Min),
		Max:    string(opt.Max),
		Offset: opt.Offset,
		Count:  opt.Count,
	}

	for i := 0; ; {
		cmd := cli.ZRangeByScore(key, rangeBy)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zrange by score failed", zap.String("key", key), zap.Reflect("zRangeBy", opt),
				zap.Int("retry", i), zap.Error(err))
			return nil, err
		}
	}
}
func (that *redisStruct) ZRangeByScore(ctx Context, key string, opt ZRangeBy) ([]string, error) {
	return that.zrangebyscore(ctx, key, opt, 1)
}

func (that *redisStruct) ZRangeByScoreRetry(ctx Context, key string, opt ZRangeBy, retry int) ([]string, error) {
	return that.zrangebyscore(ctx, key, opt, retry)
}

func (that *redisStruct) zcount(ctx Context, key string, minScore, maxScore zScore, retry int) (int64, error) {
	cli := that.client

	min := string(minScore)
	max := string(maxScore)
	for i := 0; ; {
		cmd := cli.ZCount(key, min, max)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zcount fail", zap.String("key", key),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}

}

func (that *redisStruct) ZCount(ctx Context, key string) (int64, error) {
	return that.zcount(ctx, key, ZSET_SCORE_MIN, ZSET_SCORE_MAX, 1)
}
func (that *redisStruct) ZCountRetry(ctx Context, key string, retry int) (int64, error) {
	return that.zcount(ctx, key, ZSET_SCORE_MIN, ZSET_SCORE_MAX, retry)
}
func (that *redisStruct) ZCountByScore(ctx Context, key string, min, max zScore) (int64, error) {
	return that.zcount(ctx, key, min, max, 1)
}
func (that *redisStruct) ZCountByScoreRetry(ctx Context, key string, min, max zScore, retry int) (int64, error) {
	return that.zcount(ctx, key, min, max, retry)
}

func (that *redisStruct) zrank(ctx Context, key string, member string, retry int) (int64, error) {
	cli := that.client

	for i := 0; ; {
		cmd := cli.ZRank(key, member)
		err := cmd.Err()
		if err == nil {
			return cmd.Val(), nil
		}

		if err == Nil {
			return -1, nil
		}

		if i++; i >= retry {
			ctx.LogE("redis zrank fail", zap.String("key", key),
				zap.Int("retry", i), zap.Error(err))
			return 0, err
		}
	}

}

func (that *redisStruct) ZRank(ctx Context, key string, member string) (int64, error) {
	return that.zrank(ctx, key, member, 1)
}
func (that *redisStruct) ZRankRetry(ctx Context, key string, member string, retry int) (int64, error) {
	return that.zrank(ctx, key, member, retry)
}
