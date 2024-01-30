/*
 * Copyright (C) distroy
 */

package cmp

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/distroy/ldgo/v2/ldcmp
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	randStringLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func randString(n int) string {
	letters := randStringLetters
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var testStringCompareCases map[int][][2]string

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

func getGestGetStringCompareCases(strLen int) [][2]string {
	if testStringCompareCases == nil {
		testStringCompareCases = make(map[int][][2]string)
	}
	const size = 1024
	// const strLen = 1024
	if testStringCompareCases[strLen] == nil {
		cases := make([][2]string, 0, size)
		for i := 0; i < size; i++ {
			cases = append(cases, [2]string{
				randString(strLen), randString(strLen),
			})
		}

		testStringCompareCases[strLen] = cases
	}

	return testStringCompareCases[strLen]
}

func benchmarkStrCompare_same_stringsCompare(b *testing.B, cases [][2]string) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[0]
			strings.Compare(s0, s1)
		}
	})
}

func benchmarkStrCompare_different_stringsCompare(b *testing.B, cases [][2]string) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[1]
			strings.Compare(s0, s1)
		}
	})
}

func benchmarkStrCompare_same_bytesCompare(b *testing.B, cases [][2][]byte) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[0]
			bytes.Compare(s0, s1)
		}
	})
}

func benchmarkStrCompare_different_bytesCompare(b *testing.B, cases [][2][]byte) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[1]
			bytes.Compare(s0, s1)
		}
	})
}

func benchmarkStrCompare_same_bytesCompareWithConverter(b *testing.B, cases [][2]string) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[0]
			bytes.Compare(testStrToBytesUnsafe(s0), testStrToBytesUnsafe(s1))
		}
	})
}

func benchmarkStrCompare_different_bytesCompareWithConverter(b *testing.B, cases [][2]string) {
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos = (pos + 1) % size

			s0 := v[0]
			s1 := v[1]
			bytes.Compare(testStrToBytesUnsafe(s0), testStrToBytesUnsafe(s1))
		}
	})
}

func benchmarkStrCompare(b *testing.B, strCases [][2]string) {
	caseCount := len(strCases)
	byteCases := make([][2][]byte, 0, caseCount)
	for _, v := range strCases {
		byteCases = append(byteCases, [2][]byte{
			testStrToBytesUnsafe(v[0]), testStrToBytesUnsafe(v[1]),
		})
	}

	b.Run("same", func(b *testing.B) {
		b.Run("strings.Compare", func(b *testing.B) {
			benchmarkStrCompare_same_stringsCompare(b, strCases)
		})
		b.Run("bytes.Compare", func(b *testing.B) {
			benchmarkStrCompare_same_bytesCompare(b, byteCases)
		})
		b.Run("bytes.Compare with convert", func(b *testing.B) {
			benchmarkStrCompare_same_bytesCompareWithConverter(b, strCases)
		})
	})
	b.Run("different", func(b *testing.B) {
		b.Run("strings.Compare", func(b *testing.B) {
			benchmarkStrCompare_different_stringsCompare(b, strCases)
		})
		b.Run("bytes.Compare", func(b *testing.B) {
			benchmarkStrCompare_different_bytesCompare(b, byteCases)
		})
		b.Run("bytes.Compare with convert", func(b *testing.B) {
			benchmarkStrCompare_different_bytesCompareWithConverter(b, strCases)
		})
	})
}

func BenchmarkStringsCompare_128(b *testing.B) {
	cases := getGestGetStringCompareCases(128)
	benchmarkStrCompare(b, cases)
}
func BenchmarkStringsCompare_1024(b *testing.B) {
	cases := getGestGetStringCompareCases(1024)
	benchmarkStrCompare(b, cases)
}
func BenchmarkStringsCompare_4096(b *testing.B) {
	cases := getGestGetStringCompareCases(4096)
	benchmarkStrCompare(b, cases)
}

// func BenchmarkStringsCompareSame(b *testing.B) {
// 	s1 := "benchmark strings compare"
// 	s2 := "benchmark strings compare"
//
// 	b.ResetTimer()
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			strings.Compare(s1, s2)
// 		}
// 	})
// }
//
// func BenchmarkBytesCompareSame(b *testing.B) {
// 	s1 := []byte("benchmark bytes compare")
// 	s2 := []byte("benchmark bytes compare")
//
// 	b.ResetTimer()
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			bytes.Compare(s1, s2)
// 		}
// 	})
// }
//
// func BenchmarkStringsCompareDifferent(b *testing.B) {
// 	s1 := "benchmark strings compare s1"
// 	s2 := "benchmark strings compare s2"
//
// 	b.ResetTimer()
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			strings.Compare(s1, s2)
// 		}
// 	})
// }
//
// func BenchmarkBytesCompareDifferent(b *testing.B) {
// 	s1 := []byte("benchmark bytes compare s1")
// 	s2 := []byte("benchmark bytes compare s2")
//
// 	b.ResetTimer()
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			bytes.Compare(s1, s2)
// 		}
// 	})
// }
