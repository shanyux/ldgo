/*
 * Copyright (C) distroy
 */

package ldref

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type testOject struct{}

func (*testOject) Func() {}

func TestGetFuncName(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("func", func(c convey.C) {
			f := GetFuncName(GetFuncName)
			c.So(f, convey.ShouldResemble, FuncName{
				Full:     "github.com/distroy/ldgo/ldref.GetFuncName",
				Short:    "ldref.GetFuncName",
				Path:     "github.com/distroy/ldgo",
				Package:  "ldref",
				Method:   "GetFuncName",
				Receiver: "",
			})
		})

		c.Convey("method", func(c convey.C) {
			f := GetFuncName((*testOject).Func)
			c.So(f, convey.ShouldResemble, FuncName{
				Full:     "github.com/distroy/ldgo/ldref.(*testOject).Func",
				Short:    "ldref.(*testOject).Func",
				Path:     "github.com/distroy/ldgo",
				Package:  "ldref",
				Method:   "Func",
				Receiver: "*testOject",
			})
		})

		c.Convey("method fm", func(c convey.C) {
			f := GetFuncName((&testOject{}).Func)
			c.So(f, convey.ShouldResemble, FuncName{
				Full:     "github.com/distroy/ldgo/ldref.(*testOject).Func-fm",
				Short:    "ldref.(*testOject).Func-fm",
				Path:     "github.com/distroy/ldgo",
				Package:  "ldref",
				Method:   "Func-fm",
				Receiver: "*testOject",
			})
		})

		c.Convey("testing.Main", func(c convey.C) {
			f := GetFuncName(testing.Main)
			c.So(f, convey.ShouldResemble, FuncName{
				Full:     "testing.Main",
				Short:    "testing.Main",
				Path:     "",
				Package:  "testing",
				Method:   "Main",
				Receiver: "",
			})
		})
	})
}
