/*
 * Copyright (C) distroy
 */

package ldasync

import (
	"log"
	"runtime"
	"sync"
)

func Go(fn func()) *GoPool { return GoN(1, fn) }
func GoN(n int, fn func()) *GoPool {
	p := &GoPool{}
	p.GoN(n, fn)
	return p
}

type GoPool struct {
	wg sync.WaitGroup
}

func (p *GoPool) Wait() { p.wg.Wait() }

func (p *GoPool) Go(fn func()) { p.GoN(1, fn) }
func (p *GoPool) GoN(n int, fn func()) {
	if n <= 0 {
		return
	}

	fnGo := func() {
		defer func() {
			p.wg.Done()

			if err := recover(); err != nil {
				const size = 4 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]

				// log.Println(err, ldconv.BytesToStrUnsafe(buf))
				log.Printf("[go pool] go func panic. err:%v, stack:\n%s", err, buf)
			}
		}()

		fn()
	}

	p.wg.Add(n)
	for i := 0; i < n; i++ {
		go fnGo()
	}
}
