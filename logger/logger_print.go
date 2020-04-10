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

	// "go.uber.org/zap"
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
	_DEBUG  = zapcore.DebugLevel
	_INFO   = zapcore.InfoLevel
	_WARN   = zapcore.WarnLevel
	_ERROR  = zapcore.ErrorLevel
	_DPANIC = zapcore.DPanicLevel
	_PANIC  = zapcore.PanicLevel
	_FATAL  = zapcore.FatalLevel
)

func sprintln(args []interface{}) string {
	if len(args) == 0 {
		return ""
	}

	text := fmt.Sprintln(args...)
	size := len(text)
	if size == 0 {
		return ""
	}
	if text[size-1] == '\n' {
		return text[:size-1]
	}
	return text
}
