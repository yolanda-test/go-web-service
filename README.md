# Go Web Service

本项目是一个使用 Go 和 Gin 框架构建的简单 Web 服务，注重模块化和测试。

## 项目结构

```
go-web-service
├── cmd
│   └── server
│       └── main.go          # 服务入口
├── internal
│   ├── handlers
│   │   ├── handler.go       # 路由处理函数（目前仅 /example 路由）
│   │   └── handler_test.go  # 路由处理单元测试
│   ├── models
│   │   ├── store.go         # 内存存储实现
│   │   └── store_test.go    # 存储层单元测试
│   └── integration
│       └── integration_test.go # 集成测试
├── pkg
│   └── utils.go             # 工具函数
├── go.mod                   # Go module 定义
├── go.sum                   # 依赖校验
└── README.md                # 项目说明
```

## 快速开始

### 环境要求

- Go 1.18 或更高
- Git

### 安装依赖

```sh
go mod tidy
```

### 启动服务

```sh
go run cmd/server/main.go
```
服务默认监听 `http://localhost:8080`。

### 路由说明

- `GET /example`  
  返回：`{"message": "Hello, World!"}`

### 运行测试

```sh
go test ./...
```
包括单元测试和集成测试。

## 贡献

欢迎贡献代码！请通过 issue 或 pull request 提交建议或修复。

## 许可证

MIT License，详见 LICENSE 文件。