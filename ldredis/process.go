/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/distroy/ldgo/ldlog"
	"github.com/distroy/ldgo/ldrand"
	"go.uber.org/zap"
)

const (
	ldRedisSrcPath = "/ldgo/ldredis"
	exampleSrcPath = "/ldgo/ldredis/example"
	goRedisSrcPath = "/github.com/go-redis/redis"
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

func getCallerField(caller bool) zap.Field {
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

func (c *Redis) defaultProcess(cmd Cmder) error {
	retry := c.retry
	reporter := c.reporter
	log := c.logger()
	caller := c.caller

	for i := 0; ; {
		begin := time.Now()
		c.oldProcess(cmd)
		reporter.Report(cmd, time.Since(begin))

		err := cmd.Err()
		if err == nil || err == Nil {
			log.Debug("redis cmd succ", zap.Int("retry", i), getCmdField(cmd), getCallerField(caller))
			return err
		}

		if i++; i >= retry {
			// log = ldlog.With(log, fields...)
			log.Error("redis cmd fail", zap.Int("retry", i), getCmdField(cmd), zap.Error(err), getCallerField(caller))
			return err
		}
	}
}

func (c *Redis) defaultProcessPipeline(cmds []Cmder) error {
	retry := c.retry
	reporter := c.reporter
	log := c.logger()

	caller := getCallerField(c.caller)
	log = log.With(zap.String("pipeline", hex.EncodeToString(ldrand.Bytes(8))))

	for i := 0; ; {
		begin := time.Now()
		c.oldProcessPipeline(cmds)
		reporter.ReportPipeline(cmds, time.Since(begin))

		err := c.checkPipelineError(cmds)
		if err == nil || err == Nil {
			c.printPipelineSuccLog(cmds, i, log, caller)
			log.Debug("redis pipeline cmd succ", zap.Int("retry", i), caller)
			return err
		}

		if i++; i >= retry {
			c.printPipelineFailLog(cmds, i, log, caller)
			log.Error("redis pipeline fail", zap.Int("retry", i), zap.Error(err), caller)
			return err
		}
	}
}

func (c *Redis) checkPipelineError(cmds []Cmder) error {
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil && err != Nil {
			return err
		}
	}
	return nil
}

func (c *Redis) printPipelineSuccLog(cmds []Cmder, retry int, log *ldlog.Logger, caller zap.Field) {
	for _, cmd := range cmds {
		log.Debug("redis pipeline cmd succ", zap.Int("retry", retry), getCmdField(cmd), caller)
	}
}

func (c *Redis) printPipelineFailLog(cmds []Cmder, retry int, log *ldlog.Logger, caller zap.Field) error {
	for _, cmd := range cmds {
		if err := cmd.Err(); err != nil && err != Nil {
			log.Error("redis pipeline cmd fail", zap.Int("retry", retry), getCmdField(cmd),
				zap.Error(err), caller)
			return err
		}
		log.Debug("redis pipeline cmd succ", zap.Int("retry", retry), getCmdField(cmd), caller)
	}
	return nil
}
