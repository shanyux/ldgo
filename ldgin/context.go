/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"encoding/hex"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ctxKeyType int

const (
	ctxKey_GinContext ctxKeyType = iota
)

const (
	GIN_HEADER_SEQUENCE = "x-gin-sequence"
)

const (
	GIN_KEY_BEGIN_TIME = "x-gin-begin-time"
	GIN_KEY_SEQUENCE   = "x-gin-sequence"
	GIN_KEY_CONTEXT    = "x-gin-context"

	GIN_KEY_ERROR    = "x-gin-error"
	GIN_KEY_REQUEST  = "x-gin-request"
	GIN_KEY_RENDERER = "x-gin-renderer"
	GIN_KEY_RESPONSE = "x-gin-response"
)

type (
	ginContext = *gin.Context
	ldContext  = ldcontext.Context
)

type Context interface {
	ldcontext.Context
}

func newContext(g *gin.Context) Context {
	c := ldcontext.NewContext(g)
	c = c.WithValue(ctxKey_GinContext, g)

	now := time.Now()
	g.Set(GIN_KEY_BEGIN_TIME, now)

	uuid, _ := uuid.NewRandom()
	seq := hex.EncodeToString(uuid[:])
	c = c.With(zap.String("sequence", seq))
	g.Header(GIN_HEADER_SEQUENCE, seq)
	g.Set(GIN_KEY_SEQUENCE, seq)

	g.Set(GIN_KEY_CONTEXT, c)

	return c
}

func GetContext(g *gin.Context) Context {
	v := g.Value(GIN_KEY_CONTEXT)
	c, _ := v.(ldContext)
	if c == nil {
		c = newContext(g)
	}
	return c
}

func GetGin(c Context) *gin.Context {
	v := c.Value(ctxKey_GinContext)
	t, ok := v.(*gin.Context)
	if !ok {
		return nil
	}
	return t
}

func GetBeginTime(g *gin.Context) time.Time  { return g.GetTime(GIN_KEY_BEGIN_TIME) }
func GetSequence(g *gin.Context) string      { return g.GetString(GIN_KEY_SEQUENCE) }
func GetRequest(g *gin.Context) interface{}  { return g.Value(GIN_KEY_REQUEST) }
func GetRenderer(g *gin.Context) interface{} { return g.Value(GIN_KEY_RENDERER) }

func GetError(g *gin.Context) Error {
	v := g.Value(GIN_KEY_REQUEST)
	err, _ := v.(Error)
	return err
}

func GetResponse(g *gin.Context) *commResponse {
	v := g.Value(GIN_KEY_CONTEXT)
	rsp, _ := v.(*commResponse)
	return rsp
}
