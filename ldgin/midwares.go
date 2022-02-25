/*
 * Copyright (C) distroy
 */

package ldgin

import "github.com/gin-gonic/gin"

type midwares struct {
	method   string
	path     string
	midwares []*midware
}

func (w midwares) WithMethod(method, path string) midwares {
	w.method = method
	w.path = path
	return w
}

func (w midwares) CombineMidwares(midwares []Midware) midwares {
	if len(midwares) == 0 {
		return w
	}

	finalSize := len(w.midwares) + len(midwares)
	merged := make([]*midware, len(w.midwares), finalSize)
	copy(merged, w.midwares)
	for _, m := range midwares {
		merged = append(merged, wrapMidware(m))
	}

	w.midwares = merged
	return w
}

func (w midwares) Get(handlers ...Handler) []gin.HandlerFunc {
	count := len(w.midwares)
	var h *handler
	if len(handlers) > 0 && handlers[0] != nil {
		count++
		h = wrapHandler(handlers[0])
	}

	buf := make([]gin.HandlerFunc, 0, count)
	for _, m := range w.midwares {
		c := *m
		m = &c
		m.setMethod(w.method)
		m.setPath(w.path)
		buf = append(buf, m.Do)
	}

	if h != nil {
		h.setMethod(w.method)
		h.setPath(w.path)
		buf = append(buf, h.Do)
	}
	return buf
}
