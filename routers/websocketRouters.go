package routers

import (
	"ginchat/controllers/websocketcon"

	"github.com/gin-gonic/gin"
)

func WebsocketRoutersInit(r *gin.Engine) {
	webRouters := r.Group("/websocket")
	{
		webRouters.GET("/sendmsg", websocketcon.SocketController{}.SendMsg)
		webRouters.GET("/sendUserMsg", websocketcon.SocketController{}.SendUserMsg)
	}
}
