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
	0x3b78bb846a443726,
	0x772c3f9b061acdac,
	0x62dfa820aede2c50,
	0xa51729651df84541,
	0x93c695f2456306c5,
	0x2aa1107ecc9b60f2,
	0x19fd8406df31d838,
	0xf0ea0719370c8f9e,
	0x418e4a4878a79b6b,
	0xe49476ea842fb1e4,
	0x0c09febc20b9e3df,
	0xdf63eca7538d728d,
	0x56b2d3cd92e65e77,
	0xbe5552d1e952f903,
	0xc83bc15ff1c0a4ba,
	0x8d406d33bb751a19,
}

func NewFastSource(seed int64) rand.Source64 {
	s := &fastSource{}
	s.Seed(seed)
	return s
}

type fastSource struct {
	seed uint64
}

func (s *fastSource) Seed(seed int64) {
	n := uint64(seed)
	// initFastSource(seed, s.xor)
	atomic.StoreUint64(&s.seed, s.next(&n))
}

func (s *fastSource) Int63() int64 {
	return int64(s.Uint64() << 1 >> 1) // clear sign bit
}

func (s *fastSource) Uint64() uint64 {
	return s.next(&s.seed)
}

func (_ *fastSource) next(seed *uint64) uint64 {
	n := atomic.AddUint64(seed, fastSourceStep)
	b := n & 0xf

	x := n
	x = x ^ fastSourceXor[b]
	x = x - (x & 0xf)

	b = uint64(fastSourceLast4Bits[b]) ^ ((n >> 4) & 0xf) ^ ((n >> 8) & 0xf)
	x = x | uint64(b)
	return x
}

func initFastSource(seed int64, xor [16]uint64) {
	buf := [16]byte{}
	for i := range buf {
		buf[i] = byte(i)
	}

	r := rand.New(rand.NewSource(seed))
	for i := 0; i < 16; i++ {
		r.Shuffle(len(buf), func(i, j int) { buf[i], buf[j] = buf[j], buf[i] })
		for j := range xor {
			xor[j] = (xor[j] << 4) | uint64(buf[j])
		}
	}
	// for i := range xor {
	// 	xor[i] ^= fastSourceXor[i]
	// }
}
