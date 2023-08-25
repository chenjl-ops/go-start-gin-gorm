# go-starter-gin-gorm

#### Description

符合Python开发者基于Django框架使用习惯 封装Go基于Gin框架 初始化工程 让开发者只关心 业务实现

* url 聚焦路由功能
* views 聚焦业务实现

#### Feature

* 提供携程Apollo配置，提供统一配置管理
* 提供Gin Url分组管理，提供统一路由插拔策略
* 提供Mysql，提供统一数据库操作Client 使用gorm模块取代之前xorm
* 提供test 模块，熟悉如何完成业务开发

#### 使用

```
## apollo模块为全局所有配置唯一来源
修改apollo模块中 struct中相关配置(包含：mysql相关配置，以及未来可能使用到的所有配置)

## nacos模块为阿里nacos配置管理
同apollo配置管理为同一类服务


## 设置环境变量(正常Prod环境以及Dev Test 已经有统一的环境变量)，仅适用本地开发环境：
export RUNTIME_ENV=dev && export RUNTIME_TENANT=public && export RUNTIME_APP_NAME=go-test-config && export RUNTIME_GROUP=test && export LOG_BASE=debug
go run cmd/app/main.go
```

#### Swagger
```
## install swag
go get -u github.com/swaggo/swag/cmd/swag

## gen docs
swag init -g cmd/app/main.go -o ./api
CI build 参考如下
https://github.com/hashicorp/terraform/tree/main/scripts
```
