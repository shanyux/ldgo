/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/distroy/ldgo/lderr"
	"go.uber.org/zap"
)

var chunkedBufferPool = &sync.Pool{}

type chunkedBuffer struct {
	Chunk  []byte
	Buffer *bytes.Buffer
}

func getChunkedBuffer() *chunkedBuffer {
	b, _ := chunkedBufferPool.Get().(*chunkedBuffer)
	if b == nil {
		size := chunkedBufferSize
		b = &chunkedBuffer{
			Chunk:  make([]byte, size-chunkedReserveSize),
			Buffer: bytes.NewBuffer(make([]byte, 0, size)),
		}
	}

	b.Buffer.Reset()
	return b
}

func pubChunkedBuffer(b *chunkedBuffer) {
	chunkedBufferPool.Put(b)
}

type ChunkedRenderer struct {
	Headers          map[string]string // optional.
	ContentType      string            // optional. default=not set
	TransferEncoding string            // optional. default=chunked
	Code             int               // optional. default=http.StatusOK
	Reader           io.Reader         // required.
}

func (r ChunkedRenderer) Render(c *Context) {
	reader := r.Reader

	buf := getChunkedBuffer()

	defer func() {
		pubChunkedBuffer(buf)
		if closer, _ := reader.(io.Closer); closer != nil {
			closer.Close()
		}
	}()

	r.writeHeaders(c)

	for {
		n, err := reader.Read(buf.Chunk)
		if err == io.EOF {
			e := r.writeChunk(c, buf, nil)
			c.setError(e)
			break

		} else if err != nil {
			c.LogE("[ldgin chunked] read chunk fail", zap.Error(err))
			e := lderr.WithDetail(lderr.ErrHttpRenderBody, err.Error())
			c.setError(e)
			writeError(c, e)
			return

		} else if n == 0 {
			continue
		}

		chunk := buf.Chunk[:n]
		if err := r.writeChunk(c, buf, chunk); err != nil {
			c.setError(err)
			// c.LogE("[ldgin chunked] write chunk fail", zap.ByteString("chunk", chunk), zap.Error(err))
			return
		}
	}
}

func (r ChunkedRenderer) writeChunk(c *Context, buf *chunkedBuffer, chunk []byte) lderr.Error {
	g := c.Gin()
	writer := g.Writer

	size := len(chunk)

	buf.Buffer.Reset()

	fmt.Fprintf(buf.Buffer, "%x", size)
	// fmt.Fprint(buf.Buffer, lineEnd)
	buf.Buffer.WriteString(lineEnd)
	if size > 0 {
		// fmt.Fprint(buf.Buffer, chunk)
		buf.Buffer.Write(chunk)
	}
	// fmt.Fprint(buf.Buffer, lineEnd)
	buf.Buffer.WriteString(lineEnd)

	if _, err := io.Copy(writer, buf.Buffer); err != nil {
		c.LogE("[ldgin chunked] write chunk fail", zap.Int("size", size), zap.ByteString("chunk", chunk), zap.Error(err))
		return lderr.WithDetail(lderr.ErrHttpRenderBody, err.Error())
	}

	return nil
}

func (r ChunkedRenderer) writeHeaders(c *Context) {
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

	if r.TransferEncoding != "" {
		header.Set(chunkedHeaderKey, r.TransferEncoding)
	} else {
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

	fmt.Fprintf(writer, "code: %d%s", err.Code(), lineEnd)
	fmt.Fprintf(writer, "message: %s%s", err.Error(), lineEnd)

	var details []string
	if e, _ := err.(lderr.ErrorWithDetails); e != nil {
		details = e.Details()
	}

	if len(details) == 0 {
		return
	}

	fmt.Fprintf(writer, "details:%s", lineEnd)
	for _, v := range details {
		fmt.Fprintf(writer, "    %s%s", v, lineEnd)
	}
}
