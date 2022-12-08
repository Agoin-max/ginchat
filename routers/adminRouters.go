package routers

import (
	"ginchat/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.UserController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.List)
		adminRouters.GET("/article", admin.UserController{}.Article)

		adminRouters.GET("/user/list", admin.UserController{}.UserList)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.POST("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)
	}
}
