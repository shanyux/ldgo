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

const (
	lineEnd = "\r\n"
)

var chunkedBufferPool = &sync.Pool{}

type chunkedBuffer struct {
	Chunk  []byte
	Buffer *bytes.Buffer
}

func getChunkedBuffer() *chunkedBuffer {
	b, _ := chunkedBufferPool.Get().(*chunkedBuffer)
	if b == nil {
		size := 4096
		b = &chunkedBuffer{
			Chunk:  make([]byte, size-32),
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

func (r ChunkedRenderer) Render(c Context) {
	reader := r.Reader
	closer := reader.(io.ReadCloser)

	buf := getChunkedBuffer()

	defer func() {
		pubChunkedBuffer(buf)
		if closer != nil {
			closer.Close()
		}
	}()

	r.writeHeaders(c)

	for {
		n, err := reader.Read(buf.Chunk)
		if err == io.EOF {
			r.writeChunk(c, buf, nil)
			break

		} else if err != nil {
			c.LogE("read chunk fail", zap.Error(err))
			e := lderr.WithDetail(lderr.ErrHttpRenderBody, err.Error())
			c.AbortWithError(e)
			return
		}

		chunk := buf.Chunk[:n]
		if err := r.writeChunk(c, buf, chunk); err != nil {
			// c.LogE("write chunk fail", zap.ByteString("chunk", chunk), zap.Error(err))
			return
		}
	}
}

func (r ChunkedRenderer) writeChunk(c Context, buf *chunkedBuffer, chunk []byte) lderr.Error {
	writer := c.Writer

	buf.Buffer.Reset()

	size := len(chunk)
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
		c.LogE("write chunk fail", zap.Int("size", size), zap.ByteString("chunk", chunk), zap.Error(err))
		return lderr.WithDetail(lderr.ErrHttpRenderBody, err.Error())
	}

	return nil
}

func (r ChunkedRenderer) writeHeaders(c Context) {
	header := c.Writer.Header()
	for k, v := range r.Headers {
		if v != "" {
			header.Set(k, v)
		}
	}

	if r.ContentType != "" {
		header.Set("Content-Type", r.ContentType)
	}

	if r.TransferEncoding != "" {
		header.Set("Transfer-Encoding", r.TransferEncoding)
	} else {
		header.Set("Transfer-Encoding", "chunked")
	}

	if r.Code > 0 {
		c.AbortWithStatus(r.Code)
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
