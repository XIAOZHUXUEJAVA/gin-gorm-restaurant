package v1

import (
	"net/http"
	"zhu/myrest/model"
	"zhu/myrest/proto"
	"zhu/myrest/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

func InsertTable(c *gin.Context) {
	var data proto.ReqAddTable
	// 将传过来的json绑定到data中
	_ = c.ShouldBindJSON(&data)

	// table := &model.BookTable{
	// 	BookName:   data.BookName,
	// 	BookPhone:  data.BookPhone,
	// 	BookPeople: data.BookPeople,
	// 	BookTables: data.BookTables,
	// 	UserId:     data.UserId,
	// 	BookWhen:   data.BookWhen,
	// 	BookNote:   data.BookNote,
	// }
	table := &model.BookTable{}
	deepcopier.Copy(data).To(table)
	result, code := model.InsertTable(table)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})
}
