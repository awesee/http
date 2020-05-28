package http

import "net/http"

// Head issues a HEAD to the specified URL.
func Head(url string) *Response {
	resp, err := http.Head(url)
	return &Response{err: err, resp: resp}
}
