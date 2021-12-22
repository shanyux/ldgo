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
		router: router,
	})
}

type ginWapperBase struct {
	appPath  string
	basePath string
	midwares midwares
}

func (w *ginWapperBase) setAppPath(path string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	w.appPath = path
}

func (w *ginWapperBase) BasePath() string {
	if len(w.basePath) != 0 {
		return w.basePath
	}
	return "/"
}

func (w *ginWapperBase) combineMidwares(midwares []Midware) midwares {
	return w.midwares.CombineMidwares(midwares)
}

func (w *ginWapperBase) calculateFullPath(relativePath string) string {
	return w.joinPath(w.appPath, w.basePath, relativePath)
}

func (w *ginWapperBase) calculateAbsolutePath(relativePath string) string {
	return w.joinPath(w.basePath, relativePath)
}

func (w *ginWapperBase) joinPath(elems ...string) string {
	count := 0
	for _, v := range elems {
		if len(v) > 0 {
			count++
		}
	}
	if count < len(elems) {
		i := 0
		for _, v := range elems {
			if len(v) > 0 {
				elems[i] = v
				i++
			}
		}
		elems = elems[:count]
	}

	if count == 0 {
		return "/"
	}

	last := elems[count-1]
	final := path.Join(elems...)

	if !strings.HasPrefix(final, "/") {
		final = "/" + final
	}
	if strings.HasSuffix(last, "/") && !strings.HasSuffix(final, "/") {
		final = final + "/"
	}
	return final
}

type ginWapper struct {
	ginWapperBase

	router gin.IRouter
}

func (w *ginWapper) clone() *ginWapper {
	c := *w
	return &c
}

func (w *ginWapper) WithAppPath(path string) routerBase {
	w = w.clone()
	w.setAppPath(path)
	return w
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
	fullPath := w.calculateFullPath(path)
	absPath := w.calculateAbsolutePath(path)
	midwares := w.combineMidwares(ms).WithMethod(method, fullPath)
	w.router.Handle(method, absPath, midwares.Get(h)...)
	return w
}
