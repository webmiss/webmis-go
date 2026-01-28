package main

import (
	"fmt"
	"os"
	"webmis/app/task"
)

func main() {
	// 参数
	param := os.Args
	fmt.Println(param[1], param[2])
	// 任务
	switch {
	case param[1] == "socket":
		if param[2] == "start" {
			// (&task.Socket{}).Start()
		}
	case param[1] == "logs":
		if param[2] == "log" {
			// (&task.Logs{}).Log()
		}
	default:
		(&task.Main{}).Index()
	}
}
