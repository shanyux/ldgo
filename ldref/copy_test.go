/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testCopyStruct struct {
	Id     int64  `copy:"id"`
	Ignore bool   `copy:"-"`
	Name   string `copy:""`
	Age    int64
	Female bool
}

func TestCopy(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("to **int", func() {
			convey.Convey("bool to **int", func() {
				var target *int = new(int)
				var source bool = true

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*target, convey.ShouldEqual, 1)
			})

			convey.Convey("**float to **int", func() {
				var (
					target *int     = new(int)
					source *float64 = new(float64)
				)
				*source = 100

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*target, convey.ShouldEqual, *source)
			})

			convey.Convey("**int to **int", func() {
				var (
					target *int = new(int)
					source *int = new(int)
				)
				*source = 100

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("nil to **int", func() {
				var (
					target *int
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, nil)
			})
		})

		convey.Convey("func", func() {
			convey.Convey("func to *func", func() {
				var (
					target func()
					source func() = func() {}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("func to *int", func() {
				var (
					target int    = 100
					source func() = func() {}
				)

				err := Copy(&target, source)
				convey.So(err.Error(), convey.ShouldEqual, "func() can not convert to int")
				convey.So(target, convey.ShouldEqual, 0)
			})
		})

		convey.Convey("number to *numbers", func() {
			convey.Convey("nil to *int", func() {
				var (
					target int
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 0)
			})

			convey.Convey("int to *int", func() {
				var (
					target int
					source int = 100
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("uint to *int", func() {
				var (
					target int
					source uint = 100
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("int to *uint", func() {
				var (
					target uint
					source int = 100
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("float to *bool", func() {
				var (
					target bool
					source float64 = 100
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source != 0)
			})
		})

		convey.Convey("number to *string", func() {
			convey.Convey("nil to *string", func() {
				var (
					target string
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "")
			})

			convey.Convey("bool to *string", func() {
				var (
					target string
					source bool = true
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "true")
			})

			convey.Convey("int to *string", func() {
				var (
					target string
					source int = 100
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})
		})

		convey.Convey("string to *number", func() {
			convey.Convey("string to *float", func() {
				var (
					target float64
					source string = "100"
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 100)
			})

			convey.Convey("string to *bool", func() {
				var (
					target bool
					source string = "100"
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, true)
			})
		})

		convey.Convey("string/[]byte/[]rune", func() {
			convey.Convey("[]byte to *string", func() {
				var (
					target string
					source []byte = []byte("100")
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})

			convey.Convey("[]rune to *string", func() {
				var (
					target string
					source []rune = []rune("100")
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})

			convey.Convey("string to *[]byte", func() {
				var (
					target []byte
					source string = "100"
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, []byte(source))
			})

			convey.Convey("string to *[]rune", func() {
				var (
					target []rune
					source string = "100"
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, []rune(source))
			})
		})

		convey.Convey("to *map", func() {
			convey.Convey("nil to *map[int64]string", func() {
				var (
					target map[int64]string
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldBeNil)
			})

			convey.Convey("map[int]string to *map[interface{}]interface{}", func() {
				var (
					target map[interface{}]interface{}
					source map[int]string = map[int]string{
						1: "abc",
						2: "xyz",
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[interface{}]interface{}{
					1: "abc",
					2: "xyz",
				})
			})

			convey.Convey("map[int]string to *map[int64][]byte", func() {
				var (
					target map[int64][]byte
					source map[int]string = map[int]string{
						1: "abc",
						2: "xyz",
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[int64][]byte{
					1: []byte("abc"),
					2: []byte("xyz"),
				})
			})

			convey.Convey("*struct to *map[string]string", func() {
				var (
					target map[string]string
					source *testCopyStruct = &testCopyStruct{
						Id:     123,
						Ignore: true,
						Name:   "abc",
						Age:    100,
						Female: true,
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[string]string{
					`id`:     "123",
					`Name`:   "abc",
					`Age`:    "100",
					`Female`: "true",
				})
			})

			convey.Convey("[]string to *map[string]string", func() {
				var (
					target map[string]string
					source []string = []string{
						"abc", "xyz", "123", "zzz",
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[string]string{
					`0`: "abc",
					`1`: "xyz",
					`2`: "123",
					`3`: "zzz",
				})
			})
		})

		convey.Convey("to *slice", func() {
			convey.Convey("nil to *[]string", func() {
				var (
					target []string
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldBeNil)
			})

			convey.Convey("nil to *[2]string", func() {
				var (
					target [2]string
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, [2]string{})
			})

			convey.Convey("[]int64 to *[]complex128", func() {
				var (
					target []complex128
					source []int64 = []int64{
						2, -100, 356,
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, []complex128{
					2, -100, 356,
				})
			})

		})

		convey.Convey("to *struct/**struct", func() {
			convey.Convey("nil to *struct", func() {
				var (
					target testCopyStruct
					source interface{}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{})
			})

			convey.Convey("*struct to *struct", func() {
				var (
					target testCopyStruct
					source *testCopyStruct = &testCopyStruct{
						Id:     100,
						Ignore: true,
						Name:   "abc",
						Age:    23,
						Female: true,
					}
				)

				err := Copy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{
					Id:     100,
					Name:   "abc",
					Age:    23,
					Female: true,
				})
			})

			convey.Convey("**struct to *struct", func() {
				var (
					target testCopyStruct
					source *testCopyStruct = &testCopyStruct{
						Id:     100,
						Ignore: true,
						Name:   "abc",
						Age:    23,
						Female: true,
					}
				)

				err := Copy(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{
					Id:     100,
					Name:   "abc",
					Age:    23,
					Female: true,
				})
			})

			convey.Convey("**struct to **struct", func() {
				var (
					target *testCopyStruct
					source *testCopyStruct = &testCopyStruct{
						Id:     100,
						Ignore: true,
						Name:   "abc",
						Age:    23,
						Female: true,
					}
				)

				err := Copy(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("map[string]string to *struct", func() {
				var (
					target testCopyStruct
					source map[string]string = map[string]string{
						`id`:     "100",
						`Ignore`: "true",
						`Name`:   "abc",
						`Age`:    "23",
						`Female`: "true",
					}
				)

				err := Copy(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{
					Id:     100,
					Name:   "abc",
					Age:    23,
					Female: true,
				})
			})

		})
	})
}
