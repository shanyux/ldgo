/*
 * Copyright (C) distroy
 */

package ldredis

import "time"

type Reporter interface {
	Report(cmd Cmder, d time.Duration)
}

type discardReporter struct{}

func (_ discardReporter) Report(cmd Cmder, d time.Duration) {
	// cmd.Name()
	// cmd.Err()
}
