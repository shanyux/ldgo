/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"github.com/distroy/ldgo/ldcontext"
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
// func (*gin.Context, Request) Error
// func (*gin.Context) (Response, Error)
// func (*gin.Context, Request) (Response, Error)
// func (Context)
// func (Context, Request) Error
// func (Context) (Response, Error)
// func (Context, Request) (Response, Error)
type Handler interface{}

// Midware must be:
// func (*gin.Context)
// func (*gin.Context, Request) Error
// func (Context)
// func (Context, Request) Error
type Midware interface{}

type CommResponse struct {
	ErrCode  int         `json:"code"`
	ErrMsg   string      `json:"msg"`
	Cost     string      `json:"cost"`
	Sequence string      `json:"sequence"`
	Data     interface{} `json:"data"`
}

type (
	ginContext = gin.Context
	ldContext  = ldcontext.Context
)
