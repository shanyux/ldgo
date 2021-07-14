/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldsync"

type DoneWait = ldsync.DoneWait

func NewDoneWait() DoneWait {
	return ldsync.NewDoneWait()
}
