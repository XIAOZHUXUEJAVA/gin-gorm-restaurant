package v1

import (
	"log"
	"net/http"
	"zhu/myrest/errmsg"
	"zhu/myrest/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	users, err := model.GetAllUsers()
	// for _, user := range users {
	// 	log.Println(user)
	// }
	log.Println(users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}
func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	data, code := model.GetUserByEmail(email)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
