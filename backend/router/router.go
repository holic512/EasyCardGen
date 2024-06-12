package router

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {

	//管理员路由组
	adminGroup(server)

}
