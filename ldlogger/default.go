/*
 * Copyright (C) distroy
 */

package ldlogger

var (
	defLogger = NewLogger()
	console   = NewLogger()
)

func SetDefault(l Logger) { defLogger = l }

func Default() Logger { return defLogger }
func Console() Logger { return console }
