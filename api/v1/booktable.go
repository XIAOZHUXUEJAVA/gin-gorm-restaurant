package v1

import (
	"net/http"
	"zhu/myrest/errmsg"
	"zhu/myrest/model"
	"zhu/myrest/proto"

	"github.com/gin-gonic/gin"
)

func InsertTable(c *gin.Context) {
	var data proto.ReqAddTable
	// 将传过来的json绑定到data中
	_ = c.ShouldBindJSON(&data)
	table := &model.BookTable{
		BookName:   data.BookName,
		BookPhone:  data.BookPhone,
		BookPeople: data.BookPeople,
		BookTables: data.BookTables,
		UserId:     data.UserId,
		BookWhen:   data.BookWhen,
		BookNote:   data.BookNote,
	}
	result, code := model.InsertTable(table)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})
}
