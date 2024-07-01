package router

import (
	controller_admin "easyCardGen/controller/admin"
	"github.com/gin-gonic/gin"
)

func adminGroup(server *gin.Engine) {
	admin1 := server.Group("/api/admin")
	{
		admin1.POST("/addUser", controller_admin.AddUser)
		admin1.GET("/getUserCount", controller_admin.GetUserCount)
		admin1.GET("/getUserInfo", controller_admin.GetUsersInfo)
	}

}
