/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"reflect"

	"github.com/distroy/ldgo/v2/ldconv"
	"github.com/distroy/ldgo/v2/lderr"
	"go.uber.org/zap"
)

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(b []byte) (interface{}, error)
}

// ***** redis marshal begin *****

func (c *CodecRedis) marshal(value interface{}) (interface{}, error) {
	if value == nil {
		return nil, nil
	}
	bytes, err := c.codec.Marshal(value)
	if err != nil {
		return errorMarshaler{err: err}, err
	}
	return ldconv.BytesToStrUnsafe(bytes), nil
}

func (c *CodecRedis) mustMarshal(value interface{}) interface{} {
	res, _ := c.marshal(value)
	return res
}

func (c *CodecRedis) marshalPairs(pairs []interface{}) []interface{} {
	for i, v := range pairs {
		if i%2 == 0 {
			continue
		}

		val, err := c.marshal(v)
		pairs[i] = val
		if err != nil {
			break
		}
	}
	return pairs
}

func (c *CodecRedis) marshalSlice(s []interface{}) []interface{} {
	for i, v := range s {
		val, err := c.marshal(v)
		s[i] = val
		if err != nil {
			break
		}
	}
	return s
}

func (c *CodecRedis) marshalStringInterfaceMap(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		val, err := c.marshal(v)
		m[k] = val
		if err != nil {
			break
		}
	}
	return m
}

func (c *CodecRedis) marshalZMember(v ZMember) ZMember {
	v.Member = c.mustMarshal(v.Member)
	return v
}
func (c *CodecRedis) marshalZMembers(members []ZMember) []ZMember {
	for i := range members {
		v := &members[i]
		val, err := c.marshal(v.Member)
		v.Member = val
		if err != nil {
			break
		}
	}
	return members
}

// ***** redis marshal end *****

// ***** redis unmarshal begin *****

func (c *CodecRedis) unmarshal(bytes []byte) (interface{}, error) {
	if len(bytes) == 0 {
		return nil, nil
	}
	v, err := c.codec.Unmarshal(bytes)
	if err != nil {
		c.logger().Error("redis codec unmarshal fail", zap.Error(err),
			getCaller(c.caller))
		return nil, lderr.ErrCacheUnmarshal
	}
	return v, nil
}

func (c *CodecRedis) unmarshalInterface(i interface{}) (interface{}, error) {
	if i == nil {
		return nil, nil
	}

	var bytes []byte
	switch tmp := i.(type) {
	case []byte:
		bytes = tmp
	case string:
		bytes = ldconv.StrToBytesUnsafe(tmp)
	default:
		c.logger().Error("redis codec cannot unmarshal the type",
			zap.Stringer("type", reflect.TypeOf(i)), getCaller(c.caller))
		return nil, lderr.ErrCacheUnmarshal
	}
	return c.unmarshal(bytes)
}

// ***** redis unmarshal end *****
