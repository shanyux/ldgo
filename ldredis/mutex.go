/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type MutexEvent int

const (
	MutexEvent_Deleted MutexEvent = iota + 1
)

const (
	_MUTEX_MIN_INTERVAL = time.Second
	_MUTEX_MIN_TIMEOUT  = 10 * time.Second
)

type Mutex interface {
	SetRedis(redis Redis) Mutex
	SetInterval(time.Duration) Mutex
	SetTimeout(time.Duration) Mutex

	Key() string

	Events() <-chan MutexEvent

	Lock(ctx ldcontext.Context, key string) error
	Unlock(ctx ldcontext.Context) error
}

type mutex struct {
	redis         Redis
	events        chan MutexEvent
	ctx           ldcontext.Context
	key           string
	token         string
	interval      time.Duration
	timeout       time.Duration
	lastHeartbeat time.Time
	locked        bool
}

func NewMutex(redis Redis) Mutex {
	return &mutex{
		redis:    redis,
		interval: 5 * time.Second,
		timeout:  55 * time.Minute,
	}
}

func (that *mutex) SetRedis(redis Redis) Mutex {
	that.redis = redis
	return that
}
func (that *mutex) SetInterval(d time.Duration) Mutex {
	if d < _MUTEX_MIN_INTERVAL {
		d = _MUTEX_MIN_INTERVAL
	}
	that.interval = d

	if timeout := that.getMinTimeout(d); that.timeout < timeout {
		that.timeout = timeout
	}
	return that
}

func (that *mutex) SetTimeout(d time.Duration) Mutex {
	that.timeout = that.getMinTimeout(d)
	return that
}

func (that *mutex) getMinTimeout(d time.Duration) time.Duration {
	if d < _MUTEX_MIN_TIMEOUT {
		d = _MUTEX_MIN_TIMEOUT
	}

	if t := that.interval * 3; d < t {
		d = t
	}

	return d
}

func (that *mutex) getExpiration() time.Duration {
	return that.interval + that.timeout
}

func (that *mutex) Key() string               { return that.key }
func (that *mutex) Events() <-chan MutexEvent { return that.events }

func (that *mutex) Lock(ctx ldcontext.Context, key string) error {
	ctx = ctx.With(zap.String("key", key))
	ctx = ctx.WithCancel()

	if that.locked {
		ctx.LogE("redis mutex had locked", zap.String("old", that.key))
		return ErrMutexLocked
	}

	cli := that.redis.Client()
	val := uuid.New().String()

	cmd := cli.SetNX(key, val, that.getExpiration())
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex setnx fail", zap.Error(err))
		return err
	}

	if ok := cmd.Val(); !ok {
		ctx.LogW("redis mutex is locking")
		return ErrMutexLocking
	}

	that.locked = true
	that.ctx = ctx
	that.key = key
	that.token = val
	that.events = make(chan MutexEvent, 1)

	go that.goroutine(ctx, key, val)

	ctx.LogD("redis mutex lock succ")
	return nil
}

func (that *mutex) Unlock(ctx ldcontext.Context) error {
	ctx = ctx.With(zap.String("key", that.key))
	if !that.locked {
		ctx.LogW("redis mutex has not locked")
		return nil
	}

	cli := that.redis.Client()
	key := that.key
	val := that.token

	ctx.LogD("redis mutex will be unlock")
	that.locked = false
	that.ctx.TryCancel()

	if err := that.checkToken(ctx, key, val); err != nil {
		return err
	}

	cmd := cli.Del(key)
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex unlock fail", zap.Error(err))
		return err
	}

	return nil
}

func (that *mutex) goroutine(ctx ldcontext.Context, key, val string) {
	ticker := time.NewTicker(that.interval)
	defer func() {
		ticker.Stop()
		ctx.TryCancel()
		close(that.events)
		that.events = nil
		that.locked = false
	}()

	that.lastHeartbeat = time.Now()
	for {
		select {
		case <-ctx.Done():
			return

		case now := <-ticker.C:
			that.heartbeat(ctx, now, key, val)
		}
	}
}

func (that *mutex) heartbeat(ctx ldcontext.Context, now time.Time, key, val string) {
	switch err := that.checkToken(ctx, key, val); err {
	case nil:
		that.lastHeartbeat = now
	case ErrMutexNotExists, ErrMutexNotMatch:
		that.doHeartbeatError(ctx)
	default:
		that.checkHeartbeatTime(ctx)
	}
}

func (that *mutex) checkToken(ctx ldcontext.Context, key, val string) error {
	cli := that.redis.Client()
	{
		cmd := cli.Expire(key, that.getExpiration())
		if err := cmd.Err(); err != nil {
			ctx.LogE("redis mutex expire fail", zap.Error(err))
			return err
		}

		if ok := cmd.Val(); !ok {
			ctx.LogE("redis mutex is not exists")
			return ErrMutexNotExists
		}
	}

	{
		cmd := cli.Get(key)
		if err := cmd.Err(); err != nil {
			ctx.LogE("redis mutex get fail", zap.Error(err))
			return err
		}

		if val != cmd.Val() {
			ctx.LogE("redis mutex token is not match", zap.String("old", val),
				zap.String("new", cmd.Val()))
			return ErrMutexNotMatch
		}
	}

	return nil
}

func (that *mutex) checkHeartbeatTime(ctx ldcontext.Context) {
	if time.Since(that.lastHeartbeat) < that.timeout {
		return
	}

	ctx.LogW("redis mutex fail too much")
	that.doHeartbeatError(ctx)
}

func (that *mutex) doHeartbeatError(ctx ldcontext.Context) {
	defer func() {
		that.ctx.TryCancel()
	}()
	select {
	case <-ctx.Done():
		break
	case that.events <- MutexEvent_Deleted:
		break
	default:
		break
	}
}
