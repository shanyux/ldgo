/*
 * Copyright (C) distroy
 */

package ldrate

import (
	"context"
	"time"

	"github.com/distroy/ldgo/v2/ldatomic"
	"github.com/distroy/ldgo/v2/ldctx"
	"github.com/distroy/ldgo/v2/lderr"
	"github.com/distroy/ldgo/v2/ldmath"
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

func (l *Limiter) SetOptions(opts ...Option) {
	cfg := &l.config
	for _, fn := range opts {
		fn(cfg)
	}
}

func (l *Limiter) SetName(name string)         { l.config.Name.Store(name) }
func (l *Limiter) SetBurst(burst int64)        { l.config.Burst.Store(burst) }
func (l *Limiter) SetLimit(limit int64)        { l.config.Limit.Store(limit) }
func (l *Limiter) SetInterval(d time.Duration) { l.config.Interval.Store(d) }
func (l *Limiter) SetNodeCount(n int64)        { l.config.NodeCount.Store(n) }

func (l *Limiter) Name() string            { return l.config.Name.Load() }
func (l *Limiter) Burst() int64            { return l.config.Burst.Load() }
func (l *Limiter) Limit() int64            { return l.config.Limit.Load() }
func (l *Limiter) Interval() time.Duration { return l.config.Interval.Load() }
func (l *Limiter) NodeCount() int64        { return l.config.NodeCount.Load() }

func (l *Limiter) refresh(ctx context.Context, now time.Time) error {
	cfg := &l.config

	burst := cfg.Burst.Load()
	if burst <= 0 {
		burst = ldmath.MaxInt64(l.lastBurst.Load(), 1)
	}
	if burst != l.lastBurst.Load() {
		l.lastBurst.Store(burst)
		l.limiter.SetBurstAt(now, int(burst))
		ldctx.LogI(ctx, "[limiter] refresh the burst succ", zap.String("name", l.Name()),
			zap.Int64("burst", burst))
	}

	limit := cfg.Limit.Load()
	interval := cfg.Interval.Load()
	nodeCount := cfg.NodeCount.Load()

	if limit <= 0 {
		limit = ldmath.MaxInt64(l.lastLimit.Load(), 1)
	}
	if interval < 0 {
		interval = 0
	}
	if nodeCount <= 0 {
		nodeCount = ldmath.MaxInt64(l.lastNodeCount.Load(), 1)
	}

	// if interval < 0 || limit <= 0 || nodeCount <= 0 {
	// 	ldctx.LogE(ctx, "[limiter] invalid rate every parameters", zap.Int64("limit", limit),
	// 		zap.Stringer("interval", interval), zap.Int64("nodeCount", nodeCount))
	// 	return lderr.ErrInternalServerError
	// }

	if limit == l.lastLimit.Load() && interval == l.lastInterval.Load() && nodeCount == l.lastNodeCount.Load() {
		return nil
	}

	every := interval * time.Duration(nodeCount) / time.Duration(limit)
	// if every < 0 {
	// 	ldctx.LogE(ctx, "[limiter] invalid rate every", zap.Int64("limit", limit),
	// 		zap.Stringer("interval", interval), zap.Int64("serviceCount", nodeCount))
	// 	return lderr.ErrInternalServerError
	// }

	l.limiter.SetLimitAt(now, rate.Every(every))
	l.lastLimit.Store(limit)
	l.lastInterval.Store(interval)
	l.lastNodeCount.Store(nodeCount)

	ldctx.LogI(ctx, "[limiter] refresh rate every succ", zap.String("name", l.Name()),
		zap.Int64("limit", limit), zap.Stringer("interval", interval),
		zap.Int64("serviceCount", nodeCount), zap.Stringer("every", every))
	return nil
}

func (l *Limiter) Wait(ctx context.Context) error {
	return l.WaitN(ctx, 1)
}

func (l *Limiter) WaitN(ctx context.Context, n int) error {
	now := time.Now()

	if err := l.refresh(ctx, now); err != nil {
		return err
	}

	// if err := l.limiter.WaitN(ctx, n); err != nil {
	// 	ldctx.LogE(ctx, "[limiter] wait fail", zap.Int("n", n), zap.Error(err))
	// 	if e := ldctx.GetError(ctx); e != nil {
	// 		return e
	// 	}
	// 	return lderr.ErrCtxCanceled
	// }

	if err := wait(ctx, l, n); err != nil {
		ldctx.LogE(ctx, "[limiter] wait fail", zap.String("name", l.Name()), zap.Int("n", n), zap.Error(err))
		return err
	}
	return nil
}

func (l *Limiter) Reserve(ctx context.Context) (*Reservation, error) {
	return l.ReserveN(ctx, 1)
}

func (l *Limiter) ReserveN(ctx context.Context, n int) (*Reservation, error) {
	now := time.Now()
	if err := l.refresh(ctx, now); err != nil {
		return nil, err
	}

	reservation := l.limiter.ReserveN(now, n)
	if !reservation.OK() {
		err := lderr.ErrCtxDeadlineNotEnough
		ldctx.LogE(ctx, "[limiter] reserve fail", zap.String("name", l.Name()), zap.Int("n", n), zap.Error(err))
		return nil, err
	}

	r := &Reservation{
		reservations: []*rate.Reservation{reservation},
	}
	return r, nil
}
