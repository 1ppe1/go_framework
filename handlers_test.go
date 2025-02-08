// handlers_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 正常なユーザー作成リクエストのテスト
func TestCreateUser_Success(t *testing.T) {
	userData := map[string]string{
		"name":  "Alice",
		"email": "alice@example.com",
	}
	body, _ := json.Marshal(userData)
	req := httptest.NewRequest("POST", "/user/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := NewContext(rec, req)

	err := createUser(ctx)
	assert.NoError(t, err)

	result := rec.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	var response map[string]string
	err = json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", response["message"])
	assert.Equal(t, "Alice", response["name"])
	assert.Equal(t, "alice@example.com", response["email"])
}

// 名前が短い場合（2文字）のエラーをテスト
func TestCreateUser_Failure_ShortName(t *testing.T) {
	userData := map[string]string{
		"name":  "Al", // 2文字なのでエラーになるはず
		"email": "al@example.com",
	}
	body, _ := json.Marshal(userData)
	req := httptest.NewRequest("POST", "/user/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := NewContext(rec, req)

	err := createUser(ctx)
	// ハンドラ内でエラー時にもJSONレスポンスが返されるので、errはnilであることも多い
	assert.NoError(t, err)

	result := rec.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)

	var response map[string]string
	err = json.NewDecoder(rec.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Name must be at least 3 characters")
}
