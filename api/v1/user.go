package v1

import (
	"net/http"
	"zhu/myrest/model"
	"zhu/myrest/proto"
	"zhu/myrest/utils/errmsg"
	"zhu/myrest/utils/validator"

	"github.com/ulule/deepcopier"

	"github.com/gin-gonic/gin"
)

// 测试代码
func GetAllUser(c *gin.Context) {
	if users, err := model.GetAllUsers(); err != nil {
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
	msg, code := validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	user := &model.User{}
	deepcopier.Copy(data).To(user)
	result, code := model.InsertUser(user)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})

}
