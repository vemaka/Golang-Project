# Go微服务框架简易开发

## 项目结构
    ├── client # 客户端代码
    │ └── client.go
    ├── gw # 网关相关代码
    │ └── mux.go
    ├── handler # 处理具体业务逻辑的模块
    │ ├── handler.go       # 业务实现
    │ └── server.go 
    ├── middleware # 中间件模块
    │ ├── healthcheck.go    
    │ ├── logging.go       
    │ └── middleware.go
    ├── router # 路由管理
    │ ├── discover.go    # 服务发现
    │ ├── register.go    # 服务注册
    │ └── router.go
    ├── types # 类型定义模块
    │ └── type.go
    ├── go.mod # Go 模块依赖管理文件
    ├── go.sum # Go 模块校验和文件
    └── main.go # 项目入口文件

## 中间件支持
    提供健康检查和日志记录中间件
    
## 启动服务
    go run main.go

## 客户端使用
    进入 client 目录，运行以下命令启动客户端：
    go run client.go
    客户端代码会向服务端/good 路由发送请求，并打印响应结果

## 整体代码流程：
     [统一 API 前置处理]    ------------>  [进行服务注册等操作]
                                                 |
                                                 v
    [解析请求路径以及参数]   <------------    [交给网关 GW]
             |
             v
    [将路径与参数存入上下文] ------------> [router从上下文接收参数]
                                                 |
                                                 v
       [调用相关业务逻辑]    <------------   [进行服务查找]
    

    
    
