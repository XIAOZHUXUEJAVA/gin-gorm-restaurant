package v1

import (
	"net/http"
	"strconv"
	"zhu/myrest/errmsg"
	"zhu/myrest/model"

	"github.com/gin-gonic/gin"
)

// func GetAllFood(c *gin.Context) {
// 	foods, err := model.GetAllFood()

//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//		} else {
//			log.Println("GetAllFood Success!")
//			c.JSON(http.StatusOK, foods)
//		}
//	}

// 获取所有菜品信息
func GetAllFood(c *gin.Context) {
	data, code := model.GetAllFood()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 根据id获取菜品信息
func GetFoodById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.FRONTEND_PARAMS_ERROR,
			"message": errmsg.GetErrMsg(errmsg.FRONTEND_PARAMS_ERROR),
		})
	}
	data, code := model.GetFoodById(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}
