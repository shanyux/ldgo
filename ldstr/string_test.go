/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestStrMapReplace(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		m := map[string]string{
			"a":   "A",
			"b":   "B",
			"env": "test",
		}

		convey.So(StrMapReplace("a}{{a}{abc}{b}{}.{{a}}", m, "{", "}"), convey.ShouldEqual, "a}{A{abc}B{}.{A}")
		convey.So(StrMapReplace("./config_{env}.ini", m, "{", "}"), convey.ShouldEqual, "./config_test.ini")
	})
}

func TestStrIMapReplace(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
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

func TestStrMapParse(t *testing.T) {
	type subtest struct {
		text    string
		wantErr bool
		want    map[string]string
	}

	type test struct {
		tmpl     string
		left     string
		right    string
		splits   []string
		wantErr  bool
		subtests []subtest
	}

	tests := []test{
		{
			tmpl:    "{}: {user}",
			left:    "{",
			right:   "}",
			wantErr: false,
			subtests: []subtest{
				{
					text:    "user: xxx",
					wantErr: false,
					want:    map[string]string{`user`: "xxx"},
				},
			},
		},
		{
			tmpl:    "{key}: {value}",
			left:    "{",
			right:   "}",
			wantErr: false,
			subtests: []subtest{
				{
					text:    "user: xxx",
					wantErr: false,
					want: map[string]string{
						`key`:   "user",
						`value`: "xxx",
					},
				},
			},
		},
		{
			tmpl:    "{key}{}: {value}{ignore}",
			left:    "{",
			right:   "}",
			wantErr: false,
			subtests: []subtest{
				{
					text:    "user: xxx",
					wantErr: false,
					want: map[string]string{
						`key`:    "user",
						`value`:  "xxx",
						`ignore`: "",
					},
				},
			},
		},
	}

	subtestRun := func(subtests []subtest, parser *strMapParser) {
		for _, stt := range subtests {
			convey.Convey(stt.text, func() {
				got, err := parser.Parse(stt.text)
				if stt.wantErr {
					convey.So(err, convey.ShouldNotBeNil)
				} else {
					convey.So(err, convey.ShouldBeNil)
				}
				convey.So(got, convey.ShouldResemble, stt.want)
			})
		}
	}

	convey.Convey(t.Name(), t, func() {
		for _, tt := range tests {
			convey.Convey(tt.tmpl, func() {
				parser := &strMapParser{}
				err := parser.Init(tt.tmpl, tt.left, tt.right)
				log.Printf("%s", mustMarshalJson(parser.fields))
				if tt.wantErr {
					convey.So(err, convey.ShouldNotBeNil)
					return
				}

				convey.So(err, convey.ShouldBeNil)
				subtestRun(tt.subtests, parser)
			})
		}
	})
}
