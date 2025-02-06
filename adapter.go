package main

import (
	"context"
	"net/http"
)

type Adapter interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// AWS Lambdaアダプター例（実際のイベント変換などは省略）
type LambdaAdapter struct {
	handler http.Handler
}

func (la *LambdaAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	la.handler.ServeHTTP(w, r)
}

// Lambda環境での呼び出し用メソッド（実装例：リクエスト変換、レスポンス生成などのロジックを追加）
func (la *LambdaAdapter) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	// ここでpayloadをHTTPリクエストに変換し、handlerを呼び出すロジックを実装
	return nil, nil
}
