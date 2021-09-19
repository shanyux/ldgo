/*
 * Copyright (C) distroy
 */

package ldredis

import "time"

type Reporter interface {
	Report(cmd Cmder, d time.Duration)
	ReportPipeline(cmds []Cmder, d time.Duration)
}

type discardReporter struct{}

func (_ discardReporter) Report(cmd Cmder, d time.Duration) {
	// cmd.Name()
	// cmd.Err()
}

func (_ discardReporter) ReportPipeline(cmds []Cmder, d time.Duration) {
}
