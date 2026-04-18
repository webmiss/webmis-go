package main

import (
	"net/http"
	"webmis/app/config"
	"webmis/app/modules/admin"
	"webmis/app/modules/api"
	"webmis/app/modules/web"
	"webmis/core"
)

/* 允许跨域请求 */
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Range, Content-Disposition, Content-Description, Authorization")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Max-Age", "2592000")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// 路由
	mux := http.NewServeMux()
	// 网站
	mux.HandleFunc("/", (&web.Index{}).Index)
	mux.HandleFunc("/index", (&web.Index{}).Index)
	mux.HandleFunc("/index/index", (&web.Index{}).Index)
	// API
	mux.HandleFunc("/api", (&api.Index{}).Index)
	mux.HandleFunc("/api/index", (&api.Index{}).Index)
	mux.HandleFunc("/api/index/index", (&api.Index{}).Index)
	// 后台
	mux.HandleFunc("/admin", (&admin.Index{}).Index)
	mux.HandleFunc("/admin/index", (&admin.Index{}).Index)
	mux.HandleFunc("/admin/index/index", (&admin.Index{}).Index)
	mux.HandleFunc("/admin/index/version", (&admin.Index{}).Version)
	mux.HandleFunc("/admin/index/holiday", (&admin.Index{}).Holiday)
	mux.HandleFunc("/admin/msg/list", (&admin.Msg{}).List)
	mux.HandleFunc("/admin/user/login", (&admin.User{}).Login)
	mux.HandleFunc("/admin/user/token", (&admin.User{}).Token)
	mux.HandleFunc("/admin/sys_file/list", (&admin.SysFile{}).List)
	mux.HandleFunc("/admin/sys_file/mkdir", (&admin.SysFile{}).Mkdir)
	mux.HandleFunc("/admin/sys_file/rename", (&admin.SysFile{}).Rename)
	mux.HandleFunc("/admin/sys_file/remove", (&admin.SysFile{}).Remove)
	mux.HandleFunc("/admin/sys_file/upload", (&admin.SysFile{}).Upload)
	mux.HandleFunc("/admin/sys_file/down", (&admin.SysFile{}).Down)
	mux.HandleFunc("/admin/sys_menus/get_menus_perm", (&admin.SysMenus{}).GetMenusPerm)
	// 启动
	cfg := config.Env()
	if cfg.Mode == "dev" {
		(&core.Base{}).Print("[ Server ]", "http://"+cfg.ServerHost+":"+cfg.ServerPort)
	}
	err := http.ListenAndServe(cfg.ServerHost+":"+cfg.ServerPort, corsMiddleware(mux))
	if err != nil {
		(&core.Base{}).Print("[ Server ]", err.Error())
	}
}
