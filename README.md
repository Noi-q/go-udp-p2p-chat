### UDP协议P2P通信(支持本地和局域网通信)
- 可以实现公网UDP通信

### 注:
- 运行之前请到client.go文件的第29行局域网IP改为自己的局域网IP然后打包编译

#### 编译
```go
go build main.go
```
```go
go build client.go
```

#### 代码运行
- 启动UDP服务器
```go
go run main.go
```
- 启动客户端用户1
```go
go run client.go 30303 noi1
```
- 启动客户端用户2
```go
go run client.go 30301 noi2
```
#### 本地exe运行 和 局域网exe运行
- 启动UDP服务器
```
main.exe
```
- 启动客户端用户1
```
client.exe 30303 noi1
```
- 启动客户端用户2
```
client.exe 30301 noi2
```

- 客户端运行基本参数
- go
```go
go run client.go 端口号 用户名
```
- exe
```
client.exe 端口号 用户名
```

- 退出 `ctrl+c`