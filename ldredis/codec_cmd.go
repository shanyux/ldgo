/*
 * Copyright (C) distroy
 */

package ldredis

import "github.com/distroy/ldgo/ldconv"

type errorMarshaler struct {
	err error
}

func (c errorMarshaler) MarshalBinary() ([]byte, error) {
	return nil, c.err
}

var _ Cmder = (*CodecCmd)(nil)

type CodecCmd struct {
	*StringCmd

	err error
	val interface{}
}

func (c *CodecCmd) Err() error                   { return c.err }
func (c *CodecCmd) Val() interface{}             { return c.val }
func (c *CodecCmd) Result() (interface{}, error) { return c.Val(), c.Err() }
func (c *CodecCmd) parse(codec Codec, cmd *StringCmd) {
	c.StringCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	v, err := codec.Unmarshal(ldconv.StrToBytesUnsafe(cmd.Val()))
	if err != nil {
		c.err = err
		return
	}

	c.val = v
}

var _ Cmder = (*StringCodecMapCmd)(nil)

type StringCodecMapCmd struct {
	*StringStringMapCmd

	err error
	val map[string]interface{}
}

func (c *StringCodecMapCmd) Err() error                              { return c.err }
func (c *StringCodecMapCmd) Val() map[string]interface{}             { return c.val }
func (c *StringCodecMapCmd) Result() (map[string]interface{}, error) { return c.Val(), c.Err() }
func (c *StringCodecMapCmd) parse(codec Codec, cmd *StringStringMapCmd) {
	c.StringStringMapCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	m := make(map[string]interface{}, len(cmd.Val()))
	for k, val := range cmd.Val() {
		v, err := codec.Unmarshal(ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		m[k] = v
	}
	c.val = m
}

var _ Cmder = (*CodecSliceCmd)(nil)

type CodecSliceCmd struct {
	*StringSliceCmd

	err error
	val []interface{}
}

func (c *CodecSliceCmd) Err() error                     { return c.err }
func (c *CodecSliceCmd) Val() []interface{}             { return c.val }
func (c *CodecSliceCmd) Result() ([]interface{}, error) { return c.Val(), c.Err() }
func (c *CodecSliceCmd) parse(codec Codec, cmd *StringSliceCmd) {
	c.StringSliceCmd = cmd
	c.err = cmd.Err()
	if c.err != nil {
		return
	}

	s := make([]interface{}, 0, len(cmd.Val()))
	for _, val := range cmd.Val() {
		v, err := codec.Unmarshal(ldconv.StrToBytesUnsafe(val))
		if err != nil {
			c.err = err
			return
		}
		s = append(s, v)
	}
	c.val = s
}
