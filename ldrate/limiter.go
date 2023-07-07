/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"time"

	"github.com/distroy/ldgo/ldatomic"
	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldmath"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Limiter struct {
	limiter *rate.Limiter
	config  config

	lastBurst     ldatomic.Int64
	lastLimit     ldatomic.Int64
	lastInterval  ldatomic.Duration
	lastNodeCount ldatomic.Int64
}

func NewLimiter(opts ...Option) *Limiter {
	cfg := newConfig(opts...)
	l := &Limiter{
		limiter: &rate.Limiter{},
		config:  cfg,
	}

	return l
}

func (l *Limiter) SetBurst(burst int64)        { l.config.Burst.Store(burst) }
func (l *Limiter) SetLimit(limit int64)        { l.config.Limit.Store(limit) }
func (l *Limiter) SetInterval(d time.Duration) { l.config.Interval.Store(d) }
func (l *Limiter) SetNodeCount(n int64)        { l.config.NodeCount.Store(n) }
func (l *Limiter) SetOptions(opts ...Option) {
	cfg := &l.config
	for _, fn := range opts {
		fn(cfg)
	}
}

func (l *Limiter) refresh(ctx ldctx.Context, now time.Time) lderr.Error {
	cfg := &l.config

	burst := cfg.Burst.Load()
	burst = ldmath.MaxInt64(burst, 1)
	if burst != l.lastBurst.Load() {
		l.lastBurst.Store(burst)
		l.limiter.SetBurstAt(now, int(burst))
		ctx.LogI("[distributed limiter] refresh the burst succ", zap.Int64("burst", burst))
	}

	limit := cfg.Limit.Load()
	interval := cfg.Interval.Load()
	nodeCount := cfg.NodeCount.Load()

	if interval < 0 || limit <= 0 || nodeCount <= 0 {
		ctx.LogE("[distributed limiter] invalid rate every parameters", zap.Int64("limit", limit),
			zap.Stringer("interval", interval), zap.Int64("nodeCount", nodeCount))
		return lderr.ErrInternalServerError
	}

	if limit == l.lastLimit.Load() && interval == l.lastInterval.Load() && nodeCount == l.lastNodeCount.Load() {
		return nil
	}

	every := interval * time.Duration(nodeCount) / time.Duration(limit)
	if every < 0 {
		ctx.LogE("[distributed limiter] invalid rate every", zap.Int64("limit", limit),
			zap.Stringer("interval", interval), zap.Int64("serviceCount", nodeCount))
		return lderr.ErrInternalServerError
	}

	l.limiter.SetLimitAt(now, rate.Every(every))
	l.lastLimit.Store(limit)
	l.lastInterval.Store(interval)
	l.lastNodeCount.Store(nodeCount)

	ctx.LogI("[distributed limiter] refresh rate every succ", zap.Int64("limit", limit),
		zap.Stringer("interval", interval), zap.Int64("serviceCount", nodeCount),
		zap.Stringer("every", every))
	return nil
}

func (l *Limiter) Wait(ctx ldctx.Context) lderr.Error {
	return l.WaitN(ctx, 1)
}

func (l *Limiter) WaitN(ctx ldctx.Context, n int) lderr.Error {
	now := time.Now()

	if err := l.refresh(ctx, now); err != nil {
		return err
	}

	if err := l.limiter.WaitN(ctx, n); err != nil {
		ctx.LogE("[distributed limiter]")
	}

	return nil
}

func (l *Limiter) Reserve(ctx ldctx.Context) (*Reservation, lderr.Error) {
	return l.ReserveN(ctx, 1)
}

func (l *Limiter) ReserveN(ctx ldctx.Context, n int) (*Reservation, lderr.Error) {
	now := time.Now()
	if err := l.refresh(ctx, now); err != nil {
		return nil, err
	}

	reservation := l.limiter.ReserveN(now, n)
	r := &Reservation{
		reservations: []*rate.Reservation{reservation},
	}
	return r, nil
}
