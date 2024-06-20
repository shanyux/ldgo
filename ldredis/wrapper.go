/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"reflect"
	"unsafe"

	"github.com/distroy/ldgo/v2/ldredis/internal"
	"github.com/distroy/ldgo/v2/ldref"
	redis "github.com/redis/go-redis/v9"
)

func getOptions(c *Redis) *internal.Options     { return &c.opts }
func getOptionsPointer(c *Redis) unsafe.Pointer { return unsafe.Pointer(&c.opts) }

type cmdable interface {
	Cmdable

	AddHook(hook redis.Hook)
	Clone() cmdable
}

func newRedisClient(cfg *Config) cmdable {
	cli := redis.NewClient(cfg.toClient())
	return newWrapper(cli)
}

func newRedisCluster(cfg *Config) cmdable {
	cli := redis.NewClusterClient(cfg.toCluster())
	return newWrapper(cli)
}

type cmdableInWrapper interface {
	Cmdable

	AddHook(hook redis.Hook)
}

func newWrapper(c cmdableInWrapper) wrapper { return wrapper{cmdableInWrapper: c} }

type wrapper struct {
	cmdableInWrapper
}

func (c wrapper) Clone() cmdable {
	c.cmdableInWrapper = ldref.Clone(c.cmdableInWrapper)
	v := reflect.ValueOf(&c.cmdableInWrapper).Elem().Elem()
	v0 := v
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return c
	}

	c.cloneField(v, "hooksMixin")
	c.setField(v, "cmdable", v0.MethodByName("Process"))

	return c
}

func (c wrapper) cloneField(v reflect.Value, fieldName string) {
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		return
	}

	a0 := unsafe.Pointer(f.UnsafeAddr())
	o0 := reflect.NewAt(f.Type(), a0).Elem()

	o1 := ldref.DeepClone(o0)
	o0.Set(o1)
}

func (c wrapper) setField(v reflect.Value, fieldName string, fieldValue reflect.Value) {
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		return
	}

	a0 := unsafe.Pointer(f.UnsafeAddr())
	o0 := reflect.NewAt(f.Type(), a0).Elem()

	// log.Printf(" === type 1:%s, 2:%s", o0.Type().String(), fieldValue.Type().String())
	// if o0.Type() != fieldValue.Type() {
	// 	fieldValue = fieldValue.Convert(o0.Type())
	// }

	o0.Set(fieldValue)
}
