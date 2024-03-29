package routes

import (
	"log"
	v1 "zhu/myrest/api/v1"
	"zhu/myrest/middleware"
	"zhu/myrest/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	routeV1 := r.Group("api")
	{
		// 这里我们一定要和前端的url对接好
		// user相关
		routeV1.GET("users/all", v1.GetAllUser)
		routeV1.GET("users/:email", v1.GetUserByEmail)
		routeV1.POST("users/", v1.InsertUser)

		// booktable相关
		routeV1.POST("booking", v1.InsertTable)

		// food相关
		routeV1.GET("foods", v1.GetAllFood)
		routeV1.GET("foods/:id", v1.GetFoodById)
	}
	log.Println("====================正在监听端口: 9090====================")
	r.Run(utils.HttpPort)

}
