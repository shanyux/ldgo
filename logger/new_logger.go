/*
 * Copyright (C) distroy
 */

package logger

import (
	"time"

	"github.com/distroy/ldgo/core"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(opts ...Option) Logger {
	options := newOptions()
	for _, opt := range opts {
		opt(options)
	}

	var level zapcore.Level

	if err := level.UnmarshalText(core.StrToBytesUnsafe(options.level)); err != nil {
		level.UnmarshalText(core.StrToBytesUnsafe("info"))
	}

	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	handle := zapcore.Lock(options.writer)
	zapCores := []zapcore.Core{
		// zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), handle, level),
		zapcore.NewCore(options.encoderBuilder(encoder), handle, level),
	}

	// create options with priority for our opts
	defaultOptions := []zap.Option{}
	if options.enableCaller {
		// defaultOptions = append(defaultOptions, zap.AddStacktrace(level))
		defaultOptions = append(defaultOptions, zap.AddCaller())
		defaultOptions = append(defaultOptions, zap.AddCallerSkip(1))
	}

	core := zapcore.NewTee(zapCores...)

	zlog := zap.New(core, defaultOptions...)

	return newLogger(newLoggerWrap(zlog))
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
