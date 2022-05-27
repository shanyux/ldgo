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

type testGinRouter struct {
	handleFunc func(method, path string, handler Handler, midwares ...gin.HandlerFunc)
}

func (r *testGinRouter) Handle(method string, path string, handlerAndMidWares ...gin.HandlerFunc) gin.IRoutes {
	var handler gin.HandlerFunc
	var midwares []gin.HandlerFunc
	if len(handlerAndMidWares) > 0 {
		handler = handlerAndMidWares[0]
	}
	if len(handlerAndMidWares) > 1 {
		midwares = handlerAndMidWares[1:]
	}
	r.handleFunc(method, path, handler, midwares...)
	return r
}
func (r *testGinRouter) Use(...gin.HandlerFunc) gin.IRoutes { return r }

func (r *testGinRouter) Group(string, ...gin.HandlerFunc) *gin.RouterGroup { return nil }

func (r *testGinRouter) Any(string, ...gin.HandlerFunc) gin.IRoutes     { return r }
func (r *testGinRouter) GET(string, ...gin.HandlerFunc) gin.IRoutes     { return r }
func (r *testGinRouter) POST(string, ...gin.HandlerFunc) gin.IRoutes    { return r }
func (r *testGinRouter) DELETE(string, ...gin.HandlerFunc) gin.IRoutes  { return r }
func (r *testGinRouter) PATCH(string, ...gin.HandlerFunc) gin.IRoutes   { return r }
func (r *testGinRouter) PUT(string, ...gin.HandlerFunc) gin.IRoutes     { return r }
func (r *testGinRouter) OPTIONS(string, ...gin.HandlerFunc) gin.IRoutes { return r }
func (r *testGinRouter) HEAD(string, ...gin.HandlerFunc) gin.IRoutes    { return r }

func (r *testGinRouter) StaticFile(string, string) gin.IRoutes        { return r }
func (r *testGinRouter) Static(string, string) gin.IRoutes            { return r }
func (r *testGinRouter) StaticFS(string, http.FileSystem) gin.IRoutes { return r }

func TestWrapGin(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		testRouter(func(f func(method, path string, handler Handler, midwares ...gin.HandlerFunc)) Router {
			router := &testGinRouter{
				handleFunc: f,
			}
			return WrapGin(router)
		})
	})
}
