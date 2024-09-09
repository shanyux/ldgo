/*
 * Copyright (C) distroy
 */

package ldlog

var (
	defLogger = New()
	console   = New()
	discard   = newDiscard()
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }
