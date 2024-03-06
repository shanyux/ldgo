/*
 * Copyright (C) distroy
 */

package ldchan

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testRangeIface[Data any] interface {
	Get() Data
	HasNext() bool
	Next()
}

func testRangeToSlice[T any](r testRangeIface[T]) []T {
	l := make([]T, 0, 16)
	for ; r.HasNext(); r.Next() {
		l = append(l, r.Get())
	}
	return l
}

func testSliceToChan[T any](s []T) chan T {
	ch := make(chan T, len(s))
	for _, v := range s {
		ch <- v
	}
	close(ch)
	return ch
}

func TestIterator(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("nil", func(c convey.C) {
			var ch chan int
			b := Begin(ch)
			e := End(ch)

			c.So(b, convey.ShouldResemble, e)
		})

		c.Convey("not nil", func(c convey.C) {
			ch := testSliceToChan([]int{1})
			b := Begin(ch)
			e := End(ch)

			c.So(b, convey.ShouldHaveSameTypeAs, e)
			c.So(b, convey.ShouldNotResemble, e)
			c.So(b.Get(), convey.ShouldEqual, 1)
			c.So(func() { b.Prev() }, convey.ShouldPanicWith, errMovePrev)

			b = b.Next()
			c.So(b, convey.ShouldResemble, e)
			c.So(func() { b.Next() }, convey.ShouldPanicWith, errEndMoveNext)
		})
	})
}

func TestRange(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		l := []int{1, 100, 50, 200, 70, 80, 30, 40, 20, 90, 150}
		ch := testSliceToChan(l)

		r := testRangeToSlice[int](&Range[int]{
			Begin: Begin(ch),
			End:   End(ch),
		})
		c.So(r, convey.ShouldResemble, []int{
			1, 100, 50, 200, 70, 80, 30, 40, 20, 90, 150,
		})
	})
}
