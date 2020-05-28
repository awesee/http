package http

import (
	"io"
	"net/http"
)

func Post(url, contentType string, body io.Reader) *Response {
	resp, err := http.Post(url, contentType, body)
	return &Response{err: err, resp: resp}
}
