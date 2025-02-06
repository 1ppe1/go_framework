package main

import (
	"net/http"
)

// UserCreateRequest：リクエストボディの構造体定義（コードファーストアプローチ）
type UserCreateRequest struct {
	Name  string `json:"name"`  // 本来は validate:"required,min=3" 等を付与
	Email string `json:"email"` // 本来は validate:"email" 等
}

// createUser：ユーザー作成用ハンドラ
func createUser(ctx *Context) error {
	var req UserCreateRequest
	if err := ctx.BindAndValidate(&req); err != nil {
		return err
	}
	// バリデーションチェック（簡略化：実際には詳細な検証を実施）
	if len(req.Name) < 3 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Name must be at least 3 characters"})
	}
	// ユーザー作成処理（DBとの連携等はここで実施）
	response := map[string]string{
		"message": "User created successfully",
		"name":    req.Name,
		"email":   req.Email,
	}
	return ctx.JSON(http.StatusOK, response)
}
