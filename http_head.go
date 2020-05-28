package http

import "net/http"

func Head(url string) *Response {
	resp, err := http.Head(url)
	return &Response{err: err, resp: resp}
}
