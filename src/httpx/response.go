package httpx

import (
	"io"
)

type Response struct {
	Body   io.WriteCloser
	header Header
}

func NewResponse(body io.WriteCloser) *Response {
	return &Response{
		Body:   body,
		header: Header{},
	}
}
