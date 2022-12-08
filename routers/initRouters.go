package routers

import "github.com/gin-gonic/gin"

// 初始化路由
func InitRouter(r *gin.Engine) {
	// AdminRoutersInit(r)
	ApiRoutersInit(r)
	WebsocketRoutersInit(r)
}
