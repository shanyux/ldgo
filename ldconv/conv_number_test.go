/*
 * Copyright (C) distroy
 */

package ldconv

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

const (
	_TEST_FLOAT_MUL_BASE = 100000
)

func testConvInt(s string, r int64) {
	convey.Convey(s, func() {
		n, _ := convInt(StrToBytesUnsafe(s))
		convey.So(n, convey.ShouldEqual, r)
	})
}

func testConvUint(s string, r uint64) {
	convey.Convey(s, func() {
		n, _ := convUint(StrToBytesUnsafe(s))
		convey.So(n, convey.ShouldEqual, r)
	})
}

func testConvFloat(s string, r0 float64, r1 int64) {
	name := fmt.Sprintf("%s|%g|%d", s, r0, r1)
	convey.Convey(name, func() {
		n0, _ := convFloat(StrToBytesUnsafe(s))
		v0, _ := n0.Float64()
		convey.So(v0, convey.ShouldEqual, r0)

		baseFloat := newDecimalFromInt(_TEST_FLOAT_MUL_BASE)
		n1 := n0.Mul(baseFloat)
		// x, _ := (&big.Float{}).SetString(strconv.FormatFloat(r, 'f', -1, 64))
		// x, _ := decimal.NewFromString(s)
		f1, _ := n1.BigFloat().Int64()
		convey.So(f1, convey.ShouldEqual, r1)
	})
}

func TestConvInt(t *testing.T) {
	convey.Convey("", t, func() {
		testConvInt("0", 0)
		testConvInt("123", 123)
		testConvInt("123.1", 123)
		testConvInt(".1", 0)
		testConvInt("0123", 123)
		testConvInt("0o123", 0123)
		testConvInt("0O123", 0123)
		testConvInt("+-0o123", -0123)
		testConvInt("123e+1", 1230)
		testConvInt("123e1", 1230)
		testConvInt("123e-1", 12)
		testConvInt("0x123", 0x123)
		testConvInt("--0X123", 0x123)
		testConvInt("-+-0x123", 0x123)
		testConvInt("-+-0xeFaB", 0xeFaB)
	})
}

func TestConvUint(t *testing.T) {
	convey.Convey("", t, func() {
		testConvUint("0", 0)
		testConvUint("123", 123)
		testConvUint("123.1", 123)
		testConvUint(".1", 0)
		testConvUint("0123", 123)
		testConvUint("0o123", 0123)
		// testConvUint("+-0123", uint64(int64(-0123)))
		testConvUint("123e+1", 1230)
		testConvUint("123e1", 1230)
		testConvUint("123e-1", 12)
		testConvUint("0x123", 0x123)
		testConvUint("--0X123", 0x123)
		testConvUint("-+-0x123", 0x123)
		testConvUint("-+-0xeFaB", 0xeFaB)
	})
}

func TestConvFloat(t *testing.T) {
	convey.Convey("", t, func() {
		testConvFloat("0", 0, 0*_TEST_FLOAT_MUL_BASE)
		testConvFloat("9.9", 9.9, 9.9*_TEST_FLOAT_MUL_BASE)
		testConvFloat("20.09", 20.09, 20.09*_TEST_FLOAT_MUL_BASE)
		testConvFloat("19.99", 19.99, 19.99*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0.99", 0.99, 0.99*_TEST_FLOAT_MUL_BASE)
		testConvFloat("2.32", 2.32, 2.32*_TEST_FLOAT_MUL_BASE)
		testConvFloat("123", 123, 123*_TEST_FLOAT_MUL_BASE)
		testConvFloat("123.1", 123.1, 123.1*_TEST_FLOAT_MUL_BASE)
		testConvFloat(".1", 0.1, 0.1*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0123", 123, 123*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0o123", 0123, 0123*_TEST_FLOAT_MUL_BASE)
		// testConvFloat("+-0123", uint64(int64(-0123)))
		testConvFloat("123e+1", 1230, 1230*_TEST_FLOAT_MUL_BASE)
		testConvFloat("123e1", 1230, 1230*_TEST_FLOAT_MUL_BASE)
		testConvFloat("123e-1", 12.3, 12.3*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0x123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		testConvFloat("--0X123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		testConvFloat("-+-0x123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		testConvFloat("-+-0xeFaB", 0xeFaB, 0xeFaB*_TEST_FLOAT_MUL_BASE)
		testConvFloat("---0xeFaB", -0xeFaB, -0xeFaB*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0.30129", 0.30129, 0.30129*_TEST_FLOAT_MUL_BASE)
		testConvFloat("0.30129e3", 301.29, 301.29*_TEST_FLOAT_MUL_BASE)
		testConvFloat("-0.30129E-3", -0.00030129, -30)
	})
}
