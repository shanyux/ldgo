/*
 * Copyright (C) distroy
 */

package ldmath

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSumInt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumInt(0), convey.ShouldEqual, 0)
		convey.So(SumInt(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumInt8(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumInt8(0), convey.ShouldEqual, 0)
		convey.So(SumInt8(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumInt16(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumInt16(0), convey.ShouldEqual, 0)
		convey.So(SumInt16(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumInt32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumInt32(0), convey.ShouldEqual, 0)
		convey.So(SumInt32(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumInt64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumInt64(0), convey.ShouldEqual, 0)
		convey.So(SumInt64(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumUint(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumUint(0), convey.ShouldEqual, 0)
		convey.So(SumUint(123, 237, 35), convey.ShouldEqual, int64(123+237+35))
	})
}

func TestSumUint8(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumUint8(0), convey.ShouldEqual, 0)
		convey.So(SumUint8(123, 237, 35), convey.ShouldEqual, int64(123+237+35))
	})
}

func TestSumUint16(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumUint16(0), convey.ShouldEqual, 0)
		convey.So(SumUint16(123, 237, 35), convey.ShouldEqual, int64(123+237+35))
	})
}

func TestSumUint32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumUint32(0), convey.ShouldEqual, 0)
		convey.So(SumUint32(123, 237, 35), convey.ShouldEqual, int64(123+237+35))
	})
}

func TestSumUint64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumUint64(0), convey.ShouldEqual, 0)
		convey.So(SumUint64(123, 237, 35), convey.ShouldEqual, int64(123+237+35))
	})
}

func TestSumFloat32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumFloat32(0), convey.ShouldEqual, 0)
		convey.So(SumFloat32(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}

func TestSumFloat64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(SumFloat64(0), convey.ShouldEqual, 0)
		convey.So(SumFloat64(123, -115, 35), convey.ShouldEqual, int64(123-115+35))
	})
}
