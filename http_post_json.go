package http

import (
	"bytes"
	"encoding/json"
)

// PostJson issues a POST to the specified URL, with data
// json encoding as the request body.
func PostJson(url string, data map[string]interface{}) *Response {
	body, err := json.Marshal(data)
	if err != nil {
		return &Response{err: err}
	}
	return Post(url, "application/json", bytes.NewReader(body))
}
