/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"
	"time"
	"unsafe"

	"github.com/distroy/ldgo/v2/ldredis/internal"
	redis "github.com/redis/go-redis/v9"
)

func getOptions(c *Redis) *internal.Options     { return &c.opts }
func getOptionsPointer(c *Redis) unsafe.Pointer { return unsafe.Pointer(&c.opts) }

type cmdable interface {
	Cmdable

	AddHook(hook redis.Hook)

	// MSet is like Set but accepts multiple values:
	//   - MSet("key1", "value1", "key2", "value2")
	//   - MSet([]string{"key1", "value1", "key2", "value2"})
	//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
	//   - MSet(struct), For struct types, see HSet description.
	MSet(ctx context.Context, values ...interface{}) *StatusCmd
	// MSetNX is like SetNX but accepts multiple values:
	//   - MSetNX("key1", "value1", "key2", "value2")
	//   - MSetNX([]string{"key1", "value1", "key2", "value2"})
	//   - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
	//   - MSetNX(struct), For struct types, see HSet description.
	MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
	// Set Redis `SET key value [expiration]` command.
	// Use expiration for `SETEx`-like behavior.
	//
	// Zero expiration means the key has no expiration time.
	// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
	// otherwise you will receive an error: (error) ERR syntax error.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd

	// HSet accepts values in following formats:
	//
	//   - HSet("myhash", "key1", "value1", "key2", "value2")
	//
	//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	//
	//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
	//
	//     Playing struct With "redis" tag.
	//     type MyHash struct { Key1 string `redis:"key1"`; Key2 int `redis:"key2"` }
	//
	//   - HSet("myhash", MyHash{"value1", "value2"}) Warn: redis-server >= 4.0
	//
	//     For struct, can be a structure pointer type, we only parse the field whose tag is redis.
	//     if you don't want the field to be read, you can use the `redis:"-"` flag to ignore it,
	//     or you don't need to set the redis tag.
	//     For the type of structure field, we only support simple data types:
	//     string, int/uint(8,16,32,64), float(32,64), time.Time(to RFC3339Nano), time.Duration(to Nanoseconds ),
	//     if you are other more complex or custom data types, please implement the encoding.BinaryMarshaler interface.
	//
	// Note that in older versions of Redis server(redis-server < 4.0), HSet only supports a single key-value pair.
	// redis-docs: https://redis.io/commands/hset (Starting with Redis version 4.0.0: Accepts multiple field and value arguments.)
	// If you are using a Struct type and the number of fields is greater than one,
	// you will receive an error similar to "ERR wrong number of arguments", you can use HMSet as a substitute.
	HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
}

func newRedisClient(cfg *Config) cmdable {
	return redis.NewClient(cfg.toClient())
}

func newRedisCluster(cfg *Config) cmdable {
	return redis.NewClusterClient(cfg.toCluster())
}
