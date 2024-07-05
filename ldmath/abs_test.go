/*
 * Copyright (C) distroy
 */

package ldmath

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAbsInt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsInt(0), convey.ShouldEqual, 0)
		convey.So(AbsInt(100), convey.ShouldEqual, 100)
		convey.So(AbsInt(-100), convey.ShouldEqual, 100)
	})
}

func TestAbsInt8(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsInt8(0), convey.ShouldEqual, 0)
		convey.So(AbsInt8(100), convey.ShouldEqual, 100)
		convey.So(AbsInt8(-100), convey.ShouldEqual, 100)
	})
}

func TestAbsInt16(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsInt16(0), convey.ShouldEqual, 0)
		convey.So(AbsInt16(100), convey.ShouldEqual, 100)
		convey.So(AbsInt16(-100), convey.ShouldEqual, 100)
	})
}

func TestAbsInt32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsInt32(0), convey.ShouldEqual, 0)
		convey.So(AbsInt32(100), convey.ShouldEqual, 100)
		convey.So(AbsInt32(-100), convey.ShouldEqual, 100)
	})
}

func TestAbsInt64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsInt64(0), convey.ShouldEqual, 0)
		convey.So(AbsInt64(100), convey.ShouldEqual, 100)
		convey.So(AbsInt64(-100), convey.ShouldEqual, 100)
	})
}

func TestAbsFloat32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsFloat32(0), convey.ShouldEqual, 0)
		convey.So(AbsFloat32(100), convey.ShouldEqual, 100)
		convey.So(AbsFloat32(-100), convey.ShouldEqual, 100)

		convey.So(IsNaN32(AbsFloat32(0)), convey.ShouldResemble, false)
		convey.So(IsNaN32(AbsFloat32(NaN32())), convey.ShouldResemble, true)

		convey.So(AbsFloat32(Inf32(1)), convey.ShouldEqual, Inf64(1))
		convey.So(AbsFloat32(Inf32(-1)), convey.ShouldEqual, Inf64(1))
	})
}

func TestAbsFloat64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(AbsFloat64(0), convey.ShouldEqual, 0)
		convey.So(AbsFloat64(100), convey.ShouldEqual, 100)
		convey.So(AbsFloat64(-100), convey.ShouldEqual, 100)

		convey.So(IsNaN64(AbsFloat64(0)), convey.ShouldResemble, false)
		convey.So(IsNaN64(AbsFloat64(NaN64())), convey.ShouldResemble, true)

		convey.So(AbsFloat64(Inf64(1)), convey.ShouldEqual, Inf64(1))
		convey.So(AbsFloat64(Inf64(-1)), convey.ShouldEqual, Inf64(1))
	})
}

func TestAbs(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("int", func(c convey.C) {
			c.So(Abs(0), convey.ShouldEqual, 0)
			c.So(Abs(100), convey.ShouldEqual, 100)
			c.So(Abs(-100), convey.ShouldEqual, 100)
		})
		c.Convey("float32", func(c convey.C) {
			c.So(Abs[float32](0), convey.ShouldEqual, 0)
			c.So(Abs[float32](100), convey.ShouldEqual, 100)
			c.So(Abs[float32](-100), convey.ShouldEqual, 100)

			c.So(IsNaN(Abs[float32](0)), convey.ShouldResemble, false)
			c.So(IsNaN(Abs(NaN32())), convey.ShouldResemble, true)

			c.So(IsInf(Abs(Inf32(1)), 1), convey.ShouldEqual, true)
			c.So(IsInf(Abs(Inf32(-1)), 1), convey.ShouldEqual, true)
		})
		c.Convey("float64", func(c convey.C) {
			c.So(Abs[float64](0), convey.ShouldEqual, 0)
			c.So(Abs[float64](100), convey.ShouldEqual, 100)
			c.So(Abs[float64](-100), convey.ShouldEqual, 100)

			c.So(IsNaN(AbsFloat64(0)), convey.ShouldResemble, false)
			c.So(IsNaN(Abs(NaN64())), convey.ShouldResemble, true)

			c.So(IsInf(Abs(Inf64(1)), -1), convey.ShouldEqual, false)
			c.So(IsInf(Abs(Inf64(-1)), -1), convey.ShouldEqual, false)
		})
	})
}
