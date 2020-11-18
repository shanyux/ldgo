/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_StrMapReplace(t *testing.T) {
	convey.Convey("", t, func() {
		m := map[string]string{
			"a":   "A",
			"b":   "B",
			"env": "test",
		}

		convey.So(StrMapReplace("a}{{a}{abc}{b}{}.{{a}}", m, "{", "}"), convey.ShouldEqual, "a}{A{abc}B{}.{A}")
		convey.So(StrMapReplace("./config_{env}.ini", m, "{", "}"), convey.ShouldEqual, "./config_test.ini")
	})
}

func Test_StrIMapReplace(t *testing.T) {
	convey.Convey("", t, func() {
		s0 := "a}{{a}{abc}{b}{}.{{a}}"
		m0 := map[string]string{
			"a": "A",
			"b": "B",
		}
		convey.So(StrIMapReplace(s0, m0, "{", "}"), convey.ShouldEqual, "a}{A{abc}B{}.{A}")

		m1 := map[string]int{
			"a": 1,
			"b": 2,
		}
		convey.So(StrIMapReplace(s0, m1, "{", "}"), convey.ShouldEqual, "a}{1{abc}2{}.{1}")
	})
}
