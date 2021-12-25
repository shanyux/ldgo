/*
 * Copyright (C) distroy
 */

package main

import (
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/ldlog"
	"github.com/distroy/ldgo/ldredis"
	"go.uber.org/zap"
)

func init() {
	log := ldlog.NewLogger(ldlog.Level("debug"))
	ldlog.SetDefault(log)
}

func newRedis(ctx ldctx.Context) *ldredis.Redis {
	rds := ldredis.NewByConfig(&ldredis.Config{
		Addr: "proxy.codis-toc.test.shopeemobile.com:9000",
	})

	return rds.WithContext(ctx)
}

func pipeline(ctx ldctx.Context) {
	rds := newRedis(ctx)
	rds = rds.WithRetry(3)
	// rds = rds.WithCaller(false)
	defer rds.Close()

	p := rds.Pipeline()
	p.Get("key1")
	p.Get("key2")
	p.Get("key3")
	cmds, err := p.Exec()
	ctx.LogI("pipeline return", zap.Error(err))
	for _, v := range cmds {
		cmd, _ := v.(*ldredis.StringCmd)
		ctx.LogI("pipeline res", zap.Reflect("cmd", cmd.Args()), zap.String("val", cmd.Val()), zap.Error(cmd.Err()))
	}
}

func codecStruct(ctx ldctx.Context) {
	type codecStruct struct {
		Str1 string `json:"str1"`
		Str2 string `json:"str2"`
		Int1 int64  `json:"int1"`
		Int2 int64  `json:"int2"`
	}

	rds := newRedis(ctx)
	defer rds.Close()
	key := "test:trc:abc"

	sCmd := rds.WithCodec(ldredis.NewJsonCodec()).Set(key, &codecStruct{
		Str1: "aaa",
		Str2: "bbb",
		Int1: 111,
		Int2: 222,
	}, time.Minute)
	ctx.LogI("cmd", zap.Reflect("cmd", sCmd.Args()))

	gCmd0 := rds.WithCodec(ldredis.NewJsonCodec(&codecStruct{})).Get(key)
	ctx.LogIf("type:%T, value:%v", gCmd0.Val(), gCmd0.Val())

	gCmd1 := rds.WithCodec(ldredis.NewJsonCodec()).Get(key)
	ctx.LogIf("type:%T, value:%v", gCmd1.Val(), gCmd1.Val())
}

func codeBaseType(ctx ldctx.Context) {
	rds := newRedis(ctx)
	defer rds.Close()
}

func main() {
	ctx := ldctx.Default()
	// pipeline(ctx)
	codecStruct(ctx)
}
