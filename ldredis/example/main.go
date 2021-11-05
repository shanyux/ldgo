/*
 * Copyright (C) distroy
 */

package main

import (
	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/ldlog"
	"github.com/distroy/ldgo/ldredis"
	"go.uber.org/zap"
)

func init() {
	log := ldlog.NewLogger(ldlog.Level("debug"))
	ldlog.SetDefault(log)
}

func newRedis() *ldredis.Redis {
	rds := ldredis.NewByConfig(&ldredis.Config{
		Addr: "proxy.codis-toc.test.shopeemobile.com:9000",
	})

	return rds
}

func pipeline(ctx ldctx.Context) {
	rds := newRedis()
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

func main() {
	ctx := ldctx.Default()
	pipeline(ctx)
}
