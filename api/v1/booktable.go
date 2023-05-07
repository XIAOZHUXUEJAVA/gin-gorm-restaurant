package v1

import (
	"log"
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

	table := &model.BookTable{}
	if err := deepcopier.Copy(data).To(table); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"data":    nil,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
		})
		log.Fatalf("转换预订餐桌数据出错: %s", err)
	}
	result, code := model.InsertTable(table)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})
}
