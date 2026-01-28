# webmis-go
采用Go + Redis + MariaDB开发的轻量级HMVC基础框架，目录结构清晰，支持CLI方式访问资料方便执行定时脚本。包括HMVC模块化管理、自动路由、CLI命令行、Socket通信、redis缓存、Token机制等功能，提供支付宝、微信、文件上传、图像处理、二维码等常用类。

**演示**
- 使用文档( [https://webmis.vip/](https://webmis.vip/go/install/index) )
- 网站-API( [https://go.webmis.vip/](https://go.webmis.vip/) )
- 前端-API( [https://go.webmis.vip/api/](https://go.webmis.vip/api/) )
- 后台-API( [https://go.webmis.vip/admin/](https://go.webmis.vip/admin/) )

## 安装
```bash
$ git clone https://github.com/webmiss/webmis-go.git
$ cd webmis-go
$ go clean --modcache && go get -v
```

## 运行
```bash
# Linux、MacOS
./bash serve
./bash socketServer
# Windows
.\cmd serve
.\cmd socketServer
# 测试Socket
go run cli.go socket client admin '{"type":"","msg":"\u6d4b\u8bd5"}'
# 命令行: 控制器->方法(参数...)
go run cli.go main index params
```
- 浏览器访问 http://127.0.0.1:9030/

## 项目结构
```plaintext
webmis-go/
├── app
│    ├── config                 // 配置文件
│    ├── librarys               // 第三方类
│    ├── models                 // 模型
│    └── modules                // 模块
│    │    ├── admin            // 后台
│    │    ├── api              // 应用
│    │    └── home             // 网站
│    ├── service                // 项目服务类
│    ├── task                   // 任务类
│    ├── util                   // 工具类
│    └── views                  // 视图文件
├── cli
│    └── main.go                // 命令行入口: 路由配置
├── core
│    ├── Base.go                // 基础类
│    ├── Controller.go          // 基础控制器
│    ├── Model.go               // 基础模型
│    ├── Redis.go               // 缓存数据库
│    └── View.go                // 基础视图
├── public                       // 静态资源
│    ├── upload                 // 上传目录
│    └── favicon.png            // 图标
├── tmp                          // 热加载缓存( go install github.com/air-verse/air@latest )
├── bash                         // Linux/MacOS 启动脚本
├── cmd.bat                      // Windows 启动脚本
└── main.go                      // 入口文件: 路由配置
```
