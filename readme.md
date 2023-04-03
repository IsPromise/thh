# readme.md

## 工具依赖
```
go install golang.org/x/tools/cmd/goimports@latest
```

## 一个简易的gin项目文件

参考了 `laravel` 目录

# todoList

- [ ] 优化启动逻辑，避免在全逻辑命令下的的无意义初始化
  - 方案一 逻辑懒加载
  - 方案二 入口分离
  - 方案三 抽离全逻辑，创建新项目
- [x] 日志链路追踪
- [x] 日志调用栈输出
- [x] resty 效率提升

# how to run

如果你想使用热加载进行开发。

[air](https://github.com/cosmtrek/air)

> [realize](https://github.com/oxequa/realize)
> [fresh](https://github.com/gravityblast/fresh)

git set

```shell
git config user.name 'github用户名'  
git config user.email '邮箱'
```


```shell
cd actor
npm install
npm run dev 
# or 
npm run dev  -- --host 0.0.0.0
npm run build
```

```shell
git clone https://github.com/eatmeatball/thh.git 
cd thh
go mod tidy
go install github.com/cosmtrek/air@latest
air 
```

## 测试相关
```text
go  test  ./...   
// array_test.go 依赖 array.go 和 types.go
go test ./helpers/array_test.go ./helpers/array.go ./helpers/types.go 
```

## 代码格式化相关

```shell    
gofmt -w .
goimports -w .   
```

## 编译相关

```
go build -ldflags="-w -s" .
```
windows
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=arm64
go build

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build
```

powershell
```powershell
// 设置Linux编译环境
$env:CGO_ENABLED="0"
$env:GOOS="linux"
$env:GOARCH="amd64"
 
$env:CGO_ENABLED=""
$env:GOOS=""
$env:GOARCH=""
// 开始编译
go build .

//https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/about/about_environment_variables?view=powershell-5.1
```

mac
```
go build
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

linux
```
go build
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```


## 程序覆盖测试开启及展示

```shell
go build -o ./storage/main -cover 
go build -o ./storage/main -cover -coverpkg=thh/arms,thh/app/http/controllers
mkdir -p storage/mycovdata
export GOCOVERDIR=storage/mycovdata
## 运行你的程序
## ./storage/main serve
go tool covdata percent -i=storage/mycovdata/
go tool covdata textfmt -i=storage/mycovdata/ -pkg="thh,thh/arms/...,thh/app/http/...,thh/app/models/...,thh/app/service/..." -o storage/profile.txt
go tool covdata textfmt -i=storage/mycovdata/ -o storage/profile.txt
go tool cover -func=storage/profile.txt
go tool cover -html=storage/profile.txt
```

```shell
sudo apt install nginx
sudo systemctl start nginx
sudo systemctl stop nginx
sudo systemctl restart nginx
sudo systemctl status nginx
```


ng
```conf
upstream my_server {
        # 将流量代理到你的网站所在的服务器 IP 和端口
    server host.docker.internal:90;
}

server {
    listen 80 default_server;

    # 支持 WebSocket 流量
    location /ws {
        proxy_pass http://my_server;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 86400;
    }

    # 其他 HTTP 流量
    location / {
        proxy_pass http://my_server;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```


# todo

demo收集部分

- 语音聊天室，代码已在gpt找到，有待实际运行

# 如果考虑引入插件

`github.com/traefik/yaegi/interp`
  

bbs开发部分

- thh 登录优化
  - 从登录页面改为弹框登录，登录后隐藏弹框
  - 验证码接口实现
  - 注册页面实现
    - 后续更改为邀请码注册，邀请码要有次数限制，时间限制。
    - 或加入注册后答题后/邮箱验证通过后可解锁相关功能。
- bbs列表页面改为左右视图，左宽右窄。右侧可以添加友链/跳转文章上传/讨论上传页面。
- 讨论/文章类型加入tag
- md上传页面实现
- 详情页也同bbs列表页面视图。右侧展示调整，可以增加评论。
- 后台bbs管理页面。
- 完善文件队列/引入队列
- 增加积分/等级功能
- 增加申请置顶（消耗积分）
- 增加用户设置小尾巴，外链功能
- 增加列表页相关特效
- 一键按照功能（初始化）
- 在线广播功能
- 申请展示功能（侧边留下部分可以申请上墙）
- 文档+后续维护