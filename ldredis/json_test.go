/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	"github.com/distroy/ldgo/ldptr"
	"github.com/smartystreets/goconvey/convey"
)

func TestJsonCodec(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-json-codec"
		expiration := time.Duration(0)

		convey.Convey("int64", func() {
			val := int64(100)

			set := rds.WithCodec(JsonCodec()).Set(key, val, expiration)
			convey.So(set.Err(), convey.ShouldBeNil)

			convey.Convey("get-str", func() {
				get := rds.Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, "100")
			})
			convey.Convey("get-i64", func() {
				get := rds.WithCodec(JsonCodec(int64(0))).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, int64(100))
			})
			convey.Convey("get-pi64", func() {
				get := rds.WithCodec(JsonCodec(ldptr.NewInt64(0))).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, ldptr.NewInt64(100))
			})
			convey.Convey("get-nil", func() {
				get := rds.WithCodec(JsonCodec()).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, float64(100))
			})
		})
		convey.Convey("object", func() {
			type Object struct {
				S    string  `json:"s"`
				PS   *string `json:"ps"`
				I64  int64   `json:"i64"`
				PI64 *int64  `json:"pi64"`
			}

			val := Object{
				S:    "abc",
				PS:   ldptr.NewString("xyz"),
				I64:  123,
				PI64: ldptr.NewInt64(234),
			}

			set := rds.WithCodec(JsonCodec()).Set(key, val, expiration)
			convey.So(set.Err(), convey.ShouldBeNil)

			convey.Convey("get-str", func() {
				get := rds.Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, `{"s":"abc","ps":"xyz","i64":123,"pi64":234}`)
			})
			convey.Convey("get-obj", func() {
				get := rds.WithCodec(JsonCodec(Object{})).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, Object{
					S:    "abc",
					PS:   ldptr.NewString("xyz"),
					I64:  123,
					PI64: ldptr.NewInt64(234),
				})
			})
			convey.Convey("get-ptr", func() {
				get := rds.WithCodec(JsonCodec(&Object{})).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, &Object{
					S:    "abc",
					PS:   ldptr.NewString("xyz"),
					I64:  123,
					PI64: ldptr.NewInt64(234),
				})
			})
			convey.Convey("get-nil", func() {
				get := rds.WithCodec(JsonCodec()).Get(key)
				convey.So(get.Err(), convey.ShouldBeNil)
				convey.So(get.Val(), convey.ShouldResemble, map[string]interface{}{
					"s":    "abc",
					"ps":   "xyz",
					"i64":  float64(123),
					"pi64": float64(234),
				})
			})
		})
	})
}