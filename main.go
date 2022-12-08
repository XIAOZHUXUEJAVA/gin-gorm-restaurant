package main

import (
	"log"
	"zhu/myrest/model"
	"zhu/myrest/routes"
)

func main() {
	log.Println("hello, this is a restaurant order system")
	model.InitDb()
	routes.InitRouter()
}
