/*
 * Copyright (C) distroy
 */

package implement

import (
	"github.com/distroy/ldgo/v2/lditer"
	"github.com/distroy/ldgo/v2/ldrbtree"
)

var (
	_ lditer.Iterator[int] = lditer.MakeIter[int](ldrbtree.RBTreeIterator[int]{})
	_ lditer.Iterator[int] = lditer.MakeIter[int](ldrbtree.RBTreeReverseIterator[int]{})

	_ lditer.Iterator[ldrbtree.MapNode[string, int]] = lditer.MakeIter[ldrbtree.MapNode[string, int]](ldrbtree.MapIterator[string, int]{})
	_ lditer.Iterator[ldrbtree.MapNode[string, int]] = lditer.MakeIter[ldrbtree.MapNode[string, int]](ldrbtree.MapReverseIterator[string, int]{})

	_ lditer.Range[int] = (*ldrbtree.RBTreeRange[int])(nil)
	_ lditer.Range[int] = (*ldrbtree.RBTreeReverseRange[int])(nil)

	_ lditer.Range[ldrbtree.MapNode[string, int]] = (*ldrbtree.MapRange[string, int])(nil)
	_ lditer.Range[ldrbtree.MapNode[string, int]] = (*ldrbtree.MapReverseRange[string, int])(nil)
)
