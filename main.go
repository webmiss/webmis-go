package main

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/modules/admin"
	"webmis/app/modules/api"
	"webmis/app/modules/home"
	"webmis/core"
)

func main() {
	// 路由
	mux := http.NewServeMux()
	// 网站
	mux.HandleFunc("/", (&home.Index{}).Index)
	mux.HandleFunc("/home", (&home.Index{}).Index)
	mux.HandleFunc("/home/index", (&home.Index{}).Index)
	mux.HandleFunc("/home/index/index", (&home.Index{}).Index)
	// API
	mux.HandleFunc("/api", (&api.Index{}).Index)
	mux.HandleFunc("/api/index", (&api.Index{}).Index)
	mux.HandleFunc("/api/index/index", (&api.Index{}).Index)
	// 后台
	mux.HandleFunc("/admin", (&admin.Index{}).Index)
	mux.HandleFunc("/admin/index", (&admin.Index{}).Index)
	mux.HandleFunc("/admin/index/index", (&admin.Index{}).Index)
	// 启动
	cfg := config.Env()
	if cfg.Mode == "dev" {
		(&core.Base{}).Print("[ Server ]", "http://"+cfg.ServerHost+":"+cfg.ServerPort)
	}
	err := http.ListenAndServe(cfg.ServerHost+":"+cfg.ServerPort, mux)
	if err != nil {
		(&core.Base{}).Print("[ Server ]", err.Error())
	}
}
