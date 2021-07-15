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

func test_ConvFloat(s string, r0 float64, r1 int64) {
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
		test_ConvFloat("9.9", 9.9, 9.9*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("20.09", 20.09, 20.09*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("19.99", 19.99, 19.99*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0.99", 0.99, 0.99*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("2.32", 2.32, 2.32*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("123", 123, 123*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("123.1", 123.1, 123.1*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat(".1", 0.1, 0.1*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0123", 123, 123*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0o123", 0123, 0123*_TEST_FLOAT_MUL_BASE)
		// testConvFloat("+-0123", uint64(int64(-0123)))
		test_ConvFloat("123e+1", 1230, 1230*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("123e1", 1230, 1230*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("123e-1", 12.3, 12.3*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0x123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("--0X123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("-+-0x123", 0x123, 0x123*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("-+-0xeFaB", 0xeFaB, 0xeFaB*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("---0xeFaB", -0xeFaB, -0xeFaB*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0.30129", 0.30129, 0.30129*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("0.30129e3", 301.29, 301.29*_TEST_FLOAT_MUL_BASE)
		test_ConvFloat("-0.30129E-3", -0.00030129, -30)
	})
}
