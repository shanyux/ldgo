/*
 * Copyright (C) distroy
 */

package ldrbtree

type rbtreeInterface interface {
	Compare(a, b interface{}) int
	Left(n *rbtreeNode) *rbtreeNode
	Right(n *rbtreeNode) *rbtreeNode
}

func forward(tree *RBTree) rbtreeInterface { return rbtreeForward{tree: tree} }
func reverse(tree *RBTree) rbtreeInterface { return rbtreeReverse{tree: tree} }

type rbtreeForward struct {
	tree *RBTree
}

func (p rbtreeForward) Compare(a, b interface{}) int    { return p.tree.Compare(a, b) }
func (p rbtreeForward) Left(n *rbtreeNode) *rbtreeNode  { return n.Left }
func (p rbtreeForward) Right(n *rbtreeNode) *rbtreeNode { return n.Right }

type rbtreeReverse struct {
	tree *RBTree
}

func (p rbtreeReverse) Compare(a, b interface{}) int    { return p.tree.Compare(b, a) }
func (p rbtreeReverse) Left(n *rbtreeNode) *rbtreeNode  { return n.Right }
func (p rbtreeReverse) Right(n *rbtreeNode) *rbtreeNode { return n.Left }
