/*
 * Copyright (C) distroy
 */

package ldsheet

import (
	"io"
	"testing"

	"github.com/distroy/ldgo/ldptr"
	"github.com/smartystreets/goconvey/convey"
)

func Test_Reader(t *testing.T) {
	type Object struct {
		Name    string `sheet:"name:name;notempty"`
		ID      int64  `sheet:"name:id;notempty"`
		Prefix  string `sheet:"prefix:Prefix"`
		HttpUrl *string
		Int     int
		Uint    *uint
	}

	convey.Convey(t.Name(), t, func() {
		convey.Convey("read header fail", func() {
			lines := [][]string{}

			r := &Reader{Reader: Lines(lines)}
			p := &Object{}

			err := r.Read(p)
			convey.So(err.Error(), convey.ShouldEqual, "[ldsheet] read header fail. err:EOF")
		})

		convey.Convey("header is missing field", func() {
			lines := [][]string{
				{"id"},
			}

			r := &Reader{Reader: Lines(lines)}
			p := &Object{}

			err := r.Read(p)
			convey.So(err.Error(), convey.ShouldEqual, "[ldsheet] the field is missed in header. field:Name, name:name")
		})

		convey.Convey("field is empty", func() {
			lines := [][]string{
				{"id", "name", "Int"},
				{"100", "", ""},
			}

			r := &Reader{Reader: Lines(lines)}
			p := &Object{}

			err := r.Read(p)
			convey.So(err.Error(), convey.ShouldEqual, "[ldsheet] the field must not be empty. field:Name")
		})

		convey.Convey("parse string fail", func() {
			lines := [][]string{
				{"id", "name", "Int"},
				{"100", "aaa", "aaa"},
			}

			r := &Reader{Reader: Lines(lines)}
			p := &Object{}

			err := r.Read(p)
			convey.So(err.Error(), convey.ShouldStartWith, "[ldsheet] parse field value fail. type:int,")
		})

		convey.Convey("read succ", func() {
			lines := [][]string{
				{"id", "name", "http url", "Prefix (for test)", "Int", "Uint"},
				{"100", "aaa", "http://a", "xxxx", "0x100"},
				{"200", "bbb", "http://b", "yyyy", "", "0x200"},
				{"300", "ccc", ""},
			}

			r := &Reader{Reader: Lines(lines)}

			convey.Convey("read", func() {
				p := &Object{}

				convey.So(r.Read(p), convey.ShouldBeNil)
				convey.So(p, convey.ShouldResemble, &Object{
					ID:      100,
					Name:    "aaa",
					Prefix:  "xxxx",
					HttpUrl: ldptr.NewString("http://a"),
					Int:     0x100,
				})

				convey.So(r.Read(p), convey.ShouldBeNil)
				convey.So(p, convey.ShouldResemble, &Object{
					ID:      200,
					Name:    "bbb",
					Prefix:  "yyyy",
					HttpUrl: ldptr.NewString("http://b"),
					Uint:    ldptr.NewUint(0x200),
				})

				convey.So(r.Read(p), convey.ShouldBeNil)
				convey.So(p, convey.ShouldResemble, &Object{
					ID:   300,
					Name: "ccc",
				})

				convey.So(r.Read(p), convey.ShouldEqual, io.EOF)
			})

			convey.Convey("read all (struct)", func() {
				var res []Object
				convey.So(r.ReadAll(&res), convey.ShouldBeNil)
				convey.So(res, convey.ShouldResemble, []Object{
					{
						ID:      100,
						Name:    "aaa",
						Prefix:  "xxxx",
						HttpUrl: ldptr.NewString("http://a"),
						Int:     0x100,
					}, {
						ID:      200,
						Name:    "bbb",
						Prefix:  "yyyy",
						HttpUrl: ldptr.NewString("http://b"),
						Uint:    ldptr.NewUint(0x200),
					}, {
						ID:   300,
						Name: "ccc",
					},
				})
			})

			convey.Convey("read all (*struct)", func() {
				var res []*Object
				convey.So(r.ReadAll(&res), convey.ShouldBeNil)
				convey.So(res, convey.ShouldResemble, []*Object{
					{
						ID:      100,
						Name:    "aaa",
						Prefix:  "xxxx",
						HttpUrl: ldptr.NewString("http://a"),
						Int:     0x100,
					}, {
						ID:      200,
						Name:    "bbb",
						Prefix:  "yyyy",
						HttpUrl: ldptr.NewString("http://b"),
						Uint:    ldptr.NewUint(0x200),
					}, {
						ID:   300,
						Name: "ccc",
					},
				})
			})
		})
	})
}
