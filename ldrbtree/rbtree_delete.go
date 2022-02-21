/*
 * Copyright (C) distroy
 */

package ldrbtree

// node must not be sentinel
func (rbt *RBTree) deleteNode(node *rbtreeNode) {
	root := &rbt.root
	sentinel := rbt.sentinel

	var subst, temp *rbtreeNode
	switch {
	case node.Left == sentinel:
		temp = node.Right
		subst = node

	case node.Right == sentinel:
		temp = node.Left
		subst = node

	default:
		subst = node.Right.min(sentinel, forward(rbt))
		if subst.Left != sentinel {
			temp = subst.Left
		} else {
			temp = subst.Right
		}
	}

	// ldlog.Default().Info("find the subst succ", zap.Any("node", node.Data), zap.Any("subst", subst.Data), zap.Any("temp", temp.Data))

	if subst == *root {
		*root = temp
		temp.Color = _colorBlack
		if temp != sentinel {
			temp.Parent = sentinel
		}
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
		return
	}

	// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))

	substColor := subst.Color

	rbt.deleteNodeWithSubst(node, subst, temp)
	// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))

	if substColor != _colorRed {
		// ldlog.Default().Info("==== debug ======")
		rbt.deleteFixup(temp)
	}
}

func (rbt *RBTree) deleteNodeWithSubst(node *rbtreeNode, subst, temp *rbtreeNode) {
	root := &rbt.root
	sentinel := rbt.sentinel

	// delete subst
	if subst == subst.Parent.Left {
		subst.Parent.Left = temp
	} else {
		subst.Parent.Right = temp
	}

	if subst == node {
		temp.Parent = subst.Parent
		// ldlog.Default().Info("==== debug ======")
		return
	}

	// replace node to subst
	if subst.Parent == node {
		temp.Parent = subst
	} else {
		temp.Parent = subst.Parent
	}

	subst.Left = node.Left
	subst.Right = node.Right
	subst.Parent = node.Parent
	subst.Color = node.Color

	if node == *root {
		*root = subst

	} else {
		if node == node.Parent.Left {
			node.Parent.Left = subst
		} else {
			node.Parent.Right = subst
		}
	}

	if subst.Left != sentinel {
		subst.Left.Parent = subst
	}

	if subst.Right != sentinel {
		subst.Right.Parent = subst
	}
}

func (rbt *RBTree) deleteFixup(temp *rbtreeNode) {
	root := &rbt.root
	// ldlog.Default().Info("==== fixup", zap.Reflect("temp", temp.ToMap(rbt.sentinel)))
	// ldlog.Default().Info("==== fixup", zap.Bool("temp != *root", temp != *root), zap.Bool("temp.Color == _colorBlack", temp.Color == _colorBlack))
	for temp != *root && temp.Color == _colorBlack {
		if temp == temp.Parent.Left {
			// ldlog.Default().Info("==== deleteFixupLeftNode")
			rbt.deleteFixupLeftNode(&temp)
		} else {
			// ldlog.Default().Info("==== deleteFixupRightNode")
			rbt.deleteFixupRightNode(&temp)
		}
	}

	temp.Color = _colorBlack
	// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
}

func (rbt *RBTree) deleteFixupLeftNode(temp **rbtreeNode) {
	root := &rbt.root
	sibling := (*temp).Parent.Right
	if sibling.Color == _colorRed {
		// case 1: left rotate at parent
		//     P               S
		//    / \             / \
		//   N   s    -->    p   Sr
		//      / \         / \
		//     Sl  Sr      N   Sl
		sibling.Color = _colorBlack
		(*temp).Parent.Color = _colorRed
		rbt.rotateLeft((*temp).Parent)
		sibling = (*temp).Parent.Right
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
	}

	if sibling.Left.Color == _colorBlack && sibling.Right.Color == _colorBlack {
		// case 2: sibling color flip
		//     p             p
		//    / \           / \
		//   N   S    -->  N   s
		//      / \           / \
		//     Sl  Sr        Sl  Sr
		sibling.Color = _colorRed
		(*temp) = (*temp).Parent
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
		return
	}

	if sibling.Right.Color == _colorBlack {
		// case 3: right rotate at sibling
		// (p could be either color here)
		//    p             p
		//   / \           / \
		//  N   S    -->  N   sl
		//     / \             \
		//    sl  Sr            S
		//                       \
		//                        Sr
		sibling.Left.Color = _colorBlack
		sibling.Color = _colorRed
		rbt.rotateRight(sibling)
		sibling = (*temp).Parent.Right
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
	}

	// case 4: left rotate at parent + color flips
	// (p and sl could be either color here. After rotation,
	// p becomes black, s acquires p's color, and sl keeps its color)
	//
	//      (p)             (s)
	//      / \             / \
	//     N   S     -->   P   Sr
	//        / \         / \
	//       sl  sr      N   sl
	sibling.Color = (*temp).Parent.Color
	(*temp).Parent.Color = _colorBlack
	sibling.Right.Color = _colorBlack
	rbt.rotateLeft((*temp).Parent)

	(*temp) = *root
	// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
}

func (rbt *RBTree) deleteFixupRightNode(temp **rbtreeNode) {
	root := &rbt.root

	sibling := (*temp).Parent.Left
	if sibling.Color == _colorRed {
		// case 1
		sibling.Color = _colorBlack
		(*temp).Parent.Color = _colorRed
		rbt.rotateRight((*temp).Parent)
		sibling = (*temp).Parent.Left
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
	}

	if sibling.Left.Color == _colorBlack && sibling.Right.Color == _colorBlack {
		// case 2
		sibling.Color = _colorRed
		(*temp) = (*temp).Parent
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
		return
	}

	if sibling.Left.Color == _colorBlack {
		// case 3
		sibling.Right.Color = _colorBlack
		sibling.Color = _colorRed
		rbt.rotateLeft(sibling)
		sibling = (*temp).Parent.Left
		// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
	}

	// case 4
	sibling.Color = (*temp).Parent.Color
	(*temp).Parent.Color = _colorBlack
	sibling.Left.Color = _colorBlack
	rbt.rotateRight((*temp).Parent)
	(*temp) = *root
	// ldlog.Default().Info("==== debug", zap.Reflect("tree", rbt.ToMap()))
}
