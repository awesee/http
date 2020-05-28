package http

import "net/http"

// Get issues a GET to the specified URL.
func Get(url string) *Response {
	resp, err := http.Get(url)
	return &Response{err: err, resp: resp}
}
