/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newCore(log *zap.Logger, sugar *zap.SugaredLogger) core {
	return core{
		log:   log,
		sugar: sugar,
	}
}

type core struct {
	log      *zap.Logger
	sugar    *zap.SugaredLogger
	sequence string
	rate     rateConfig
}

func (l *core) Sync() { l.log.Sync() }

func (l *core) Enabled(lvl zapcore.Level) bool { return l.enabled(lvl) }

func (l *core) Core() *zap.Logger         { return l.log }
func (l *core) Sugar() *zap.SugaredLogger { return l.sugar }

func (l *core) zCore(skip int) *zap.Logger {
	if !l.checkRateOrInterval(skip + 1) {
		return Discard().Core()
	}
	return l.log
}
func (l *core) zSugar(skip int) *zap.SugaredLogger {
	if !l.checkRateOrInterval(skip + 1) {
		return Discard().Sugar()
	}
	return l.sugar
}

func (l *core) enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }
