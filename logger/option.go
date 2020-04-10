/*
 * Copyright (C) distroy
 */

package logger

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var _ = bytes.NewBuffer
var _ = fmt.Sprintf
var _ = log.New
var _ = math.Abs
var _ = os.Exit
var _ = strconv.Itoa
var _ = strings.Replace
var _ = sync.NewCond
var _ = time.Now

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
