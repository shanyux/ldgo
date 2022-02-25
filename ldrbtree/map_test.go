/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"testing"

	"github.com/distroy/ldgo/ldrand"
	"github.com/smartystreets/goconvey/convey"
)

func TestMap_Insert(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		m := testNewMap()

		convey.Convey("insert", func() {
			convey.Convey("len", func() {
				convey.So(m.Len(), convey.ShouldEqual, len(_nums))
			})

			convey.Convey("check parent", func() {
				convey.So(m.tree.root.checkParent(m.tree.sentinel), convey.ShouldBeTrue)
			})

			convey.Convey("check color", func() {
				convey.So(m.tree.root.checkColor(m.tree.sentinel), convey.ShouldBeTrue)
			})

			convey.Convey("clear", func() {
				m.Clear()
				convey.So(m.Len(), convey.ShouldEqual, 0)
				convey.So(m.Begin(), convey.ShouldResemble, m.End())
			})

			convey.Convey("count", func() {
				n := ldrand.Intn(_count)
				count := m.Count(n)
				m.Insert(n, 0)
				m.Insert(n, 1)
				m.Insert(n, 2)
				convey.So(m.Count(n), convey.ShouldEqual, count+3)
			})
		})

		convey.Convey("insert or search", func() {
			convey.So(m.Len(), convey.ShouldEqual, len(_nums))

			convey.Convey("insert when exists", func() {
				n := ldrand.Intn(_count)
				it := m.InsertOrSearch(n, -100)
				convey.So(it.Key(), convey.ShouldEqual, n)
				convey.So(it.Value(), convey.ShouldEqual, n)
				convey.So(m.Len(), convey.ShouldEqual, len(_nums))
			})

			convey.Convey("insert when not exists", func() {
				it := m.InsertOrSearch(-100, -100)
				convey.So(it.Key(), convey.ShouldEqual, -100)
				convey.So(it.Value(), convey.ShouldEqual, -100)
				convey.So(m.Len(), convey.ShouldEqual, len(_nums)+1)
			})
		})

		convey.Convey("insert or assign", func() {
			m := &Map{}
			convey.So(m.Len(), convey.ShouldEqual, 0)

			m.InsertOrAssign(100, 0)
			convey.So(m.Len(), convey.ShouldEqual, 1)
			convey.So(m.Begin().Key(), convey.ShouldResemble, 100)
			convey.So(m.Begin().Value(), convey.ShouldResemble, 0)
			convey.So(m.Begin().Next(), convey.ShouldResemble, m.End())

			m.InsertOrAssign(100, 2)
			convey.So(m.Len(), convey.ShouldEqual, 1)
			convey.So(m.Begin().Key(), convey.ShouldResemble, 100)
			convey.So(m.Begin().Value(), convey.ShouldResemble, 2)
			convey.So(m.Begin().Next(), convey.ShouldResemble, m.End())
		})
	})
}
func TestMap_DuplicateData(t *testing.T) {
	fnDeleteAll := testMapDeleteAll

	convey.Convey(t.Name(), t, func() {
		rbtree := testNewMap()

		convey.Convey("search", func() {
			d := ldrand.Intn(_count)
			fnDeleteAll(rbtree, d)

			rbtree.Insert(d, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d, 2)

			it := rbtree.Search(d)

			convey.So(it.Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)

			convey.So(it.Value(), convey.ShouldEqual, 0)
			convey.So(it.Next().Value(), convey.ShouldEqual, 1)
			convey.So(it.Next().Next().Value(), convey.ShouldEqual, 2)
		})

		convey.Convey("lower bound", func() {
			d := ldrand.Intn(_count)
			fnDeleteAll(rbtree, d)

			rbtree.Insert(d, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d, 2)

			it := rbtree.LowerBound(d)

			convey.So(it.Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)

			convey.So(it.Value(), convey.ShouldEqual, 0)
			convey.So(it.Next().Value(), convey.ShouldEqual, 1)
			convey.So(it.Next().Next().Value(), convey.ShouldEqual, 2)
		})

		convey.Convey("upper bound", func() {
			d := ldrand.Intn(_count)

			rbtree.Insert(d-1, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d+1, 2)

			it := rbtree.UpperBound(d)

			convey.So(it.Key(), convey.ShouldBeGreaterThan, d)
			convey.So(it.Prev().Key(), convey.ShouldBeLessThanOrEqualTo, d)
			// convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)
		})
	})
}

func TestMap_ReverseDuplicateData(t *testing.T) {
	fnDeleteAll := testMapRDeleteAll

	convey.Convey(t.Name(), t, func() {
		rbtree := testNewMap()

		convey.Convey("search", func() {
			d := ldrand.Intn(_count)
			fnDeleteAll(rbtree, d)

			rbtree.Insert(d, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d, 2)

			it := rbtree.RSearch(d)

			convey.So(it.Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)

			convey.So(it.Value(), convey.ShouldEqual, 2)
			convey.So(it.Next().Value(), convey.ShouldEqual, 1)
			convey.So(it.Next().Next().Value(), convey.ShouldEqual, 0)
		})

		convey.Convey("lower bound", func() {
			d := ldrand.Intn(_count)
			fnDeleteAll(rbtree, d)

			rbtree.Insert(d, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d, 2)

			it := rbtree.RLowerBound(d)

			convey.So(it.Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Key(), convey.ShouldEqual, d)
			convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)

			convey.So(it.Value(), convey.ShouldEqual, 2)
			convey.So(it.Next().Value(), convey.ShouldEqual, 1)
			convey.So(it.Next().Next().Value(), convey.ShouldEqual, 0)
		})

		convey.Convey("upper bound", func() {
			d := ldrand.Intn(_count)

			rbtree.Insert(d-1, 0)
			rbtree.Insert(d, 1)
			rbtree.Insert(d+1, 2)

			it := rbtree.RUpperBound(d)

			convey.So(it.Key(), convey.ShouldBeLessThan, d)
			convey.So(it.Prev().Key(), convey.ShouldBeGreaterThanOrEqualTo, d)
			// convey.So(it.Next().Next().Key(), convey.ShouldEqual, d)
		})
	})
}
