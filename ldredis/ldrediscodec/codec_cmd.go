/*
 * Copyright (C) distroy
 */

package ldrediscodec

import (
	"context"
	"encoding"

	"github.com/distroy/ldgo/v2/ldconv"
	"github.com/distroy/ldgo/v2/ldredis"
)

var _ encoding.BinaryMarshaler = (*errorMarshaler)(nil)

type errorMarshaler struct {
	Err error
}

func (c errorMarshaler) MarshalBinary() ([]byte, error) {
	return nil, c.Err
}

func newAnyCmd[T comparable](cc context.Context, cli *Redis[T], cmd *StringCmd) *AnyCmd[T] {
	c := &AnyCmd[T]{
		base:      cli.base,
		StringCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type AnyCmd[T comparable] struct {
	base[T]
	*StringCmd

	val T
}

func (c *AnyCmd[T]) Val() T             { return c.val }
func (c *AnyCmd[T]) Result() (T, error) { return c.Val(), c.Err() }
func (c *AnyCmd[T]) Parse(cc context.Context) error {
	cmd := c.StringCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	v, err := c.unmarshal(cc, ldconv.StrToBytesUnsafe(cmd.Val()))
	if err != nil {
		c.SetErr(err)
		return err
	}

	c.val = v
	return nil
}

func newMapStringAnyCmd[T comparable](cc context.Context, cli *Redis[T], cmd *MapStringStringCmd) *MapStringAnyCmd[T] {
	c := &MapStringAnyCmd[T]{
		base:               cli.base,
		MapStringStringCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type MapStringAnyCmd[T comparable] struct {
	base[T]
	*MapStringStringCmd

	val map[string]T
}

func (c *MapStringAnyCmd[T]) Val() map[string]T             { return c.val }
func (c *MapStringAnyCmd[T]) Result() (map[string]T, error) { return c.Val(), c.Err() }
func (c *MapStringAnyCmd[T]) Parse(cc context.Context) error {
	cmd := c.MapStringStringCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	m := make(map[string]T, len(cmd.Val()))
	for k, val := range cmd.Val() {
		v, err := c.unmarshal(cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.SetErr(err)
			return err
		}
		m[k] = v
	}
	c.val = m
	return nil
}

func newAnySliceCmd[T comparable](cc context.Context, cli *Redis[T], cmd *StringSliceCmd) *AnySliceCmd[T] {
	c := &AnySliceCmd[T]{
		base:           cli.base,
		StringSliceCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type AnySliceCmd[T comparable] struct {
	base[T]
	*StringSliceCmd

	val []T
}

func (c *AnySliceCmd[T]) Val() []T             { return c.val }
func (c *AnySliceCmd[T]) Result() ([]T, error) { return c.Val(), c.Err() }
func (c *AnySliceCmd[T]) Parse(cc context.Context) error {
	cmd := c.StringSliceCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	s := make([]T, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := c.unmarshal(cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.SetErr(err)
			return err
		}
		s = append(s, v)
	}
	c.val = s
	return nil
}

func newAnySetCmd[T comparable](cc context.Context, cli *Redis[T], cmd *StringSetCmd) *AnySetCmd[T] {
	c := &AnySetCmd[T]{
		base:         cli.base,
		StringSetCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type AnySetCmd[T comparable] struct {
	base[T]
	*StringSetCmd

	val map[T]struct{}
}

func (c *AnySetCmd[T]) Val() map[T]struct{}             { return c.val }
func (c *AnySetCmd[T]) Result() (map[T]struct{}, error) { return c.Val(), c.Err() }
func (c *AnySetCmd[T]) Parse(cc context.Context) error {
	cmd := c.StringSetCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	s := make(map[T]struct{}, len(cmd.Val()))
	for val := range cmd.Val() {
		v, err := c.unmarshal(cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.SetErr(err)
			return err
		}
		s[v] = struct{}{}
	}
	c.val = s
	return nil
}

func newSliceCmd[T comparable](cc context.Context, cli *Redis[T], cmd *ldredis.SliceCmd) *SliceCmd[T] {
	c := &SliceCmd[T]{
		base:     cli.base,
		SliceCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type SliceCmd[T any] struct {
	base[T]
	*ldredis.SliceCmd

	val []T
}

func (c *SliceCmd[T]) Val() []T             { return c.val }
func (c *SliceCmd[T]) Result() ([]T, error) { return c.Val(), c.Err() }
func (c *SliceCmd[T]) Parse(cc context.Context) error {
	cmd := c.SliceCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	s := make([]T, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := c.unmarshalInterface(cc, val)
		if err != nil {
			c.SetErr(err)
			return err
		}
		s = append(s, v)
	}
	c.val = s
	return nil
}

func newZMemberSliceCmd[T comparable](cc context.Context, cli *Redis[T], cmd *ZSliceCmd) *ZMemberSliceCmd[T] {
	c := &ZMemberSliceCmd[T]{
		base:      cli.base,
		ZSliceCmd: cmd,
	}
	c.Parse(cc)
	return c
}

type ZMemberSliceCmd[T any] struct {
	base[T]
	*ZSliceCmd

	val []ZMember[T]
}

func (c *ZMemberSliceCmd[T]) Val() []ZMember[T]             { return c.val }
func (c *ZMemberSliceCmd[T]) Result() ([]ZMember[T], error) { return c.Val(), c.Err() }
func (c *ZMemberSliceCmd[T]) Parse(cc context.Context) error {
	cmd := c.ZSliceCmd
	if err := cmd.Err(); err != nil {
		return err
	}

	members := make([]ZMember[T], 0, len(cmd.Val()))
	for _, v := range cmd.Val() {
		val, err := c.unmarshalInterface(cc, v.Member)
		if err != nil {
			c.SetErr(err)
			return err
		}
		members = append(members, ZMember[T]{
			Score:  v.Score,
			Member: val,
		})
	}

	c.val = members
	return nil
}
