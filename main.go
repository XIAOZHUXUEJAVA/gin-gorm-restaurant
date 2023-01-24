package main

import (
	"log"
	"zhu/myrest/model"
	"zhu/myrest/routes"
)

func main() {
	log.Println("====================欢迎使用餐馆点餐系统====================")
	model.InitDb()
	routes.InitRouter()
}
