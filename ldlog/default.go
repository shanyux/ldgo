/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defLogger = New()
	console   = New()
	discard   = newDiscard()
)

func SetDefault(l *Logger) { defLogger = l }

func Default() *Logger { return defLogger }
func Console() *Logger { return console }
func Discard() *Logger { return discard }

func zCoreByLogger(l *Logger, lvl zapcore.Level, skip int) *zap.Logger {
	return l.zCore(lvl, skip+1)
}
func zSugarByLogger(l *Logger, lvl zapcore.Level, skip int) *zap.SugaredLogger {
	return l.zSugar(lvl, skip+1)
}
