/*
 * Copyright (C) distroy
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/distroy/ldgo/ldctx"
	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldgin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	StdContext = context.Context
)

var (
	ErrTestOneError = lderr.New(http.StatusOK, 1, "test 1")
)

func testOneError(c *ldgin.Context) ldgin.Error {
	c.LogI("", zap.String("method", c.GetMethod()), zap.String("path", c.GetPath()),
		zap.String("handler", c.GetHandler()))
	return ErrTestOneError
}

type testBindReq struct {
	Uri      string `uri:"uri"`
	Query1   string `form:"query1"`
	Query2   int64  `form:"query2"`
	Language string `header:"accept-language"`
}

func testBind(ctx StdContext, req *testBindReq) (*testBindReq, ldgin.Error) {
	g := ldgin.GetGin(ctx)
	c := ldgin.GetContext(g)
	c.LogI("", zap.String("method", c.GetMethod()), zap.String("path", c.GetPath()),
		zap.String("handler", c.GetHandler()))
	return req, nil
}

type testRenderer struct{}

func (_ *testRenderer) Render(c *gin.Context) {
	c.JSON(http.StatusOK, "abc")
}

func testRender(c *ldgin.Context) (*testRenderer, ldgin.Error) {
	return &testRenderer{}, nil
}

type testValidateReq struct {
	Valid int64 `form:"valid"`
}

func (req *testValidateReq) Validate(c StdContext) ldgin.Error {
	if req.Valid != 0 {
		return lderr.New(http.StatusOK, 111, fmt.Sprintf("invalid requet. valid=%v", req.Valid))
	}
	return nil
}

func testValidate(c *ldgin.Context, req *testValidateReq) ldgin.Error {
	return nil
}

type testParseReq struct {
	Query1 string `form:"query1"`
	Query2 int64  `form:"query2"`
}

func (req *testParseReq) Parse(c *ldgin.Context) ldgin.Error {
	if err := c.ShouldBindQuery(req); err != nil {
		c.LogE("ShouldBindQuery() fail", zap.Error(err))
		return lderr.ErrParseRequest
	}
	return nil
}

func testParse(c *ldgin.Context, req *testParseReq) (*testParseReq, ldgin.Error) {
	return req, nil
}

func testSucc(c *ldgin.Context) ldgin.Error {
	return nil
}

func testPanic(c *ldgin.Context) ldgin.Error {
	c.LogI("", zap.String("method", c.GetMethod()), zap.String("path", c.GetPath()),
		zap.String("handler", c.GetHandler()))

	var p *int
	*p = 1
	return nil
}

func initRouter(c ldctx.Context, router gin.IRouter) {
	r := ldgin.WrapGin(router)

	r = r.WithAppPath("aaaaaa/")
	r = r.Group("/test", midware1)
	r.GET("panic", testPanic)

	r = r.Use(midware2)
	r.GET("/one_error", testOneError)
	r.GET("/bind/:uri", testBind)
	r.GET("/render/", testRender)
	r.GET("/validate", testValidate)
	r.GET("/parse", testParse)
	r.GET("/midware_error", testSucc, midware3)
}

func main() {
	ctx := ldctx.Default()

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
