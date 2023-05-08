/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"encoding/hex"
	"time"
	"unsafe"

	"github.com/distroy/ldgo/ldatomic"
	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldrand"
	"go.uber.org/zap"
)

type MutexEvent int

const (
	MutexEvent_Deleted MutexEvent = iota + 1
)

var (
	closedMutextEventsChan chan MutexEvent
)

const (
	mutexMinHeartbeatInterval = 1 * time.Second
	mutexMinHeartbeatTimeout  = 10 * time.Second
	mutexMinLockForceInterval = 1 * time.Millisecond
)

type mutexContext struct {
	ctx ldctx.Context

	key           string
	token         string
	lastHeartbeat time.Time
	lockTime      ldatomic.Int64 // if equal 0, has not locked
	events        chan MutexEvent
}

type atomicMutexContext struct {
	d ldatomic.Pointer
}

func (a *atomicMutexContext) toData(p unsafe.Pointer) *mutexContext {
	return (*mutexContext)(p)
}

func (a *atomicMutexContext) Load() *mutexContext   { return a.toData(a.d.Load()) }
func (a *atomicMutexContext) Store(d *mutexContext) { a.d.Store(unsafe.Pointer(d)) }

func (a *atomicMutexContext) Swap(new *mutexContext) (old *mutexContext) {
	return a.toData(a.d.Swap(unsafe.Pointer(new)))
}

func (a *atomicMutexContext) CompareAndSwap(old, new *mutexContext) (swapped bool) {
	return a.d.CompareAndSwap(unsafe.Pointer(old), unsafe.Pointer(new))
}

type Mutex struct {
	redis             *Redis
	heartbeatInterval time.Duration
	heartbeatTimeout  time.Duration
	lockForce         bool
	lockForceInterval time.Duration
	lockForceTimeout  time.Duration
	unlockDelay       time.Duration

	// if equal nil, has not locked
	// but when mutex has been cloned. maybe the lockTime is equal 0, the ctx is not equal nil
	ctx atomicMutexContext
}

func NewMutex(redis *Redis) *Mutex {
	m := &Mutex{
		redis:             redis,
		heartbeatInterval: 5 * time.Second,
		heartbeatTimeout:  2 * time.Minute,
		lockForceInterval: 1 * time.Second,
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

// func (m *Mutex) WithLogger(l ldlog.Logger) *Mutex {
// 	m = m.clone()
// 	m.redis = m.redis.WithLogger(l)
// 	return m
// }

// WithLockForce returns the redis mutex with lock force
// but if the context is cancelled, the lock force is not available
//
// WithLockForce should be called like these:
//
//	WithLockForce(false)
//	WithLockForce(true, {interval})
//	WithLockForce(true, {interval}, {timeout})
func (m *Mutex) WithLockForce(force bool, intervalAndTimeout ...time.Duration) *Mutex {
	m = m.clone()
	m.lockForce = force

	m.lockForceInterval = 0
	m.lockForceInterval = mutexMinLockForceInterval
	if len(intervalAndTimeout) > 0 {
		m.lockForceInterval = intervalAndTimeout[0]
	}
	if m.lockForceInterval < mutexMinLockForceInterval {
		m.lockForceInterval = mutexMinLockForceInterval
	}

	m.lockForceTimeout = 0
	if len(intervalAndTimeout) > 1 {
		m.lockForceTimeout = intervalAndTimeout[1]
	}

	return m
}

// WithUnlockDelay returns the redis mutex with unlock delay
func (m *Mutex) WithUnlockDelay(delay ...time.Duration) *Mutex {
	m = m.clone()
	m.unlockDelay = 0
	if len(delay) > 0 && delay[0] > 0 {
		m.unlockDelay = delay[0]
	}
	return m
}

func (m *Mutex) WithInterval(d time.Duration) *Mutex {
	m = m.clone()

	if d < mutexMinHeartbeatInterval {
		d = mutexMinHeartbeatInterval
	}
	m.heartbeatInterval = d

	if timeout := m.getMinTimeout(d); m.heartbeatTimeout < timeout {
		m.heartbeatTimeout = timeout
	}
	return m
}

func (m *Mutex) WithTimeout(d time.Duration) *Mutex {
	m = m.clone()
	m.heartbeatTimeout = m.getMinTimeout(d)
	return m
}

func (m *Mutex) getMinTimeout(d time.Duration) time.Duration {
	if d < mutexMinHeartbeatTimeout {
		d = mutexMinHeartbeatTimeout
	}

	if t := m.heartbeatInterval * 3; d < t {
		d = t
	}

	return d
}

func (m *Mutex) mustGetContext() *mutexContext {
	mc := m.ctx.Load()
	if mc == nil {
		mc = &mutexContext{
			events: closedMutextEventsChan,
		}
	}
	return mc
}

func (m *Mutex) getExpiration() time.Duration {
	return m.heartbeatInterval + m.heartbeatTimeout
}

func (m *Mutex) Key() string               { return m.mustGetContext().key }
func (m *Mutex) Events() <-chan MutexEvent { return m.mustGetContext().events }

func (m *Mutex) Lock(key string) lderr.Error {
	ctx := m.redis.context()
	ctx = ldctx.WithCancel(ctx)

	mc := m.ctx.Load()
	if mc != nil && mc.lockTime.Load() != 0 {
		ctx.LogE("redis mutex has been locked", zap.String("key", key), zap.String("old", mc.key),
			getCaller(m.redis.caller))
		return lderr.ErrCacheMutexLocked
	}
	if mc != nil {
		m.ctx.CompareAndSwap(mc, nil)
	}

	token := hex.EncodeToString(ldrand.Bytes(16))

	log := m.redis.logger()
	log = log.With(zap.String("key", key), zap.String("token", token))
	ctx = ldctx.WithLogger(ctx, log)

	if err := m.internalLockOrLockForce(ctx, key, token); err != nil {
		return err
	}

	now := time.Now().UnixNano()
	mc = &mutexContext{
		ctx:      ctx,
		key:      key,
		token:    token,
		lockTime: ldatomic.Int64(now),
		events:   make(chan MutexEvent, 1),
	}
	if ok := m.ctx.CompareAndSwap(nil, mc); !ok {
		// cli := m.redis.Client()
		// cli.Del(key)
		return lderr.ErrCacheMutexLocked
	}

	go m.goroutine(mc, now)

	ctx.LogD("redis mutex lock succ", getCaller(m.redis.caller))
	return nil
}

func (m *Mutex) internalLockOrLockForce(ctx Context, key, token string) lderr.Error {
	err := m.internalLock(ctx, key, token)
	if err == nil {
		return nil
	}

	if !m.lockForce {
		return err
	}

	timeout := m.lockForceTimeout
	if timeout > 0 {
		ctx = ldctx.WithTimeout(ctx, timeout)
	}

	ticker := time.NewTicker(m.lockForceInterval)
	defer ticker.Stop()
	for err != nil {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return err

		case <-ticker.C:
			err = m.internalLock(ctx, key, token)
		}
	}

	return nil
}

func (m *Mutex) internalLock(ctx Context, key, token string) lderr.Error {
	cli := m.redis.Client()

	cmd := cli.SetNX(key, token, m.getExpiration())
	if err := cmd.Err(); err != nil {
		ctx.LogE("redis mutex setnx fail", zap.Error(err), getCaller(m.redis.caller))
		return lderr.Wrap(err, lderr.ErrCacheRead)
	}

	if ok := cmd.Val(); !ok {
		ctx.LogW("redis mutex has been locked by another goroutine/process", getCaller(m.redis.caller))
		return lderr.ErrCacheMutexLocked
	}

	return nil
}

func (m *Mutex) Unlock() lderr.Error {
	d := m.unlockDelay
	if d <= 0 {
		return m.unlock()
	}

	go func() {
		time.Sleep(d)
		m.unlock()
	}()

	return nil
}

func (m *Mutex) unlock() lderr.Error {
	ctx := m.redis.context()

	mc := m.ctx.Load()
	if mc == nil {
		ctx.LogW("redis mutex has not been locked", getCaller(m.redis.caller))
		return nil
	}

	lockTime := mc.lockTime.Load()
	if lockTime == 0 {
		ctx.LogW("redis mutex has not been locked", getCaller(m.redis.caller))
		return nil
	}

	ctx = mc.ctx
	cli := m.redis.Client()
	key := mc.key
	val := mc.token

	log := m.redis.logger()
	log = log.With(zap.String("key", key), zap.String("token", val))
	ctx = ldctx.WithLogger(ctx, log)

	if ok := mc.lockTime.CompareAndSwap(lockTime, 0); !ok {
		ctx.LogW("redis mutex has been unlocked by another goroutine", getCaller(m.redis.caller))
		return nil
	}
	m.ctx.CompareAndSwap(mc, nil)

	ctx.LogD("redis mutex will be unlocked", getCaller(m.redis.caller))

	ldctx.TryCancel(ctx)
	if err := m.checkToken(mc); err != nil {
		return err
	}

	cmd := cli.Del(key)
	if err := cmd.Err(); err != nil {
		ctx.LogW("redis mutex del fail", zap.Error(err), getCaller(m.redis.caller))
		return lderr.Wrap(err, lderr.ErrCacheWrite)
	}

	return nil
}

func (m *Mutex) goroutine(mc *mutexContext, lockTime int64) {
	ctx := mc.ctx
	ctx.LogD("redis mutex goroutine start")
	ticker := time.NewTicker(m.heartbeatInterval)

	defer func() {
		ctx.LogD("redis mutex goroutine stop")

		ldctx.TryCancel(ctx)
		mc.lockTime.CompareAndSwap(lockTime, 0)

		close(mc.events)
		m.ctx.CompareAndSwap(mc, nil)

		ticker.Stop()
	}()

	mc.lastHeartbeat = time.Now()
	for running := true; running; {
		select {
		case now := <-ticker.C:
			running = m.heartbeat(mc, now)

		case <-ctx.Done():
			c := ldctx.WithLogger(ldctx.Default(), ldctx.GetLogger(ctx))
			m.WithContext(c).unlock()
			return
		}
	}
}

func (m *Mutex) heartbeat(mc *mutexContext, now time.Time) bool {
	switch err := m.checkToken(mc); err {
	case nil:
		mc.lastHeartbeat = now

	case lderr.ErrCacheMutexNotExists, lderr.ErrCacheMutexNotMatch:
		m.doHeartbeatError(mc)
		return false

	default:
		return m.checkHeartbeatTime(mc)
	}

	return true
}

func (m *Mutex) checkToken(mc *mutexContext) lderr.Error {
	ctx := mc.ctx
	cli := m.redis.Client()
	key := mc.key
	val := mc.token
	{
		cmd := cli.Expire(key, m.getExpiration())
		if err := cmd.Err(); err != nil {
			ctx.LogE("redis mutex expire fail", zap.Error(err))
			return lderr.Wrap(err, lderr.ErrCacheWrite)
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
			return lderr.Wrap(err, lderr.ErrCacheRead)
		}

		if val != cmd.Val() {
			ctx.LogE("redis mutex token is not match", zap.String("old", val),
				zap.String("new", cmd.Val()))
			return lderr.ErrCacheMutexNotMatch
		}
	}

	return nil
}

func (m *Mutex) checkHeartbeatTime(mc *mutexContext) bool {
	ctx := mc.ctx
	if time.Since(mc.lastHeartbeat) < m.heartbeatTimeout {
		return true
	}

	ctx.LogW("redis mutex fail too much")
	m.doHeartbeatError(mc)
	return false
}

func (m *Mutex) doHeartbeatError(mc *mutexContext) {
	ctx := mc.ctx
	select {
	case <-ctx.Done():
	case mc.events <- MutexEvent_Deleted:
	}
}

func init() {
	closedMutextEventsChan = make(chan MutexEvent)
	close(closedMutextEventsChan)
}
