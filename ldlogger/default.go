/*
 * Copyright (C) distroy
 */

package ldlogger

import (
	"io"

	"github.com/distroy/ldgo/ldlog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug  = zapcore.DebugLevel
	LevelInfo   = zapcore.InfoLevel
	LevelWarn   = zapcore.WarnLevel
	LevelError  = zapcore.ErrorLevel
	LevelDpanic = zapcore.DPanicLevel
	LevelPanic  = zapcore.PanicLevel
	LevelFatal  = zapcore.FatalLevel
)

type (
	Logger        = ldlog.LoggerInterface
	LoggerWrapper = ldlog.WrapperInterface
	Option        = ldlog.Option

	encoderBuilder = func(cfg zapcore.EncoderConfig) zapcore.Encoder
)

func SetDefault(l Logger) { ldlog.SetDefault(ldlog.GetLogger(l.Core())) }

func Default() Logger { return ldlog.Default() }
func Console() Logger { return ldlog.Console() }
func Discard() Logger { return ldlog.Discard() }

func NewLogger(opts ...Option) Logger {
	return ldlog.NewLogger(opts...)
}

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	return ldlog.GetLogger(log, fields...)
}

func GetWrapper(l Logger) LoggerWrapper               { return ldlog.GetWrapper(l) }
func With(l Logger, fields ...zap.Field) Logger       { return ldlog.With(l, fields...) }
func WithOptions(l Logger, opts ...zap.Option) Logger { return ldlog.WithOptions(l, opts...) }

func Writer(w io.Writer) Option  { return ldlog.Writer(w) }
func Level(l string) Option      { return ldlog.Level(l) }
func EnableCaller(e bool) Option { return ldlog.EnableCaller(e) }

func Encoder(e encoderBuilder) Option { return ldlog.Encoder(e) }
func JsonEncoder() Option             { return ldlog.JsonEncoder() }

// NewLoggerEncoder
// log format: 2013-04-08 15:30:42.621|FATAL|0x7f865e4b9720|GdpProcessor.cpp(35)|CGdpPrcessor::Init|init_db_client_fail|id=95,db=gpp_db,type=1
func NewLoggerEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return ldlog.NewLoggerEncoder(cfg)
}
