/*
 * Copyright (C) distroy
 */

package ldlog

import "github.com/distroy/ldgo/v2/ldio"

var (
	defLogger = New()
	console   = New()
	discard   = New(Writer(ldio.Discard()), Level("dpanic"))
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }
