/*
 * Copyright (C) distroy
 */

package ldtopk

func TopkStringsAdd(b []string, k int, x string) ([]string, bool) {
	return TopkAdd[string](b, k, x)
}
