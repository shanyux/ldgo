/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"

	"github.com/distroy/ldgo/v2/ldatomic"
)

type ErrGroup struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	nodes    *asyncNode
	capacity int32
	count    int32
	ch       chan func() error
	err      ldatomic.Error
}

func (p *ErrGroup) Start(concurrency int) {
	n := concurrency
	if n <= 0 {
		n = defaultConcurrency
	}
	p.Reset(n)
}

func (p *ErrGroup) Reset(concurrency int) {
	n := concurrency
	if n <= 0 {
		return
	}
	if p.getCap() == n {
		return
	}
	p.mu.Lock()
	p.start(n)
	p.mu.Unlock()
}

func (p *ErrGroup) Capacity() int { return p.getCap() }
func (p *ErrGroup) Running() int  { return int(atomic.LoadInt32(&p.count)) }

func (p *ErrGroup) getCap() int  { return int(atomic.LoadInt32(&p.capacity)) }
func (p *ErrGroup) setCap(n int) { atomic.StoreInt32(&p.capacity, int32(n)) }

func (p *ErrGroup) init() {
	if p.getCap() > 0 {
		return
	}

	p.mu.Lock()
	if p.getCap() <= 0 {
		p.start(defaultConcurrency)
	}
	p.mu.Unlock()
}

func (p *ErrGroup) start(n int) {
	num := p.getCap()
	if num == 0 {
		p.ch = make(chan func() error, defaultChannelSize)
	}

	if n == num {
		return
	}
	if n > num {
		p.incrNodesTo(n)
		return
	}
	p.decrNodesTo(n)
}

func (p *ErrGroup) incrNodesTo(to int) {
	delta := to - p.getCap()
	p.setCap(to)

	p.wg.Add(delta)
	for i := 0; i < delta; i++ {
		node := getAsyncNode()

		node.next = p.nodes
		p.nodes = node

		go p.main(node)
	}
}

func (p *ErrGroup) decrNodesTo(to int) {
	delta := p.getCap() - to
	p.setCap(to)

	for i := 0; i < delta; i++ {
		node := p.nodes
		if node == nil {
			return
		}
		p.nodes = node.next
		node.next = nil
		close(node.ch)
	}
}

func (p *ErrGroup) Wait() error {
	p.wg.Wait()
	return p.err.Load()
}

func (p *ErrGroup) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	num := p.getCap()
	if num == 0 {
		return
	}

	p.setCap(0)
	p.nodes = nil
	close(p.ch)
}

func (p *ErrGroup) Async() chan<- func() error {
	p.init()
	return p.ch
}

func (p *ErrGroup) main(node *asyncNode) {
	defer p.wg.Done()

	for {
		select {
		case <-node.ch:
			putAsyncNode(node)
			return

		case fn, ok := <-p.ch:
			if !ok {
				putAsyncNode(node)
				return
			}

			atomic.AddInt32(&p.count, 1)
			err := p.doWithRecover(fn)
			p.err.CompareAndSwap(nil, err)
			atomic.AddInt32(&p.count, -1)
		}
	}
}

func (p *ErrGroup) doWithRecover(fn func() error) (_err error) {
	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			log.Printf("[error group] do async func panic. err:%v, stack:\n%s", err, buf)
			_err = fmt.Errorf("async func panic. err:%v", err)
		}
	}()

	return fn()
}
