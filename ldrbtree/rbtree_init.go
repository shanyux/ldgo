/*
 * Copyright (C) distroy
 */

package ldrbtree

func (rbt *RBTree) init() {
	if rbt.sentinel == nil {
		rbt.sentinel = getRBTreeNode(nil)
	}
	if rbt.root == nil {
		rbt.root = rbt.sentinel
	}
	if rbt.Compare == nil {
		rbt.Compare = defaultCompare
	}
}

func (rbt *RBTree) beginIterator(iface rbtreeInterface) rbtreeIterator {
	root := rbt.root
	sentinel := rbt.sentinel

	return rbtreeIterator{
		tree: rbt,
		node: root.min(sentinel, iface),
	}
}

func (rbt *RBTree) endIterator(iface rbtreeInterface) rbtreeIterator {
	sentinel := rbt.sentinel

	return rbtreeIterator{
		tree: rbt,
		node: sentinel,
	}
}
