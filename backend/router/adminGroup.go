package router

import "github.com/gin-gonic/gin"

func adminGroup(server *gin.Engine) {
	admin := server.Group("/api/admin")
	{
		admin.POST("/addUser")

	}

}
