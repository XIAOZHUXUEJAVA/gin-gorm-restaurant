package utils

import (
	"log"

	"github.com/spf13/viper"
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
	// viper.SetConfigName("config")  // 配置文件名（不带扩展名）
	// viper.AddConfigPath("config/") // 配置文件路径
	// viper.SetConfigType("ini")     // 配置文件类型
	viper.SetConfigFile("config/config.yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("====================读取文件失败====================%s\n", err)
	}
	LoadServer()
	LoadDate()
}

// 使用viper

func LoadServer() {
	AppMode = viper.GetString("server.AppMode")
	HttpPort = viper.GetString("server.HttpPort")
}

func LoadDate() {
	Db = viper.GetString("database.Db")
	DbHost = viper.GetString("database.DbHost")
	DbPort = viper.GetString("database.DbPort")
	DbUser = viper.GetString("database.DbUser")
	DbPassWord = viper.GetString("database.DbPassWord")
	DbName = viper.GetString("database.DbName")
	// fmt.Println("Db:", viper.GetString("database.Db"))
	// fmt.Println("DbHost:", viper.GetString("database.DbHost"))
	// fmt.Println("DbPort:", viper.GetString("database.DbPort"))
	// fmt.Println("DbUser:", viper.GetString("database.DbUser"))
	// fmt.Println("DbPassWord:", viper.GetString("database.DbPassWord"))
	// fmt.Println("DbName:", viper.GetString("database.DbName"))
}

// 使用ini.Load()
// 	file, err := ini.Load("config/config.ini")
// 	log.Println("====================开始读取config配置文件====================")
// 	if err != nil {
// 		log.Fatalln("====================读取文件失败====================\n", err)
// 	}
// 	LoadServer(file)
// 	LoadDate(file)
// }

// func LoadServer(file *ini.File) {
// 	AppMode = file.Section("server").Key("AppMode").MustString("debug")
// 	HttpPort = file.Section("server").Key("HttpPort").MustString(":80")
// }

// func LoadDate(file *ini.File) {
// 	Db = file.Section("database").Key("Db").MustString("mysql")
// 	DbHost = file.Section("database").Key("DbHost").MustString("172.18.0.3")
// 	DbPort = file.Section("database").Key("DbPort").MustString(":3306")
// 	DbUser = file.Section("database").Key("DbUser").MustString("root")
// 	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
// 	DbName = file.Section("database").Key("DbName").MustString("db_restaurant")
// }
