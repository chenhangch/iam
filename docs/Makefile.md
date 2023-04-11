# Makefile 常见功能

| Makefile目标名称          | 说明                           |
|-----------------------|------------------------------|
| lint                  | 静态代码检查，推荐`golangci-lint`     |
| test                  | 单元测试，运行 `go test ./...`      |
| build/build.multiarch | 编译源码，支持不同平台，不同的CPU架构         |
| image/image.multiarch | 构建Docker镜像                   |
| push/push.multiarch   | 构建Docker镜像，并推送到镜像仓库          |
| clean                 | 清理临时文件或者编译后的产物               |
| gen                   | 代码生成，例如要编译生成Protobuf pg.go   |
| deploy                | 部署（可选），一键部署功能，方便测试           |
| release               | 发布功能，比如发布到Docker Hub、GitHub等 |
| help                  | 帮助，告知Makefile有哪些功能，如果执行这些功能  |
| add-copyright         | 版权声明                         |
| *release              | 发布版本                         |
| swagger               | 生成swagger格式的api文档            |
| format                | 格式化代码                        |
| install               | 一键部署项目到测试环境                  |