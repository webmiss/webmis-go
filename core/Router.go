package core

import (
	"fmt"
	"net/http"
)

/* 路由类 */
type Router struct {
	// Base
}

/* 请求 */
func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// (&api.Index{}).Index
	mux := http.NewServeMux()
	mux.HandleFunc("/api/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Service :8080")
	})
	switch r.URL.Path {
	case "/":
		fmt.Println("Home")
	case "/api":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Println("Api")
	case "/admin":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Println("Admin")
	default:
		fmt.Println("Not Found")
	}

}
