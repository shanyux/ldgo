/*
 * Copyright (C) distroy
 */

package ldrbtree

type color int

const (
	colorBlack color = iota
	colorRed
)

// var (
// 	rbtreePoolMap = &sync.Map{}
// )

func initRBTreeNode[T any](n *rbtreeNode[T], sentinel *rbtreeNode[T]) {
	*n = rbtreeNode[T]{
		Parent: sentinel,
		Right:  sentinel,
		Left:   sentinel,
		Color:  colorBlack,
	}
}

// func getPool[T any]() *sync.Pool {
// 	var key T
// 	m := rbtreePoolMap
// 	i, ok := m.Load(key)
// 	if ok {
// 		return i.(*sync.Pool)
// 	}
//
// 	p := &sync.Pool{New: func() any { return &rbtreeNode[T]{} }}
// 	i, loaded := m.LoadOrStore(key, p)
// 	if loaded {
// 		p = i.(*sync.Pool)
// 	}
// 	return p
// }

func getRBTreeNode[T any](sentinel *rbtreeNode[T]) *rbtreeNode[T] {
	// n := getPool[T]().Get().(*rbtreeNode[T])
	n := &rbtreeNode[T]{}
	initRBTreeNode[T](n, sentinel)
	return n
}

func putRBTreeNode[T any](p *rbtreeNode[T]) {
	initRBTreeNode[T](p, nil)
	// getPool[T]().Put(p)
}

type rbtreeNode[T any] struct {
	Parent *rbtreeNode[T] `json:"-"`
	Left   *rbtreeNode[T] `json:"left"`
	Right  *rbtreeNode[T] `json:"right"`
	Color  color          `json:"color"`
	Data   T              `json:"data"`
}

func (n *rbtreeNode[T]) min(iface rbtreeInterface[T]) *rbtreeNode[T] {
	sentinel := iface.Sentinel()

	if n == sentinel {
		return sentinel
	}
	for iface.Left(n) != sentinel {
		n = iface.Left(n)
	}
	return n
}

func (n *rbtreeNode[T]) max(iface rbtreeInterface[T]) *rbtreeNode[T] {
	return n.min(iface.Reverse())
}

func (n *rbtreeNode[T]) next(iface rbtreeInterface[T]) *rbtreeNode[T] {
	sentinel := iface.Sentinel()
	node := n

	if node == sentinel {
		return sentinel
	}

	if iface.Right(node) != sentinel {
		return iface.Right(node).min(iface)
	}

	for node.Parent != sentinel {
		if node == iface.Left(node.Parent) {
			return node.Parent
		}

		// node == node.Parent.Right
		node = node.Parent
	}

	return sentinel
}

func (n *rbtreeNode[T]) prev(iface rbtreeInterface[T]) *rbtreeNode[T] {
	return n.next(iface.Reverse())
}

func (n *rbtreeNode[T]) toDebugMap(sentinel *rbtreeNode[T]) map[string]interface{} {
	if n == sentinel {
		return nil
	}
	color := "black"
	if n.Color == colorRed {
		color = "red"
	}
	return map[string]interface{}{
		"parent": n.Parent.Data,
		"left":   n.Left.toDebugMap(sentinel),
		"right":  n.Right.toDebugMap(sentinel),
		"color":  color,
		"data":   n.Data,
	}
}
