/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

type TopkInterface = ldcore.TopkInterface

func TopkAdd(b TopkInterface, k int, x interface{}) bool {
	return ldcore.TopkAdd(b, k, x)
}
