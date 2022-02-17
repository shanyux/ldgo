/*
 * Copyright (C) distroy
 */

package ldlog

var (
	defLogger = NewLogger()
	console   = NewLogger()
	discard   = NewLogger(Writer(writerDiscard{}))
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }

type writerDiscard struct{}

func (writerDiscard) Write(p []byte) (int, error)       { return len(p), nil }
func (writerDiscard) WriteString(s string) (int, error) { return len(s), nil }
