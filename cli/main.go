package main

import (
	"os"
	"webmis/app/task"
)

func main() {
	// 参数
	args := os.Args
	param := map[string]string{}
	param["c"] = "main"
	param["a"] = "index"
	param["p"] = "index"
	if len(args) > 1 {
		param["c"] = args[1]
	}
	if len(args) > 2 {
		param["a"] = args[2]
	}
	if len(args) > 3 {
		param["a"] = args[3]
	}
	// 任务
	switch {
	case param["c"] == "socket":
		if param["a"] == "start" {
			// (&task.Socket{}).Start()
		}
	case param["c"] == "logs":
		if param["a"] == "log" {
			// (&task.Logs{}).Log()
		}
	default:
		(&task.Main{}).Index()
	}
}
