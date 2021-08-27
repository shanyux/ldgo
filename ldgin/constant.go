/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type ctxKeyType int

const (
	_CTX_KEY_GIN_CONTEXT ctxKeyType = iota
	_CTX_KEY_BEGIN_TIME
	_CTX_KEY_SEQUENCE
)

const (
	GIN_HEADER_SEQUENCE = "x-ldgin-sequence"
)

const (
	GIN_KEY_CONTEXT = "x-ldgin-context"

	GIN_KEY_ERROR    = "x-ldgin-error"
	GIN_KEY_REQUEST  = "x-ldgin-request"
	GIN_KEY_RENDERER = "x-ldgin-renderer"
	GIN_KEY_RESPONSE = "x-ldgin-response"
)

var (
	_TYPE_OF_INTERFACE = reflect.TypeOf((*interface{})(nil)).Elem()

	_TYPE_OF_GIN_HANDLER       = reflect.TypeOf((*gin.HandlerFunc)(nil)).Elem()
	_TYPE_OF_GIN_HANDLER_CHAIN = reflect.TypeOf((*gin.HandlersChain)(nil)).Elem()

	_TYPE_OF_GIN_CONTEXT = reflect.TypeOf((*gin.Context)(nil))
	_TYPE_OF_CONTEXT     = reflect.TypeOf((*Context)(nil)).Elem()

	_TYPE_OF_COMM_ERROR = reflect.TypeOf((*error)(nil)).Elem()
	_TYPE_OF_ERROR      = reflect.TypeOf((*Error)(nil)).Elem()

	_TYPE_OF_PARSER          = reflect.TypeOf((*Parser)(nil)).Elem()
	_TYPE_OF_VALIDATER       = reflect.TypeOf((*Validator)(nil)).Elem()
	_TYPE_OF_PARSE_VALIDATOR = reflect.TypeOf((*ParseValidator)(nil)).Elem()

	_TYPE_OF_GIN_PARSER          = reflect.TypeOf((*GinParser)(nil)).Elem()
	_TYPE_OF_GIN_VALIDATER       = reflect.TypeOf((*GinValidator)(nil)).Elem()
	_TYPE_OF_GIN_PARSE_VALIDATOR = reflect.TypeOf((*GinParseValidator)(nil)).Elem()

	_TYPE_OF_RENDERER = reflect.TypeOf((*Renderer)(nil)).Elem()

	_TYPE_OF_GIN_RENDERER = reflect.TypeOf((*GinRenderer)(nil)).Elem()
)
