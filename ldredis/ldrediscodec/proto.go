/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	protov2 "google.golang.org/protobuf/proto"
)

func ProtoCodec(m ...proto.Message) Codec { return ProtoV1Codec(m...) }

func ProtoV1Codec(m ...proto.Message) Codec {
	if len(m) == 0 || m[0] == nil {
		return protoV1Codec{}
	}

	return protoV1Codec{
		t: reflect.TypeOf(m[0]),
	}
}

func ProtoV2Codec(m ...protov2.Message) Codec {
	if len(m) == 0 || m[0] == nil {
		return protoV2Codec{}
	}

	return protoV2Codec{
		m: m[0],
	}
}

type protoV1Codec struct {
	t reflect.Type
}

func (c protoV1Codec) Marshal(v interface{}) ([]byte, error) {
	m, _ := v.(proto.Message)
	if m == nil {
		return nil, fmt.Errorf("the object for marshal must be `proto.Message`. type:%T", v)
	}
	return proto.Marshal(m)
}

func (c protoV1Codec) Unmarshal(b []byte) (interface{}, error) {
	if c.t == nil {
		return nil, fmt.Errorf("cannot new object before unmarshal. type:<nil>")
	}

	if c.t.Kind() != reflect.Ptr || c.t.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("cannot new object before unmarshal. type:%s", c.t.String())
	}

	p := reflect.New(c.t.Elem()).Interface()
	m, _ := p.(proto.Message)
	if m == nil {
		return nil, fmt.Errorf("the object for unmarshal must be `proto.Message`. type:%T", p)
	}

	return m, proto.Unmarshal(b, m)
}

type protoV2Codec struct {
	m protov2.Message
}

func (c protoV2Codec) Marshal(v interface{}) ([]byte, error) {
	m, _ := v.(protov2.Message)
	if m == nil {
		return nil, fmt.Errorf("the object for marshal must be `protoV2.Message`. type:%T", v)
	}
	return protov2.Marshal(m)
}

func (c protoV2Codec) Unmarshal(b []byte) (interface{}, error) {
	if c.m == nil {
		return nil, fmt.Errorf("cannot new object before unmarshal. type:<nil>")
	}

	m := c.m.ProtoReflect().New().Interface()
	return m, protov2.Unmarshal(b, m)
}
