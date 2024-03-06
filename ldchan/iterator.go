/*
 * Copyright (C) distroy
 */

package ldchan

import "fmt"

var (
	errMovePrev    = fmt.Errorf("the chan iterator can not move prev")
	errEndMoveNext = fmt.Errorf("the chan iterator is already at the end, can not move next")
)

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

type Range[T comparable] struct {
	Begin Iterator[T]
	End   Iterator[T]
}

func (r *Range[T]) Get() T        { return r.Begin.Get() }
func (r *Range[T]) HasNext() bool { return r.Begin.ch != nil }
func (r *Range[T]) Next()         { r.Begin = r.Begin.Next() }

// Begin will read the first item in chan, can not be called multiple times
func Begin[T comparable](ch <-chan T) Iterator[T] {
	if ch == nil {
		return End(ch)
	}
	it := Iterator[T]{ch: ch}
	return it.Next()
}

func End[T comparable](ch <-chan T) Iterator[T] { return Iterator[T]{} }
