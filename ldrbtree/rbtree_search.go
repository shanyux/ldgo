/*
 * Copyright (C) distroy
 */

package ldrbtree

// lowerBound returns first node >= d
func (rbt *RBTree) lowerBound(d interface{}, iface rbtreeInterface) *rbtreeNode {
	sentinel := rbt.sentinel

	node := rbt.root
	res := sentinel
	for node != sentinel {
		r := iface.Compare(node.Data, d)
		if r == 0 {
			// node == d
			res = node
			node = iface.Left(node)

		} else if r > 0 {
			// node > d
			res = node
			node = iface.Left(node)

		} else {
			// node < d
			node = iface.Right(node)
		}
	}
	return res
}

// upperBound returns first node > d
func (rbt *RBTree) upperBound(d interface{}, iface rbtreeInterface) *rbtreeNode {
	sentinel := rbt.sentinel

	node := rbt.root
	res := sentinel
	for node != sentinel {
		r := rbt.Compare(node.Data, d)
		if r == 0 {
			// node == 0
			node = iface.Right(node)

		} else if r > 0 {
			// node > d
			res = node
			node = iface.Left(node)

		} else {
			// node < d
			node = iface.Right(node)
		}
	}
	return res
}
