package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	// server相关
	AppMode  string
	HttpPort string

	// 数据库相关

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {

	file, err := ini.Load("config/config.ini")
	fmt.Println("==========开始读取config配置文件==========")
	if err != nil {
		fmt.Println("==========读取文件失败==========", err)
	}
	LoadServer(file)
	LoadDate(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":80")
}

func LoadDate(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("172.18.0.3")
	DbPort = file.Section("database").Key("DbPort").MustString(":3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("db_restaurant")
}
