/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/distroy/ldgo/lderr"
	"go.uber.org/zap"
)

type ReaderRenderer struct {
	Headers       map[string]string // optional.
	Code          int               // optional. default=http.StatusOK
	Chunked       bool              // optional. default=false
	ContentLength int64             // optional. default=0
	ContentType   string            // optional. default=not set
	Reader        io.Reader         // required.
}

func (r ReaderRenderer) Render(c *Context) {
	reader := r.Reader
	defer func() {
		if closer, _ := reader.(io.Closer); closer != nil {
			closer.Close()
		}
	}()

	r.writeHeaders(c)

	writer := c.Gin().Writer
	_, err := io.Copy(writer, reader)
	if err == nil {
		return
	}

	c.LogE("[ldgin] render from reader fail", zap.Error(err))
	e := lderr.WithDetail(lderr.ErrHttpRenderBody, err.Error())
	c.setError(e)

	if r.Chunked || c.Gin().Writer.Header().Get(chunkedHeaderKey) == chunkedHeaderValue {
		writeError(c, e)
		c.CloseConn()
	}
}

func (r ReaderRenderer) writeHeaders(c *Context) {
	g := c.Gin()

	header := g.Writer.Header()
	for k, v := range r.Headers {
		if v != "" {
			header.Set(k, v)
		}
	}

	if r.ContentType != "" {
		header.Set(headerContentType, r.ContentType)
	}

	if r.ContentLength > 0 {
		header.Set(headerContentLength, strconv.FormatInt(r.ContentLength, 10))
	}

	// 设置了 chunked header, http 官方库会处理 chunked 格式，不需要上层处理
	if r.Chunked {
		header.Set(chunkedHeaderKey, chunkedHeaderValue)
	}

	if r.Code > 0 {
		g.AbortWithStatus(r.Code)
	} else {
		g.AbortWithStatus(http.StatusOK)
	}
}

func writeError(c *Context, err lderr.Error) {
	g := c.Gin()
	writer := g.Writer
	defer writer.Flush()

	fmt.Fprint(writer, crlf)
	fmt.Fprint(writer, crlf)

	fmt.Fprintf(writer, "server happened some errors%s", crlf)
	fmt.Fprintf(writer, "code: %d%s", err.Code(), crlf)
	fmt.Fprintf(writer, "message: %s%s", err.Error(), crlf)

	var details []string
	if e, _ := err.(lderr.ErrorWithDetails); e != nil {
		details = e.Details()
	}

	if len(details) == 0 {
		return
	}

	fmt.Fprintf(writer, "details:%s", crlf)
	for _, v := range details {
		fmt.Fprintf(writer, "\t%s%s", v, crlf)
	}
}
