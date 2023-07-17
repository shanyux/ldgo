/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldrand"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	StdContext = context.Context
)

var (
	_ context.Context = &Context{}
	_ ldctx.Context   = &Context{}
)

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
	if ctx := getCtxByCommCtx(c); ctx != nil {
		return ctx.beginTime
	}
	return time.Time{}
}

func GetSequence(c context.Context) string {
	if ctx := getCtxByCommCtx(c); ctx != nil {
		return ctx.sequence
	}
	return ""
}

func GetRequest(c context.Context) interface{}  { return GetGin(c).Value(GinKeyRequest) }
func GetRenderer(c context.Context) interface{} { return GetGin(c).Value(GinKeyRenderer) }

func GetError(c StdContext) Error {
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
	return hex.EncodeToString(ldrand.Bytes(16))
}

func getCtxByCommCtx(child context.Context) *Context {
	if g, ok := child.(*gin.Context); ok {
		return getCtxByGinCtx(g)
	}

	r, _ := child.Value(ctxKeyContext).(*Context)
	return r
}

func getCtxByGinCtx(g *gin.Context) *Context {
	c, ok := g.Value(GinKeyContext).(*Context)
	if ok {
		return c
	}

	return nil
}

func newCtxIfNotExists(g *gin.Context) *Context {
	c := getCtxByGinCtx(g)
	if c == nil {
		c = newContext(g)
	}
	return c
}

func newContext(g *gin.Context) *Context {
	now := time.Now()
	seq := newSequence(g)

	ctx := ldctx.NewContext(g, zap.String("sequence", seq))

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

	handler   string
	method    string
	path      string
	beginTime time.Time
	sequence  string
}

func (c *Context) String() string { return ldctx.ContextName(c.ldCtx) + ".WithGin" }

func (c *Context) clone() *Context {
	copy := *c
	return &copy
}

func (c *Context) Copy() *Context {
	c = c.clone()
	c.ginCtx = c.ginCtx.Copy()

	c.ginCtx.Set(GinKeyContext, c)
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

func (c *Context) AbortWithData(data interface{}) {
	c.AbortWithErrorData(lderr.ErrSuccess, data)
}

func (c *Context) AbortWithError(err Error) {
	c.AbortWithErrorData(err, struct{}{})
}

func (c *Context) AbortWithErrorData(err Error, data interface{}) {
	if data == nil {
		data = struct{}{}
	}

	if err == nil {
		err = lderr.ErrSuccess
	}

	response := &CommResponse{
		Sequence: c.sequence,
		Cost:     time.Since(c.beginTime).String(),
		ErrCode:  err.Code(),
		ErrMsg:   err.Error(),
		Data:     data,
	}

	if e, ok := err.(lderr.ErrorWithDetails); ok {
		response.ErrDetails = e.Details()
	}

	if err != lderr.ErrSuccess {
		c.Set(GinKeyError, err)
	}
	c.Set(GinKeyResponse, response)
	c.AbortWithStatusJSON(err.Status(), response)
}

func (c *Context) setHandler(h string) { c.handler = h }
func (c *Context) GetHandler() string  { return c.handler }

func (c *Context) setPath(p string) { c.path = p }
func (c *Context) GetPath() string  { return c.path }

func (c *Context) setMethod(m string) { c.method = m }
func (c *Context) GetMethod() string  { return c.method }

func (c *Context) GetBeginTime() time.Time { return c.beginTime }
func (c *Context) GetSequence() string     { return c.sequence }

func (c *Context) GetError() Error {
	v := c.Gin().Value(GinKeyError)
	r, _ := v.(Error)
	return r
}

func (c *Context) GetResponse() *CommResponse {
	v := c.Gin().Value(GinKeyResponse)
	r, _ := v.(*CommResponse)
	return r
}
