/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldtopk"

func TopkStringsAdd(b []string, k int, x string) ([]string, bool) {
	return ldtopk.TopkStringsAdd(b, k, x)
}
