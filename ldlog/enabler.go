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

// Enable based on probability(rate).
//   - Rate should be in [0, 1.0].
//   - Always enable log levels higher than error.
func RateEnabler(rate float64) Enabler { return rateEnabler{rate: rate} }

// Enable based on time interval.
//   - Calculate the time interval separately at each invocation location.
//   - Always enable log levels higher than error.
func IntervalEnabler(dur time.Duration) Enabler { return intervalEnabler{interval: dur} }

type defaultEnabler struct{}

func (p defaultEnabler) Enable(lvl zapcore.Level, skips ...int) bool { return true }

type falseEnabler struct{}

func (p falseEnabler) Enable(lvl zapcore.Level, skips ...int) bool { return lvl >= zapcore.ErrorLevel }

type rateEnabler struct {
	rate float64
}

func (p rateEnabler) Enable(lvl zapcore.Level, skips ...int) bool {
	if (falseEnabler{}).Enable(lvl, skips...) {
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
	if (falseEnabler{}).Enable(lvl, skips...) {
		return true
	}
	skip := 1
	if len(skips) > 0 && skips[0] > 0 {
		skip = skips[0]
	}
	return hitInterval(p.interval, skip+1)
}
