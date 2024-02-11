/*
 * Copyright (C) distroy
 */

package internal

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CmderWithParse interface {
	redis.Cmder

	Parse(c context.Context) error
}
