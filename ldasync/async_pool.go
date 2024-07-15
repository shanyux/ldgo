/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"

	"github.com/distroy/ldgo/v2/ldsync"
)

const (
	defaultConcurrency = 1
	defaultChannelSize = 1
)

type AsyncPool struct {
	wg       sync.WaitGroup
	mu       sync.Mutex
	nodes    *asyncNode
	capacity int32
	count    int32
	ch       chan func()
}

func NewAsyncPool(concurrency int) *AsyncPool {
	p := &AsyncPool{}
	p.Start(concurrency)
	return p
}

func (p *AsyncPool) Start(concurrency int) {
	n := concurrency
	if n <= 0 {
		n = defaultConcurrency
	}
	p.Reset(n)
}

func (p *AsyncPool) Reset(concurrency int) {
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

func (p *AsyncPool) Capacity() int { return p.getCap() }
func (p *AsyncPool) Running() int  { return int(atomic.LoadInt32(&p.count)) }

func (p *AsyncPool) getCap() int  { return int(atomic.LoadInt32(&p.capacity)) }
func (p *AsyncPool) setCap(n int) { atomic.StoreInt32(&p.capacity, int32(n)) }

func (p *AsyncPool) init() {
	if p.getCap() > 0 {
		return
	}

	p.mu.Lock()
	if p.getCap() <= 0 {
		p.start(defaultConcurrency)
	}
	p.mu.Unlock()
}

func (p *AsyncPool) start(n int) {
	num := p.getCap()
	if num == 0 {
		p.ch = make(chan func(), defaultChannelSize)
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

func (p *AsyncPool) incrNodesTo(to int) {
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

func (p *AsyncPool) decrNodesTo(to int) {
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

func (p *AsyncPool) Wait() { p.wg.Wait() }
func (p *AsyncPool) Close() {
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

func (p *AsyncPool) Async() chan<- func() {
	p.init()
	return p.ch
}

func (p *AsyncPool) main(node *asyncNode) {
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
			// fn()
			p.doWithRecover(fn)
			atomic.AddInt32(&p.count, -1)
		}
	}
}

func (p *AsyncPool) doWithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			log.Printf("[async pool] do async func panic. err:%v, stack:\n%s", err, buf)
		}
	}()

	fn()
}

var (
	asyncNodePool = &ldsync.Pool[*asyncNode]{}
)

type asyncNode struct {
	ch   chan struct{}
	next *asyncNode
}

func getAsyncNode() *asyncNode {
	p := asyncNodePool

	n := p.Get()
	if n == nil {
		n = &asyncNode{}
	}

	n.ch = make(chan struct{})
	n.next = nil
	return n
}

func putAsyncNode(n *asyncNode) {
	p := asyncNodePool
	if n == nil {
		return
	}
	p.Put(n)
}
