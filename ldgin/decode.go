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
func shouldBind(ctx Context, req interface{}) Error {
	g := ctx.Gin()

	reqV := reflect.ValueOf(req)
	if reqV.Kind() != reflect.Ptr || reqV.Elem().Kind() != reflect.Struct {
		ctx.LogE("input request type must be pointer to struct", zap.Stringer("type", reqV.Kind()))
		return lderr.ErrInvalidRequestType
	}

	reqV = reqV.Elem()
	reqT := getRequestType(reqV.Type())

	if fields := reqT.FormFields; len(fields) != 0 {
		reqNew := newRequest(reqT)
		if err := g.ShouldBindQuery(reqNew.Interface()); err != nil {
			ctx.LogE("ShouldBindQuery() fail", zap.Error(err))
			delRequest(reqT, reqNew)
			return lderr.ErrParseRequest
		}

		fillHttpRequestByFeilds(reqV, reqNew.Elem(), fields)
		delRequest(reqT, reqNew)
	}

	if fields := reqT.JsonFields; len(fields) != 0 {
		reqNew := newRequest(reqT)
		if err := g.ShouldBindJSON(reqNew.Interface()); err != nil {
			ctx.LogE("ShouldBindJSON() fail", zap.Error(err))
			delRequest(reqT, reqNew)
			return lderr.ErrParseRequest
		}

		fillHttpRequestByFeilds(reqV, reqNew.Elem(), fields)
		delRequest(reqT, reqNew)
	}

	if fields := reqT.UriFields; len(fields) != 0 {
		reqNew := newRequest(reqT)
		if err := g.ShouldBindUri(reqNew.Interface()); err != nil {
			ctx.LogE("ShouldBindUri() fail", zap.Error(err))
			delRequest(reqT, reqNew)
			return lderr.ErrParseRequest
		}

		fillHttpRequestByFeilds(reqV, reqNew.Elem(), fields)
		delRequest(reqT, reqNew)
	}

	if fields := reqT.HeaderFields; len(fields) != 0 {
		reqNew := newRequest(reqT)
		if err := g.ShouldBindHeader(reqNew.Interface()); err != nil {
			ctx.LogE("ShouldBindHeader() fail", zap.Error(err))
			delRequest(reqT, reqNew)
			return lderr.ErrParseRequest
		}

		fillHttpRequestByFeilds(reqV, reqNew.Elem(), fields)
		delRequest(reqT, reqNew)
	}

	return nil
}
