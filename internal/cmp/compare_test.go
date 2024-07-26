/*
 * Copyright (C) distroy
 */

package cmp

import (
	"math"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestCompareInterface(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("aType != bType", func(c convey.C) {
			c.So(CompareInterface(nil, false), convey.ShouldEqual, -1)
			c.So(CompareInterface(true, 0), convey.ShouldEqual, -1)

			c.So(CompareInterface((*int)(nil), (*uint)(nil)), convey.ShouldEqual, -1)
			c.So(CompareInterface((*int8)(nil), (*uint)(nil)), convey.ShouldEqual, -1)
			c.So(CompareInterface((*int)(nil), (*int)(nil)), convey.ShouldEqual, 0)
			c.So(CompareInterface(0, []int{}), convey.ShouldEqual, -1)
		})

		c.Convey("bool", func(c convey.C) {
			c.So(CompareInterface(false, true), convey.ShouldEqual, -1)
			c.So(CompareInterface(true, false), convey.ShouldEqual, 1)
			c.So(CompareInterface(false, false), convey.ShouldEqual, 0)
			c.So(CompareInterface(true, true), convey.ShouldEqual, 0)
		})

		c.Convey("int", func(c convey.C) {
			c.So(CompareInterface(uint64(99), 100), convey.ShouldEqual, -1)
			c.So(CompareInterface(uint64(math.MaxInt64+1), 100), convey.ShouldEqual, 1)
			c.So(CompareInterface(200, int64(100)), convey.ShouldEqual, 1)
			c.So(CompareInterface(200, uint64(200)), convey.ShouldEqual, 0)
			c.So(CompareInterface(int64(-200), 100), convey.ShouldEqual, -1)
			c.So(CompareInterface(int64(-200), uint(100)), convey.ShouldEqual, -1)
		})

		c.Convey("float", func(c convey.C) {
			c.So(CompareInterface(float64(100.0), float32(100.0)), convey.ShouldEqual, 0)
			c.So(CompareInterface(99.0, 100.0), convey.ShouldEqual, -1)
			c.So(CompareInterface(99.0, math.NaN()), convey.ShouldEqual, 1)
			c.So(CompareInterface(-99.0, math.NaN()), convey.ShouldEqual, 1)
			c.So(CompareInterface(-99.0, float32(math.NaN())), convey.ShouldEqual, 1)
			c.So(CompareInterface(float32(math.NaN()), 100.0), convey.ShouldEqual, -1)

			c.So(CompareInterface(math.NaN(), math.NaN()), convey.ShouldEqual, 0)
			c.So(CompareInterface(float32(math.NaN()), float32(math.NaN())), convey.ShouldEqual, 0)

			c.So(CompareInterface(-99.0, math.Inf(1)), convey.ShouldEqual, -1)
			c.So(CompareInterface(-99.0, math.Inf(-1)), convey.ShouldEqual, 1)

			c.So(CompareInterface(float64(4503599627370496.1), float64(4503599627370496)), convey.ShouldEqual, 0) // exceeds the precision of float64

			c.Convey("int", func(c convey.C) {
				c.So(CompareInterface(float64(4503599627370496), int64(4503599627370496)), convey.ShouldEqual, 0)    //
				c.So(CompareInterface(float64(4503599627370496.1), int64(4503599627370496)), convey.ShouldEqual, 0)  // exceeds the precision of float64
				c.So(CompareInterface(float64(503599627370496.1), int64(503599627370496)), convey.ShouldEqual, 1)    //
				c.So(CompareInterface(float64(45035996273704960), int64(45035996273704961)), convey.ShouldEqual, -1) //
				c.So(CompareInterface(float64(45035996273704961), int64(45035996273704960)), convey.ShouldEqual, 0)  // exceeds the precision of float64
				c.So(CompareInterface(float64(math.MaxInt64)*2, int64(45035996273704960)), convey.ShouldEqual, 1)    //

				c.So(CompareInterface(float64(-4503599627370496), int64(-4503599627370496)), convey.ShouldEqual, 0)   //
				c.So(CompareInterface(float64(-4503599627370496.1), int64(-4503599627370496)), convey.ShouldEqual, 0) // exceeds the precision of float64
				c.So(CompareInterface(float64(-503599627370496.1), int64(-503599627370496)), convey.ShouldEqual, -1)  //
				c.So(CompareInterface(float64(-45035996273704960), int64(-45035996273704961)), convey.ShouldEqual, 1) //
				c.So(CompareInterface(float64(-45035996273704961), int64(-45035996273704960)), convey.ShouldEqual, 0) // exceeds the precision of float64
				c.So(CompareInterface(float64(math.MinInt64)*2, int64(45035996273704960)), convey.ShouldEqual, -1)    //
			})

			c.Convey("uint", func(c convey.C) {
				c.So(CompareInterface(float64(4503599627370496), uint64(4503599627370496)), convey.ShouldEqual, 0)    //
				c.So(CompareInterface(float64(4503599627370496.1), uint64(4503599627370496)), convey.ShouldEqual, 0)  // exceeds the precision of float64
				c.So(CompareInterface(float64(503599627370496.1), uint64(503599627370496)), convey.ShouldEqual, 1)    //
				c.So(CompareInterface(float64(45035996273704960), uint64(45035996273704961)), convey.ShouldEqual, -1) //
				c.So(CompareInterface(float64(45035996273704961), uint64(45035996273704960)), convey.ShouldEqual, 0)  // exceeds the precision of float64
				c.So(CompareInterface(float64(math.MaxUint)*2, uint64(45035996273704960)), convey.ShouldEqual, 1)     //
				c.So(CompareInterface(float64(-1), uint64(0)), convey.ShouldEqual, -1)                                //
			})
		})

		c.Convey("number", func(c convey.C) {
			c.So(CompareInterface(99, float64(100)), convey.ShouldEqual, -1)
			c.So(CompareInterface(99, float64(99)), convey.ShouldEqual, 0)
		})

		c.Convey("complex", func(c convey.C) {
			c.So(CompareInterface(complex(100, 200), complex64(complex(100, 200))), convey.ShouldEqual, 0)

			c.So(CompareInterface(complex(100, 200), complex(11, 300)), convey.ShouldEqual, 1)
			c.So(CompareInterface(complex(100, 200), complex(111, -300)), convey.ShouldEqual, -1)
			c.So(CompareInterface(complex(100, 200), complex(100, 300)), convey.ShouldEqual, -1)
			c.So(CompareInterface(complex(100, 200), complex(100, 150)), convey.ShouldEqual, 1)
		})

		c.Convey("string", func(c convey.C) {
			c.So(CompareInterface("", `abc`), convey.ShouldEqual, -1)
			c.So(CompareInterface("aaa", `a`), convey.ShouldEqual, 1)
			c.So(CompareInterface("bbb", `aaaaaa`), convey.ShouldEqual, 1)
		})

		c.Convey("map", func(c convey.C) {
			c.So(CompareInterface(map[int]int{0: 0}, map[interface{}]int{0: 0}), convey.ShouldEqual, -1)
			c.So(CompareInterface(map[int]int{0: 0}, map[int]int{0: 0}), convey.ShouldEqual, 0)
			c.So(CompareInterface(map[int]int{0: 0}, map[int]int{}), convey.ShouldEqual, 1)
			c.So(CompareInterface(map[int]int{0: 0}, map[int]int{1: 0}), convey.ShouldEqual, 1)
			c.So(CompareInterface(map[int]int{1: 1}, map[int]int{1: 0}), convey.ShouldEqual, 1)
			c.So(CompareInterface(map[int]int{1: 1}, map[int]int{0: 0, 1: 0}), convey.ShouldEqual, -1)
			c.So(CompareInterface(map[int]int{1: 1}, map[int]int{1: 1, 2: 0}), convey.ShouldEqual, -1)
			c.So(CompareInterface(map[int]int{0: 0, 1: 0}, map[int]int{0: 0}), convey.ShouldEqual, 1)
		})

		c.Convey("slice", func(c convey.C) {
			c.So(CompareInterface(
				[]interface{}{100, uint(200), float32(300)},
				[]interface{}{100, float64(200), ""},
			), convey.ShouldEqual, -1)

			c.So(CompareInterface(
				[]interface{}{100, uint(200), ""},
				[]interface{}{100, float64(200), ""},
			), convey.ShouldEqual, 0)
		})
	})
}

func TestCompareBool(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareBool(false, false), convey.ShouldEqual, 0)
		c.So(CompareBool(true, true), convey.ShouldEqual, 0)
		c.So(CompareBool(false, true), convey.ShouldEqual, -1)
		c.So(CompareBool(true, false), convey.ShouldEqual, 1)
	})
}

func TestCompareByte(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareByte(0, 0), convey.ShouldEqual, 0)
		c.So(CompareByte(123, 123), convey.ShouldEqual, 0)
		c.So(CompareByte(0, 123), convey.ShouldEqual, -1)
		c.So(CompareByte(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareRune(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareRune(0, 0), convey.ShouldEqual, 0)
		c.So(CompareRune(123, 123), convey.ShouldEqual, 0)
		c.So(CompareRune(0, 123), convey.ShouldEqual, -1)
		c.So(CompareRune(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareInt(0, 0), convey.ShouldEqual, 0)
		c.So(CompareInt(123, 123), convey.ShouldEqual, 0)
		c.So(CompareInt(0, 123), convey.ShouldEqual, -1)
		c.So(CompareInt(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt8(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareInt8(0, 0), convey.ShouldEqual, 0)
		c.So(CompareInt8(123, 123), convey.ShouldEqual, 0)
		c.So(CompareInt8(0, 123), convey.ShouldEqual, -1)
		c.So(CompareInt8(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt16(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareInt16(0, 0), convey.ShouldEqual, 0)
		c.So(CompareInt16(123, 123), convey.ShouldEqual, 0)
		c.So(CompareInt16(0, 123), convey.ShouldEqual, -1)
		c.So(CompareInt16(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt32(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareInt32(0, 0), convey.ShouldEqual, 0)
		c.So(CompareInt32(123, 123), convey.ShouldEqual, 0)
		c.So(CompareInt32(0, 123), convey.ShouldEqual, -1)
		c.So(CompareInt32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareInt64(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareInt64(0, 0), convey.ShouldEqual, 0)
		c.So(CompareInt64(123, 123), convey.ShouldEqual, 0)
		c.So(CompareInt64(0, 123), convey.ShouldEqual, -1)
		c.So(CompareInt64(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareUint(0, 0), convey.ShouldEqual, 0)
		c.So(CompareUint(123, 123), convey.ShouldEqual, 0)
		c.So(CompareUint(0, 123), convey.ShouldEqual, -1)
		c.So(CompareUint(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint8(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareUint8(0, 0), convey.ShouldEqual, 0)
		c.So(CompareUint8(123, 123), convey.ShouldEqual, 0)
		c.So(CompareUint8(0, 123), convey.ShouldEqual, -1)
		c.So(CompareUint8(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint16(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareUint16(0, 0), convey.ShouldEqual, 0)
		c.So(CompareUint16(123, 123), convey.ShouldEqual, 0)
		c.So(CompareUint16(0, 123), convey.ShouldEqual, -1)
		c.So(CompareUint16(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint32(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareUint32(0, 0), convey.ShouldEqual, 0)
		c.So(CompareUint32(123, 123), convey.ShouldEqual, 0)
		c.So(CompareUint32(0, 123), convey.ShouldEqual, -1)
		c.So(CompareUint32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareUint64(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareUint64(0, 0), convey.ShouldEqual, 0)
		c.So(CompareUint64(123, 123), convey.ShouldEqual, 0)
		c.So(CompareUint64(0, 123), convey.ShouldEqual, -1)
		c.So(CompareUint64(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareFloat32(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareFloat32(float32(math.NaN()), float32(math.NaN())), convey.ShouldEqual, 0)
		c.So(CompareFloat32(0, float32(math.NaN())), convey.ShouldEqual, 1)
		c.So(CompareFloat32(float32(math.NaN()), 0), convey.ShouldEqual, -1)

		c.So(CompareFloat32(float32(math.Inf(-1)), float32(math.NaN())), convey.ShouldEqual, 1)
		c.So(CompareFloat32(float32(math.NaN()), float32(math.Inf(-1))), convey.ShouldEqual, -1)

		c.So(CompareFloat32(float32(math.Inf(1)), float32(math.Inf(1))), convey.ShouldEqual, 0)
		c.So(CompareFloat32(float32(math.Inf(-1)), float32(math.Inf(-1))), convey.ShouldEqual, 0)

		c.So(CompareFloat32(float32(math.Inf(-1)), float32(math.Inf(1))), convey.ShouldEqual, -1)
		c.So(CompareFloat32(float32(math.Inf(1)), float32(math.Inf(-1))), convey.ShouldEqual, 1)

		c.So(CompareFloat32(float32(math.Inf(-1)), 0), convey.ShouldEqual, -1)
		c.So(CompareFloat32(0, float32(math.Inf(-1))), convey.ShouldEqual, 1)
		c.So(CompareFloat32(float32(math.Inf(1)), 0), convey.ShouldEqual, 1)
		c.So(CompareFloat32(0, float32(math.Inf(1))), convey.ShouldEqual, -1)

		c.So(CompareFloat32(0, 0), convey.ShouldEqual, 0)
		c.So(CompareFloat32(123, 123), convey.ShouldEqual, 0)
		c.So(CompareFloat32(0, 123), convey.ShouldEqual, -1)
		c.So(CompareFloat32(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareFloat64(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareFloat64(math.NaN(), math.NaN()), convey.ShouldEqual, 0)
		c.So(CompareFloat64(0, math.NaN()), convey.ShouldEqual, 1)
		c.So(CompareFloat64(math.NaN(), 0), convey.ShouldEqual, -1)

		c.So(CompareFloat64(math.Inf(-1), math.NaN()), convey.ShouldEqual, 1)
		c.So(CompareFloat64(math.NaN(), math.Inf(-1)), convey.ShouldEqual, -1)

		c.So(CompareFloat64(math.Inf(1), math.Inf(1)), convey.ShouldEqual, 0)
		c.So(CompareFloat64(math.Inf(-1), math.Inf(-1)), convey.ShouldEqual, 0)

		c.So(CompareFloat64(math.Inf(-1), math.Inf(1)), convey.ShouldEqual, -1)
		c.So(CompareFloat64(math.Inf(1), math.Inf(-1)), convey.ShouldEqual, 1)

		c.So(CompareFloat64(math.Inf(-1), 0), convey.ShouldEqual, -1)
		c.So(CompareFloat64(0, math.Inf(-1)), convey.ShouldEqual, 1)
		c.So(CompareFloat64(math.Inf(1), 0), convey.ShouldEqual, 1)
		c.So(CompareFloat64(0, math.Inf(1)), convey.ShouldEqual, -1)

		c.So(CompareFloat64(0, 0), convey.ShouldEqual, 0)
		c.So(CompareFloat64(123, 123), convey.ShouldEqual, 0)
		c.So(CompareFloat64(0, 123), convey.ShouldEqual, -1)
		c.So(CompareFloat64(123, 0), convey.ShouldEqual, 1)

		c.So(CompareFloat64(4503599627370496.1, 4503599627370496), convey.ShouldEqual, 0) // exceeds the precision of float64
	})
}

func TestString(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareString(``, ""), convey.ShouldEqual, 0)
		c.So(CompareString(`aaa`, "aaa"), convey.ShouldEqual, 0)

		c.So(CompareString(`aaa`, ""), convey.ShouldEqual, 1)
		c.So(CompareString(``, "aaa"), convey.ShouldEqual, -1)
		c.So(CompareString(`abc`, "aaa"), convey.ShouldEqual, 1)
		c.So(CompareString(`aa`, "aaa"), convey.ShouldEqual, -1)
	})
}

func TestCompareDuration(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareDuration(0, 0), convey.ShouldEqual, 0)
		c.So(CompareDuration(123, 123), convey.ShouldEqual, 0)
		c.So(CompareDuration(0, 123), convey.ShouldEqual, -1)
		c.So(CompareDuration(123, 0), convey.ShouldEqual, 1)
	})
}

func TestCompareTime(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(CompareTime(time.Unix(0, 0), time.Unix(0, 0)), convey.ShouldEqual, 0)
		c.So(CompareTime(time.Unix(123, 0), time.Unix(123, 0)), convey.ShouldEqual, 0)
		c.So(CompareTime(time.Unix(0, 0), time.Unix(123, 0)), convey.ShouldEqual, -1)
		c.So(CompareTime(time.Unix(123, 0), time.Unix(0, 0)), convey.ShouldEqual, 1)
	})
}
