/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"encoding/hex"
	"sync/atomic"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldlogger"
	"github.com/distroy/ldgo/ldrand"
	"go.uber.org/zap"
)

type MutexEvent int

const (
	MutexEvent_Deleted MutexEvent = iota + 1
)

const (
	mutexMinInterval = time.Second
	mutexMinTimeout  = 10 * time.Second
)

type Mutex struct {
	redis     *Redis
	interval  time.Duration
	timeout   time.Duration
	lockForce bool

	key           string
	token         string
	lastHeartbeat time.Time
	lockTime      int64 // if equal 0, has not locked
	events        chan MutexEvent
}

func NewMutex(redis *Redis) *Mutex {
	m := &Mutex{
		redis:    redis,
		interval: 5 * time.Second,
		timeout:  2 * time.Minute,
	}
	return m
}

func (m *Mutex) clone() *Mutex {
	c := *m
	return &c
}

func (m *Mutex) WithRedis(redis *Redis) *Mutex {
	m = m.clone()
	m.redis = redis
	return m
}

func (m *Mutex) WithContext(ctx Context) *Mutex {
	m = m.clone()
	m.redis = m.redis.WithContext(ctx)
	return m
}

// func (m *Mutex) WithLogger(l ldlogger.Logger) *Mutex {
// 	m = m.clone()
// 	m.redis = m.redis.WithLogger(l)
// 	return m
// }

func (m *Mutex) WithLockForce(force bool) *Mutex {
	m = m.clone()
	m.lockForce = force
	return m
}

func (m *Mutex) WithInterval(d time.Duration) *Mutex {
	m = m.clone()

	if d < mutexMinInterval {
		d = mutexMinInterval
	}
	m.interval = d

	if timeout := m.getMinTimeout(d); m.timeout < timeout {
		m.timeout = timeout
	}
	return m
}

func (m *Mutex) WithTimeout(d time.Duration) *Mutex {
	m = m.clone()
	m.timeout = m.getMinTimeout(d)
	return m
}

func (m *Mutex) getMinTimeout(d time.Duration) time.Duration {
	if d < mutexMinTimeout {
		d = mutexMinTimeout
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

func (m *Mutex) Lock(key string) error {
	ctx := m.redis.context()
	ctx = ldcontext.WithCancel(ctx)

	if atomic.LoadInt64(&m.lockTime) != 0 {
		ctx.LogE("redis mutex has been locked", zap.String("key", key), zap.String("old", m.key),
			getCaller(m.redis.caller))
		return lderr.ErrCacheMutexLocked
	}

	token := hex.EncodeToString(ldrand.Bytes(16))

	log := m.redis.logger()
	log = ldlogger.With(log, zap.String("key", key), zap.String("token", token))
	ctx = ldcontext.WithLogger(ctx, log)

	if err := m.internalLock(ctx, key, token); err != nil {
		if !m.lockForce {
			return err
		}

		for err != nil {
			time.Sleep(m.interval)
			err = m.internalLock(ctx, key, token)
		}
	}

	now := time.Now().UnixNano()
	if ok := atomic.CompareAndSwapInt64(&m.lockTime, 0, now); !ok {
		// cli := m.redis.Client()
		// cli.Del(key)
		return lderr.ErrCacheMutexLocked
	}

	m.key = key
	m.token = token
	m.events = make(chan MutexEvent, 1)

	go m.goroutine(ctx, key, token, now)

	ctx.LogD("redis mutex lock succ", getCaller(m.redis.caller))
	return nil
}

func (m *Mutex) internalLock(ctx Context, key, token string) error {
	cli := m.redis.Client()

	cmd := cli.SetNX(key, token, m.getExpiration())
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex setnx fail", zap.Error(err), getCaller(m.redis.caller))
		return err
	}

	if ok := cmd.Val(); !ok {
		ctx.LogW("redis mutex has been locked by another goroutine/process", getCaller(m.redis.caller))
		return lderr.ErrCacheMutexLocking
	}

	return nil
}

func (m *Mutex) Unlock() error {
	ctx := m.redis.context()

	lockTime := atomic.LoadInt64(&m.lockTime)
	if lockTime == 0 {
		ctx.LogW("redis mutex has not been locked", getCaller(m.redis.caller))
		return nil
	}

	cli := m.redis.Client()
	key := m.key
	val := m.token

	log := m.redis.logger()
	log = ldlogger.With(log, zap.String("key", key), zap.String("token", val))
	ctx = ldcontext.WithLogger(ctx, log)

	if ok := atomic.CompareAndSwapInt64(&m.lockTime, lockTime, 0); !ok {
		ctx.LogW("redis mutex has been unlocked by another goroutine", getCaller(m.redis.caller))
		return nil
	}

	ctx.LogD("redis mutex will be unlocked", getCaller(m.redis.caller))

	ldcontext.TryCancel(ctx)
	if err := m.checkToken(ctx, key, val); err != nil {
		return err
	}

	cmd := cli.Del(key)
	if err := cmd.Err(); err != nil {
		ctx.LogW("redis mutex del fail", zap.Error(err), getCaller(m.redis.caller))
		return err
	}

	return nil
}

func (m *Mutex) goroutine(ctx ldcontext.Context, key, val string, lockTime int64) {
	ctx.LogD("redis mutex goroutine start")
	ticker := time.NewTicker(m.interval)

	defer func() {
		ctx.LogD("redis mutex goroutine stop")

		ldcontext.TryCancel(ctx)
		atomic.CompareAndSwapInt64(&m.lockTime, lockTime, 0)

		close(m.events)
		m.events = nil

		ticker.Stop()
	}()

	m.lastHeartbeat = time.Now()
	for running := true; running; {
		select {
		case <-ctx.Done():
			return

		case now := <-ticker.C:
			running = m.heartbeat(ctx, now, key, val)
		}
	}
}

func (m *Mutex) heartbeat(ctx ldcontext.Context, now time.Time, key, val string) bool {
	switch err := m.checkToken(ctx, key, val); err {
	case nil:
		m.lastHeartbeat = now

	case lderr.ErrCacheMutexNotExists, lderr.ErrCacheMutexNotMatch:
		m.doHeartbeatError(ctx)
		return false

	default:
		return m.checkHeartbeatTime(ctx)
	}

	return true
}

func (m *Mutex) checkToken(ctx ldcontext.Context, key, val string) error {
	cli := m.redis.Client()
	{
		cmd := cli.Expire(key, m.getExpiration())
		if err := cmd.Err(); err != nil {
			ctx.LogE("redis mutex expire fail", zap.Error(err))
			return err
		}

		if ok := cmd.Val(); !ok {
			ctx.LogE("redis mutex is not exists")
			return lderr.ErrCacheMutexNotExists
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
			return lderr.ErrCacheMutexNotMatch
		}
	}

	return nil
}

func (m *Mutex) checkHeartbeatTime(ctx ldcontext.Context) bool {
	if time.Since(m.lastHeartbeat) < m.timeout {
		return true
	}

	ctx.LogW("redis mutex fail too much")
	m.doHeartbeatError(ctx)
	return false
}

func (m *Mutex) doHeartbeatError(ctx ldcontext.Context) {
	select {
	case <-ctx.Done():
	case m.events <- MutexEvent_Deleted:
	}
}
