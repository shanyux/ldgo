/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"context"

	"github.com/distroy/ldgo/v2/ldconv"
)

type errorMarshaler struct {
	err error
}

func (c errorMarshaler) MarshalBinary() ([]byte, error) {
	return nil, c.err
}

func newCodecCmd(cc context.Context, cli *CodecRedis, cmd *StringCmd) *CodecCmd {
	c := &CodecCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type CodecCmd struct {
	*StringCmd

	err error
	val interface{}
}

func (c *CodecCmd) Err() error                   { return c.err }
func (c *CodecCmd) Val() interface{}             { return c.val }
func (c *CodecCmd) Result() (interface{}, error) { return c.Val(), c.Err() }
func (c *CodecCmd) parse(cc context.Context, cli *CodecRedis, cmd *StringCmd) {
	c.StringCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	v, err := cli.unmarshal(cc, ldconv.StrToBytesUnsafe(cmd.Val()))
	if err != nil {
		c.err = err
		return
	}

	c.val = v
}

func newStringCodecMapCmd(cc context.Context, cli *CodecRedis, cmd *StringStringMapCmd) *StringCodecMapCmd {
	c := &StringCodecMapCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type StringCodecMapCmd struct {
	*StringStringMapCmd

	err error
	val map[string]interface{}
}

func (c *StringCodecMapCmd) Err() error                              { return c.err }
func (c *StringCodecMapCmd) Val() map[string]interface{}             { return c.val }
func (c *StringCodecMapCmd) Result() (map[string]interface{}, error) { return c.Val(), c.Err() }
func (c *StringCodecMapCmd) parse(cc context.Context, cli *CodecRedis, cmd *StringStringMapCmd) {
	c.StringStringMapCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	m := make(map[string]interface{}, len(cmd.Val()))
	for k, val := range cmd.Val() {
		v, err := cli.unmarshal(cc, cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		m[k] = v
	}
	c.val = m
}

func newCodecsCmd(cc context.Context, cli *CodecRedis, cmd *StringsCmd) *CodecsCmd {
	c := &CodecsCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type CodecsCmd struct {
	*StringsCmd

	err error
	val []interface{}
}

func (c *CodecsCmd) Err() error                     { return c.err }
func (c *CodecsCmd) Val() []interface{}             { return c.val }
func (c *CodecsCmd) Result() ([]interface{}, error) { return c.Val(), c.Err() }
func (c *CodecsCmd) parse(cc context.Context, cli *CodecRedis, cmd *StringsCmd) {
	c.StringsCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make([]interface{}, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := cli.unmarshal(cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		s = append(s, v)
	}
	c.val = s
}

func newCodecSetCmd(cc context.Context, cli *CodecRedis, cmd *StringSetCmd) *CodecSetCmd {
	c := &CodecSetCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type CodecSetCmd struct {
	*StringSetCmd

	err error
	val map[interface{}]struct{}
}

func (c *CodecSetCmd) Err() error                                { return c.err }
func (c *CodecSetCmd) Val() map[interface{}]struct{}             { return c.val }
func (c *CodecSetCmd) Result() (map[interface{}]struct{}, error) { return c.Val(), c.Err() }
func (c *CodecSetCmd) parse(cc context.Context, cli *CodecRedis, cmd *StringSetCmd) {
	c.StringSetCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make(map[interface{}]struct{}, len(cmd.Val()))
	for val := range cmd.Val() {
		v, err := cli.unmarshal(cc, ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		s[v] = struct{}{}
	}
	c.val = s
}

func newCodecSliceCmd(cc context.Context, cli *CodecRedis, cmd *SliceCmd) *CodecSliceCmd {
	c := &CodecSliceCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type CodecSliceCmd struct {
	*SliceCmd

	err error
	val []interface{}
}

func (c *CodecSliceCmd) Err() error                     { return c.err }
func (c *CodecSliceCmd) Val() []interface{}             { return c.val }
func (c *CodecSliceCmd) Result() ([]interface{}, error) { return c.Val(), c.Err() }
func (c *CodecSliceCmd) parse(cc context.Context, cli *CodecRedis, cmd *SliceCmd) {
	c.SliceCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make([]interface{}, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := cli.unmarshalInterface(cc, val)
		if err != nil {
			c.err = err
			return
		}
		s = append(s, v)
	}
	c.val = s
}

func newZCodecSliceCmd(cc context.Context, cli *CodecRedis, cmd *ZSliceCmd) *ZCodecSliceCmd {
	c := &ZCodecSliceCmd{}
	c.parse(cc, cli, cmd)
	return c
}

type ZCodecSliceCmd struct {
	*ZSliceCmd

	err error
	val []ZMember
}

func (c *ZCodecSliceCmd) Err() error                 { return c.err }
func (c *ZCodecSliceCmd) Val() []ZMember             { return c.val }
func (c *ZCodecSliceCmd) Result() ([]ZMember, error) { return c.Val(), c.Err() }
func (c *ZCodecSliceCmd) parse(cc context.Context, cli *CodecRedis, cmd *ZSliceCmd) {
	c.ZSliceCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	members := make([]ZMember, 0, len(cmd.Val()))
	for _, v := range cmd.Val() {
		val, err := cli.unmarshalInterface(cc, v.Member)
		if err != nil {
			return
		}
		v.Member = val
		members = append(members, v)
	}

	c.val = members
}
