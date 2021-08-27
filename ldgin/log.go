/*
 * Copyright (C) distroy
 */

package ldgin

import "github.com/distroy/ldgo/ldlogger"

var logger ldlogger.Logger

func SetLogger(l ldlogger.Logger) {
	logger = l
}

func log() ldlogger.Logger {
	if logger != nil {
		return logger
	}
	return ldlogger.Default()
}
