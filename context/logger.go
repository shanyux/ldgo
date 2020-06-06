/*
 * Copyright (C) distroy
 */

package context

import (
	"github.com/distroy/ldgo/logger"

	"go.uber.org/zap"
)

var DefaultLogger = logger.NewLogger()

type log struct {
	logger.Logger
}

func newLogger(l logger.Logger) log {
	return log{
		Logger: l,
	}
}

func (that *log) With(fields ...zap.Field) log {
	return newLogger(that.Logger.With(fields...))
}

func (that *log) LogSync() { that.Sync() }

func (that *log) LogD(msg string, fields ...zap.Field) { that.Core().Debug(msg, fields...) }
func (that *log) LogI(msg string, fields ...zap.Field) { that.Core().Info(msg, fields...) }
func (that *log) LogW(msg string, fields ...zap.Field) { that.Core().Warn(msg, fields...) }
func (that *log) LogE(msg string, fields ...zap.Field) { that.Core().Error(msg, fields...) }
func (that *log) LogF(msg string, fields ...zap.Field) { that.Core().Fatal(msg, fields...) }

func (that *log) LogDf(fmt string, args ...interface{}) { that.Sugar().Debugf(fmt, args...) }
func (that *log) LogIf(fmt string, args ...interface{}) { that.Sugar().Infof(fmt, args...) }
func (that *log) LogWf(fmt string, args ...interface{}) { that.Sugar().Warnf(fmt, args...) }
func (that *log) LogEf(fmt string, args ...interface{}) { that.Sugar().Errorf(fmt, args...) }
func (that *log) LogFf(fmt string, args ...interface{}) { that.Sugar().Fatalf(fmt, args...) }
