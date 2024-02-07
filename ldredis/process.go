/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/distroy/ldgo/v2/ldctx"
	"github.com/distroy/ldgo/v2/ldlog"
	"github.com/distroy/ldgo/v2/ldrand"
	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	ldRedisSrcPath = "/ldredis"
	exampleSrcPath = "/ldredis/example"
	goRedisSrcPath = "/github.com/redis/go-redis"
)

func isCallerFilePath(file string) bool {
	if strings.Contains(file, goRedisSrcPath) {
		return false
	}
	if !strings.Contains(file, ldRedisSrcPath) {
		return true
	}
	if strings.HasSuffix(file, "_test.go") {
		return true
	}
	if strings.Contains(file, exampleSrcPath) {
		return true
	}
	return false
}

func getCaller(caller bool) zap.Field {
	if !caller {
		return zap.Skip()
	}
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !isCallerFilePath(file) {
			continue
		}
		return zap.String("caller", fmt.Sprintf("%s:%d", file, line))
	}

	return zap.String("caller", "overflow")
}

func getCmdField(cmd Cmder) zap.Field {
	return zap.Reflect("cmd", cmd.Args())
}

func newHook(rds *Redis) redis.Hook {
	return hook{
		Redis: rds,
	}
}

type hook struct {
	Redis *Redis
}

func (h hook) DialHook(next redis.DialHook) redis.DialHook { return nil }
func (h hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(c context.Context, cmd redis.Cmder) error {
		return h.Process(c, cmd, next)
	}
}
func (h hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(c context.Context, cmds []redis.Cmder) error {
		return h.ProcessPipeline(c, cmds, next)
	}
}

func (h hook) Process(c context.Context, cmd Cmder, next redis.ProcessHook) error {
	if isInProcess(c) {
		return next(c, cmd)
	}

	var (
		ctx           = newContext(c)
		retry         = h.Redis.retry
		retryInterval = h.Redis.retryInterval
		reporter      = h.Redis.reporter
		logger        = ldctx.GetLogger(ctx)
		caller        = h.Redis.caller
	)

	for i := 0; ; {
		begin := time.Now()
		err := next(ctx, cmd)
		cmd.SetErr(err)
		reporter.Report(cmd, time.Since(begin))

		if isErrNil(err) {
			logger.Debug("redis cmd succ", zap.Int("retry", i), getCmdField(cmd), getCaller(caller))
			return err
		}

		if i++; i >= retry {
			logger.Error("redis cmd fail", zap.Int("retry", i), getCmdField(cmd), zap.Error(err), getCaller(caller))
			return err
		}

		time.Sleep(retryInterval)
	}
}

func (h hook) ProcessPipeline(c context.Context, cmds []Cmder, next redis.ProcessPipelineHook) error {
	if isInProcess(c) {
		return next(c, cmds)
	}
	var (
		ctx           = newContext(c)
		retry         = h.Redis.retry
		retryInterval = h.Redis.retryInterval
		reporter      = h.Redis.reporter
		logger        = ldctx.GetLogger(ctx)
		caller        = getCaller(h.Redis.caller)
	)
	logger = logger.With(zap.String("pipeline", hex.EncodeToString(ldrand.Bytes(8))))

	for i := 0; ; {
		begin := time.Now()
		next(ctx, cmds) // nolint
		reporter.ReportPipeline(cmds, time.Since(begin))

		err := h.checkPipelineError(cmds)
		if isErrNil(err) {
			h.printPipelineSuccLog(cmds, i, logger, caller)
			logger.Debug("redis pipeline cmd succ", zap.Int("retry", i), caller)
			return err
		}

		if i++; i >= retry {
			h.printPipelineFailLog(cmds, i, logger, caller)
			logger.Error("redis pipeline fail", zap.Int("retry", i), zap.Error(err), caller)
			return err
		}

		time.Sleep(retryInterval)
	}
}

func (h hook) checkPipelineError(cmds []Cmder) error {
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil && err != Nil {
			return err
		}
	}
	return nil
}

func (h hook) printPipelineSuccLog(cmds []Cmder, retry int, logger *ldlog.Logger, caller zap.Field) {
	for _, cmd := range cmds {
		logger.Debug("redis pipeline cmd succ", zap.Int("retry", retry), getCmdField(cmd), caller)
	}
}

func (h hook) printPipelineFailLog(cmds []Cmder, retry int, logger *ldlog.Logger, caller zap.Field) {
	for _, cmd := range cmds {
		if err := cmd.Err(); !isErrNil(err) {
			logger.Error("redis pipeline cmd fail", zap.Int("retry", retry), getCmdField(cmd),
				zap.Error(err), caller)
			break
		}
		logger.Debug("redis pipeline cmd succ", zap.Int("retry", retry), getCmdField(cmd), caller)
	}
}
