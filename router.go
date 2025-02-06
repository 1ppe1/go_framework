package main

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context) error

// ノード構造体（各パスセグメントごとのルーティング情報）
type node struct {
	path     string
	handler  HandlerFunc
	children []*node
}

// Router構造体
type Router struct {
	root *node
}

func NewRouter() *Router {
	return &Router{root: &node{}}
}

// Add()：指定されたメソッド・パスに対するハンドラ登録（ここではメソッドは内部で利用していませんが、拡張時のための引数）
func (r *Router) Add(method, path string, handler HandlerFunc) {
	current := r.root
	segments := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segments {
		var child *node
		for _, c := range current.children {
			if c.path == seg {
				child = c
				break
			}
		}
		if child == nil {
			child = &node{path: seg}
			current.children = append(current.children, child)
		}
		current = child
	}
	current.handler = handler
}

// ServeHTTP()：http.Handlerとしてのエントリーポイント
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.Trim(req.URL.Path, "/")
	segments := strings.Split(path, "/")
	current := r.root
	for _, seg := range segments {
		found := false
		for _, child := range current.children {
			if child.path == seg {
				current = child
				found = true
				break
			}
		}
		if !found {
			http.NotFound(w, req)
			return
		}
	}
	if current.handler == nil {
		http.NotFound(w, req)
		return
	}
	// コンテキスト生成してハンドラ実行
	ctx := NewContext(w, req)
	if err := current.handler(ctx); err != nil {
		// エラー発生時のシンプルなエラーハンドリング
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
