package httpx

import (
	"io"
	"net/url"
)

type Request struct {
	Method string
	URL    *url.URL
	Body   io.Reader
}

func NewRequest(method string, url_ string, body io.Reader) (*Request, error) {
	parsedUrl, err := url.Parse(url_)
	if err != nil {
		return nil, err
	}
	return &Request{
		Method: method,
		URL:    parsedUrl,
		Body:   body,
	}, nil
}
