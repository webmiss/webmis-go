# webmis-go
采用Go + Redis + MariaDB开发的轻量级HMVC基础框架，目录结构清晰，支持CLI方式访问资料方便执行定时脚本。包括HMVC模块化管理、自动路由、CLI命令行、Socket通信、redis缓存、Token机制等功能，提供支付宝、微信、文件上传、图像处理、二维码等常用类。

**演示**
- 使用文档( [https://webmis.vip/](https://webmis.vip/go/install/index) )
- 网站-API( [https://go.webmis.vip/](https://go.webmis.vip/) )
- 前端-API( [https://go.webmis.vip/api](https://go.webmis.vip/api) )
- 后台-API( [https://go.webmis.vip/admin](https://go.webmis.vip/admin) )

## 安装
```bash
# 下载
$ git clone https://github.com/webmiss/webmis-go.git
$ cd webmis-go

# Linux、MacOS
./bash install

# Windows 11
.\cmd install
```

## 开发环境
```bash
# Linux、MacOS
./bash serve
./bash socketServer

# Windows 11
.\cmd serve
.\cmd socketServer
```
- 浏览器访问 http://127.0.0.1:9030/

## 生产环境
### 交换分区( 编译时内存不足 )
```bash
# 创建文件
fallocate -l 4G /swapfile
# 设置权限
chmod 600 /swapfile
# 格式化
mkswap /swapfile
# 启用
swapon /swapfile
# 优化
echo 'vm.swappiness=10' >> /etc/sysctl.conf
sysctl -p
# 查看交换空间信息
free -m
```
- 开机启动: 编辑 /etc/fstab 文件，添加内容: /swapfile none swap sw 0 0

### Nginx
```bash
upstream go {
    server localhost:9030;
}
server {
    server_name  go.webmis.vip;
    set $root_path /home/www/webmis/go/public;
    root $root_path;
    index index.html;

    location / {
        proxy_pass http://go;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    location ~* ^/(upload|favicon.png)/(.+)$ {
        root $root_path;
        add_header 'Access-Control-Allow-Origin' '*';
        add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
        add_header 'Access-Control-Allow-Headers' 'Content-Type, Authorization';
        if ($request_method = 'OPTIONS') { return 204; }
    }

}
```

## 项目结构
```plaintext
webmis-go/
├── app
│    ├── config                   // 配置文件
│    ├── librarys                 // 第三方类
│    ├── models                   // 模型
│    └── modules                  // 模块
│    │    ├── admin              // 后台
│    │    ├── api                // 应用
│    │    └── home               // 网站
│    ├── service                  // 项目服务类
│    ├── task                     // 任务类
│    ├── util                     // 工具类
│    └── views                    // 视图文件
├── cli
│    └── main.go                  // 命令行入口: 路由配置
├── core
│    ├── Base.go                  // 基础类
│    ├── Controller.go            // 基础控制器
│    ├── Model.go                 // 基础模型
│    ├── MySQLConnectionPool.go   // MySQL 连接池
│    └── Redis.go                 // 缓存数据库( 连接池 )
├── public                         // 静态资源
│    ├── upload                   // 上传目录
│    └── favicon.png              // 图标
├── tmp                            // 热加载缓存( go install github.com/air-verse/air@latest )
├── bash                           // Linux/MacOS 启动脚本
├── cmd.bat                        // Windows 启动脚本
└── main.go                        // 入口文件: 路由配置
```
