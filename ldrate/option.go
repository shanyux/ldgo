/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"time"

	"github.com/distroy/ldgo/ldatomic"
)

type Option = func(cfg *config)

type config struct {
	Burst     ldatomic.Int64
	Limit     ldatomic.Int64
	Interval  ldatomic.Duration
	NodeCount ldatomic.Int64
}

func WithBurst(burst int64) Option          { return func(cfg *config) { cfg.Burst.Store(burst) } }
func WithLimit(limit int64) Option          { return func(cfg *config) { cfg.Limit.Store(limit) } }
func WithInterval(dur time.Duration) Option { return func(cfg *config) { cfg.Interval.Store(dur) } }
func WithNodeCount(n int64) Option          { return func(cfg *config) { cfg.NodeCount.Store(n) } }

func newConfig(opts ...Option) config {
	cfg := config{
		Burst:     1,
		Limit:     1,
		Interval:  0,
		NodeCount: 1,
	}

	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}
