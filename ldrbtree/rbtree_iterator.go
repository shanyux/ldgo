/*
 * Copyright (C) distroy
 */

package ldrbtree

import "fmt"

type rbtreeIterator struct {
	tree *RBTree
	node *rbtreeNode
}

func (it rbtreeIterator) next(name string, iface rbtreeInterface) rbtreeIterator {
	if it.tree == nil {
		panic(fmt.Sprintf("the %s not bind any red-black tree, can not move next", name))
	}

	sentinel := it.tree.sentinel
	if it.node == sentinel {
		panic(fmt.Sprintf("the %s is already at the end of the red-black tree, can not move next", name))
	}

	node := it.nextNode(iface)
	return rbtreeIterator{
		tree: it.tree,
		node: node,
	}
}

func (it rbtreeIterator) prev(name string, iface rbtreeInterface) rbtreeIterator {
	if it.tree == nil {
		panic(fmt.Sprintf("the %s does not bind any red-black tree, can not move prev", name))
	}

	sentinel := it.tree.sentinel
	root := it.tree.root

	var node *rbtreeNode
	if it.node == sentinel {
		node = root.max(sentinel, iface)
	} else {
		node = it.prevNode(iface)
	}

	if node == sentinel {
		panic(fmt.Sprintf("the %s is already at the begin of the red-black tree, can not move prev", name))
	}

	return rbtreeIterator{
		tree: it.tree,
		node: node,
	}
}

func (it rbtreeIterator) nextNode(iface rbtreeInterface) *rbtreeNode {
	sentinel := it.tree.sentinel
	node := it.node

	if node == sentinel {
		return sentinel
	}

	if iface.Right(node) != sentinel {
		return iface.Right(node).min(sentinel, iface)
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

func (it rbtreeIterator) prevNode(iface rbtreeInterface) *rbtreeNode {
	return it.nextNode(iface.Reverse())
}

type RBTreeIterator rbtreeIterator

func (i RBTreeIterator) Data() interface{} {
	return i.node.Data
}

func (i RBTreeIterator) base() rbtreeIterator {
	return rbtreeIterator(i)
}

func (i RBTreeIterator) Next() RBTreeIterator {
	return RBTreeIterator(i.base().next("iterator", forward(i.tree)))
}

func (i RBTreeIterator) Prev() RBTreeIterator {
	return RBTreeIterator(i.base().prev("iterator", forward(i.tree)))
}

type RBTreeReverseIterator rbtreeIterator

func (i RBTreeReverseIterator) Data() interface{} {
	return i.node.Data
}

func (i RBTreeReverseIterator) base() rbtreeIterator {
	return rbtreeIterator(i)
}

func (i RBTreeReverseIterator) Next() RBTreeReverseIterator {
	return RBTreeReverseIterator(i.base().next("reverse iterator", reverse(i.tree)))
}

func (i RBTreeReverseIterator) Prev() RBTreeReverseIterator {
	return RBTreeReverseIterator(i.base().prev("reverse iterator", reverse(i.tree)))
}
