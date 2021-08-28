/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ context.Context = &Context{}
var _ ldcontext.Context = &Context{}

func GetContext(g *gin.Context) *Context {
	return newCtxIfNotExists(g)
}

func GetGin(c context.Context) *gin.Context {
	if g, ok := c.(*gin.Context); ok && g != nil {
		return g
	}

	if v, ok := c.Value(ctxKeyContext).(*Context); ok {
		return v.Gin()
	}

	return nil
}

func GetBeginTime(c context.Context) time.Time {
	if ctx := getCtxByCtx(c); ctx != nil {
		return ctx.beginTime
	}
	return time.Time{}
}

func GetSequence(c context.Context) string {
	if ctx := getCtxByCtx(c); ctx != nil {
		return ctx.sequence
	}
	return ""
}

func GetRequest(c context.Context) interface{}  { return GetGin(c).Value(GinKeyRequest) }
func GetRenderer(c context.Context) interface{} { return GetGin(c).Value(GinKeyRenderer) }

func GetError(c context.Context) Error {
	v := GetGin(c).Value(GinKeyError)
	r, _ := v.(Error)
	return r
}

func GetResponse(c context.Context) *CommResponse {
	v := GetGin(c).Value(GinKeyResponse)
	r, _ := v.(*CommResponse)
	return r
}

func newSequence(g *gin.Context) string {
	uuid, _ := uuid.NewRandom()
	return hex.EncodeToString(uuid[:])
}

func getCtxByCtx(c context.Context) *Context {
	var v interface{}
	if g, ok := c.(*gin.Context); ok {
		v = g.Value(GinKeyContext)
	} else {
		v = c.Value(ctxKeyContext)
	}

	r, _ := v.(*Context)
	return r
}

func newCtxIfNotExists(g *gin.Context) *Context {
	v := g.Value(GinKeyContext)
	c, ok := v.(*Context)
	if !ok {
		c = newContext(g)
	}
	return c
}

func newContext(g *gin.Context) *Context {
	now := time.Now()
	seq := newSequence(g)

	ctx := ldcontext.NewContext(g, zap.String("sequence", seq))

	c := &Context{
		ginCtx:    g,
		ldCtx:     ctx,
		beginTime: now,
		sequence:  seq,
	}

	g.Header(GinHeaderSequence, seq)
	g.Set(GinKeyContext, c)
	return c
}

type Context struct {
	*ginCtx
	ldCtx

	beginTime time.Time
	sequence  string
}

func (c *Context) String() string { return ldcontext.ContextName(c.ldCtx) + ".WithGin" }

func (c *Context) clone() *Context {
	copy := *c
	return &copy
}

func (c *Context) Copy() *Context {
	c = c.clone()
	c.ginCtx = c.ginCtx.Copy()
	return c
}

func (c *Context) Gin() *gin.Context { return c.ginCtx }

func (c *Context) Err() error                  { return c.ldCtx.Err() }
func (c *Context) Done() <-chan struct{}       { return c.ldCtx.Done() }
func (c *Context) Deadline() (time.Time, bool) { return c.ldCtx.Deadline() }

func (c *Context) Value(key interface{}) interface{} {
	if key == ctxKeyContext {
		return c
	}
	return c.ldCtx.Value(key)
}
