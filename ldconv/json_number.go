/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"encoding/json"
	"math/big"
	"strconv"
)

type jsonNumber json.Number

func (p *jsonNumber) String() string {
	if p == nil {
		return ""
	}
	return (*json.Number)(p).String()
}

func (p *jsonNumber) Int64() (int64, error) {
	s := p.String()
	if s == "" {
		return 0, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

func (p *jsonNumber) Uint64() (uint64, error) {
	s := p.String()
	if s == "" {
		return 0, nil
	}
	return strconv.ParseUint(s, 10, 64)
}

func (p *jsonNumber) Float64() (float64, error) {
	s := p.String()
	if s == "" {
		return 0, nil
	}

	return strconv.ParseFloat(s, 64)
}

func (p *jsonNumber) Float32() (float32, error) {
	n, err := p.Float64()
	if err != nil {
		return float32(n), err
	}

	// if n > math.MaxFloat32 {
	// 	return float32(n), lderr.ErrNumberOverflow
	// }

	return float32(n), nil
}

func (p *jsonNumber) Bool() (bool, error) {
	s := p.String()
	if s == "" {
		return false, nil
	}

	if s[0] == '-' {
		n, err := p.Int64()
		if err == nil {
			return n != 0, nil
		}

	} else {
		n, err := p.Uint64()
		if err == nil {
			return n != 0, nil
		}
	}

	n, err := p.Decimal()
	return !n.IsZero(), err

	// n, err := p.Float64()
	// return n != 0, err
}

func (p *jsonNumber) BigFloat() (*big.Float, error) {
	s := p.String()
	if s == "" {
		return newBigFloatZero(), nil
	}

	r := &big.Float{}
	r, _, err := r.Parse(s, 10)
	return r, err
}

func (p *jsonNumber) Decimal() (decimalNumber, error) {
	s := p.String()
	if s == "" {
		return newDecimalZero(), nil
	}

	return newDecimalFromFloatStr(s)
}
