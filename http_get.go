package http

import "net/http"

func Get(url string) *Response {
	resp, err := http.Get(url)
	return &Response{err: err, resp: resp}
}
