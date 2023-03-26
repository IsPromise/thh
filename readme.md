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
 
// 开始编译
go build .
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