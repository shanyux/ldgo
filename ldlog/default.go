/*
 * Copyright (C) distroy
 */

package ldlog

import "github.com/distroy/ldgo/ldio"

var (
	defLogger = NewLogger()
	console   = NewLogger()
	discard   = NewLogger(Writer(ldio.Discard()))
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }
