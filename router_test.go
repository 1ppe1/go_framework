// router_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// dummyHandlerは単純なレスポンスを返すハンドラです
func dummyHandler(ctx *Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"result": "success"})
}

// 正常なルートがマッチするかテストします
func TestRouter_ServeHTTP_Found(t *testing.T) {
	router := NewRouter()
	router.Add("GET", "/test", dummyHandler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	// レスポンスボディに "success" が含まれているかを確認
	bodyBytes := rec.Body.Bytes()
	assert.Contains(t, string(bodyBytes), "success")
}

// 存在しないルートの場合、404が返るかテストします
func TestRouter_ServeHTTP_NotFound(t *testing.T) {
	router := NewRouter()
	router.Add("GET", "/test", dummyHandler)

	req := httptest.NewRequest("GET", "/nonexistent", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := rec.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}
