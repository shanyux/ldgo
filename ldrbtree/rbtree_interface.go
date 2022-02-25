/*
 * Copyright (C) distroy
 */

package ldrbtree

type rbtreeInterface interface {
	Tree() *RBTree
	Root() *rbtreeNode
	Sentinel() *rbtreeNode

	Compare(a, b interface{}) int
	Left(n *rbtreeNode) *rbtreeNode
	Right(n *rbtreeNode) *rbtreeNode
	Reverse() rbtreeInterface
}

func forward(tree *RBTree) rbtreeInterface { return rbtreeForward{tree: tree} }
func reverse(tree *RBTree) rbtreeInterface { return rbtreeReverse{tree: tree} }

type rbtreeForward struct {
	tree *RBTree
}

func (p rbtreeForward) Tree() *RBTree                   { return p.tree }
func (p rbtreeForward) Root() *rbtreeNode               { return p.tree.root }
func (p rbtreeForward) Sentinel() *rbtreeNode           { return p.tree.sentinel }
func (p rbtreeForward) Compare(a, b interface{}) int    { return p.tree.Compare(a, b) }
func (p rbtreeForward) Left(n *rbtreeNode) *rbtreeNode  { return n.Left }
func (p rbtreeForward) Right(n *rbtreeNode) *rbtreeNode { return n.Right }
func (p rbtreeForward) Reverse() rbtreeInterface        { return reverse(p.tree) }

type rbtreeReverse struct {
	tree *RBTree
}

func (p rbtreeReverse) Tree() *RBTree                   { return p.tree }
func (p rbtreeReverse) Root() *rbtreeNode               { return p.tree.root }
func (p rbtreeReverse) Sentinel() *rbtreeNode           { return p.tree.sentinel }
func (p rbtreeReverse) Compare(a, b interface{}) int    { return p.tree.Compare(b, a) }
func (p rbtreeReverse) Left(n *rbtreeNode) *rbtreeNode  { return n.Right }
func (p rbtreeReverse) Right(n *rbtreeNode) *rbtreeNode { return n.Left }
func (p rbtreeReverse) Reverse() rbtreeInterface        { return forward(p.tree) }
