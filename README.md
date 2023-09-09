# gin-gorm-restaurant
使用go + gorm + gin +vue3 构建的一个简单的点餐系统前台


## 库或工具

## cosmtrek/air

`cosmtrek/air` 是 Go 开发中常用的热加载工具，其通过监视 Go 源文件的变化来实现热加载，从而编译 Go 代码并重新启动应用程序。`cosmtrek/air` 的主要作用在于加快开发流程中的编译、构建和部署过程，使代码的修改、编译、测试和部署更加顺畅和高效。

## spf13/viper

`spf13/viper` 是 Go 语言中一个功能强大的配置管理库，支持多种配置文件格式和解析方式，包括 JSON、YAML、TOML、HCL、INI 等等。它可以用于读取和解析应用程序的配置文件，使配置管理更加简单、易用和灵活。

## go-playground/validator/v10

`go-playground/validator/v10` 是 Go 语言中一个功能强大的数据验证库，用于对数据结构进行验证和校验，包括验证字符串、数字、时间、结构体等多种数据类型。它可以有效提高应用程序的数据质量和数据安全性。


## 使用
说明： 数据库的使用为MySQL5，其他版本没有测试

1. 修改config/config.yaml文件中的数据库信息:
```yaml
server:
  AppMode: debug
  HttpPort: ":9090"
database:
  Db: mysql
  DbHost: 172.18.0.3
  DbPort: ":3306"
  DbUser: root
  DbPassWord: "123456"
  DbName: db_restaurant
```

2. 本地启动，在根目录下

```shell
$ go mod tidy
$ go run .
$ cd frontend
$ npm install
$ npm run serve
```
