/*
 * Copyright (C) distroy
 */

package ldgin

import "net/http"

type routerWapper struct {
	router routerBase
}

func wrapRouter(r routerBase) Router {
	return routerWapper{
		router: r,
	}
}

func (w routerWapper) Group(relativePath string, midwares ...Midware) Router {
	w.router = w.router.Group(relativePath, midwares...)
	return w
}

func (w routerWapper) Use(midwares ...Midware) Router {
	w.router = w.router.Use(midwares...)
	return w
}

func (w routerWapper) BasePath() string {
	return w.router.BasePath()
}

func (w routerWapper) Handle(method, path string, h Handler, ms ...Midware) Router {
	w.router = w.router.Handle(method, path, h, ms...)
	return w
}

func (w routerWapper) GET(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodGet, path, h, ms...)
}

func (w routerWapper) DELETE(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodDelete, path, h, ms...)
}

func (w routerWapper) HEAD(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodHead, path, h, ms...)
}

func (w routerWapper) OPTIONS(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodConnect, path, h, ms...)
}

func (w routerWapper) PATCH(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodPatch, path, h, ms...)
}

func (w routerWapper) POST(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodPost, path, h, ms...)
}

func (w routerWapper) PUT(path string, h Handler, ms ...Midware) Router {
	return w.Handle(http.MethodPut, path, h, ms...)
}
