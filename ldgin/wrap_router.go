/*
 * Copyright (C) distroy
 */

package ldgin

import "net/http"

// Router is http router
type Router struct {
	router routerAdapter
}

func wrapRouter(r routerAdapter) Router {
	return Router{
		router: r,
	}
}

func (w Router) Group(relativePath string, midwares ...Midware) Router {
	w.router = w.router.Group(relativePath, midwares...)
	return w
}

func (w Router) Use(midwares ...Midware) Router {
	w.router = w.router.Use(midwares...)
	return w
}

func (w Router) WithAppPath(path string) Router {
	w.router = w.router.WithAppPath(path)
	return w
}

func (w Router) BasePath() string {
	return w.router.BasePath()
}

func (w Router) Handle(method, path string, h Handler, ms ...Midware) {
	w.router = w.router.Handle(method, path, h, ms...)
}

func (w Router) GET(p string, h Handler, ms ...Midware)    { w.Handle(http.MethodGet, p, h, ms...) }
func (w Router) DELETE(p string, h Handler, ms ...Midware) { w.Handle(http.MethodDelete, p, h, ms...) }
func (w Router) HEAD(p string, h Handler, ms ...Midware)   { w.Handle(http.MethodHead, p, h, ms...) }
func (w Router) PATCH(p string, h Handler, ms ...Midware)  { w.Handle(http.MethodPatch, p, h, ms...) }
func (w Router) POST(p string, h Handler, ms ...Midware)   { w.Handle(http.MethodPost, p, h, ms...) }
func (w Router) PUT(p string, h Handler, ms ...Midware)    { w.Handle(http.MethodPut, p, h, ms...) }
func (w Router) OPTIONS(p string, h Handler, ms ...Midware) {
	w.Handle(http.MethodConnect, p, h, ms...)
}
