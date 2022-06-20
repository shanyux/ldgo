/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"github.com/distroy/ldgo/ldconv"
	"github.com/distroy/ldgo/ldrand"
	"github.com/smartystreets/goconvey/convey"
)

func TestCompareInterface(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("aType != bType", func() {
			convey.So(CompareInterface(nil, false), convey.ShouldEqual, -1)
			convey.So(CompareInterface(true, 0), convey.ShouldEqual, -1)

			convey.So(CompareInterface((*int)(nil), (*uint)(nil)), convey.ShouldEqual, -1)
			convey.So(CompareInterface((*int8)(nil), (*uint)(nil)), convey.ShouldEqual, -1)
			convey.So(CompareInterface((*int)(nil), (*int)(nil)), convey.ShouldEqual, 0)
			convey.So(CompareInterface(0, []int{}), convey.ShouldEqual, -1)
		})

		convey.Convey("bool", func() {
			convey.So(CompareInterface(false, true), convey.ShouldEqual, -1)
			convey.So(CompareInterface(true, false), convey.ShouldEqual, 1)
			convey.So(CompareInterface(false, false), convey.ShouldEqual, 0)
			convey.So(CompareInterface(true, true), convey.ShouldEqual, 0)
		})

		convey.Convey("int", func() {
			convey.So(CompareInterface(uint64(99), 100), convey.ShouldEqual, -1)
			convey.So(CompareInterface(uint64(math.MaxInt64+1), 100), convey.ShouldEqual, 1)
			convey.So(CompareInterface(200, int64(100)), convey.ShouldEqual, 1)
			convey.So(CompareInterface(200, uint64(200)), convey.ShouldEqual, 0)
			convey.So(CompareInterface(int64(-200), 100), convey.ShouldEqual, -1)
		})

		convey.Convey("float", func() {
			convey.So(CompareInterface(float64(100.0), float32(100.0)), convey.ShouldEqual, 0)
			convey.So(CompareInterface(99.0, 100.0), convey.ShouldEqual, -1)
			convey.So(CompareInterface(99.0, math.NaN()), convey.ShouldEqual, 1)
			convey.So(CompareInterface(-99.0, math.NaN()), convey.ShouldEqual, 1)
			convey.So(CompareInterface(-99.0, float32(math.NaN())), convey.ShouldEqual, 1)
			convey.So(CompareInterface(float32(math.NaN()), 100.0), convey.ShouldEqual, -1)

			convey.So(CompareInterface(math.NaN(), math.NaN()), convey.ShouldEqual, 0)
			convey.So(CompareInterface(float32(math.NaN()), float32(math.NaN())), convey.ShouldEqual, 0)

			convey.So(CompareInterface(-99.0, math.Inf(1)), convey.ShouldEqual, -1)
			convey.So(CompareInterface(-99.0, math.Inf(-1)), convey.ShouldEqual, 1)
		})

		convey.Convey("number", func() {
			convey.So(CompareInterface(99, float64(100)), convey.ShouldEqual, -1)
			convey.So(CompareInterface(99, float64(99)), convey.ShouldEqual, 0)
		})

		convey.Convey("complex", func() {
			convey.So(CompareInterface(complex(100, 200), complex64(complex(100, 200))), convey.ShouldEqual, 0)

			convey.So(CompareInterface(complex(100, 200), complex(11, 300)), convey.ShouldEqual, 1)
			convey.So(CompareInterface(complex(100, 200), complex(111, -300)), convey.ShouldEqual, -1)
			convey.So(CompareInterface(complex(100, 200), complex(100, 300)), convey.ShouldEqual, -1)
			convey.So(CompareInterface(complex(100, 200), complex(100, 150)), convey.ShouldEqual, 1)
		})

		convey.Convey("string", func() {
			convey.So(CompareInterface("", `abc`), convey.ShouldEqual, -1)
			convey.So(CompareInterface("aaa", `a`), convey.ShouldEqual, 1)
			convey.So(CompareInterface("bbb", `aaaaaa`), convey.ShouldEqual, 1)
		})

		convey.Convey("map", func() {
			convey.So(CompareInterface(map[int]int{0: 0}, map[interface{}]int{0: 0}), convey.ShouldEqual, -1)
			convey.So(CompareInterface(map[int]int{0: 0}, map[int]int{0: 0}), convey.ShouldEqual, 0)
			convey.So(CompareInterface(map[int]int{0: 0}, map[int]int{}), convey.ShouldEqual, 1)
			convey.So(CompareInterface(map[int]int{0: 0}, map[int]int{1: 0}), convey.ShouldEqual, -1)
			convey.So(CompareInterface(map[int]int{1: 1}, map[int]int{1: 0}), convey.ShouldEqual, 1)
		})

		convey.Convey("slice", func() {
			convey.So(CompareInterface(
				[]interface{}{100, uint(200), float32(300)},
				[]interface{}{100, float64(200), ""},
			), convey.ShouldEqual, -1)

			convey.So(CompareInterface(
				[]interface{}{100, uint(200), ""},
				[]interface{}{100, float64(200), ""},
			), convey.ShouldEqual, 0)
		})
	})
}

var testStringCompareCases [][2]string

func testGetStringCompareCases(t testing.TB) [][2]string {
	const size = 1024
	const strLen = 128
	if testStringCompareCases == nil {
		cases := make([][2]string, 0, size)
		for i := 0; i < size; i++ {
			cases = append(cases, [2]string{
				// ldconv.StrToBytesUnsafe(ldrand.String(128)), ldconv.StrToBytesUnsafe(ldrand.String(128)),
				ldrand.String(128), ldrand.String(128),
			})
		}

		testStringCompareCases = cases
	}

	return testStringCompareCases
}

func BenchmarkStringsCompare(b *testing.B) {
	cases := testGetStringCompareCases(b)
	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos++
			if pos >= size {
				pos -= size
			}

			a, b := v[0], v[1]
			strings.Compare(a, b)
			// strings.Compare(ldconv.BytesToStrUnsafe(a), ldconv.BytesToStrUnsafe(b))
		}
	})
}

func BenchmarkBytesCompare(b *testing.B) {
	cases := testGetStringCompareCases(b)
	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		size := len(cases)
		pos := 0
		for p.Next() {
			v := cases[pos]
			pos++
			if pos >= size {
				pos -= size
			}

			a, b := v[0], v[1]
			bytes.Compare(ldconv.StrToBytesUnsafe(a), ldconv.StrToBytesUnsafe(b))
			// bytes.Compare(a, b)
		}
	})
}
