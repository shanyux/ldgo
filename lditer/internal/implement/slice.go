/*
 * Copyright (C) distroy
 */

package implement

import (
	"github.com/distroy/ldgo/v2/lditer"
	"github.com/distroy/ldgo/v2/ldslice"
)

var (
	_ lditer.Iterator[int] = lditer.MakeIter[int](ldslice.Begin[int]([]int{}))
	_ lditer.Iterator[int] = lditer.MakeIter[int](ldslice.End[int]([]int{}))

	_ lditer.Iterator[int] = lditer.MakeIter[int](ldslice.RBegin[int]([]int{}))
	_ lditer.Iterator[int] = lditer.MakeIter[int](ldslice.REnd[int]([]int{}))

	_ lditer.Range[int] = (*ldslice.Range[int])(nil)
	_ lditer.Range[int] = (*ldslice.ReverseRange[int])(nil)
)
