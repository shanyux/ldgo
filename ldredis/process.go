/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"go.uber.org/zap"
)

func (c *Redis) process(cmd Cmder) error {
	retry := c.retry
	reporter := c.reporter
	log := c.logger

	for i := 0; ; {
		begin := time.Now()
		c.oldProcess(cmd)
		reporter.Report(cmd, time.Since(begin))

		err := cmd.Err()
		if isErrorNil(err) {
			log.Debug("redis cmd succ", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()))
			return err
		}

		if i++; i >= retry {
			// log = ldlogger.With(log, fields...)
			log.Error("redis cmd fail", zap.Int("retry", i), zap.Reflect("cmd", cmd.Args()), zap.Error(err))
			return err
		}
	}
}
