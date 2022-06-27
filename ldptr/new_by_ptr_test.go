/*
 * Copyright (C) distroy
 */

package ldptr

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewBoolByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewBoolByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewBoolByPtr(NewBool(true)), convey.ShouldResemble, NewBool(true))
	})
}

func TestNewByteByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewByteByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewByteByPtr(NewByte(1)), convey.ShouldResemble, NewByte(1))
	})
}

func TestNewRuneByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewRuneByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewRuneByPtr(NewRune(1)), convey.ShouldResemble, NewRune(1))
	})
}

func TestNewIntByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewIntByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewIntByPtr(NewInt(1)), convey.ShouldResemble, NewInt(1))
	})
}

func TestNewInt8ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewInt8ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewInt8ByPtr(NewInt8(1)), convey.ShouldResemble, NewInt8(1))
	})
}

func TestNewInt16ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewInt16ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewInt16ByPtr(NewInt16(1)), convey.ShouldResemble, NewInt16(1))
	})
}

func TestNewInt32ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewInt32ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewInt32ByPtr(NewInt32(1)), convey.ShouldResemble, NewInt32(1))
	})
}

func TestNewInt64ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewInt64ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewInt64ByPtr(NewInt64(1)), convey.ShouldResemble, NewInt64(1))
	})
}

func TestNewUintByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUintByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUintByPtr(NewUint(1)), convey.ShouldResemble, NewUint(1))
	})
}

func TestNewUint8ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUint8ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUint8ByPtr(NewUint8(1)), convey.ShouldResemble, NewUint8(1))
	})
}

func TestNewUint16ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUint16ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUint16ByPtr(NewUint16(1)), convey.ShouldResemble, NewUint16(1))
	})
}

func TestNewUint32ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUint32ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUint32ByPtr(NewUint32(1)), convey.ShouldResemble, NewUint32(1))
	})
}

func TestNewUint64ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUint64ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUint64ByPtr(NewUint64(1)), convey.ShouldResemble, NewUint64(1))
	})
}

func TestNewUintptrByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewUintptrByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewUintptrByPtr(NewUintptr(1)), convey.ShouldResemble, NewUintptr(1))
	})
}

func TestNewFloat32ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewFloat32ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewFloat32ByPtr(NewFloat32(1)), convey.ShouldResemble, NewFloat32(1))
	})
}

func TestNewFloat64ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewFloat64ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewFloat64ByPtr(NewFloat64(1)), convey.ShouldResemble, NewFloat64(1))
	})
}

func TestNewStringByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewStringByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewStringByPtr(NewString("1")), convey.ShouldResemble, NewString("1"))
	})
}

func TestNewComplex64ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewComplex64ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewComplex64ByPtr(NewComplex64(1)), convey.ShouldResemble, NewComplex64(1))
	})
}

func TestNewComplex128ByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewComplex128ByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewComplex128ByPtr(NewComplex128(1)), convey.ShouldResemble, NewComplex128(1))
	})
}

func TestNewTimeByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewTimeByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewTimeByPtr(NewTime(time.Unix(100, 0))), convey.ShouldResemble, NewTime(time.Unix(100, 0)))
	})
}

func TestNewDurationByPtr(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(NewDurationByPtr(nil), convey.ShouldEqual, nil)
		convey.So(NewDurationByPtr(NewDuration(100)), convey.ShouldResemble, NewDuration(100))
	})
}
