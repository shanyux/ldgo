/*
 * Copyright (C) distroy
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldgin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	ErrTestOneError = lderr.NewError(http.StatusOK, 1, "test 1")
)

func testOneError(ctx ldgin.Context) ldgin.Error {
	return ErrTestOneError
}

type testBindReq struct {
	Uri      string `uri:"uri"`
	Query1   string `form:"query1"`
	Query2   int64  `form:"query2"`
	Language string `header:"accept-language"`
}

func testBind(ctx context.Context, req *testBindReq) (*testBindReq, ldgin.Error) {
	return req, nil
}

type testRenderer struct{}

func (_ *testRenderer) Render(ctx ldgin.Context) {
	ctx.JSON(http.StatusOK, "abc")
}

func testRender(ctx ldgin.Context) (*testRenderer, ldgin.Error) {
	return &testRenderer{}, nil
}

type testValidateReq struct {
	Valid int64 `form:"valid"`
}

func (req *testValidateReq) Validate(ctx ldgin.Context) ldgin.Error {
	if req.Valid != 0 {
		return lderr.NewError(http.StatusOK, 111, fmt.Sprintf("invalid requet. valid=%v", req.Valid))
	}
	return nil
}

func testValidate(ctx ldgin.Context, req *testValidateReq) ldgin.Error {
	return nil
}

type testParseReq struct {
	Query1 string `form:"query1"`
	Query2 int64  `form:"query2"`
}

func (req *testParseReq) Parse(ctx ldgin.Context) ldgin.Error {
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.LogI("ShouldBindQuery() fail", zap.Error(err))
		return lderr.ErrParseRequest
	}
	return nil
}

func testParse(ctx ldgin.Context, req *testParseReq) (*testParseReq, ldgin.Error) {
	return req, nil
}

func initRouter(ctx ldcontext.Context, r gin.IRouter) {
	r.GET("/test/one_error", ldgin.Wrap(testOneError))
	r.GET("/test/bind/:uri", ldgin.Wrap(testBind))
	r.GET("/test/render/", ldgin.Wrap(testRender))
	r.GET("/test/validate", ldgin.Wrap(testValidate))
	r.GET("/test/parse", ldgin.Wrap(testParse))
}

func main() {
	ctx := ldcontext.Default()

	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	initRouter(ctx, engine)

	server := &http.Server{
		Addr:    ":8000",
		Handler: engine,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe() fail. err:%v", err)
	}
}
