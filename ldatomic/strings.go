/*
 * Copyright (C) distroy
 */

package ldatomic

type (
	Strings = Any[[]string]
	Bytes   = Any[[]byte]
)

func NewStrings(d []string) *Strings { return NewAny[[]string](d) }
func NewBytes(d []byte) *Bytes       { return NewAny[[]byte](d) }
