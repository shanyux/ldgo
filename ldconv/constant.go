/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"fmt"
	"math/big"
)

var (
	_ERR_INVALID_SYNTAX = fmt.Errorf("invalid syntax")
	_ERR_UNKOWN_TYPE    = fmt.Errorf("unkown type")

	_BIG_FLOAT_ZERO = (&big.Float{}).SetInt64(0)
)

func newBigFloatZero() *big.Float {
	return (&big.Float{}).Copy(_BIG_FLOAT_ZERO)
}
