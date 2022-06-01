//go:build js && wasm
// +build js,wasm

package httpx

import (
	"syscall/js"
)

type responseWriter struct {
	jsResp js.Value
	*Response
	flushFunc func() error
}

func (resp *responseWriter) Write(b []byte) (int, error) {
	return resp.Body.Write(b)
}

func (resp *responseWriter) WriteHeader(s int) {
	resp.jsResp.Set("status", s)
}

func (resp *responseWriter) Header() Header {
	return resp.header
}

func NewResponseWriter(jresp js.Value) ResponseWriter {
	buf := NewBuffer()
	buf.closerFunc = func(data js.Value) error {
		jresp.Set("body", data)
		return nil
	}
	resp := NewResponse(buf)
	return &responseWriter{
		jsResp:   jresp,
		Response: resp,
		flushFunc: func() (err error) {
			if err = buf.Close(); err != nil {
				return
			}
			resp.header.applyTo(jresp.Get("headers"))
			return
		},
	}
}

func (resp *responseWriter) Close() error {
	if f := resp.flushFunc; f != nil {
		return f()
	}
	return nil
}
