/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	protov2 "google.golang.org/protobuf/proto"
)

func ProtoCodec(m ...proto.Message) Codec { return ProtoV1Codec(m...) }

func ProtoV1Codec(m ...proto.Message) Codec {
	c := protoCodec{
		marshal: func(p interface{}) ([]byte, error) {
			m, _ := p.(proto.Message)
			if m == nil {
				return nil, fmt.Errorf("the object for marshal must be `proto.Message`. type:%T", p)
			}
			return proto.Marshal(m)
		},
		unmarshal: func(b []byte, p interface{}) error {
			m, _ := p.(proto.Message)
			if m == nil {
				return fmt.Errorf("the object for unmarshal must be `proto.Message`. type:%T", p)
			}
			return proto.Unmarshal(b, m)
		},
	}

	if len(m) > 0 && m[0] != nil {
		c.typ = reflect.TypeOf(m[0])
	}

	return c
}

func ProtoV2Codec(m ...protov2.Message) Codec {
	c := protoCodec{
		marshal: func(p interface{}) ([]byte, error) {
			m, _ := p.(protov2.Message)
			if m == nil {
				return nil, fmt.Errorf("the object for marshal must be `protoV2.Message`. type:%T", p)
			}
			return protov2.Marshal(m)
		},
		unmarshal: func(b []byte, p interface{}) error {
			m, _ := p.(protov2.Message)
			if m == nil {
				return fmt.Errorf("the object for unmarshal must be `protoV2.Message`. type:%T", p)
			}
			return protov2.Unmarshal(b, m)
		},
	}

	if len(m) > 0 && m[0] != nil {
		c.typ = reflect.TypeOf(m[0])
	}

	return c
}

type protoCodec struct {
	typ       reflect.Type
	marshal   func(v interface{}) ([]byte, error)
	unmarshal func(b []byte, p interface{}) error
}

func (c protoCodec) Marshal(v interface{}) ([]byte, error) {
	return c.marshal(v)
}

func (c protoCodec) Unmarshal(b []byte) (interface{}, error) {
	if c.typ == nil {
		return nil, fmt.Errorf("cannot new object before unmarshal. type:<nil>")
	}

	switch c.typ.Kind() {
	case reflect.Struct:
		v := reflect.New(c.typ).Elem().Interface()
		return v, c.unmarshal(b, &v)

	case reflect.Ptr:
		p := reflect.New(c.typ.Elem()).Interface()
		return p, c.unmarshal(b, p)

	default:
		return nil, fmt.Errorf("cannot new object before unmarshal. type:%s", c.typ.Name())
	}
}
