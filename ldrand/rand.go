/*
 * Copyright (C) distroy
 */

package ldrand

import (
	"math/rand"
	"time"

	"github.com/distroy/ldgo/ldconv"
)

var randStringLetters = ldconv.StrToBytesUnsafe("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var globalRand = rand.New(NewFastSource(time.Now().UnixNano()))

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt() int     { return Int() }
func RandInt31() int32 { return Int31() }
func RandInt63() int64 { return Int63() }

func RandIntn(n int) int       { return Intn(n) }
func RandInt31n(n int32) int32 { return Int31n(n) }
func RandInt63n(n int64) int64 { return Int63n(n) }

func RandUint() uint     { return Uint() }
func RandUint32() uint32 { return Uint32() }
func RandUint64() uint64 { return Uint64() }

// RandFloat64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func RandFloat64() float64 { return Float64() }

// RandFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func RandFloat32() float32 { return Float32() }

func RandString(n int) string { return String(n) }

func Int() int     { return globalRand.Int() }
func Int31() int32 { return globalRand.Int31() }
func Int63() int64 { return globalRand.Int63() }

func Intn(n int) int       { return globalRand.Intn(n) }
func Int31n(n int32) int32 { return globalRand.Int31n(n) }
func Int63n(n int64) int64 { return globalRand.Int63n(n) }

func Uint() uint     { return uint(globalRand.Uint64()) }
func Uint32() uint32 { return globalRand.Uint32() }
func Uint64() uint64 { return globalRand.Uint64() }

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Float64() float64 { return globalRand.Float64() }

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func Float32() float32 { return globalRand.Float32() }

func String(n int) string {
	letters := randStringLetters

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[RandIntn(len(letters))]
	}

	return ldconv.BytesToStrUnsafe(b)
}

func Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	b := make([]byte, n)
	globalRand.Read(b)
	return b
}

// A Rand is a source of random numbers.
type Rand interface {
	Int() int
	Int31() int32
	Int63() int64

	Uint() uint
	Uint32() uint32
	Uint64() uint64

	Intn(n int) int
	Int31n(n int32) int32
	Int63n(n int64) int64

	// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
	Float32() float32
	// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
	Float64() float64

	Shuffle(n int, swap func(i, j int))

	Read(p []byte) (int, error)
}

func New(src rand.Source) Rand {
	if r, ok := src.(Rand); ok {
		return r
	}
	return randWrapper{Rand: rand.New(src)}
}
