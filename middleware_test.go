// middleware_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	counter := 0
	// ミドルウェア1: カウンターを1増やす
	mw1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			counter++
			next.ServeHTTP(w, r)
		})
	}
	// ミドルウェア2: カウンターをさらに1増やす
	mw2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			counter++
			next.ServeHTTP(w, r)
		})
	}

	// 最終ハンドラ: "done" を返す
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("done"))
	})
	wrappedHandler := Chain(finalHandler, mw1, mw2)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(rec, req)

	assert.Equal(t, "done", rec.Body.String())
	// 2つのミドルウェアが正しく実行されると、counterは2になるはず
	assert.Equal(t, 2, counter)
}
