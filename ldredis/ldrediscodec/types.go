/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"github.com/distroy/ldgo/v2/ldredis"
)

type (
	client   = ldredis.Redis
	Reporter = ldredis.Reporter
)

type ZMember[T any] struct {
	Score  float64
	Member T
}

type Pair[T any] struct {
	First  string // key or field
	Second T      // value
}

type Cmdable[T comparable] interface {
	StringCmdable[T]
	HashCmdable[T]
	ListCmdable[T]
	SetCmdable[T]
	SortedSetCmdable[T]

	// XAdd(a *XAddArgs) *StringCmd
}
