package router

import (
	controlleradminstoreaccess "easyCardGen/controller/admin/store_access"
	controlleradminuser "easyCardGen/controller/admin/user"

	"github.com/gin-gonic/gin"
)

func adminGroup(server *gin.Engine) {
	admin1 := server.Group("/api/admin")
	{
		admin1.POST("/addUser", controlleradminuser.AddUser)
		admin1.GET("/getUserCount", controlleradminuser.GetUserCount)
		admin1.GET("/getUserInfo", controlleradminuser.GetUsersInfo)
		admin1.GET("/getUserList", controlleradminuser.GetUserList)

		admin1.POST("/addAccess", controlleradminstoreaccess.AddAccess)
		admin1.GET("/accessList", controlleradminstoreaccess.GetAccessList)
	}

}
