/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"sync"
)

type color int

const (
	_colorBlack color = iota
	_colorRed
)

var (
	_rbTreeNodeLeaf = &rbtreeNode{}
	_rbTreeNodePool = sync.Pool{New: func() interface{} { return &rbtreeNode{} }}
)

func initRBTreeNode(n *rbtreeNode, sentinel *rbtreeNode) {
	*n = rbtreeNode{
		Parent: sentinel,
		Right:  sentinel,
		Left:   sentinel,
		Color:  _colorBlack,
	}
}

func getRBTreeNode(sentinel *rbtreeNode) *rbtreeNode {
	n := _rbTreeNodePool.Get().(*rbtreeNode)
	initRBTreeNode(n, sentinel)
	return n
}

func putRBTreeNode(p *rbtreeNode) {
	initRBTreeNode(p, nil)
	_rbTreeNodePool.Put(p)
}

type rbtreeNode struct {
	Parent *rbtreeNode `json:"-"`
	Left   *rbtreeNode `json:"left"`
	Right  *rbtreeNode `json:"right"`
	Color  color       `json:"color"`
	Data   interface{} `json:"data"`
}

func (n *rbtreeNode) min(sentinel *rbtreeNode, iface rbtreeInterface) *rbtreeNode {
	if n == sentinel {
		return sentinel
	}
	for iface.Left(n) != sentinel {
		n = iface.Left(n)
	}
	return n
}

func (n *rbtreeNode) max(sentinel *rbtreeNode, iface rbtreeInterface) *rbtreeNode {
	if n == sentinel {
		return sentinel
	}
	for iface.Right(n) != sentinel {
		n = iface.Right(n)
	}
	return n
}

func (n *rbtreeNode) ToMap(sentinel *rbtreeNode) map[string]interface{} {
	if n == sentinel {
		return nil
	}
	color := "black"
	if n.Color == _colorRed {
		color = "red"
	}
	m := map[string]interface{}{
		"parent": n.Parent.Data,
		"color":  color,
		"data":   n.Data,
	}

	left := n.Left.ToMap(sentinel)
	if left != nil {
		m["left"] = left
	}

	right := n.Right.ToMap(sentinel)
	if right != nil {
		m["right"] = right
	}

	return m
}
