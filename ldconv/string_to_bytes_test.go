/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/smartystreets/goconvey/convey"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/distroy/ldgo/ldconv
// cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
// BenchmarkStrToBytesUnsafe1
// BenchmarkStrToBytesUnsafe1-8    1000000000               1.360 ns/op
// BenchmarkStrToBytesUnsafe2
// BenchmarkStrToBytesUnsafe2-8    1000000000               0.7977 ns/op

func testStrToBytesUnsafe1(s string) []byte {
	type bytesFromString reflect.SliceHeader
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &bytesFromString{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  0,
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}

func testStrToBytesUnsafe2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func TestStrToBytesUnsafe1(t *testing.T) {
	fn := testStrToBytesUnsafe1

	convey.Convey(t.Name(), t, func() {
		convey.So(fn(""), convey.ShouldResemble, []byte(nil))
		convey.So(len(fn("abc")), convey.ShouldEqual, 3)
		convey.So(cap(fn("abc")), convey.ShouldEqual, 0)
	})
}

func TestStrToBytesUnsafe2(t *testing.T) {
	fn := testStrToBytesUnsafe2

	convey.Convey(t.Name(), t, func() {
		convey.So(fn(""), convey.ShouldResemble, []byte(nil))
		convey.So(len(fn("abc")), convey.ShouldEqual, 3)
		// convey.So(cap(fn("abc")), convey.ShouldEqual, 0)
	})
}

func BenchmarkStrToBytesUnsafe1(b *testing.B) {
	fn := testStrToBytesUnsafe1

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			fn("benchmark str to bytes unsafe 1")
		}
	})
}

func BenchmarkStrToBytesUnsafe2(b *testing.B) {
	fn := testStrToBytesUnsafe2

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			fn("benchmark str to bytes unsafe 2")
		}
	})
}
