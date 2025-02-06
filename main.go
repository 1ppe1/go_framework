package main

import (
	"log"
	"net/http"
)

func main() {
	// ルーターの初期化
	router := NewRouter()

	// ルート登録（メソッド情報は実装上簡略化のため未利用ですが、拡張性として保持）
	router.Add("POST", "/user/create", createUser)

	// ミドルウェアの定義例
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Request:", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			if r.Method == http.MethodOptions {
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// ルーターを http.Handler に変換してミドルウェアチェーンを適用
	finalHandler := http.HandlerFunc(router.ServeHTTP)
	wrappedHandler := Chain(finalHandler, loggingMiddleware, corsMiddleware)

	// サーバ起動
	log.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", wrappedHandler); err != nil {
		log.Fatal(err)
	}
}
