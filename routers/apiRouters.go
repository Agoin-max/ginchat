package routers

import (
	"ginchat/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.POST("/user/register", api.ApiController{}.Register)
		apiRouters.POST("/user/login", api.ApiController{}.Login)
		apiRouters.GET("/user/searchFriends", api.ApiController{}.SearchFriends)
	}
}
