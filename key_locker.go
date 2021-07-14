/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldsync"

type KeyLocker = ldsync.KeyLocker

func NewKeyLocker() KeyLocker {
	return ldsync.NewKeyLocker()
}
