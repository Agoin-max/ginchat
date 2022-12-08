package main

import (
	"ginchat/docs"
	"ginchat/routers"
	"ginchat/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 加载配置
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 全局中间件
	// r.Use(middlewares.InitMiddlewareOne, middlewares.InitMiddlewareTwo)

	// 路由
	routers.InitRouter(r)

	// 启动服务
	r.Run("0.0.0.0:8000")
}
