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

	for i := 0; ; {
		begin := time.Now()
		c.oldProcess(cmd)
		reporter.Report(cmd, time.Since(begin))

		err := cmd.Err()
		if isErrorNil(err) {
			return err
		}

		if i++; i >= retry {
			log := c.logger
			// log = ldlogger.With(log, fields...)
			log.Error("redis cmd fail", zap.Int("retry", i), zap.Error(err), zap.Reflect("cmd", cmd.Args()))
			return err
		}
	}
}
