Go 微服务框架 简易开发

项目结构
├── client # 客户端代码
│ └── client.go
├── gw # 网关相关代码
│ └── mux.go
├── handler # 处理具体业务逻辑的模块
│ ├── handler.go
│ └── server.go
├── middleware # 中间件模块，用于处理跨切面关注点
│ ├── healthcheck.go
│ ├── logging.go
│ └── middleware.go
├── router # 路由相关代码
│ ├── discover.go
│ ├── register.go
│ └── router.go
├── types # 类型定义模块
│ └── type.go
├── go.mod # Go 模块依赖管理文件
├── go.sum # Go 模块校验和文件
└── main.go # 项目入口文件

中间件支持：提供健康检查和日志记录中间件

安装依赖
go mod download

启动服务
go run main.go

客户端使用
进入 client 目录，运行以下命令启动客户端：
go run client.go

客户端代码会向服务端/good 路由发送请求，并打印响应结果。

整体代码流程：
进入 统一 API --前置处理，进行服务注册等操作
-> 交给 GW 处理 --解析请求路径以及参数，并存放到上下文中
-> router 从上下文中接收参数，并且进行服务查找
-> 查找调用相关业务
