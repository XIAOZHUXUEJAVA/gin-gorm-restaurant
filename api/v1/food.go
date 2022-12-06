package v1

import (
	"fmt"
	"net/http"
	"zhu/myrest/model"

	"github.com/gin-gonic/gin"
)

func GetAllFood(c *gin.Context) {
	foods, err := model.GetAllFood()
	for _, food := range foods {
		fmt.Println(food)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, foods)
	}
}
