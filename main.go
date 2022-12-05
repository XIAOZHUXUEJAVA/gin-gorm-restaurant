package main

import (
	"fmt"
	"zhu/myrest/model"
	"zhu/myrest/routes"
)

func main() {
	fmt.Println("hello, my gin_gorm_restaurant")
	model.InitDb()
	routes.InitRouter()
}
