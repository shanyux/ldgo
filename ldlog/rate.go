/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"runtime"
	"sync"
	"time"

	"github.com/distroy/ldgo/v2/ldatomic"
	"github.com/distroy/ldgo/v2/ldrand"
	"go.uber.org/zap"
)

type rateConfig struct {
	Rate     float64
	Interval time.Duration
}

const (
	ctxCallerCacheSizeMax = 100 * 10000
)

var (
	ctxCallerCache     = &sync.Map{}
	ctxCallerCacheSize = ldatomic.NewInt(0)
)

type caller struct {
	File string
	Line int
}

type ctxCallerInfo struct {
	LastLogTime ldatomic.Time // for log at intervals
}

func (l *Logger) WithRate(rate float64) *Logger {
	l = l.clone()
	l.rate = rateConfig{Rate: rate + 1}
	return l
}

func (l *Logger) WithInterval(d time.Duration) *Logger {
	l = l.clone()
	l.rate = rateConfig{Interval: d}
	return l
}

func (l *core) CheckRateOrInterval(skip int) bool { return l.checkRateOrInterval(skip + 1) }
func (l *core) checkRateOrInterval(skip int) bool {
	c := l
	if d := c.rate.Interval; d > 0 {
		return hitInterval(d, skip+1)
	}
	if rate := c.rate.Rate - 1; rate >= 0 {
		return hitRate(rate)
	}
	return true
}

func hitRate(rate float64) bool {
	if rate >= 1 {
		return true
	}
	if rate <= 0 {
		return false
	}
	return ldrand.Float64() < rate
}

func hitInterval(dur time.Duration, skip int) bool {
	if dur <= 0 {
		return true
	}
	caller := getCaller(skip + 1)
	if caller == nil {
		return false
	}

	lastLogTime := caller.LastLogTime.Load()
	nextLogTime := lastLogTime.Add(dur)
	now := time.Now()
	if now.After(nextLogTime) {
		return caller.LastLogTime.CompareAndSwap(lastLogTime, now)
	}

	return false
}

func getCaller(skip int) *ctxCallerInfo {
	_, file, line, _ := runtime.Caller(skip + 1)
	key := caller{
		File: file,
		Line: line,
	}

	i, _ := ctxCallerCache.Load(key)
	v, _ := i.(*ctxCallerInfo)
	if v != nil {
		return v
	}

	if ctxCallerCacheSize.Load() >= ctxCallerCacheSizeMax {
		Default().Error("reached the limit of interval callers, drop it",
			zap.String("file", file), zap.Int("line", line))
		return nil
	}

	v = &ctxCallerInfo{}
	v.LastLogTime.Store(time.Time{})
	i, loaded := ctxCallerCache.LoadOrStore(key, v)
	if loaded {
		v, _ = i.(*ctxCallerInfo)
		return v
	}

	n := ctxCallerCacheSize.Add(1)
	if n <= ctxCallerCacheSizeMax {
		return v
	}

	ctxCallerCache.Delete(key)
	Default().Error("reached the limit of interval callers, drop it",
		zap.String("file", file), zap.Int("line", line))
	return nil
}
