/*
 * Copyright (C) distroy
 */

package ldgo

import "github.com/distroy/ldgo/ldcore"

func TopkStringsAdd(b []string, k int, x string) ([]string, bool) {
	return ldcore.TopkStringsAdd(b, k, x)
}
