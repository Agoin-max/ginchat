package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 响应说明
	Data interface{} `json:"data"` // 数据结构体
}

func Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

func Fail(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": -1,
		"msg":  msg,
		"data": data,
	})
}
