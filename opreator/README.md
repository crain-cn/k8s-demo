```
$ operator-sdk version
operator-sdk version: "v1.1.0", commit: "9d27e224efac78fcc9354ece4e43a50eb30ea968", kubernetes version: "v1.18.2", go version: "go1.15 darwin/amd64", GOOS: "darwin", GOARCH: "amd64"
$ go version
go version go1.15.3 darwin/amd64
```

operator-sdk init --domain ydzs.io --license apache2 --owner "zhaoyu"


使用 `operator-sdk init` 命令创建新的 Operator 项目后，项目目录就包含了很多生成的文件夹和文件。

- go.mod/go.sum  - Go Modules 包管理清单，用来描述当前 Operator 的依赖包。
- main.go 文件，使用 operator-sdk API 初始化和启动当前 Operator 的入口。
- deploy - 包含一组用于在 Kubernetes 集群上进行部署的通用的 Kubernetes 资源清单文件。
- pkg/apis - 包含定义的 API 和自定义资源（CRD）的目录树，这些文件允许 sdk 为 CRD 生成代码并注册对应的类型，以便正确解码自定义资源对象。
- pkg/controller - 用于编写所有的操作业务逻辑的地方
- version - 版本定义
- build - Dockerfile 定义目录


operator-sdk create api --group app --version v1beta1 --kind AppService