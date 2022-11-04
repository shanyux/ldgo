/*
 * Copyright (C) distroy
 */

package ldgorm

import "strings"

var (
	escapeReplacerForLike *strings.Replacer
)

func init() {
	replaceStrings := []string{
		"%", `\%`,
		"_", `\_`,
		"\\", `\\`,
	}

	escapeReplacerForLike = strings.NewReplacer(replaceStrings...)
}

func escapeForLike(s string) string {
	return escapeReplacerForLike.Replace(s)
}
