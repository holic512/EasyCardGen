package router

import (
	controlleradmin "easyCardGen/controller/admin"
	"github.com/gin-gonic/gin"
)

func adminGroup(server *gin.Engine) {
	admin1 := server.Group("/api/admin")
	{
		admin1.POST("/addUser", controlleradmin.AddUser)
		admin1.GET("/getUserCount", controlleradmin.GetUserCount)
		admin1.GET("/getUserInfo", controlleradmin.GetUsersInfo)

		admin1.POST("/addAccess", controlleradmin.AddAccess)
	}

}
