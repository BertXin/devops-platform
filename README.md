# 运维平台2.0
## 目录结构

``` bash
.
├── bin
│   └── deploy-system
├── build.sh
├── cmd
│   └── deploy-system
│       ├── init.go
│       └── main.go
├── config
│   └── dev.toml
├── docs
│   ├── docs.go
│   ├── package.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── common
│   │   ├── config                                    # 应用配置
│   │   ├── database                                  # 数据库源
│   │   ├── init                                      # common初始化，包含common下所有子模块的初始化
│   │   │   └── init.go
│   │   ├── log                                       # 日志
│   │   ├── login-user.go
│   │   ├── repository                                # common-repository
│   │   │   └── repostory.go
│   │   ├── service                                   # common-service
│   │   │   └── service.go
│   │   ├── swagger                                   # API文档
│   │   └── web                                       # gin web框架
│   ├── deploy-system
│   │   ├── app                                       # 应用
│   │   ├── build                                     # 构建
│   │   ├── deploy                                    # 部署
│   │   ├── gitlab
│   │   ├── init                                      # 模块初始化，包含deploy-system下所有子模块的初始化
│   │   │   └── init.go  
│   │   ├── middleware                                # 中间件
│   │   ├── productline                               # 产品线
│   │   ├── server                                    # 集群服务器(kubernetes)
│   │   ├── task                                      # 任务
│   │   └── user                                      # 用户
│   │       ├── export.go
│   │       ├── init
│   │       │   └── init.go                           # user模块初始化
│   │       └── internal
│   │           ├── controller                        # controller和路由
│   │           │   ├── controller.go
│   │           │   └── routing.go
│   │           ├── domain                            # 数据库实体、模块常量
│   │           │   ├── constants.go
│   │           │   └── domain.go
│   │           ├── repository                        # 数据库操作
│   │           │   ├── repostory.go
│   │           │   └── repostory_test.go
│   │           └── service                           # 服务
│   │               ├── service.go
│   │               └── service_test.go
│   └── pkg
│       ├── module
│       │   └── module.go
│       └── security
│           └── security.go
├── pkg
│   ├── beans
│   │   ├── factory.go
│   │   └── lifecycle.go
│   ├── common
│   │   └── error.go
│   └── types
│       ├── date.go
│       ├── env.go
│       ├── ints
│       │   ├── array.go
│       │   └── ints.go
│       ├── long.go
│       ├── longs
│       │   ├── array.go
│       │   └── longs.go
│       ├── pagination.go
│       ├── time.go
│       ├── times
│       │   └── times.go
│       └── types_test.go
└── test
    └── user-create_test.go

```
## 文档生成

Package routes 生成swagger文档

文档规则请参考：`https://github.com/swaggo/swag#declarative-comments-format`

### 使用方式：
```
# 下载工具
go get -u github.com/swaggo/swag/cmd/swag@v1.6.7  
(更高的版本会出问题)

# 执行命令生成或更新文档，根据操作系统选择命令

mac:
    swag init --generalInfo  cmd/deploy-system/main.go
windows:
    swag.exe init -g  cmd\deploy-system\main.go
```
默认访问 `http://localhost/swagger/index.html?docExpansion=none`

### 单元测试需要设置GoLand环境变量
Run/Debug Configurations --> Templates --> Go Test

Environment:
```
DEPLOY_SYSTEM_CONFIG_PATH=/Users/xin/Work/devops-platform/config/dev.toml
```
