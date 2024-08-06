/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"runtime"
	"sync"
	"time"

	"github.com/distroy/ldgo/v2/ldatomic"
	"go.uber.org/zap"
)

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
	cache := ctxCallerCache
	_, file, line, _ := runtime.Caller(skip + 1)
	key := caller{
		File: file,
		Line: line,
	}

	i, _ := cache.Load(key)
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
	i, loaded := cache.LoadOrStore(key, v)
	if loaded {
		v, _ = i.(*ctxCallerInfo)
		return v
	}

	n := ctxCallerCacheSize.Add(1)
	if n <= ctxCallerCacheSizeMax {
		return v
	}

	cache.Delete(key)
	Default().Error("reached the limit of interval callers, drop it",
		zap.String("file", file), zap.Int("line", line))
	return nil
}
