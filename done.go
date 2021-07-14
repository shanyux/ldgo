/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldsync"

type Done = ldsync.Done

func NewDone() Done { return ldsync.NewDone() }
