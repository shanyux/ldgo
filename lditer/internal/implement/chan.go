/*
 * Copyright (C) distroy
 */

package implement

import (
	"github.com/distroy/ldgo/v2/ldchan"
	"github.com/distroy/ldgo/v2/lditer"
)

var (
	_ lditer.ConstIterator[int] = lditer.ConstIter[int](ldchan.Begin((chan int)(nil)))
	_ lditer.ConstIterator[int] = lditer.ConstIter[int](ldchan.End((chan int)(nil)))

	_ lditer.ConstRange[int] = (*ldchan.Range[int])(nil)
)
