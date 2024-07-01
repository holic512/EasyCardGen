package router

import "github.com/gin-gonic/gin"
import "easyCardGen/controller/admin"

func adminGroup(server *gin.Engine) {
	admin1 := server.Group("/api/admin")
	{
		admin1.POST("/addUser", admin.AddUser)
		admin1.GET("/getUserCount", admin.GetUserCount)
		admin1.GET("/getUserInfo", admin.GetUsersInfo)
	}

}
