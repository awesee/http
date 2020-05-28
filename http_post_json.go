package http

import (
	"bytes"
	"encoding/json"
)

func PostJson(url string, data map[string]interface{}) *Response {
	body, err := json.Marshal(data)
	if err != nil {
		return &Response{err: err}
	}
	return Post(url, "application/json", bytes.NewReader(body))
}
