/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"sync"

	"github.com/distroy/ldgo/ldcore"
)

func NewBoolLocker(b bool) sync.Locker {
	return ldcore.NewBoolLocker(b)
}
