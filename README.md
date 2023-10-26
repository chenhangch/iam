> ciam学习自github.com/marmotedu/iam
> 
# ciam
ciam 一个Web服务，用于给第三方用于提供资源访问控制服务。

<!-- 简短的话描述项目 -->

## 功能特性

<!-- 核心功能 -->

## 软件架构
```protobuf
├─api
│  ├─openapi
│  └─swagger
├─build
│  ├─ci
│  ├─docker
│  │  └─iam-api
│  └─package
├─cmd
│  └─iam-api
├─configs
├─docs
│  ├─docker
│  ├─guide
│  │  └─zh-CN
│  │      └─api
│  └─images
├─examples
│  ├─errors_ex
│  └─gopractise-demo
│      ├─gorm
│      └─swagger
│          ├─api
│          └─docs
├─internal
│  ├─apiserver
│  │  ├─api
│  │  ├─config
│  │  ├─controller
│  │  │  └─v1
│  │  │      └─secret
│  │  ├─options
│  │  ├─service
│  │  │  └─v1
│  │  ├─store
│  │  │  ├─fake
│  │  │  └─mysql
│  │  └─testing
│  └─pkg
│      ├─code
│      ├─logger
│      ├─middleware
│      │  └─auth
│      ├─options
│      │  └─serverOptions
│      ├─server
│      ├─util
│      └─validation
├─pkg
│  ├─core
│  ├─db
│  └─log
│      ├─cronlog
│      ├─distribution
│      ├─example
│      │  ├─context
│      │  ├─simple
│      │  └─vlevel
│      ├─klog
│      └─logrus
├─scripts
│  ├─install
│  ├─lib
│  └─make-rules
├─test
├─third_party
└─tools

```

<!-- 项目的构架 -->

## 快速开始
### 依赖检查

<!-- 项目的依赖，依赖的包，工具等 -->

### 构建

<!-- 如何构建项目 -->

### 运行

<!-- 如何运行项目 -->

## 使用指南

<!-- 如何使用该项目 -->
启动api-server服务：
```bash
go run /cmd/iam/iam-api.go
```

## 如何贡献

<!-- 告诉其他开发者如何共享项目 -->

## 社区（可选）

## 关于作者

<!-- 项目的作者 -->

## 许可证

<!-- 项目的开源许可证 -->