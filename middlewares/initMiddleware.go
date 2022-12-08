package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitMiddlewareOne(ctx *gin.Context) {
	fmt.Println("1-中间件-one")
	//调用该请求的剩余处理程序
	ctx.Next()
	fmt.Println("2-中间件-one")
}

func InitMiddlewareTwo(ctx *gin.Context) {
	fmt.Println("1-中间件-two")
	//调用该请求的剩余处理程序
	ctx.Next()
	fmt.Println("2-中间件-two")

}
