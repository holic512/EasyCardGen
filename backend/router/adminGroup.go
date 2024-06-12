package router

import "github.com/gin-gonic/gin"

func adminGroup(server *gin.Engine) {
	admin := server.Group("/api/admin")
	{

		admin.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{})
		})

	}

}
