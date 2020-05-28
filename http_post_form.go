package http

import (
	"net/http"
	"net/url"
)

func PostForm(url string, data url.Values) *Response {
	resp, err := http.PostForm(url, data)
	return &Response{err: err, resp: resp}
}
