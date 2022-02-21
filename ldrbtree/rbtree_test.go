/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"testing"

	"github.com/distroy/ldgo/ldrand"
	"github.com/distroy/ldgo/ldsort"
	"github.com/smartystreets/goconvey/convey"
)

func TestRBTree(t *testing.T) {
	const count = 100

	numsUnordered := ldrand.Perm(count)
	nums := make([]int, 0, len(numsUnordered))
	nums = append(nums, numsUnordered...)
	ldsort.SortInts(nums)

	t.Logf("unordered: %v", numsUnordered)
	t.Logf("nums: %v", nums)

	newRBTree := func() *RBTree {
		rbtree := &RBTree{}
		for _, n := range numsUnordered {
			rbtree.Insert(n)
		}
		return rbtree
	}

	convey.Convey(t.Name(), t, func() {

		convey.Convey("insert", func() {
			rbtree := &RBTree{}
			for _, n := range numsUnordered {
				rbtree.Insert(n)
				convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
				convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
			}
		})

		convey.Convey("len", func() {
			rbtree := newRBTree()
			convey.So(rbtree.Len(), convey.ShouldEqual, len(nums))
		})

		convey.Convey("check parent", func() {
			rbtree := newRBTree()
			convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
		})
		convey.Convey("check color", func() {
			rbtree := newRBTree()
			convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
		})

		convey.Convey("iterator", func() {
			convey.Convey("next", func() {
				rbtree := newRBTree()
				it := rbtree.Begin()
				for _, n := range nums {
					convey.So(n, convey.ShouldEqual, it.Data())
					it = it.Next()
				}
				convey.So(it, convey.ShouldResemble, rbtree.End())
				convey.So(func() { it.Next() }, convey.ShouldPanic)
			})

			convey.Convey("prev", func() {
				rbtree := newRBTree()
				i := len(nums)
				for begin, it := rbtree.Begin(), rbtree.End(); it != begin; {
					it = it.Prev()
					i--
					n := nums[i]
					convey.So(n, convey.ShouldEqual, it.Data())
				}
			})

			convey.Convey("delete ordered", func() {
				rbtree := newRBTree()
				for it, end := rbtree.Begin(), rbtree.End(); it != end; {
					it = rbtree.Delete(it)
					convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
					convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
				}

				convey.So(rbtree.Begin(), convey.ShouldResemble, rbtree.End())
				convey.So(rbtree.Len(), convey.ShouldEqual, 0)
			})

			convey.Convey("delete unordered", func() {
				rbtree := newRBTree()
				for _, n := range numsUnordered {
					it := rbtree.Search(n)
					convey.So(it.Data(), convey.ShouldResemble, n)

					rbtree.Delete(it)
					convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
					convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)

					// it = rbtree.Search(n)
					// convey.So(it, convey.ShouldResemble, rbtree.End())
				}

				convey.So(rbtree.Begin(), convey.ShouldResemble, rbtree.End())
				convey.So(rbtree.Len(), convey.ShouldEqual, 0)
			})
		})

		convey.Convey("reveiterator", func() {
			convey.Convey("reverse iterator next", func() {
				rbtree := newRBTree()
				it := rbtree.RBegin()
				for i := len(nums) - 1; i >= 0; i-- {
					n := nums[i]
					convey.So(n, convey.ShouldEqual, it.Data())
					it = it.Next()
				}
				convey.So(it, convey.ShouldResemble, rbtree.REnd())
				convey.So(func() { it.Next() }, convey.ShouldPanic)
			})

			convey.Convey("prev", func() {
				rbtree := newRBTree()
				i := 0
				for begin, it := rbtree.RBegin(), rbtree.REnd(); it != begin; {
					it = it.Prev()
					n := nums[i]
					i++
					convey.So(n, convey.ShouldEqual, it.Data())
				}
			})

			convey.Convey("delete ordered", func() {
				rbtree := newRBTree()
				for it, end := rbtree.RBegin(), rbtree.REnd(); it != end; {
					it = rbtree.RDelete(it)
					convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
					convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)
				}

				convey.So(rbtree.Begin(), convey.ShouldResemble, rbtree.End())
				convey.So(rbtree.Len(), convey.ShouldEqual, 0)
			})

			convey.Convey("delete unordered", func() {
				rbtree := newRBTree()
				for _, n := range numsUnordered {
					it := rbtree.RSearch(n)
					convey.So(it.Data(), convey.ShouldResemble, n)

					rbtree.RDelete(it)
					convey.So(rbtree.root.checkParent(rbtree.sentinel), convey.ShouldBeTrue)
					convey.So(rbtree.root.checkColor(rbtree.sentinel), convey.ShouldBeTrue)

					// it = rbtree.Search(n)
					// convey.So(it, convey.ShouldResemble, rbtree.End())
				}

				convey.So(rbtree.Begin(), convey.ShouldResemble, rbtree.End())
				convey.So(rbtree.Len(), convey.ShouldEqual, 0)
			})
		})

		convey.Convey("duplicate data", func() {
			const retry = 20
			convey.Convey("search", func() {
				for i := 0; i < retry; i++ {

					rbtree := newRBTree()

					d := ldrand.Intn(count)
					if it := rbtree.Search(d); it != rbtree.End() {
						rbtree.Delete(it)
					}

					rbtree.Insert(d)
					rbtree.Insert(d)
					rbtree.Insert(d)

					it := rbtree.Search(d)
					// ldlog.Default().Info("", zap.Int("data", d))
					// ldlog.Default().Info("", zap.Reflect("node", it.node.ToMap(it.tree.sentinel)))
					// ldlog.Default().Info("", zap.Reflect("node", it.Next().node.ToMap(it.tree.sentinel)))
					// ldlog.Default().Info("", zap.Reflect("node", it.Next().Next().node.ToMap(it.tree.sentinel)))

					convey.So(it.Data(), convey.ShouldEqual, d)
					convey.So(it.Next().Data(), convey.ShouldEqual, d)
					convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
				}
			})

			convey.Convey("rsearch", func() {
				for i := 0; i < retry; i++ {
					rbtree := newRBTree()

					d := ldrand.Intn(count)
					if rit := rbtree.RSearch(d); rit != rbtree.REnd() {
						rbtree.RDelete(rit)
					}

					rbtree.Insert(d)
					rbtree.Insert(d)
					rbtree.Insert(d)

					it := rbtree.Search(d)
					// ldlog.Default().Info("", zap.Int("data", d))
					// ldlog.Default().Info("", zap.Reflect("node", it.node.ToMap(it.tree.sentinel)))
					// ldlog.Default().Info("", zap.Reflect("node", it.Next().node.ToMap(it.tree.sentinel)))
					// ldlog.Default().Info("", zap.Reflect("node", it.Next().Next().node.ToMap(it.tree.sentinel)))

					convey.So(it.Data(), convey.ShouldEqual, d)
					convey.So(it.Next().Data(), convey.ShouldEqual, d)
					convey.So(it.Next().Next().Data(), convey.ShouldEqual, d)
				}
			})
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
			// ldlog.Default().Info("the children of red node is not black", zap.Reflect("node", n.ToMap(sentinel)))
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
