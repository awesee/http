package http

import (
	"net/http"
	"net/url"
)

// PostForm issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body.
func PostForm(url string, data url.Values) *Response {
	resp, err := http.PostForm(url, data)
	return &Response{err: err, resp: resp}
}
