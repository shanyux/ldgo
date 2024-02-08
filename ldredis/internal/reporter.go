/*
 * Copyright (C) distroy
 */

package internal

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Reporter interface {
	Report(cmd redis.Cmder, d time.Duration)
	ReportPipeline(cmds []redis.Cmder, d time.Duration)
}

type DiscardReporter struct{}

func (_ DiscardReporter) Report(cmd redis.Cmder, d time.Duration) {
	// cmd.Name()
	// cmd.Err()
}

func (_ DiscardReporter) ReportPipeline(cmds []redis.Cmder, d time.Duration) {
}
