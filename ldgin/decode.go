/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"reflect"

	"github.com/distroy/ldgo/lderr"
	"go.uber.org/zap"
)

// shouldBind will decode header/uri/json/query(form)
func shouldBind(ctx *Context, req interface{}) Error {
	g := ctx.Gin()

	reqV := reflect.ValueOf(req)
	if reqV.Kind() != reflect.Ptr || reqV.Elem().Kind() != reflect.Struct {
		ctx.LogE("input request type must be pointer to struct", zap.Stringer("type", reqV.Kind()))
		return lderr.ErrInvalidRequestType
	}

	reqV = reqV.Elem()
	reqT := getRequestType(reqV.Type())

	for _, bind := range reqT.Binds {
		if len(bind.Fields) == 0 {
			continue
		}

		reqNew := newRequest(reqT)
		if err := bind.Func(g, reqNew.Interface()); err != nil {
			ctx.LogE("ShouldBind() fail", zap.String("tag", bind.Tag), zap.Error(err))
			delRequest(reqT, reqNew)
			return lderr.ErrParseRequest
		}

		fillHttpRequestByFeilds(reqV, reqNew.Elem(), bind.Fields)
		delRequest(reqT, reqNew)
	}

	return nil
}
