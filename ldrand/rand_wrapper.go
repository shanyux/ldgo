/*
 * Copyright (C) distroy
 */

package ldrand

import "math/rand"

type randWrapper struct {
	*rand.Rand
}

func (r randWrapper) Uint() uint { return uint(r.Rand.Uint64()) }
