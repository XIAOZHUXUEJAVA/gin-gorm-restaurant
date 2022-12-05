package routes

import (
	v1 "zhu/myrest/api/v1"
	"zhu/myrest/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())

	routeV1 := r.Group("api/users")
	{
		routeV1.GET("/all", v1.GetAllUser)
	}
	r.Run(utils.HttpPort)

}
