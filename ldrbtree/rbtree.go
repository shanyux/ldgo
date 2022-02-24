/*
 * Copyright (C) distroy
 */

package ldrbtree

type CompareFunc = func(a, b interface{}) int

// RBTree is red-black tree
type RBTree struct {
	Compare  CompareFunc
	root     *rbtreeNode
	sentinel *rbtreeNode
	count    int
}

func (rbt *RBTree) Len() int {
	return rbt.count
}

func (rbt *RBTree) Insert(d interface{}) RBTreeIterator {
	rbt.init()

	root := &rbt.root

	node := getRBTreeNode(rbt.sentinel)
	node.Data = d
	it := RBTreeIterator{
		tree: rbt,
		node: node,
	}

	if *root == rbt.sentinel {
		*root = node
		rbt.count++
		return it
	}

	rbt.insertNode(node)
	rbt.insertFixup(node)
	rbt.count++
	return it
}

func (rbt *RBTree) InsertOrSearch(d interface{}) RBTreeIterator {
	rbt.init()

	root := &rbt.root

	node := getRBTreeNode(rbt.sentinel)
	node.Data = d

	if *root == rbt.sentinel {
		*root = node
		rbt.count++
		return RBTreeIterator{
			tree: rbt,
			node: node,
		}
	}

	temp := rbt.insertOrSearchNode(node)
	if temp != node {
		// d already exists
		putRBTreeNode(node)

	} else {
		// d inserted just now
		rbt.insertFixup(node)
		rbt.count++
	}

	return RBTreeIterator{
		tree: rbt,
		node: temp,
	}
}

func (rbt *RBTree) InsertOrAssign(d interface{}) RBTreeIterator {
	rbt.init()

	root := &rbt.root

	node := getRBTreeNode(rbt.sentinel)
	node.Data = d

	if *root == rbt.sentinel {
		*root = node
		rbt.count++
		return RBTreeIterator{
			tree: rbt,
			node: node,
		}
	}

	temp := rbt.insertOrSearchNode(node)
	if temp != node {
		// d already exists
		temp.Data = d
		putRBTreeNode(node)

	} else {
		// d inserted just now
		rbt.insertFixup(node)
		rbt.count++
	}

	return RBTreeIterator{
		tree: rbt,
		node: temp,
	}
}

func (rbt *RBTree) Clear() {
	rbt.init()

	node := rbt.root
	sentinel := rbt.sentinel

	rbt.root = sentinel
	rbt.count = 0

	for node != sentinel {
		if node.Left != sentinel && node.Right != sentinel {
			node = node.Left
			continue
		}

		// node.Left == sentinel || node.Right == sentinel
		parent := node.Parent
		child := node.Left
		if node.Left == sentinel {
			child = node.Right
		}

		if node == parent.Left {
			parent.Left = child
		} else if node == parent.Right {
			parent.Right = child
		}

		if child != sentinel {
			child.Parent = parent
		}

		// ldlog.Default().Info("*** clear", zap.Any("data", node.Data))
		putRBTreeNode(node)

		if child != sentinel {
			node = child
		} else {
			node = parent
		}
	}
}

func (rbt *RBTree) Delete(it RBTreeIterator) RBTreeIterator {
	rbt.init()

	if it.tree != rbt {
		panic("the iterator does not belong to the red-black tree, can not delete")
	}

	sentinel := rbt.sentinel
	node := it.node
	if node == sentinel {
		panic("the iterator is already at the end of the red-black tree, can not delete")
	}

	it = it.Next()
	rbt.deleteNode(node)
	rbt.count--

	putRBTreeNode(node)
	return it
}

// Search returns the first node == d
func (rbt *RBTree) Search(d interface{}) RBTreeIterator {
	rbt.init()

	sentinel := rbt.sentinel

	node := rbtreeLowerBound(d, forward(rbt))
	if node != sentinel && rbt.Compare(d, node.Data) != 0 {
		node = sentinel
	}
	return RBTreeIterator{
		tree: rbt,
		node: node,
	}
}

// SearchRange returns the node range [first node == d, first node > d)
func (rbt *RBTree) SearchRange(d interface{}) *RBTreeRange {
	rbt.init()

	sentinel := rbt.sentinel

	begin := rbtreeLowerBound(d, forward(rbt))
	if begin != sentinel && rbt.Compare(d, begin.Data) != 0 {
		return &RBTreeRange{
			begin: RBTreeIterator{tree: rbt, node: sentinel},
			end:   RBTreeIterator{tree: rbt, node: sentinel},
		}
	}

	end := rbtreeUpperBound(d, forward(rbt))
	return &RBTreeRange{
		begin: RBTreeIterator{tree: rbt, node: begin},
		end:   RBTreeIterator{tree: rbt, node: end},
	}
}

// LowerBound returns the first node >= d
func (rbt *RBTree) LowerBound(d interface{}) RBTreeIterator {
	rbt.init()
	return RBTreeIterator{
		tree: rbt,
		node: rbtreeLowerBound(d, forward(rbt)),
	}
}

// UpperBound returns the first node > d
func (rbt *RBTree) UpperBound(d interface{}) RBTreeIterator {
	rbt.init()
	return RBTreeIterator{
		tree: rbt,
		node: rbtreeUpperBound(d, forward(rbt)),
	}
}

func (rbt *RBTree) Range() *RBTreeRange {
	rbt.init()

	return &RBTreeRange{
		begin: rbt.Begin(),
		end:   rbt.End(),
	}
}

func (rbt *RBTree) Begin() RBTreeIterator {
	rbt.init()
	return RBTreeIterator(rbtreeBeginIterator(forward(rbt)))
}

func (rbt *RBTree) End() RBTreeIterator {
	rbt.init()
	return RBTreeIterator(rbtreeEndIterator(forward(rbt)))
}

// RDelete is reverse delete
func (rbt *RBTree) RDelete(it RBTreeReverseIterator) RBTreeReverseIterator {
	rbt.init()

	if it.tree != rbt {
		panic("the iterator does not belong to the red-black tree, can not delete")
	}

	sentinel := rbt.sentinel
	node := it.node
	if node == sentinel {
		panic("the iterator is already at the end of the red-black tree, can not delete")
	}

	it = it.Next()
	rbt.deleteNode(node)
	rbt.count--

	putRBTreeNode(node)
	return it
}

// RSearch is reverse search
// RSearch returns the last node == d
func (rbt *RBTree) RSearch(d interface{}) RBTreeReverseIterator {
	rbt.init()

	sentinel := rbt.sentinel

	node := rbtreeLowerBound(d, reverse(rbt))
	if node != sentinel && rbt.Compare(d, node.Data) != 0 {
		node = sentinel
	}
	return RBTreeReverseIterator{
		tree: rbt,
		node: node,
	}
}

// RSearchRange is reverse search range
// RSearchRange returns the node range [last node == d, last node < d)
func (rbt *RBTree) RSearchRange(d interface{}) *RBTreeReverseRange {
	rbt.init()

	sentinel := rbt.sentinel

	begin := rbtreeLowerBound(d, reverse(rbt))
	if begin != sentinel && rbt.Compare(d, begin.Data) != 0 {
		return &RBTreeReverseRange{
			begin: RBTreeReverseIterator{tree: rbt, node: sentinel},
			end:   RBTreeReverseIterator{tree: rbt, node: sentinel},
		}
	}

	end := rbtreeUpperBound(d, reverse(rbt))
	return &RBTreeReverseRange{
		begin: RBTreeReverseIterator{tree: rbt, node: begin},
		end:   RBTreeReverseIterator{tree: rbt, node: end},
	}
}

// RLowerBound is reverse lower bound
// RLowerBound returns the last node <= d
func (rbt *RBTree) RLowerBound(d interface{}) RBTreeReverseIterator {
	rbt.init()
	return RBTreeReverseIterator{
		tree: rbt,
		node: rbtreeLowerBound(d, reverse(rbt)),
	}
}

// RUpperBound is reverse upper bound
// RUpperBound returns the last node < d
func (rbt *RBTree) RUpperBound(d interface{}) RBTreeReverseIterator {
	rbt.init()
	return RBTreeReverseIterator{
		tree: rbt,
		node: rbtreeUpperBound(d, reverse(rbt)),
	}
}

// RRange is reverse range
func (rbt *RBTree) RRange() *RBTreeReverseRange {
	rbt.init()

	return &RBTreeReverseRange{
		begin: rbt.RBegin(),
		end:   rbt.REnd(),
	}
}

// RBegin is reverse begin
func (rbt *RBTree) RBegin() RBTreeReverseIterator {
	rbt.init()
	return RBTreeReverseIterator(rbtreeBeginIterator(reverse(rbt)))
}

// REnd is reverse end
func (rbt *RBTree) REnd() RBTreeReverseIterator {
	rbt.init()
	return RBTreeReverseIterator(rbtreeEndIterator(reverse(rbt)))
}
