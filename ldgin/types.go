/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"github.com/distroy/ldgo/v2/ldctx"
	"github.com/gin-gonic/gin"
)

type Parser interface {
	Parse(*Context) error
}

type Validator interface {
	Validate(*Context) error
}

type ParseValidator interface {
	Parser
	Validator
}

type Renderer interface {
	Render(*Context)
}

type GinParser interface {
	Parse(*gin.Context) error
}

type GinValidator interface {
	Validate(*gin.Context) error
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
// func (*gin.Context) error
// func (*gin.Context, Request) error
// func (*gin.Context) (Response, error)
// func (*gin.Context, Request) (Response, error)
// func (*Context)
// func (*Context) error
// func (*Context, Request) error
// func (*Context) (Response, error)
// func (*Context, Request) (Response, error)
type Handler interface{}

// Midware must be:
// func (*gin.Context)
// func (*gin.Context, Request) error
// func (*Context)
// func (*Context, Request) error
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

	// calculateAbsolutePath(relativePath string) string
	calculateFullPath(relativePath string) string
}
