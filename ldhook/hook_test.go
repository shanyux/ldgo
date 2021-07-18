/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func testFuncForConvert(_ ...interface{}) int { return 1001 }

type testMethodForConvert struct{}

func (_ *testMethodForConvert) Test(_ ...interface{}) interface{} {
	return 2001
}

func TestHookConvert(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		patches := NewPatches()
		defer patches.Reset()
		defer patches.Reset()

		convey.Convey("FuncHook", func() {
			patches.Apply(FuncHook{
				Target: testFuncForConvert,
				Double: Values{int64(1101)},
			})
			convey.So(testFuncForConvert(1, ""), convey.ShouldEqual, 1101)
			convey.So(testFuncForConvert(1, ""), convey.ShouldEqual, 1101)
		})
		convey.Convey("MethodHook", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForConvert)(nil),
				Method: "Test",
				Double: Values{1701},
			})
			obj := &testMethodForConvert{}
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 1701)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 1701)
		})
	})
}
func testFuncForHook(_ ...interface{}) interface{} { return 1001 }

func TestFuncHook(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		patches := NewPatches()
		defer patches.Reset()

		convey.Convey("function", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: func(_ ...interface{}) interface{} {
					return 1101
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1101)
		})
		convey.Convey("[]OutputCell", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: []OutputCell{
					{
						Values: Values{(1201)},
					},
					{
						Times:  2,
						Values: Values{1202},
					},
					{
						Times:  1,
						Values: Values{1203},
					},
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1201)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1202)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1202)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1203)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1203)
		})
		convey.Convey("[]*OutputCell", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: []*OutputCell{
					{
						Times:  2,
						Values: Values{1301},
					},
					{
						Times:  0,
						Values: Values{1302},
					},
					{
						Times:  1,
						Values: Values{1303},
					},
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1301)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1301)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1302)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1303)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1303)
		})
		convey.Convey("OutputCell", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: OutputCell{
					Times:  2,
					Values: Values{1401},
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1401)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1401)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1401)
		})
		convey.Convey("*OutputCell", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: &OutputCell{
					Times:  2,
					Values: Values{1501},
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1501)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1501)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1501)
		})
		convey.Convey("[]Values", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: []Values{
					{1601},
					{1602},
					{1603},
				},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1601)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1602)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1603)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1603)
		})
		convey.Convey("Values", func() {
			patches.Apply(FuncHook{
				Target: testFuncForHook,
				Double: Values{1701},
			})
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1701)
			convey.So(testFuncForHook(1, ""), convey.ShouldEqual, 1701)
		})
	})
}

type testMethodForHook struct{}

func (_ *testMethodForHook) Test(_ ...interface{}) interface{} {
	return 2001
}

func TestMethodHook(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		patches := NewPatches()
		obj := &testMethodForHook{}
		defer patches.Reset()

		convey.Convey("function", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: func(_ *testMethodForHook, _ ...interface{}) interface{} {
					return 2101
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2101)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2101)
		})
		convey.Convey("[]OutputCell", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: []OutputCell{
					{
						Values: Values{2201},
					},
					{
						Times:  2,
						Values: Values{2202},
					},
					{
						Times:  1,
						Values: Values{2203},
					},
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2201)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2202)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2202)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2203)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2203)
		})
		convey.Convey("[]*OutputCell", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: []*OutputCell{
					{
						Times:  2,
						Values: Values{2301},
					},
					{
						Times:  0,
						Values: Values{2302},
					},
					{
						Times:  1,
						Values: Values{2303},
					},
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2301)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2301)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2302)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2303)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2303)
		})
		convey.Convey("OutputCell", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: OutputCell{
					Times:  2,
					Values: Values{2401},
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2401)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2401)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2401)
		})
		convey.Convey("*OutputCell", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: &OutputCell{
					Times:  2,
					Values: Values{2501},
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2501)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2501)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2501)
		})
		convey.Convey("[]Values", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: []Values{
					{2601},
					{2602},
					{2603},
				},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2601)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2602)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2603)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 2603)
		})
		convey.Convey("Values", func() {
			patches.Apply(MethodHook{
				Target: (*testMethodForHook)(nil),
				Method: "Test",
				Double: Values{1701},
			})
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 1701)
			convey.So(obj.Test(1, ""), convey.ShouldEqual, 1701)
		})
	})
}
