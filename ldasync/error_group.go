/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/distroy/ldgo/v2/ldatomic"
)

type ErrGroup struct {
	asyncBase[func() error]

	err ldatomic.Error
}

func (p *ErrGroup) Start(concurrency int) {
	p.asyncBase.start(concurrency, p.doWithRecover)
}

func (p *ErrGroup) Reset(concurrency int) {
	p.reset(concurrency, p.doWithRecover)
}

func (p *ErrGroup) Capacity() int { return p.asyncBase.getCap() }
func (p *ErrGroup) Running() int  { return p.asyncBase.getLen() }

func (p *ErrGroup) init() { p.asyncBase.init(p.doWithRecover) }

func (p *ErrGroup) Wait() error {
	p.asyncBase.wait()
	return p.err.Load()
}

func (p *ErrGroup) Close() { p.asyncBase.close() }

func (p *ErrGroup) Async() chan<- func() error {
	p.init()
	return p.asyncBase.async()
}

func (p *ErrGroup) doWithRecover(fn func() error) {
	defer func() {
		if err := recover(); err != nil {
			buf := debug.Stack()
			log.Printf("[error group] do async func panic. err:%v, stack:\n%s", err, buf)
			err := fmt.Errorf("async func panic. err:%v", err)
			p.err.CompareAndSwap(nil, err)
		}
	}()

	err := fn()
	if err != nil {
		p.err.CompareAndSwap(nil, err)
	}
}
