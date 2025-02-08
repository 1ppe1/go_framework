// context_test.go
package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext_JSON(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	ctx := NewContext(rec, req)

	// テスト用データ
	data := map[string]string{"key": "value"}
	err := ctx.JSON(200, data)
	assert.NoError(t, err)

	// 出力されたJSONが正しいか確認
	var result map[string]string
	err = json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "value", result["key"])
}
