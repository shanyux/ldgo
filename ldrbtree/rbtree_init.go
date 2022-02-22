/*
 * Copyright (C) distroy
 */

package ldrbtree

func (rbt *RBTree) init() {
	if rbt.sentinel == nil {
		rbt.sentinel = getRBTreeNode(nil)
		rbt.root = rbt.sentinel
	}
	if rbt.Compare == nil {
		rbt.Compare = DefaultCompare
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

func (rbt *RBTree) toMap() map[string]interface{} {
	return map[string]interface{}{
		"count": rbt.count,
		"root":  rbt.root.toMap(rbt.sentinel),
	}
}
