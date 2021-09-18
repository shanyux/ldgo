/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"sync/atomic"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldlogger"
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

type Mutex struct {
	redis         *Redis
	events        chan MutexEvent
	ctx           ldcontext.Context
	key           string
	token         string
	interval      time.Duration
	timeout       time.Duration
	lastHeartbeat time.Time
	locked        int32
}

func NewMutex(redis *Redis) *Mutex {
	m := &Mutex{
		interval: 5 * time.Second,
		timeout:  2 * time.Minute,
	}
	m = m.SetRedis(redis)
	return m
}

func (m *Mutex) clone() *Mutex {
	c := *m
	return &c
}

func (m *Mutex) SetRedis(redis *Redis) *Mutex {
	m = m.clone()

	redis = redis.WithLogger(ldlogger.Discard())
	redis = redis.WithRetry(1)

	m.redis = redis
	return m
}
func (m *Mutex) SetInterval(d time.Duration) *Mutex {
	m = m.clone()

	if d < _MUTEX_MIN_INTERVAL {
		d = _MUTEX_MIN_INTERVAL
	}
	m.interval = d

	if timeout := m.getMinTimeout(d); m.timeout < timeout {
		m.timeout = timeout
	}
	return m
}

func (m *Mutex) SetTimeout(d time.Duration) *Mutex {
	m = m.clone()
	m.timeout = m.getMinTimeout(d)
	return m
}

func (m *Mutex) getMinTimeout(d time.Duration) time.Duration {
	if d < _MUTEX_MIN_TIMEOUT {
		d = _MUTEX_MIN_TIMEOUT
	}

	if t := m.interval * 3; d < t {
		d = t
	}

	return d
}

func (m *Mutex) getExpiration() time.Duration {
	return m.interval + m.timeout
}

func (m *Mutex) Key() string               { return m.key }
func (m *Mutex) Events() <-chan MutexEvent { return m.events }

func (m *Mutex) Lock(ctx ldcontext.Context, key string) error {
	ctx = ldcontext.WithLogger(ctx, ldcontext.GetLogger(ctx), zap.String("key", key))
	ctx = ldcontext.WithCancel(ctx)

	if atomic.LoadInt32(&m.locked) != 0 {
		ctx.LogE("redis mutex had locked", zap.String("old", m.key))
		return ErrMutexLocked
	}

	cli := m.redis
	val := uuid.New().String()

	cmd := cli.SetNX(key, val, m.getExpiration())
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex setnx fail", zap.Error(err))
		return err
	}

	if ok := cmd.Val(); !ok {
		ctx.LogW("redis mutex is locking")
		return ErrMutexLocking
	}

	if ok := atomic.CompareAndSwapInt32(&m.locked, 0, 1); !ok {
		cli.Del(key)
		return ErrMutexLocked
	}

	m.ctx = ctx
	m.key = key
	m.token = val
	m.events = make(chan MutexEvent, 1)

	go m.goroutine(ctx, key, val)

	ctx.LogD("redis mutex lock succ")
	return nil
}

func (m *Mutex) Unlock(ctx ldcontext.Context) error {
	ctx = ldcontext.WithLogger(ctx, ldcontext.GetLogger(ctx), zap.String("key", m.key))

	locked := atomic.LoadInt32(&m.locked)
	if locked == 0 {
		ctx.LogW("redis mutex has not been locked")
		return nil
	}

	cli := m.redis
	key := m.key
	val := m.token

	if ok := atomic.CompareAndSwapInt32(&m.locked, locked, 0); !ok {
		ctx.LogW("redis mutex has been unlocked by another")
		return nil
	}

	ctx.LogD("redis mutex will be unlock")

	ldcontext.TryCancel(ctx)
	if err := m.checkToken(ctx, key, val); err != nil {
		return err
	}

	cmd := cli.Del(key)
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex unlock fail", zap.Error(err))
		return err
	}

	return nil
}

func (m *Mutex) goroutine(ctx ldcontext.Context, key, val string) {
	ctx.LogD("redis mutex goroutine start")
	ticker := time.NewTicker(m.interval)

	defer func() {
		ctx.LogD("redis mutex goroutine stop")

		ldcontext.TryCancel(ctx)
		atomic.StoreInt32(&m.locked, 0)

		close(m.events)
		m.events = nil

		ticker.Stop()
	}()

	m.lastHeartbeat = time.Now()
	for {
		select {
		case <-ctx.Done():
			return

		case now := <-ticker.C:
			m.heartbeat(ctx, now, key, val)
		}
	}
}

func (m *Mutex) heartbeat(ctx ldcontext.Context, now time.Time, key, val string) {
	switch err := m.checkToken(ctx, key, val); err {
	case nil:
		m.lastHeartbeat = now
	case ErrMutexNotExists, ErrMutexNotMatch:
		m.doHeartbeatError(ctx)
	default:
		m.checkHeartbeatTime(ctx)
	}
}

func (m *Mutex) checkToken(ctx ldcontext.Context, key, val string) error {
	cli := m.redis
	{
		cmd := cli.Expire(key, m.getExpiration())
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

func (m *Mutex) checkHeartbeatTime(ctx ldcontext.Context) {
	if time.Since(m.lastHeartbeat) < m.timeout {
		return
	}

	ctx.LogW("redis mutex fail too much")
	m.doHeartbeatError(ctx)
}

func (m *Mutex) doHeartbeatError(ctx ldcontext.Context) {
	select {
	case <-ctx.Done():
		break
	case m.events <- MutexEvent_Deleted:
		break
	default:
		break
	}
}
