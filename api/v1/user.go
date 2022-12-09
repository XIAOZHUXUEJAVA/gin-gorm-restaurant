package v1

import (
	"log"
	"net/http"
	"zhu/myrest/errmsg"
	"zhu/myrest/model"
	"zhu/myrest/proto"

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
func InsertUser(c *gin.Context) {
	var data proto.ReqAddUser
	_ = c.ShouldBindJSON(&data)
	user := &model.User{
		UserName:     data.UserName,
		UserEmail:    data.Email,
		UserPhone:    data.Phone,
		UserPassword: data.Password,
		UserBirth:    data.Birth,
		UserGender:   data.Gender,
	}
	result, code := model.InsertUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})

}
