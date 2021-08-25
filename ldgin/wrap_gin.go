/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func WrapGin(router gin.IRouter) Router {
	return wrapRouter(&ginWapper{
		ginWapperBase: ginWapperBase{
			basePath: "/",
		},
		router: router,
	})
}

type ginWapperBase struct {
	basePath string
	midwares []gin.HandlerFunc
}

func (w *ginWapperBase) combineMidwares(midwares []Midware) []gin.HandlerFunc {
	if len(midwares) == 0 {
		return w.midwares
	}

	finalSize := len(w.midwares) + len(midwares)

	merged := make([]gin.HandlerFunc, len(w.midwares), finalSize)
	copy(merged, w.midwares)
	for _, m := range midwares {
		merged = append(merged, WrapMidware(m))
	}
	return merged
}

func (w *ginWapperBase) combineHandlerMidwares(handler Handler, midwares []Midware) []gin.HandlerFunc {
	finalSize := len(w.midwares) + len(midwares) + 1

	merged := make([]gin.HandlerFunc, len(w.midwares), finalSize)
	copy(merged, w.midwares)
	for _, m := range midwares {
		merged = append(merged, WrapMidware(m))
	}
	merged = append(merged, WrapHandler(handler))
	return merged
}

func (w *ginWapperBase) calculateAbsolutePath(relativePath string) string {
	absolutePath := w.basePath

	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if strings.HasSuffix(relativePath, "/") && !strings.HasSuffix(finalPath, "/") {
		return finalPath + "/"
	}

	return finalPath
}

type ginWapper struct {
	ginWapperBase

	router gin.IRouter
}

func (w *ginWapper) clone() *ginWapper {
	c := *w
	return &c
}

func (w *ginWapper) Group(relativePath string, midwares ...Midware) routerBase {
	w = w.clone()
	w.basePath = w.calculateAbsolutePath(relativePath)
	w.midwares = w.combineMidwares(midwares)
	return w
}

func (w *ginWapper) Use(midwares ...Midware) routerBase {
	w = w.clone()
	w.midwares = w.combineMidwares(midwares)
	return w
}

func (w *ginWapper) Handle(method, path string, h Handler, ms ...Midware) routerBase {
	w.router.Handle(method, w.calculateAbsolutePath(path), w.combineHandlerMidwares(h, ms)...)
	return w
}
