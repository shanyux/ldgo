/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldtopk"

type TopkInterface = ldtopk.TopkInterface

func TopkAdd(b TopkInterface, k int, x interface{}) bool {
	return ldtopk.TopkAdd(b, k, x)
}
