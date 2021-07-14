/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"sync"

	"github.com/distroy/ldgo/ldsync"
)

func NewBoolLocker(b bool) sync.Locker {
	return ldsync.NewBoolLocker(b)
}
