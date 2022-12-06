package routes

import (
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
		routeV1.GET("users/all", v1.GetAllUser)
		routeV1.GET("foods", v1.GetAllFood)
	}
	r.Run(utils.HttpPort)
}
