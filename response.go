package http

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Response represents the response from an HTTP request.
type Response struct {
	err  error
	resp *http.Response
	body []byte
}

// Err returns the response err.
func (r *Response) Err() error {
	return r.err
}

// Response returns the response.
func (r *Response) Response() *http.Response {
	return r.resp
}

// Val returns the response body as []byte.
func (r *Response) Val() []byte {
	if r.err == nil && r.body == nil {
		r.body, r.err = ioutil.ReadAll(r.resp.Body)
		defer r.resp.Body.Close()
	}
	return r.body
}

// String returns the response body as string.
func (r *Response) String() string {
	return string(r.Val())
}

// Result returns the response val and err.
func (r *Response) Result() ([]byte, error) {
	return r.Val(), r.err
}

// JsonUnmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func (r *Response) JsonUnmarshal(v interface{}) error {
	data := r.Val()
	if r.err != nil {
		return r.err
	}
	return json.Unmarshal(data, v)
}

// XmlUnmarshal parses the XML-encoded data and stores the result in
// the value pointed to by v.
func (r *Response) XmlUnmarshal(v interface{}) error {
	data := r.Val()
	if r.err != nil {
		return r.err
	}
	return xml.Unmarshal(data, v)
}
