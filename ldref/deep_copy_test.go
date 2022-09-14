/*
 * Copyright (C) distroy
 */

package ldref

import (
	"reflect"
	"sort"
	"testing"

	"github.com/distroy/ldgo/lderr"
	"github.com/smartystreets/goconvey/convey"
)

func TestDeepCopy(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey("int to int", func() {
			var (
				target int
				source int = 100
			)

			err := DeepCopy(target, source)
			convey.So(err, convey.ShouldEqual, lderr.ErrReflectTargetNotPtr)
			convey.So(target, convey.ShouldEqual, 0)
		})

		convey.Convey("to reflect.Value", func() {
			convey.Convey("int to reflect.Value - int (can not addr)", func() {
				var (
					target reflect.Value = reflect.ValueOf(int(0))
					source int           = 100
				)

				err := DeepCopy(target, source)
				convey.So(err, convey.ShouldEqual, lderr.ErrReflectTargetNotPtr)
				convey.So(target.Interface(), convey.ShouldEqual, 0)
			})

			convey.Convey("int to reflect.Value - *int", func() {
				var (
					xxxx   *int          = new(int)
					target reflect.Value = reflect.ValueOf(xxxx)
					source int           = 100
				)

				err := DeepCopy(target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target.Interface(), convey.ShouldResemble, &source)
			})

			convey.Convey("int to reflect.Value - *int (can not addr)", func() {
				var (
					xxxx   *int
					target reflect.Value = reflect.ValueOf(xxxx)
					source int           = 100
				)

				err := DeepCopy(target, source)
				// convey.So(err, convey.ShouldBeNil)
				// convey.So(target.Interface(), convey.ShouldEqual, 100)
				convey.So(err, convey.ShouldEqual, lderr.ErrReflectTargetNilPtr)
				convey.So(target.Interface(), convey.ShouldBeNil)
			})

			convey.Convey("int to reflect.Value - *int (elem of **int)", func() {
				var (
					xxxx   **int         = new(*int)
					target reflect.Value = reflect.ValueOf(xxxx).Elem()
					source int           = 100
				)

				err := DeepCopy(target, source)
				// convey.So(err, convey.ShouldBeNil)
				// convey.So(target.Interface(), convey.ShouldEqual, 100)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target.Interface(), convey.ShouldResemble, &source)
			})

			convey.Convey("int to reflect.Value - **int", func() {
				var (
					xxxx   *int
					target reflect.Value = reflect.ValueOf(&xxxx)
					source int           = 100
				)

				err := DeepCopy(target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(xxxx, convey.ShouldResemble, &source)
			})
		})

		convey.Convey("to *interface{}", func() {
			convey.Convey("nil to *interface{}", func() {
				var (
					target interface{}
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, nil)
			})

			convey.Convey("int to *interface{}", func() {
				var (
					target interface{}
					source int = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("*int to *interface{}", func() {
				var (
					target interface{}
					source *int = new(int)
				)
				*source = 100

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldNotEqual, source)
				convey.So(target, convey.ShouldResemble, source)
			})

			convey.Convey("**int to *interface{}", func() {
				var (
					target interface{}
					xxxx0  int   = 100
					xxxx1  *int  = &xxxx0
					source **int = &xxxx1
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldNotEqual, source)
				convey.So(target, convey.ShouldResemble, *source)
			})
		})

		convey.Convey("to **int", func() {
			convey.Convey("bool to **int", func() {
				var target *int = new(int)
				var source bool = true

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*target, convey.ShouldEqual, 1)
			})

			convey.Convey("**float to **int", func() {
				var (
					target *int     = new(int)
					source *float64 = new(float64)
				)
				*source = 100

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(*target, convey.ShouldEqual, *source)
			})

			convey.Convey("**int to **int", func() {
				var (
					target *int = new(int)
					source *int = new(int)
				)
				*source = 100

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, source)
				convey.So(target, convey.ShouldNotEqual, source)
			})

			convey.Convey("nil to **int", func() {
				var (
					target *int
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldNotEqual, nil)
				convey.So(*target, convey.ShouldEqual, 0)
			})
		})

		convey.Convey("func", func() {
			convey.Convey("func to *func", func() {
				var (
					target func()
					source func() = func() {}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source)
			})

			convey.Convey("func to *int", func() {
				var (
					target int    = 100
					source func() = func() {}
				)

				err := DeepCopy(&target, source)
				convey.So(err.Error(), convey.ShouldEqual, "func() can not copy to *int")
				convey.So(target, convey.ShouldEqual, 0)
			})

			convey.Convey("func to *string", func() {
				var (
					target string
					source func(interface{}) bool = IsZero
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "github.com/distroy/ldgo/ldref.IsZero")
			})
		})

		convey.Convey("to *bool", func() {
			convey.Convey("nil to *bool", func() {
				var (
					target bool
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, false)
			})

			convey.Convey("float to *bool", func() {
				var (
					target bool
					source float64 = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source != 0)
			})
			convey.Convey("int to *bool", func() {
				var (
					target bool
					source int = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source != 0)
			})
			convey.Convey("uint to *bool", func() {
				var (
					target bool
					source uint = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source != 0)
			})
			convey.Convey("complex to *bool", func() {
				var (
					target bool
					source complex128 = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, source != 0)
			})
		})

		convey.Convey("number to *numbers", func() {
			convey.Convey("to *int", func() {
				convey.Convey("nil to *int", func() {
					var (
						target int
						source interface{}
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, 0)
				})

				convey.Convey("int to *int", func() {
					var (
						target int
						source int = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("uint to *int", func() {
					var (
						target int
						source uint = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("float64 to *int", func() {
					var (
						target int
						source float64 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("complex128 to *int", func() {
					var (
						target int
						source complex128 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})
			})
			convey.Convey("to *uint", func() {
				convey.Convey("nil to *uint", func() {
					var (
						target uint
						source interface{}
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, 0)
				})

				convey.Convey("int to *uint", func() {
					var (
						target uint
						source int = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("uint to *uint", func() {
					var (
						target uint
						source uint = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("float64 to *uint", func() {
					var (
						target uint
						source float64 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("complex128 to *uint", func() {
					var (
						target uint
						source complex128 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})
			})
			convey.Convey("to *float64", func() {
				convey.Convey("nil to *float64", func() {
					var (
						target float64
						source interface{}
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, 0)
				})

				convey.Convey("int to *float64", func() {
					var (
						target float64
						source int = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("uint to *float64", func() {
					var (
						target float64
						source uint = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("float64 to *float64", func() {
					var (
						target float64
						source float64 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("complex128 to *float64", func() {
					var (
						target float64
						source complex128 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})
			})
			convey.Convey("to *complex128", func() {
				convey.Convey("nil to *complex128", func() {
					var (
						target complex128
						source interface{}
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, 0)
				})

				convey.Convey("int to *complex128", func() {
					var (
						target complex128
						source int = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("uint to *complex128", func() {
					var (
						target complex128
						source uint = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("float64 to *complex128", func() {
					var (
						target complex128
						source float64 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})

				convey.Convey("complex128 to *complex128", func() {
					var (
						target complex128
						source complex128 = 100
					)

					err := DeepCopy(&target, source)
					convey.So(err, convey.ShouldBeNil)
					convey.So(target, convey.ShouldEqual, source)
				})
			})
		})

		convey.Convey("number to *string", func() {
			convey.Convey("nil to *string", func() {
				var (
					target string
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "")
			})

			convey.Convey("bool to *string", func() {
				var (
					target string
					source bool = true
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "true")
			})

			convey.Convey("int to *string", func() {
				var (
					target string
					source int = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})
			convey.Convey("uint to *string", func() {
				var (
					target string
					source uint = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})
			convey.Convey("float64 to *string", func() {
				var (
					target string
					source float64 = 100
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})

			// convey.Convey("complex128 to *string", func() {
			// 	var (
			// 		target string
			// 		source complex128 = 100
			// 	)
			//
			// 	err := DeepCopy(&target, source)
			// 	convey.So(err, convey.ShouldBeNil)
			// 	convey.So(target, convey.ShouldEqual, "(100+0i)")
			// })
		})

		convey.Convey("string to *number", func() {
			convey.Convey("string to *float", func() {
				var (
					target float64
					source string = "100"
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 100)
			})
			convey.Convey("string to *bool", func() {
				var (
					target bool
					source string = "100"
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, true)
			})
			convey.Convey("string to *int", func() {
				var (
					target int
					source string = "100"
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 100)
			})
			convey.Convey("string to *uint", func() {
				var (
					target uint
					source string = "100"
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, 100)
			})

			// convey.Convey("string to *complex128", func() {
			// 	var (
			// 		target complex128
			// 		source string = "100"
			// 	)
			//
			// 	err := DeepCopy(&target, source)
			// 	convey.So(err, convey.ShouldBeNil)
			// 	convey.So(target, convey.ShouldEqual, complex128(100))
			// })
		})

		convey.Convey("string/[]byte/[]rune", func() {
			convey.Convey("[]byte to *string", func() {
				var (
					target string
					source []byte = []byte("100")
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})

			convey.Convey("[]rune to *string", func() {
				var (
					target string
					source []rune = []rune("100")
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldEqual, "100")
			})

			convey.Convey("string to *[]byte", func() {
				var (
					target []byte
					source string = "100"
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, []byte(source))
			})

			convey.Convey("string to *[]rune", func() {
				var (
					target []rune
					source string = "100"
				)

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[string]string{
					`0`: "abc",
					`1`: "xyz",
					`2`: "123",
					`3`: "zzz",
				})
			})

			convey.Convey("[]string to *map[string]struct{}", func() {
				var (
					target map[string]struct{}
					source []string = []string{
						"abc", "xyz", "123", "zzz",
					}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, map[string]struct{}{
					"abc": {},
					"xyz": {},
					"123": {},
					"zzz": {},
				})
			})
		})

		convey.Convey("to *slice", func() {
			convey.Convey("nil to *[]string", func() {
				var (
					target []string
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldBeNil)
			})

			convey.Convey("nil to *[2]string", func() {
				var (
					target [2]string
					source interface{}
				)

				err := DeepCopy(&target, source)
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

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, []complex128{
					2, -100, 356,
				})
			})

			convey.Convey("map[string]struct{} to *[]string", func() {
				var (
					target []string
					source map[string]struct{} = map[string]struct{}{
						"abc": {},
						"xyz": {},
						"AAA": {},
						"123": {},
					}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)

				sort.Strings(target)
				convey.So(target, convey.ShouldResemble, []string{
					"123", "AAA", "abc", "xyz",
				})
			})

		})

		convey.Convey("to *struct/**struct", func() {
			convey.Convey("nil to *struct", func() {
				var (
					target testCopyStruct
					source interface{}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{})
			})

			convey.Convey("*struct to *struct", func() {
				var (
					target testCopyStruct
					source *testCopyStruct = &testCopyStruct{
						Id:        100,
						Ignore:    true,
						Name:      "abc",
						Age:       23,
						Female:    true,
						unexpored: new(int),
					}
				)

				err := DeepCopy(&target, source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(&target, convey.ShouldResemble, &testCopyStruct{
					Id:        100,
					Name:      "abc",
					Age:       23,
					Female:    true,
					unexpored: new(int),
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

				err := DeepCopy(&target, &source)
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

				err := DeepCopy(&target, &source)
				convey.So(err, convey.ShouldBeNil)
				convey.So(target, convey.ShouldResemble, &testCopyStruct{
					Id:     100,
					Name:   "abc",
					Age:    23,
					Female: true,
				})
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

				err := DeepCopy(&target, &source)
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
