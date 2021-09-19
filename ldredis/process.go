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

	"github.com/distroy/ldgo/ldlogger"
	"github.com/distroy/ldgo/ldrand"
	"go.uber.org/zap"
)

const (
	ldRedisSrcPath = "/ldgo/ldredis/"
	exampleSrcPath = "/ldgo/ldredis/example/"
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

func getCaller() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !isCallerFilePath(file) {
			continue
		}
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
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

func (c *Redis) defaultProcessPipeline(cmds []Cmder) error {
	retry := c.retry
	reporter := c.reporter
	log := c.logger()
	caller := c.caller
	callerStr := ""

	log = ldlogger.With(log, zap.String("pipeline", hex.EncodeToString(ldrand.Bytes(16))))

	if caller {
		callerStr = getCaller()
	}

	for i := 0; ; {
		begin := time.Now()
		err := c.oldProcessPipeline(cmds)
		cost := time.Since(begin)

		if err == nil || err == Nil {
			for _, cmd := range cmds {
				reporter.Report(cmd, cost)
				if caller {
					log.Debug("redis pipeline cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.String("caller", callerStr))
				} else {
					log.Debug("redis pipeline cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()))
				}
			}
			if caller {
				log.Debug("redis pipeline cmd succ", zap.Int("retry", i), zap.String("caller", callerStr))
			} else {
				log.Debug("redis pipeline cmd succ", zap.Int("retry", i))
			}
			return err
		}

		for _, cmd := range cmds {
			reporter.Report(cmd, cost)
			if err := cmd.Err(); err != nil && err != Nil {
				if caller {
					log.Error("redis pipeline cmd fail", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.Error(err), zap.String("caller", callerStr))
				} else {
					log.Error("redis pipeline cmd fail", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.Error(err))
				}
				break
			}

			if caller {
				log.Debug("redis pipeline cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.String("caller", callerStr))
			} else {
				log.Debug("redis pipeline cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()))
			}
		}

		if i++; i >= retry {
			if caller {
				log.Error("redis pipeline fail", zap.Int("retry", i), zap.Error(err), zap.String("caller", callerStr))
			} else {
				log.Error("redis pipeline fail", zap.Int("retry", i), zap.Error(err))
			}
			return err
		}
	}
}
