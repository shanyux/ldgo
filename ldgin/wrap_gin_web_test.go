/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

// type GinWebRouter interface {
// 	Handle(method, path string, handler Handler, midwares ...gin.HandlerFunc)
// }

type testGinWebRouterStruct struct {
	handleFunc func(method, path string, handler Handler, midwares ...gin.HandlerFunc)
}

func (r *testGinWebRouterStruct) Handle(method, path string, handler Handler, midwares ...gin.HandlerFunc) {
	r.handleFunc(method, path, handler, midwares...)
}

func TestGinWebRouter(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		var method string
		var path string
		var handler interface{}
		var midwares []gin.HandlerFunc
		router := &testGinWebRouterStruct{
			handleFunc: func(m, p string, h Handler, ms ...gin.HandlerFunc) {
				method, path, handler, midwares = m, p, h, ms
			},
		}

		r := WrapGinWeb(router)
		convey.So(r.BasePath(), convey.ShouldEqual, "/")

		r = r.Group("test")
		convey.So(r.BasePath(), convey.ShouldEqual, "/test")

		r.DELETE("a", func(Context) {})
		convey.So(method, convey.ShouldEqual, http.MethodDelete)
		convey.So(path, convey.ShouldEqual, "/test/a")
		convey.So(handler, convey.ShouldNotBeNil)
		convey.So(midwares, convey.ShouldHaveLength, 0)

		r.GET("b", func(Context) {}, func(Context) {})
		convey.So(method, convey.ShouldEqual, http.MethodGet)
		convey.So(path, convey.ShouldEqual, "/test/b")
		convey.So(handler, convey.ShouldNotBeNil)
		convey.So(midwares, convey.ShouldHaveLength, 1)

		r = r.Use(func(Context) {})
		r.GET("c", func(Context) {}, func(Context) {})
		convey.So(method, convey.ShouldEqual, http.MethodGet)
		convey.So(path, convey.ShouldEqual, "/test/c")
		convey.So(handler, convey.ShouldNotBeNil)
		convey.So(midwares, convey.ShouldHaveLength, 2)
	})
}
