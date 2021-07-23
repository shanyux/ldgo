/*
 * Copyright (C) distroy
 */

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldgin"
	"github.com/gin-gonic/gin"
)

var (
	ERR_TEST1 = ldgin.NewError(http.StatusOK, 1, "test 1")
)

func test1(ctx ldgin.Context) ldgin.Error {
	return ERR_TEST1
}

type test2Req struct {
	Uri      string `uri:"uri"`
	Query1   string `form:"query1"`
	Query2   int64  `form:"query2"`
	Language string `header:"accept-language"`
}

func test2(ctx context.Context, req *test2Req) (*test2Req, ldgin.Error) {
	return req, nil
}

type test3Renderer struct {
}

func (_ *test3Renderer) Render(ctx ldgin.Context) {
	ctx.JSON(http.StatusOK, "abc")
}

func test3(ctx ldgin.Context) (*test3Renderer, ldgin.Error) {
	return &test3Renderer{}, nil
}

func initRouter(ctx ldcontext.Context, r gin.IRouter) {
	r.GET("/test1/:xxx", ldgin.Wrap(test1))
	r.GET("/test2/:uri", ldgin.Wrap(test2))
	r.GET("/test3/:renderer", ldgin.Wrap(test3))
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
