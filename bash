#!/bin/bash

# 配置
s=$1
name='webmis'                 # 项目名称
version='3.0.0'               # 版本
index='main.go'               # 入口文件
cli='cli/main.go'             # Cli命令行
log='public/server.log'       # 运行日志

# 运行
if [ "$s" == "serve" ]; then
  air
# 安装
elif [ "$s" == "install" ]; then
  go clean --modcache && go get -v
# 清理
elif [ "$s" == "clear" ]; then
  go mod tidy
# 打包
elif [ "$s" == "build" ]; then
  go build && mv $name "$name$version"
# 预览
elif [ "$s" == "http" ]; then
  "./$name$version"
# Server-启动
elif [ "$s" == "start" ]; then
  nohup "./$name$version" > $log &
# Server-停止
elif [ "$s" == "stop" ]; then
  ps -aux | grep "./$name$version" | grep -v grep | awk {'print $2'} | xargs kill
# Socket-运行
elif [ "$s" == "socket" ]; then
  go run $cli socket start
# Socket-启动
elif [ "$s" == "socketStart" ]; then
  go run $cli socket start &
# Socket-停止
elif [ "$s" == "socketStop" ]; then
  ps -aux | grep "$cli socket start" | grep -v grep | awk {'print $2'} | xargs kill
else
  echo "----------------------------------------------------"
  echo "[use] ./bash <command>"
  echo "----------------------------------------------------"
  echo "  <command>"
  echo "    serve                 运行: air"
  echo "    install               安装依赖包: go get -v"
  echo "    clear                 清理依赖包: go mod tidy"
  echo "    build                 打包: go build"
  echo "    http                  预览: ./$name"
  echo "  <Server>"
  echo "    start                 启动"
  echo "    stop                  停止"
  echo "  <WebSocket>"
  echo "    socket                运行"
  echo "    socketStart           启动"
  echo "    socketStop            停止"
  echo "----------------------------------------------------"
fi
