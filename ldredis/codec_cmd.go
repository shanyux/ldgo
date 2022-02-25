/*
 * Copyright (C) distroy
 */

package ldredis

import (
	"github.com/distroy/ldgo/ldconv"
)

type errorMarshaler struct {
	err error
}

func (c errorMarshaler) MarshalBinary() ([]byte, error) {
	return nil, c.err
}

func newCodecCmd(cli *CodecRedis, cmd *StringCmd) *CodecCmd {
	c := &CodecCmd{}
	c.parse(cli, cmd)
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
func (c *CodecCmd) parse(cli *CodecRedis, cmd *StringCmd) {
	c.StringCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	v, err := cli.unmarshal(ldconv.StrToBytesUnsafe(cmd.Val()))
	if err != nil {
		c.err = err
		return
	}

	c.val = v
}

func newStringCodecMapCmd(cli *CodecRedis, cmd *StringStringMapCmd) *StringCodecMapCmd {
	c := &StringCodecMapCmd{}
	c.parse(cli, cmd)
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
func (c *StringCodecMapCmd) parse(cli *CodecRedis, cmd *StringStringMapCmd) {
	c.StringStringMapCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	m := make(map[string]interface{}, len(cmd.Val()))
	for k, val := range cmd.Val() {
		v, err := cli.unmarshal(ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		m[k] = v
	}
	c.val = m
}

func newCodecsCmd(cli *CodecRedis, cmd *StringsCmd) *CodecsCmd {
	c := &CodecsCmd{}
	c.parse(cli, cmd)
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
func (c *CodecsCmd) parse(cli *CodecRedis, cmd *StringsCmd) {
	c.StringsCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make([]interface{}, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := cli.unmarshal(ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		s = append(s, v)
	}
	c.val = s
}

func newCodecSliceCmd(cli *CodecRedis, cmd *SliceCmd) *CodecSliceCmd {
	c := &CodecSliceCmd{}
	c.parse(cli, cmd)
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
func (c *CodecSliceCmd) parse(cli *CodecRedis, cmd *SliceCmd) {
	c.SliceCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make([]interface{}, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := cli.unmarshalInterface(val)
		if err != nil {
			c.err = err
			return
		}
		s = append(s, v)
	}
	c.val = s
}

func newZCodecSliceCmd(cli *CodecRedis, cmd *ZSliceCmd) *ZCodecSliceCmd {
	c := &ZCodecSliceCmd{}
	c.parse(cli, cmd)
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
func (c *ZCodecSliceCmd) parse(cli *CodecRedis, cmd *ZSliceCmd) {
	c.ZSliceCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	members := make([]ZMember, 0, len(cmd.Val()))
	for _, v := range cmd.Val() {
		val, err := cli.unmarshalInterface(v.Member)
		if err != nil {
			return
		}
		v.Member = val
		members = append(members, v)
	}

	c.val = members
}
