@echo off
CHCP 65001 >nul

REM 配置
set s=%1%
set name=webmis-java
set version=3.0.0
set index=main.go
set cli=cli/main.go
set go_dir=D:\soft\go
set go_url=https://golang.google.cn/dl/go1.26.0.windows-amd64.msi
set go_proxy=https://goproxy.cn,direct

@REM Go环境
go version >nul 2>&1
if %errorLevel% neq 0 (
  @REM 是否存在目录
  if not exist "%go_dir%\" (
    md "%go_dir%" >nul 2>&1
    if exist "%go_dir%\" (
      echo [✓] 创建目录: %go_dir%
    ) else (
      echo [✗] 创建目录: %go_dir%
    )
  )
  if not exist "%go_dir%\bin\go.exe" (
    @REM 下载文件
    echo [✓] 下载文件: %go_url%
    curl -L "%go_url%" -o go.msi
    @REM 安装
    echo [✓] 正在解压: go.msi
    msiexec /a "go.msi" /qn TARGETDIR="%go_dir%"
    xcopy "%go_dir%\Go\*" "%go_dir%\" /e /y >nul
    rmdir /s /q "%go_dir%\Go" >nul
    del "%go_dir%\go.msi"
    echo [✓] 解压文件: go.msi 到 %go_dir%
    @REM 清除文件
    del go.msi >nul 2>&1
  )
  @REM 配置环境变量
  echo [✓] 安装成功：请手动添加环境变量
  echo GOROOT %go_dir%
  echo GOPROXY %go_proxy%
  echo GOBIN %go_dir%\bin
  echo PATH %go_dir%\bin
  pause
  @REM 验证
  go version >nul 2>&1
  if %errorLevel% neq 0 (
    @REM 临时环境变量
    set GOROOT=%go_dir%
    set GOPROXY=%go_proxy%
    set GOBIN=%go_dir%\bin
    set PATH=%PATH%;%go_dir%\bin
    @REM 查看版本
    go version
  )
)

REM 运行
if "%s%"=="serve" (
  air
REM 安装
) else if "%s%"=="install" (
  go clean --modcache && go get -v
  go install github.com/air-verse/air@latest
  echo [✓] 运行: .\cmd serve
REM 清理
) else if "%s%"=="clear" (
  go mod tidy
REM 打包
) else if "%s%"=="build" (
  go build
REM 预览
) else if "%s%"=="http" (
  .\%name%%version%
REM Socket-运行
) else if "%s%"=="socket" (
  go run %cli% socket start
) else (
  echo ----------------------------------------------------
  echo [use] .\cmd ^<command^>
  echo ----------------------------------------------------
  echo ^<command^>
  echo   serve         运行: air
  echo   install       安装依赖包: go get -v
  echo   clear         清理依赖包: go mod tidy
  echo   build         打包: go build
  echo   http          预览: .\%name%
  echo ^<WebSocket^>
  echo   socket        运行
  echo ----------------------------------------------------
)