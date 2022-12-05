package v1

import (
	"fmt"
	"net/http"
	"zhu/myrest/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users, err := model.GetAllUsers()
	for _, user := range users {
		fmt.Println(user)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}
