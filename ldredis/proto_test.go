/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"testing"
	"time"

	"github.com/distroy/ldgo/ldconv"
	"github.com/distroy/ldgo/ldptr"
	"github.com/distroy/ldgo/proto/ldtestproto"
	"github.com/golang/protobuf/proto"
	"github.com/smartystreets/goconvey/convey"
	protov2 "google.golang.org/protobuf/proto"
)

var (
	_ proto.Message   = (*ldtestproto.TestProto)(nil)
	_ protov2.Message = proto.MessageV2(&ldtestproto.TestProto{})
)

func testMustProtoV1Marsha(m proto.Message) string {
	b, _ := proto.Marshal(m)
	return ldconv.BytesToStrUnsafe(b)
}

func testMustProtoV2Marsha(m protov2.Message) string {
	b, _ := protov2.Marshal(m)
	return ldconv.BytesToStrUnsafe(b)
}

func testGetProtoV2Values(m protov2.Message) map[string]interface{} {
	fields := make(map[string]interface{})
	for i, l := 0, m.ProtoReflect().Descriptor().Fields().Len(); i < l; i++ {
		f := m.ProtoReflect().Descriptor().Fields().Get(i)
		v := m.ProtoReflect().Get(f)
		fields[string(f.Name())] = v.Interface()
	}
	return fields
}

func TestProtoV1Codec(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-proto-v1-codec"
		expiration := time.Duration(0)
		val := &ldtestproto.TestProto{
			Str: ldptr.NewString("abc"),
			I64: ldptr.NewInt64(234),
			F64: ldptr.NewFloat64(100.234),
		}

		set := rds.WithCodec(ProtoV1Codec()).Set(key, val, expiration)
		convey.So(set.Err(), convey.ShouldBeNil)

		convey.Convey("get-str", func() {
			get := rds.Get(key)
			convey.So(get.Err(), convey.ShouldBeNil)
			convey.So(get.Val(), convey.ShouldResemble, testMustProtoV1Marsha(val))
		})
		convey.Convey("get-obj", func() {
			get := rds.WithCodec(ProtoV1Codec(&ldtestproto.TestProto{})).Get(key)
			convey.So(get.Err(), convey.ShouldBeNil)
			convey.So(get.Val(), convey.ShouldResemble, &ldtestproto.TestProto{
				Str: ldptr.NewString("abc"),
				I64: ldptr.NewInt64(234),
				F64: ldptr.NewFloat64(100.234),
			})
		})
	})
}

func TestProtoV2Codec(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		rds := testMemoryRedis()
		defer rds.Close()

		key := "test-proto-v2-codec"
		expiration := time.Duration(0)
		val := proto.MessageV2(&ldtestproto.TestProto{
			Str: ldptr.NewString("abc"),
			I64: ldptr.NewInt64(234),
			F64: ldptr.NewFloat64(100.234),
		})

		set := rds.WithCodec(ProtoV2Codec()).Set(key, val, expiration)
		convey.So(set.Err(), convey.ShouldBeNil)

		convey.Convey("get-str", func() {
			get := rds.Get(key)
			convey.So(get.Err(), convey.ShouldBeNil)
			convey.So(get.Val(), convey.ShouldResemble, testMustProtoV2Marsha(val))
		})
		convey.Convey("get-obj", func() {
			get := rds.WithCodec(ProtoV2Codec(proto.MessageV2(&ldtestproto.TestProto{}))).Get(key)
			convey.So(get.Err(), convey.ShouldBeNil)

			message := get.Val().(protov2.Message)
			convey.So(testGetProtoV2Values(message), convey.ShouldResemble, map[string]interface{}{
				"str": "abc",
				"i64": int64(234),
				"f64": float64(100.234),
			})
		})
	})
}
