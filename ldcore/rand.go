/*
 * Copyright (C) distroy
 */

package ldcore

import (
	"math/rand"
	"time"
)

var (
	_RAND_STRING_LETTERS = StrToBytesUnsafe("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt() int     { return rand.Int() }
func RandInt31() int32 { return rand.Int31() }
func RandInt63() int64 { return rand.Int63() }

func RandIntn(n int) int       { return rand.Intn(n) }
func RandInt31n(n int32) int32 { return rand.Int31n(n) }
func RandInt63n(n int64) int64 { return rand.Int63n(n) }

func RandUint() uint     { return uint(rand.Uint64()) }
func RandUint32() uint32 { return rand.Uint32() }
func RandUint64() uint64 { return rand.Uint64() }

// RandFloat64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func RandFloat64() float64 { return rand.Float64() }

// RandFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func RandFloat32() float32 { return rand.Float32() }

func RandString(n int) string {
	letters := _RAND_STRING_LETTERS
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[RandIntn(len(letters))]
	}

	return BytesToStrUnsafe(b)
}
