/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"time"

	"github.com/distroy/ldgo/v2/ldrand"
	"go.uber.org/zap/zapcore"
)

type Enabler interface {
	Enable(lvl zapcore.Level, skip ...int) bool
}

func RateEnabler(rate float64) Enabler          { return rateEnabler{rate: rate} }
func IntervalEnabler(dur time.Duration) Enabler { return intervalEnabler{interval: dur} }

type defaultEnabler struct{}

func (p defaultEnabler) Enable(lvl zapcore.Level, skips ...int) bool { return true }

type rateEnabler struct {
	rate float64
}

func (p rateEnabler) Enable(lvl zapcore.Level, skips ...int) bool {
	if lvl >= zapcore.ErrorLevel {
		return true
	}

	rate := p.rate
	if rate <= 0 {
		return false
	}
	if rate >= 1 {
		return true
	}
	return ldrand.Float64() < rate
}

type intervalEnabler struct {
	interval time.Duration
}

func (p intervalEnabler) Enable(lvl zapcore.Level, skips ...int) bool {
	skip := 1
	if len(skips) > 0 && skips[0] > 0 {
		skip = skips[0]
	}
	if lvl >= zapcore.ErrorLevel {
		return true
	}
	return hitInterval(p.interval, skip+1)
}
