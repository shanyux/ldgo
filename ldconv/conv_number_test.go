/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func test_ConvInt(s string, r int64) {
	convey.Convey(s, func() {
		n, _ := convInt(StrToBytesUnsafe(s))
		convey.So(n, convey.ShouldEqual, r)
	})
}

func test_ConvUint(s string, r uint64) {
	convey.Convey(s, func() {
		n, _ := convUint(StrToBytesUnsafe(s))
		convey.So(n, convey.ShouldEqual, r)
	})
}

func test_ConvFloat(s string, r float64) {
	convey.Convey(s, func() {
		n, _ := convFloat(StrToBytesUnsafe(s))
		f, _ := n.Float64()
		convey.So(f, convey.ShouldEqual, r)

		const base = 100000
		baseFloat := (&big.Float{}).SetInt64(base)
		n1 := (&big.Float{}).Mul(n, baseFloat)
		x, _ := (&big.Float{}).SetString(strconv.FormatFloat(r, 'f', -1, 64))
		x = x.Mul(x, baseFloat)

		f1, _ := n1.Int64()
		r1, _ := x.Int64()
		convey.So(f1, convey.ShouldEqual, r1)
	})
}

func Test_ConvInt(t *testing.T) {
	convey.Convey("", t, func() {
		test_ConvInt("123", 123)
		test_ConvInt("123.1", 123)
		test_ConvInt(".1", 0)
		test_ConvInt("0123", 123)
		test_ConvInt("0o123", 0123)
		test_ConvInt("0O123", 0123)
		test_ConvInt("+-0o123", -0123)
		test_ConvInt("123e+1", 1230)
		test_ConvInt("123e1", 1230)
		test_ConvInt("123e-1", 12)
		test_ConvInt("0x123", 0x123)
		test_ConvInt("--0X123", 0x123)
		test_ConvInt("-+-0x123", 0x123)
		test_ConvInt("-+-0xeFaB", 0xeFaB)
	})
}

func Test_ConvUint(t *testing.T) {
	convey.Convey("", t, func() {
		test_ConvUint("123", 123)
		test_ConvUint("123.1", 123)
		test_ConvUint(".1", 0)
		test_ConvUint("0123", 123)
		test_ConvUint("0o123", 0123)
		// testConvUint("+-0123", uint64(int64(-0123)))
		test_ConvUint("123e+1", 1230)
		test_ConvUint("123e1", 1230)
		test_ConvUint("123e-1", 12)
		test_ConvUint("0x123", 0x123)
		test_ConvUint("--0X123", 0x123)
		test_ConvUint("-+-0x123", 0x123)
		test_ConvUint("-+-0xeFaB", 0xeFaB)
	})
}

func Test_ConvFloat(t *testing.T) {
	convey.Convey("", t, func() {
		test_ConvFloat("2.32", 2.32)
		test_ConvFloat("123", 123)
		test_ConvFloat("123.1", 123.1)
		test_ConvFloat(".1", 0.1)
		test_ConvFloat("0123", 123)
		test_ConvFloat("0o123", 0123)
		// testConvFloat("+-0123", uint64(int64(-0123)))
		test_ConvFloat("123e+1", 1230)
		test_ConvFloat("123e1", 1230)
		test_ConvFloat("123e-1", 12.3)
		test_ConvFloat("0x123", 0x123)
		test_ConvFloat("--0X123", 0x123)
		test_ConvFloat("-+-0x123", 0x123)
		test_ConvFloat("-+-0xeFaB", 0xeFaB)
		test_ConvFloat("0.30129", 0.30129)
		test_ConvFloat("0.30129e3", 301.29)
	})
}
