/*
 * Copyright (C) distroy
 */

package ldchan

import (
	"fmt"
)

var (
	errMovePrev    = fmt.Errorf("the chan iterator can not move prev")
	errEndMoveNext = fmt.Errorf("the chan iterator is already at the end, can not move next")
)

// Begin will read the first item in chan, can not be called multiple times
func Begin[T comparable](ch <-chan T) Iterator[T] {
	if ch == nil {
		return End(ch)
	}
	it := Iterator[T]{ch: ch}
	return it.Next()
}

func End[T comparable](ch <-chan T) Iterator[T] { return Iterator[T]{} }

type Iterator[T comparable] struct {
	ch   <-chan T
	data T
}

func (it Iterator[T]) Get() T            { return it.data }
func (it Iterator[T]) Prev() Iterator[T] { panic(errMovePrev) }
func (it Iterator[T]) Next() Iterator[T] {
	if it.ch == nil {
		panic(errEndMoveNext)
	}
	v, ok := <-it.ch
	if !ok {
		return End(it.ch)
	}
	return Iterator[T]{
		ch:   it.ch,
		data: v,
	}
}

func MakeRange[T comparable](ch <-chan T) *Range[T] {
	return &Range[T]{
		ch: ch,
	}
}

type Range[T comparable] struct {
	ch     <-chan T
	begin  Iterator[T]
	end    Iterator[T]
	inited bool
}

func (r *Range[T]) init() {
	if r.inited {
		return
	}
	r.begin = Begin(r.ch)
	r.end = End(r.ch)
	r.inited = true
}

func (r *Range[T]) Get() T {
	r.init()
	return r.begin.Get()
}
func (r *Range[T]) HasNext() bool {
	r.init()
	return r.begin.ch != nil
}
func (r *Range[T]) Next() {
	r.init()
	r.begin = r.begin.Next()
}
