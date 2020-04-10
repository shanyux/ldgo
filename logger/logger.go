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

	"go.uber.org/zap"
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

func newLogger(log loggerWrap) logger {
	return logger{
		loggerWrap: log,
	}
}

type Logger interface {
	Sync()

	Wrap() LoggerWrap
	Core() *zap.Logger
	Sugar() *zap.SugaredLogger

	With(fields ...zap.Field) Logger
	WithOptions(opts ...zap.Option) Logger

	Debug(msg string, fields ...zap.Field)
	Debugf(fmt string, args ...interface{})
	Debugln(args ...interface{})

	Info(msg string, fields ...zap.Field)
	Infof(fmt string, args ...interface{})
	Infoln(args ...interface{})

	Warn(msg string, fields ...zap.Field)
	Warnf(fmt string, args ...interface{})
	Warnln(args ...interface{})

	Error(msg string, fields ...zap.Field)
	Errorf(fmt string, args ...interface{})
	Errorln(args ...interface{})

	Fatal(msg string, fields ...zap.Field)
	Fatalf(fmt string, args ...interface{})
	Fatalln(args ...interface{})
}

func GetLogger(log *zap.Logger, fields ...zap.Field) Logger {
	if len(fields) > 0 {
		log = log.With(fields...)
	}
	return newLogger(newLoggerWrap(log))
}

type logger struct {
	loggerWrap
}

func (l logger) With(fields ...zap.Field) Logger {
	log := l.log.With(fields...)
	return newLogger(newLoggerWrap(log))
}

func (l logger) WithOptions(opts ...zap.Option) Logger {
	log := l.log.WithOptions(opts...)
	return newLogger(newLoggerWrap(log))
}

func (l logger) Wrap() LoggerWrap { return l.loggerWrap }

func (l logger) Debug(msg string, fields ...zap.Field) { l.log.Debug(msg, fields...) }
func (l logger) Info(msg string, fields ...zap.Field)  { l.log.Info(msg, fields...) }
func (l logger) Warn(msg string, fields ...zap.Field)  { l.log.Warn(msg, fields...) }
func (l logger) Error(msg string, fields ...zap.Field) { l.log.Error(msg, fields...) }
func (l logger) Fatal(msg string, fields ...zap.Field) { l.log.Fatal(msg, fields...) }
