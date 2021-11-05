/*
 * Copyright (C) distroy
 */

package ldlog

import "io"

var (
	defLogger = NewLogger()
	console   = NewLogger()
	discard   = NewLogger(Writer(io.Discard))
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }
