package main

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
	}
}

// JSON()：レスポンスにJSONを出力するためのヘルパー
func (c *Context) JSON(status int, data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	return json.NewEncoder(c.Response).Encode(data)
}

// BindAndValidate()：リクエストボディをデコード（本来はバリデーションロジックを追加すべき）
func (c *Context) BindAndValidate(v interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}
