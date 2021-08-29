/*
 * Copyright (C) distroy
 */

package ldrand

import (
	"math/rand"
	"sync/atomic"
)

const (
	fastSourceStep = 0x1753715715313157
)

var fastSourceLast4Bits [16]byte = [...]byte{
	0x5,
	0x1,
	0x3,
	0x0,
	0x6,
	0xe,
	0xf,
	0x7,
	0xa,
	0x2,
	0x9,
	0xd,
	0x8,
	0xb,
	0x4,
	0xc,
}

var fastSourceXor [16]uint64 = [...]uint64{
	0x3b78bb846a443720,
	0x772c3f9b061acda0,
	0x62dfa820aede2c50,
	0xa51729651df84540,
	0x93c695f2456306c0,
	0x2aa1107ecc9b60f0,
	0x19fd8406df31d830,
	0xf0ea0719370c8f90,
	0x418e4a4878a79b60,
	0xe49476ea842fb1e0,
	0x0c09febc20b9e3d0,
	0xdf63eca7538d7280,
	0x56b2d3cd92e65e70,
	0xbe5552d1e952f900,
	0xc83bc15ff1c0a4b0,
	0x8d406d33bb751a10,
}

func NewFastSource(seed int64) rand.Source64 {
	s := &fastSource{seed: uint64(seed)}
	s.Seed(s.Int63())
	return s
}

type fastSource struct {
	seed uint64
}

func (s *fastSource) Seed(seed int64) {
	t := &fastSource{seed: uint64(seed)}
	atomic.StoreUint64(&s.seed, t.Uint64())
}

func (s *fastSource) Int63() int64 {
	return int64(s.Uint64() << 1 >> 1) // clear sign bit
}

func (s *fastSource) Uint64() uint64 {
	x := atomic.AddUint64(&s.seed, fastSourceStep)
	b := x & 0xf

	x = x ^ fastSourceXor[b]

	x = x - (x & 0xf)
	b = uint64(fastSourceLast4Bits[b]) ^ ((x & 0xf0) >> 4)
	x = x | uint64(b&0xf)
	return x
}
