/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	ldRedisSrcPath = "/ldgo/ldredis/"
	goRedisSrcPath = "/github.com/go-redis/redis"
)

func getCaller() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.Contains(file, ldRedisSrcPath) || strings.Contains(file, goRedisSrcPath) {
			continue
		}
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}

func (c *Redis) process(cmd Cmder) error {
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
			if caller {
				log.Debug("redis cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.String("caller", getCaller()))
			} else {
				log.Debug("redis cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()))
			}
			return err
		}

		if i++; i >= retry {
			// log = ldlogger.With(log, fields...)
			if caller {
				log.Error("redis cmd fail", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.Error(err), zap.String("caller", getCaller()))
			} else {
				log.Error("redis cmd fail", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.Error(err))
			}
			return err
		}
	}
}
