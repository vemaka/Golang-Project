# Go微服务框架简易开发

## 项目结构
.
├── client/                 # 客户端代码
│   ├── client.go
│   ├── gw/                 # 网关相关代码
│   ├── mux.go
│   ├── handler/            # 处理具体业务逻辑的模块
│   │   └── handler.go
│   ├── server.go
│   ├── middleware/         # 中间件模块，用于处理跨切面关注点
│   │   ├── healthcheck.go
│   │   ├── logging.go
│   │   └── middleware.go
│   └── router/             # 路由相关代码
│       ├── discover.go
│       ├── register.go
│       └── router.go
├── types/                  # 类型定义模块
│   └── type.go
├── go.mod                  # Go 模块依赖管理文件
├── go.sum                  # Go 模块校验和文件
└── main.go                 # 项目入口文件
