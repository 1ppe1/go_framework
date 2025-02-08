// adapter_test.go
package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaAdapter_Invoke(t *testing.T) {
	// Invoke メソッドのテスト用に LambdaAdapter を生成
	adapter := &LambdaAdapter{}

	// Invoke を呼び出す
	result, err := adapter.Invoke(context.Background(), []byte("test payload"))

	// 現在のスタブ実装では、常に nil, nil を返すはずなので、それを検証する
	assert.Nil(t, result, "Invoke should return nil result")
	assert.Nil(t, err, "Invoke should return nil error")
}
