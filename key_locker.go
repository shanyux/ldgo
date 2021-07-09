/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"github.com/distroy/ldgo/ldcore"
)

type KeyLocker = ldcore.KeyLocker

func NewKeyLocker() KeyLocker {
	return ldcore.NewKeyLocker()
}
