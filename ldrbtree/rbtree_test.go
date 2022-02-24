/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"testing"

	"github.com/distroy/ldgo/ldrand"
	"github.com/smartystreets/goconvey/convey"
)

func TestRBTree_Insert(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("insert", func() {
			rbtree := &RBTree{}
			for _, n := range _numsUnordered {
				rbtree.Insert(n)
				convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
				convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
			}

			convey.Convey("len", func() {
				convey.So(rbtree.Len(), convey.ShouldEqual, len(_nums))
			})

			convey.Convey("check parent", func() {
				convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
			})

			convey.Convey("check color", func() {
				convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
			})

			convey.Convey("clear", func() {
				rbtree.Clear()
				convey.So(rbtree.Len(), convey.ShouldEqual, 0)
				convey.So(rbtree.Begin(), convey.ShouldResemble, rbtree.End())
			})
		})

		convey.Convey("insert or search", func() {
			rbtree := testNewRBTree()
			convey.So(rbtree.Len(), convey.ShouldEqual, len(_nums))

			convey.Convey("insert when exists", func() {
				for _, n := range _numsUnordered {
					it := rbtree.InsertOrSearch(n)
					convey.So(it.Data(), convey.ShouldEqual, n)
					convey.So(rbtree.Len(), convey.ShouldEqual, len(_nums))
				}
				it := rbtree.Begin()
				for _, n := range _nums {
					convey.So(n, convey.ShouldEqual, it.Data())
					it = it.Next()
				}
				convey.So(it, convey.ShouldResemble, rbtree.End())
			})

			convey.Convey("insert when not exists", func() {
				it := rbtree.InsertOrSearch(-100)
				convey.So(it.Data(), convey.ShouldEqual, -100)
				convey.So(rbtree.Len(), convey.ShouldEqual, len(_nums)+1)
			})
		})

		convey.Convey("insert or assign", func() {
			rbtree := &RBTree{
				Compare: func(a, b interface{}) int {
					aa, bb := a.([2]int), b.([2]int)
					return compareInt(int64(aa[0]), int64(bb[0]))
				},
			}
			convey.So(rbtree.Len(), convey.ShouldEqual, 0)

			rbtree.InsertOrAssign([2]int{100, 0})
			convey.So(rbtree.Len(), convey.ShouldEqual, 1)
			convey.So(rbtree.Begin().Data(), convey.ShouldResemble, [2]int{100, 0})
			convey.So(rbtree.Begin().Next(), convey.ShouldResemble, rbtree.End())

			rbtree.InsertOrAssign([2]int{100, 2})
			convey.So(rbtree.Len(), convey.ShouldEqual, 1)
			convey.So(rbtree.Begin().Data(), convey.ShouldResemble, [2]int{100, 2})
			convey.So(rbtree.Begin().Next(), convey.ShouldResemble, rbtree.End())
		})
	})
}

func TestRBTree_DuplicateData(t *testing.T) {
	fnDeleteAll := testRBTreeDeleteAll

	convey.Convey(t.Name(), t, func() {
		const retry = 20
		convey.Convey("search", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)
				fnDeleteAll(rbtree, d)

				rbtree.Insert(d)
				rbtree.Insert(d)
				rbtree.Insert(d)

				it := rbtree.Search(d)

				convey.So(it.Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})

		convey.Convey("lower bound", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)
				fnDeleteAll(rbtree, d)

				rbtree.Insert(d)
				rbtree.Insert(d)
				rbtree.Insert(d)

				it := rbtree.LowerBound(d)

				convey.So(it.Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})

		convey.Convey("upper bound", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)

				rbtree.Insert(d - 1)
				rbtree.Insert(d)
				rbtree.Insert(d + 1)

				it := rbtree.UpperBound(d)

				convey.So(it.Data(), convey.ShouldBeGreaterThan, d)
				convey.So(it.Prev().Data(), convey.ShouldBeLessThanOrEqualTo, d)
				// convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})
	})
}

func TestRBTree_Count(t *testing.T) {
	const retry = 20
	convey.Convey(t.Name(), t, func() {
		rbtree := testNewRBTree()
		for i := 0; i < retry; i++ {
			d := ldrand.Intn(_count)

			count0 := 0
			for it := rbtree.SearchRange(d); it.HasNext(); it.Next() {
				count0++
			}

			convey.So(rbtree.Count(d), convey.ShouldEqual, count0)

			rbtree.Insert(d)
			rbtree.Insert(d)
			rbtree.Insert(d)

			convey.So(rbtree.Count(d), convey.ShouldEqual, count0+3)
		}

		convey.So(rbtree.Count(-100), convey.ShouldEqual, 0)
	})
}

func TestRBTree_DuplicateDataReverse(t *testing.T) {
	fnDeleteAll := testRBTreeRDeleteAll
	const retry = 20

	convey.Convey(t.Name(), t, func() {
		convey.Convey("search", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)
				fnDeleteAll(rbtree, d)

				rbtree.Insert(d)
				rbtree.Insert(d)
				rbtree.Insert(d)

				it := rbtree.RSearch(d)

				convey.So(it.Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})

		convey.Convey("lower bound", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)
				fnDeleteAll(rbtree, d)

				rbtree.Insert(d)
				rbtree.Insert(d)
				rbtree.Insert(d)

				it := rbtree.RLowerBound(d)

				convey.So(it.Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Data(), convey.ShouldEqual, d)
				convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})

		convey.Convey("upper bound", func() {
			for i := 0; i < retry; i++ {
				rbtree := testNewRBTree()

				d := ldrand.Intn(_count)
				rbtree.Insert(d - 1)
				rbtree.Insert(d)
				rbtree.Insert(d + 1)

				it := rbtree.RUpperBound(d)

				convey.So(it.Data(), convey.ShouldBeLessThan, d)
				convey.So(it.Prev().Data(), convey.ShouldBeGreaterThanOrEqualTo, d)
				// convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
			}
		})
	})
}

func (n *rbtreeNode) checkParent(sentinel *rbtreeNode) bool {
	if n == sentinel {
		return true
	}
	if n.Parent != sentinel {
		if n.Parent.Left != n && n.Parent.Right != n {
			// ldlog.Default().Info("check parent fail", zap.Any("node", n.Data), zap.Any("parent", n.Parent.Data), zap.Uintptr("sentinel", uintptr(unsafe.Pointer(sentinel))))
			return false
		}
	}

	if n.Left != sentinel {
		if n.Left.Parent != n {
			// ldlog.Default().Info("check parent fail", zap.Any("node", n.Data), zap.Any("left", n.Left.Data))
			return false
		} else if !n.Left.checkParent(sentinel) {
			return false
		}
	}

	if n.Right != sentinel {
		if n.Right.Parent != n {
			// ldlog.Default().Info("check parent fail", zap.Any("node", n.Data), zap.Any("right", n.Right.Data))
			return false
		} else if !n.Right.checkParent(sentinel) {
			return false
		}
	}

	return true
}

func (n *rbtreeNode) checkColor(sentinel *rbtreeNode) bool {
	blacks := n.getBlackCount(sentinel)
	return n.walkColor(sentinel, blacks, 0)
}

func (n *rbtreeNode) walkColor(sentinel *rbtreeNode, wantBlacks, currentBlacks int) bool {
	if n.Color == _colorRed {
		if n.Left == nil || n.Right == nil || n.Left.Color != _colorBlack || n.Right.Color != _colorBlack {
			// ldlog.Default().Info("the children of red node is not black", zap.Reflect("node", n.toMap(sentinel)))
			return false
		}
	}

	if n.Color == _colorBlack {
		currentBlacks++
	}

	if n.Left == nil {
		return currentBlacks == wantBlacks
	}

	if !n.Left.walkColor(sentinel, wantBlacks, currentBlacks) {
		return false
	}
	if !n.Right.walkColor(sentinel, wantBlacks, currentBlacks) {
		return false
	}

	return true
}

func (n *rbtreeNode) getBlackCount(sentinel *rbtreeNode) int {
	blacks := 0
	for node := n; node != nil; node = node.Left {
		if node.Color == _colorBlack {
			blacks++
		}
	}
	return blacks
}
