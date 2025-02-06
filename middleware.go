package main

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain()：複数のミドルウェアをチェーンして1つのhttp.Handlerへ変換
func Chain(final http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		final = middlewares[i](final)
	}
	return final
}
