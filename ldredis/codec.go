/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"github.com/distroy/ldgo/ldconv"
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte) (interface{}, error)
}

var _ CodecCmdable = (*CodecRedis)(nil)

type CodecRedis struct {
	redis *Redis
	codec Codec
}

func (c *CodecRedis) Get(key string) *CodecCmd {
	res := &CodecCmd{}
	cmd := c.redis.Get(key)
	res.parse(c.codec, cmd)
	return res
}

func (c *CodecRedis) newMarshaler(value interface{}) interface{} {
	bytes, err := c.codec.Marshal(value)
	if err != nil {
		return errorMarshaler{err: err}
	}
	return ldconv.BytesToStrUnsafe(bytes)
}

func (c *CodecRedis) GetSet(key string, value interface{}) *CodecCmd {
	val := c.newMarshaler(value)
	res := &CodecCmd{}
	cmd := c.redis.GetSet(key, val)
	res.parse(c.codec, cmd)
	return res
}

func (c *CodecRedis) Set(key string, value interface{}, expiration time.Duration) *StatusCmd {
	val := c.newMarshaler(value)
	return c.redis.Set(key, val, expiration)
}

func (c *CodecRedis) SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd {
	val := c.newMarshaler(value)
	return c.redis.SetNX(key, val, expiration)
}

func (c *CodecRedis) SetXX(key string, value interface{}, expiration time.Duration) *BoolCmd {
	val := c.newMarshaler(value)
	return c.redis.SetXX(key, val, expiration)
}

func (c *CodecRedis) HGet(key, field string) *CodecCmd {
	res := &CodecCmd{}
	cmd := c.redis.HGet(key, field)
	res.parse(c.codec, cmd)
	return res
}

func (c *CodecRedis) HGetAll(key string) *StringCodecMapCmd {
	res := &StringCodecMapCmd{}
	cmd := c.redis.HGetAll(key)
	res.parse(c.codec, cmd)
	return res
}

// HMGet(key string, fields ...string) *SliceCmd
func (c *CodecRedis) HMSet(key string, fields map[string]interface{}) *StatusCmd {
	m := make(map[string]interface{}, len(fields))
	for f, v := range fields {
		m[f] = c.newMarshaler(v)
	}
	return c.redis.HMSet(key, m)
}

func (c *CodecRedis) HSet(key, field string, value interface{}) *BoolCmd {
	val := c.newMarshaler(value)
	return c.redis.HSet(key, field, val)
}

func (c *CodecRedis) HSetNX(key, field string, value interface{}) *BoolCmd {
	val := c.newMarshaler(value)
	return c.redis.HSetNX(key, field, val)
}

func (c *CodecRedis) HVals(key string) *CodecSliceCmd {
	res := &CodecSliceCmd{}
	cmd := c.redis.HVals(key)
	res.parse(c.codec, cmd)
	return res
}
