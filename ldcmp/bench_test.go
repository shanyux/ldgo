/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/distroy/ldgo/ldrand"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/distroy/ldgo/ldcmp
// cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// BenchmarkStringsCompareRand
// BenchmarkStringsCompareRand-8                   1000000000               6.996 ns/op
// BenchmarkBytesCompareRandWithCoverter
// BenchmarkBytesCompareRandWithCoverter-8         1000000000               8.374 ns/op
// BenchmarkBytesCompareRand
// BenchmarkBytesCompareRand-8                     1000000000               6.411 ns/op
// BenchmarkStringsCompareSame
// BenchmarkStringsCompareSame-8                   1000000000               0.7949 ns/op
// BenchmarkBytesCompareSame
// BenchmarkBytesCompareSame-8                     1000000000               1.671 ns/op
// BenchmarkStringsCompareDifferent
// BenchmarkStringsCompareDifferent-8              1000000000               2.995 ns/op
// BenchmarkBytesCompareDifferent
// BenchmarkBytesCompareDifferent-8                1000000000               1.895 ns/op

var testStringCompareCases [][2]string

func testStrToBytesUnsafe(s string) []byte {
	// return *(*[]byte)(unsafe.Pointer(&s))

	type bytesFromString reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &bytesFromString{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}

func testGetStringCompareCases(t testing.TB) [][2]string {
	const size = 1024
	const strLen = 1024
	if testStringCompareCases == nil {
		cases := make([][2]string, 0, size)
		for i := 0; i < size; i++ {
			cases = append(cases, [2]string{
				// testStrToBytesUnsafe(ldrand.String(128)), testStrToBytesUnsafe(ldrand.String(128)),
				ldrand.String(128), ldrand.String(128),
			})
		}

		testStringCompareCases = cases
	}

	return testStringCompareCases
}

func BenchmarkStringsCompareRand(b *testing.B) {
	cases := testGetStringCompareCases(b)

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s1, s2 := v[0], v[1]
			strings.Compare(s1, s2)
		}
	})
}

func BenchmarkBytesCompareRandWithCoverter(b *testing.B) {
	cases := testGetStringCompareCases(b)

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s1, s2 := v[0], v[1]
			bytes.Compare(testStrToBytesUnsafe(s1), testStrToBytesUnsafe(s2))
		}
	})
}

func BenchmarkBytesCompareRand(b *testing.B) {
	cases0 := testGetStringCompareCases(b)
	cases := make([][2][]byte, 0, len(cases0))
	for _, v := range cases0 {
		cases = append(cases, [2][]byte{
			testStrToBytesUnsafe(v[0]), testStrToBytesUnsafe(v[1]),
		})
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s1, s2 := v[0], v[1]
			bytes.Compare(s1, s2)
		}
	})
}

func BenchmarkStringsCompareSame(b *testing.B) {
	s1 := "benchmark strings compare"
	s2 := "benchmark strings compare"

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			strings.Compare(s1, s2)
		}
	})
}

func BenchmarkBytesCompareSame(b *testing.B) {
	s1 := []byte("benchmark bytes compare")
	s2 := []byte("benchmark bytes compare")

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			bytes.Compare(s1, s2)
		}
	})
}

func BenchmarkStringsCompareDifferent(b *testing.B) {
	s1 := "benchmark strings compare s1"
	s2 := "benchmark strings compare s2"

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			strings.Compare(s1, s2)
		}
	})
}

func BenchmarkBytesCompareDifferent(b *testing.B) {
	s1 := []byte("benchmark bytes compare s1")
	s2 := []byte("benchmark bytes compare s2")

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			bytes.Compare(s1, s2)
		}
	})
}
