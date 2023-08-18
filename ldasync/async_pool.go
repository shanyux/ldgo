/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"log"
	"runtime"
	"sync"
)

type AsyncPool struct {
	wg      sync.WaitGroup
	started sync.Once
	closed  sync.Once
	ch      chan func()
}

func NewAsyncPool(concurrency int) *AsyncPool {
	p := &AsyncPool{}
	p.Start(concurrency)
	return p
}

func (p *AsyncPool) Start(concurrency int) {
	if concurrency <= 0 {
		return
	}

	p.started.Do(func() { p.start(concurrency) })
}

func (p *AsyncPool) Wait()  { p.wg.Wait() }
func (p *AsyncPool) Close() { p.closed.Do(p.close) }

func (p *AsyncPool) Async() chan<- func() { return p.ch }

func (p *AsyncPool) start(concurrency int) {
	if p.ch == nil {
		p.ch = make(chan func(), 1)
	}

	p.wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go p.main()
	}
}

func (p *AsyncPool) main() {
	defer p.wg.Done()

	for fn := range p.ch {
		p.doWithRecover(fn)
	}
}

func (p *AsyncPool) close() {
	close(p.ch)
}

func (p *AsyncPool) doWithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			const size = 4 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			// log.Println(err, ldconv.BytesToStrUnsafe(buf))
			log.Printf("[async pool] do async func panic. err:%v, stack:\n%s", err, buf)
		}
	}()

	fn()
}
