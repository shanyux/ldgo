/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"math/big"

	"github.com/shopspring/decimal"
)

type decimalNumber = decimal.Decimal

var decimalZero = newDecimalFromInt(0)

func newDecimalZero() decimalNumber {
	return decimalZero
}

func newDecimalFromFloat(n float64) decimalNumber {
	return decimal.NewFromFloat(n)
}

func newDecimalFromFloatStr(s string) (decimalNumber, error) {
	return decimal.NewFromString(s)
}

func newDecimalFromBigFloat(n *big.Float) decimalNumber {
	r, _ := decimal.NewFromString(n.String())
	return r
}

func newDecimalFromInt(n int64) decimalNumber {
	return decimal.NewFromInt(n)
}

func newDecimalFromUint(n uint64) decimalNumber {
	return decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(n)), 0)
}

func newDecimalFromIntStr(s string, negative bool, base int) (decimalNumber, error) {
	i, _ := new(big.Int).SetString(s, base)
	if negative {
		i = i.Neg(i)
	}
	return decimal.NewFromBigInt(i, 0), nil
}
