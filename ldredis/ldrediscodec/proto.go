/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
)

func Proto[T protoadapt.MessageV1](d ...T) Codec[T]   { return ProtoCodec[T]{} }
func ProtoV1[T protoadapt.MessageV1](d ...T) Codec[T] { return ProtoV1Codec[T]{} }
func ProtoV2[T protoadapt.MessageV2](d ...T) Codec[T] { return ProtoV2Codec[T]{} }

type ProtoCodec[T protoadapt.MessageV1] struct {
	ProtoV1Codec[T]
}

type ProtoV1Codec[T protoadapt.MessageV1] struct {
	codecBase[T]
}

func (c ProtoV1Codec[T]) Marshal(v T) ([]byte, error) {
	vv := protoadapt.MessageV2Of(v)
	return proto.Marshal(vv)
}

func (c ProtoV1Codec[T]) Unmarshal(b []byte) (T, error) {
	var v T
	c.fillIfPointer(&v)
	vv := protoadapt.MessageV2Of(v)
	return v, proto.Unmarshal(b, vv)
}

type ProtoV2Codec[T protoadapt.MessageV2] struct {
	codecBase[T]
}

func (c ProtoV2Codec[T]) Marshal(v T) ([]byte, error) {
	vv := v
	return proto.Marshal(vv)
}

func (c ProtoV2Codec[T]) Unmarshal(b []byte) (T, error) {
	var v T
	c.fillIfPointer(&v)
	vv := v
	return v, proto.Unmarshal(b, vv)
}
