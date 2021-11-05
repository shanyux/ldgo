/*
 * Copyright (C) distroy
 */

package logger

import (
	"github.com/distroy/ldgo/ldlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger = ldlog.LoggerInterface
	Option = ldlog.Option
)

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	return ldlog.GetLogger(log, fields...)
}

func NewLogger(opts ...Option) Logger {
	return ldlog.NewLogger(opts...)
}

func Default() Logger { return ldlog.Default() }
func Console() Logger { return ldlog.Console() }

// NewLoggerEncoder
// log format: 2013-04-08 15:30:42.621|FATAL|0x7f865e4b9720|GdpProcessor.cpp(35)|CGdpPrcessor::Init|init_db_client_fail|id=95,db=gpp_db,type=1
func NewLoggerEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return ldlog.NewLoggerEncoder(cfg)
}

type LoggerWrap = ldlog.WrapperInterface
