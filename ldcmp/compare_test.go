/*
 * Copyright (C) distroy
 */

package ldcmp

import (
	"math"
	"testing"
	"time"

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

func TestCompareBool(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareBool(false, false), convey.ShouldEqual, 0)
		convey.So(CompareBool(true, true), convey.ShouldEqual, 0)
		convey.So(CompareBool(false, true), convey.ShouldEqual, -1)
		convey.So(CompareBool(true, false), convey.ShouldEqual, 1)
	})
}

func TestCompareByte(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareByte(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareByte(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareByte(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareByte(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareRune(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareRune(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareRune(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareRune(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareRune(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareInt(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareInt(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareInt(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareInt(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt8(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareInt8(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareInt8(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareInt8(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareInt8(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt16(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareInt16(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareInt16(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareInt16(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareInt16(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareInt32(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareInt32(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareInt32(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareInt32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareInt64(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareInt64(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareInt64(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareInt64(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareUint(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareUint(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareUint(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareUint(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint8(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareUint8(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareUint8(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareUint8(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareUint8(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint16(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareUint16(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareUint16(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareUint16(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareUint16(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareUint32(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareUint32(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareUint32(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareUint32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareUint64(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareUint64(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareUint64(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareUint64(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareFloat32(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareFloat32(float32(math.NaN()), float32(math.NaN())), convey.ShouldEqual, 0)
		convey.So(CompareFloat32(0, float32(math.NaN())), convey.ShouldEqual, 1)
		convey.So(CompareFloat32(float32(math.NaN()), 0), convey.ShouldEqual, -1)

		convey.So(CompareFloat32(float32(math.Inf(-1)), float32(math.NaN())), convey.ShouldEqual, 1)
		convey.So(CompareFloat32(float32(math.NaN()), float32(math.Inf(-1))), convey.ShouldEqual, -1)

		convey.So(CompareFloat32(float32(math.Inf(1)), float32(math.Inf(1))), convey.ShouldEqual, 0)
		convey.So(CompareFloat32(float32(math.Inf(-1)), float32(math.Inf(-1))), convey.ShouldEqual, 0)

		convey.So(CompareFloat32(float32(math.Inf(-1)), float32(math.Inf(1))), convey.ShouldEqual, -1)
		convey.So(CompareFloat32(float32(math.Inf(1)), float32(math.Inf(-1))), convey.ShouldEqual, 1)

		convey.So(CompareFloat32(float32(math.Inf(-1)), 0), convey.ShouldEqual, -1)
		convey.So(CompareFloat32(0, float32(math.Inf(-1))), convey.ShouldEqual, 1)
		convey.So(CompareFloat32(float32(math.Inf(1)), 0), convey.ShouldEqual, 1)
		convey.So(CompareFloat32(0, float32(math.Inf(1))), convey.ShouldEqual, -1)

		convey.So(CompareFloat32(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareFloat32(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareFloat32(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareFloat32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareFloat64(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareFloat64(math.NaN(), math.NaN()), convey.ShouldEqual, 0)
		convey.So(CompareFloat64(0, math.NaN()), convey.ShouldEqual, 1)
		convey.So(CompareFloat64(math.NaN(), 0), convey.ShouldEqual, -1)

		convey.So(CompareFloat64(math.Inf(-1), math.NaN()), convey.ShouldEqual, 1)
		convey.So(CompareFloat64(math.NaN(), math.Inf(-1)), convey.ShouldEqual, -1)

		convey.So(CompareFloat64(math.Inf(1), math.Inf(1)), convey.ShouldEqual, 0)
		convey.So(CompareFloat64(math.Inf(-1), math.Inf(-1)), convey.ShouldEqual, 0)

		convey.So(CompareFloat64(math.Inf(-1), math.Inf(1)), convey.ShouldEqual, -1)
		convey.So(CompareFloat64(math.Inf(1), math.Inf(-1)), convey.ShouldEqual, 1)

		convey.So(CompareFloat64(math.Inf(-1), 0), convey.ShouldEqual, -1)
		convey.So(CompareFloat64(0, math.Inf(-1)), convey.ShouldEqual, 1)
		convey.So(CompareFloat64(math.Inf(1), 0), convey.ShouldEqual, 1)
		convey.So(CompareFloat64(0, math.Inf(1)), convey.ShouldEqual, -1)

		convey.So(CompareFloat64(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareFloat64(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareFloat64(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareFloat64(123, 0), convey.ShouldEqual, 1)
	})
}

func TestString(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareString(``, ""), convey.ShouldEqual, 0)
		convey.So(CompareString(`aaa`, "aaa"), convey.ShouldEqual, 0)

		convey.So(CompareString(`aaa`, ""), convey.ShouldEqual, 1)
		convey.So(CompareString(``, "aaa"), convey.ShouldEqual, -1)
		convey.So(CompareString(`abc`, "aaa"), convey.ShouldEqual, 1)
		convey.So(CompareString(`aa`, "aaa"), convey.ShouldEqual, -1)
	})
}

func TestCompareDuration(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareDuration(0, 0), convey.ShouldEqual, 0)
		convey.So(CompareDuration(123, 123), convey.ShouldEqual, 0)
		convey.So(CompareDuration(0, 123), convey.ShouldEqual, -1)
		convey.So(CompareDuration(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareTime(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.So(CompareTime(time.Unix(0, 0), time.Unix(0, 0)), convey.ShouldEqual, 0)
		convey.So(CompareTime(time.Unix(123, 0), time.Unix(123, 0)), convey.ShouldEqual, 0)
		convey.So(CompareTime(time.Unix(0, 0), time.Unix(123, 0)), convey.ShouldEqual, -1)
		convey.So(CompareTime(time.Unix(123, 0), time.Unix(0, 0)), convey.ShouldEqual, 1)
	})
}
