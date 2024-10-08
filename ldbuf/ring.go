/*
 * Copyright (C) distroy
 */

package ldbuf

func NewRing[T any](n int) *Ring[T] {
	return &Ring[T]{
		data:     make([]T, n),
		capacity: n,
	}
}

type Ring[T any] struct {
	data     []T
	capacity int
	start    int
	size     int
}

func (b *Ring[T]) Cap() int  { return b.capacity }
func (b *Ring[T]) Size() int { return b.size }

func (b *Ring[T]) Put(d T) { b.put(d) }
func (b *Ring[T]) Pop() T  { r, _ := b.pop(); return r }

func (b *Ring[T]) pop() (T, bool) {
	if b.size == 0 {
		var x T
		return x, false
	}
	pos := b.start
	b.start++
	if b.start >= b.capacity {
		b.start -= b.capacity
	}
	b.size--
	return b.data[pos], true
}

func (b *Ring[T]) put(d T) bool {
	if b.size >= b.capacity {
		return false
	}
	pos := b.start + b.size
	if pos >= b.capacity {
		pos -= b.capacity
	}
	b.data[pos] = d
	b.size++
	return true
}
