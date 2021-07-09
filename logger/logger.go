/*
 * Copyright (C) distroy
 */

package logger

import (
	"context"

	"github.com/distroy/ldgo/ldlogger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = ldlogger.Logger

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	return ldlogger.GetLogger(log, fields...)
}

const (
	LOG_LEVEL         = ldlogger.LOG_LEVEL
	LOG_ENABLE_CALLER = ldlogger.LOG_ENABLE_CALLER
)

type Option = ldlogger.Option

func NewLogger(opts ...Option) Logger {
	return ldlogger.NewLogger(opts...)
}

func Default() Logger { return ldlogger.Default() }
func Console() Logger { return ldlogger.Console() }

func NewContext(ctx context.Context, l Logger) context.Context {
	return ldlogger.NewContext(ctx, l)
}

func FromContext(ctx context.Context) Logger {
	return ldlogger.FromContext(ctx)
}

// NewLoggerEncoder
// log format: 2013-04-08 15:30:42.621|FATAL|0x7f865e4b9720|GdpProcessor.cpp(35)|CGdpPrcessor::Init|init_db_client_fail|id=95,db=gpp_db,type=1
func NewLoggerEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return ldlogger.NewLoggerEncoder(cfg)
}

type LoggerWrap = ldlogger.LoggerWrap
