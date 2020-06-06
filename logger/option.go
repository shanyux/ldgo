/*
 * Copyright (C) distroy
 */

package logger

import (
	"io"
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

func writeSyncer(w io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(w)
}

func Writer(w io.Writer) Option  { return func(o *options) { o.writer = writeSyncer(w) } }
func Level(l string) Option      { return func(o *options) { o.level = l } }
func EnableCaller(e bool) Option { return func(o *options) { o.enableCaller = e } }

func Encoder(e encoderBuilder) Option { return func(o *options) { o.encoderBuilder = e } }
func JsonEncoder() Option             { return Encoder(zapcore.NewJSONEncoder) }
