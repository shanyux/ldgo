/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/gin-gonic/gin"
)

type Error = lderr.Error

type Parser interface {
	Parse(Context) Error
}

type Validator interface {
	Validate(Context) Error
}

type ParseValidator interface {
	Parser
	Validator
}

type Renderer interface {
	Render(Context)
}

type GinParser interface {
	Parse(*gin.Context) Error
}

type GinValidator interface {
	Validate(*gin.Context) Error
}

type GinParseValidator interface {
	GinParser
	GinValidator
}

type GinRenderer interface {
	Render(*gin.Context)
}

// Request must be:
// ParseValidator
// Parser
// Validator
// GinParseValidator
// GinParser
// GinValidator
// interface{}
type Request interface{}

// Response must be:
// Renderer
// GinRenderer
// interface{}
type Response interface{}

// Handler must be:
// func (*gin.Context)
// func (*gin.Context) Error
// func (*gin.Context, Request) Error
// func (*gin.Context) (Response, Error)
// func (*gin.Context, Request) (Response, Error)
// func (*Context)
// func (*Context) Error
// func (*Context, Request) Error
// func (*Context) (Response, Error)
// func (*Context, Request) (Response, Error)
type Handler interface{}

// Midware must be:
// func (*gin.Context)
// func (*gin.Context, Request) Error
// func (*Context)
// func (*Context, Request) Error
type Midware interface{}

type CommResponse struct {
	ErrCode    int         `json:"code"`
	ErrMsg     string      `json:"msg"`
	ErrDetails []string    `json:"details,omitempty"`
	Cost       string      `json:"cost"`
	Sequence   string      `json:"sequence"`
	Data       interface{} `json:"data"`
}

type (
	ginCtx = gin.Context
	ldCtx  = ldctx.Context
)

type routerAdapter interface {
	Group(relativePath string, midwares ...Midware) routerAdapter
	Use(midwares ...Midware) routerAdapter

	WithAppPath(path string) routerAdapter
	BasePath() string

	Handle(method, path string, handler Handler, midwares ...Midware) routerAdapter
}
