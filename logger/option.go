/*
 * Copyright (C) distroy
 */

package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

const (
	LOG_LEVEL         = "INFO"
	LOG_ENABLE_CALLER = true
)

func newOptions() *options {
	return &options{
		writer:         os.Stdout,
		level:          LOG_LEVEL,
		enableCaller:   LOG_ENABLE_CALLER,
		encoderBuilder: NewLoggerEncoder,
	}
}

type encoderBuilder = func(cfg zapcore.EncoderConfig) zapcore.Encoder

type options struct {
	writer         zapcore.WriteSyncer
	level          string
	enableCaller   bool
	encoderBuilder encoderBuilder
}

type Option func(*options)

func Writer(w zapcore.WriteSyncer) Option { return func(o *options) { o.writer = w } }
func Level(l string) Option               { return func(o *options) { o.level = l } }
func EnableCaller(e bool) Option          { return func(o *options) { o.enableCaller = e } }

func Encoder(e encoderBuilder) Option { return func(o *options) { o.encoderBuilder = e } }
func JsonEncoder() Option             { return Encoder(zapcore.NewJSONEncoder) }
