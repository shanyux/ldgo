/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldlogger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Context struct {
	*ginContext
	ldContext
}

func GetContext(g *gin.Context) Context {
	v := g.Value(GIN_KEY_CONTEXT)
	c, ok := v.(Context)
	if !ok {
		c = newContext(g)
	}
	return c

}

func GetGin(c context.Context) *gin.Context {
	if g, ok := c.(*gin.Context); ok && g != nil {
		return g
	}

	if v, ok := c.(Context); ok {
		return v.Gin()
	}

	v := c.Value(_CTX_KEY_GIN_CONTEXT)
	g, _ := v.(*gin.Context)
	return g
}

func newSequence(g *gin.Context) string {
	uuid, _ := uuid.NewRandom()
	return hex.EncodeToString(uuid[:])
}

func newContext(g *gin.Context) Context {
	c := ldcontext.NewContext(g)
	c = c.WithValue(_CTX_KEY_GIN_CONTEXT, g)

	now := time.Now()
	c = c.WithValue(_CTX_KEY_BEGIN_TIME, now)

	seq := newSequence(g)
	c = c.With(zap.String("sequence", seq))
	g.Header(GIN_HEADER_SEQUENCE, seq)
	c = c.WithValue(_CTX_KEY_SEQUENCE, seq)

	ctx := Context{
		ginContext: g,
		ldContext:  c,
	}
	g.Set(GIN_KEY_CONTEXT, ctx)
	return ctx
}

func (c Context) IsValid() bool {
	return c == Context{}
}

func (c Context) Gin() *gin.Context          { return c.ginContext }
func (c Context) context() ldcontext.Context { return c }

func (c Context) Request() interface{}  { return c.Gin().Value(GIN_KEY_REQUEST) }
func (c Context) Renderer() interface{} { return c.Gin().Value(GIN_KEY_RENDERER) }

func (c Context) BeginTime() time.Time {
	v := c.Value(_CTX_KEY_BEGIN_TIME)
	t, _ := v.(time.Time)
	return t
}

func (c Context) Sequence() string {
	v := c.Value(_CTX_KEY_SEQUENCE)
	t, _ := v.(string)
	return t
}

func (c Context) Error() Error {
	v := c.Gin().Value(GIN_KEY_REQUEST)
	err, _ := v.(Error)
	return err
}

func (c Context) Response() *CommResponse {
	v := c.Gin().Value(GIN_KEY_CONTEXT)
	rsp, _ := v.(*CommResponse)
	return rsp
}

func (c Context) Copy() Context {
	c.ginContext = c.Gin().Copy()
	c.ldContext = c.ldContext.WithValue(_CTX_KEY_GIN_CONTEXT, c.ginContext)
	return c
}

func (c Context) Err() error                      { return c.ldContext.Err() }
func (c Context) Done() <-chan struct{}           { return c.ldContext.Done() }
func (c Context) Deadline() (time.Time, bool)     { return c.ldContext.Deadline() }
func (c Context) Value(k interface{}) interface{} { return c.ldContext.Value(k) }

func (c Context) With(fields ...zap.Field) ldcontext.Context {
	c.ldContext = c.ldContext.With(fields...)
	return c
}

func (c Context) WithLogger(l ldlogger.Logger) ldcontext.Context {
	c.ldContext = c.ldContext.WithLogger(l)
	return c
}

func (c Context) WithValue(k, v interface{}) ldcontext.Context {
	c.ldContext = c.ldContext.WithValue(k, v)
	return c
}

func (c Context) WithCancel() ldcontext.Context {
	c.ldContext = c.ldContext.WithCancel()
	return c
}

func (c Context) WithDeadline(deadline time.Time) ldcontext.Context {
	c.ldContext = c.ldContext.WithDeadline(deadline)
	return c
}

func (c Context) WithTimeout(timeout time.Duration) ldcontext.Context {
	c.ldContext = c.ldContext.WithTimeout(timeout)
	return c
}
